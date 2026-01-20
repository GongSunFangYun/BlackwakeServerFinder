package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/eiannone/keyboard"
)

// Attention!!! Hardcoding your Steam Web API key is very very very unsafe!!! 
// In practice, if you want to open source a project that contains sensitive keys
// please use environment variables to pass them! 
// Or delete the relevant content before committing and pushing!!!
const apiKey = "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"

type ServerInfo struct {
	Name    string
	Players int
	Max     int
	Addr    string
}

const (
	staticLines = 4
)

func main() {

	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer keyboard.Close()

	clearScreen()
	displayStaticHeader()
	displayDynamicContent()

	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			break
		}

		if key == keyboard.KeyEsc || key == keyboard.KeyCtrlC {
			break
		}

		if char == 'r' || char == 'R' {
			displayDynamicContent()
		}
	}
}

func clearScreen() {
	fmt.Print("\033[2J\033[H")
}

func displayStaticHeader() {
	fmt.Println("Blackwake Server Finder r1.1")
	fmt.Println("Author: GongSunFangYun")
	fmt.Println("Press [R] To Re-Fetch Data | Press [Esc] To Exit")
	fmt.Println("-----------------------------------------------------------")
}

func displayDynamicContent() {
	fmt.Printf("\033[%d;1H", staticLines+1)
	fmt.Print("\033[J")

	done := make(chan bool)
	go func() {
		showLoadingAnimation(done, staticLines)
	}()

	servers, totalPlayers, err := fetchServerData()

	done <- true
	time.Sleep(50 * time.Millisecond)

	fmt.Printf("\033[%d;1H", staticLines+1)
	fmt.Print("\033[J")

	if err != nil {
		fmt.Println("[ERROR] Failed To Fetch Data!")
		fmt.Println("-----------------------------------------------------------")
		return
	}

	displayServerTree(servers, totalPlayers)
}

func fetchServerData() ([]ServerInfo, int, error) {
	url := fmt.Sprintf("https://api.steampowered.com/IGameServersService/GetServerList/v1/?key=%s&filter=appid\\420290", apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()

	var data map[string]interface{}
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&data); err != nil {
		return nil, 0, err
	}

	response, ok := data["response"].(map[string]interface{})
	if !ok {
		return nil, 0, fmt.Errorf("invalid API response format")
	}

	serversRaw, ok := response["servers"].([]interface{})
	if !ok {
		return []ServerInfo{}, 0, nil
	}

	var servers []ServerInfo
	totalPlayers := 0

	for _, s := range serversRaw {
		server, ok := s.(map[string]interface{})
		if !ok {
			continue
		}

		name, _ := server["name"].(string)
		players, _ := server["players"].(float64)
		maxPlayers, _ := server["max_players"].(float64)
		addr, _ := server["addr"].(string)

		name = cleanName(name)

		servers = append(servers, ServerInfo{
			Name:    name,
			Players: int(players),
			Max:     int(maxPlayers),
			Addr:    addr,
		})

		totalPlayers += int(players)
	}

	sort.Slice(servers, func(i, j int) bool {
		return servers[i].Players > servers[j].Players
	})

	return servers, totalPlayers, nil
}

func displayServerTree(servers []ServerInfo, totalPlayers int) {
	if len(servers) == 0 {
		fmt.Println("Community Server List")
		fmt.Println("-----------------------------------------------------------")
		fmt.Println("No servers found")
		return
	}

	fmt.Println("Community Server List")

	for i, s := range servers {
		treePrefix := "├─"
		if i == len(servers)-1 {
			treePrefix = "└─"
		}

		fmt.Printf("%s %s\n", treePrefix, s.Name)

		childPrefix := "│  "
		if i == len(servers)-1 {
			childPrefix = "   "
		}
		fmt.Printf("%s ├─ Player(s): %d/%d\n", childPrefix, s.Players, s.Max)
		fmt.Printf("%s └─ Address: %s", childPrefix, s.Addr)

		if i < len(servers)-1 {
			fmt.Println()
		}
	}

	fmt.Println("\n-----------------------------------------------------------")
	fmt.Printf("Total Server Count: %d\n", len(servers))
	fmt.Printf("Total Player Count: %d\n", totalPlayers)
	fmt.Printf("Data Uptime: %s\n", time.Now().Format("15:04:05"))
}

func showLoadingAnimation(done chan bool, startLine int) {
	frames := []string{"|", "/", "-", "\\"}
	i := 0

	for {
		select {
		case <-done:
			return
		default:
			fmt.Printf("\033[%d;1H", startLine+1)
			fmt.Print("\033[K")

			fmt.Printf("%s Fetching Data... ", frames[i])
			i = (i + 1) % len(frames)
			time.Sleep(150 * time.Millisecond)
		}
	}
}

func cleanName(name string) string {
	if name == "" {
		return "Unknown Server"
	}

	name = strings.ReplaceAll(name, "|3.9|a", "")
	name = strings.ReplaceAll(name, "d::", "")

	for {
		idx := strings.Index(name, "|")
		if idx == -1 || idx >= len(name)-1 {
			break
		}

		if len(name) > idx+1 && (name[idx+1] >= 'a' && name[idx+1] <= 'z' || name[idx+1] >= 'A' && name[idx+1] <= 'Z') {
			name = name[:idx] + name[idx+2:]
		} else {
			name = name[:idx] + name[idx+1:]
		}
	}

	return strings.TrimSpace(name)
}
