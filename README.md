# Blackwake Server Finder

A lightweight terminal-based utility for finding and monitoring Blackwake game servers in real-time. This tool fetches server information from the Steam API and displays it in an organized tree view, making it easy to find active servers and their player counts.

## Features

- **Real-time Server List**: Fetches live server data from Steam API
- **Clean Tree View**: Displays servers in an organized hierarchical format
- **Auto-refresh**: Manual refresh with the R key
- **Player Count Sorting**: Servers are sorted by player count (highest first)
- **Server Name Cleaning**: Removes unwanted tags and formatting from server names
- **Error Handling**: Robust error handling with informative messages


## Installation

### Prerequisites
- Go 1.16 or higher

### Build from Source
```bash
# Clone the repository
git clone <repository-url>
cd blackwake-server-finder

# Build the binary
go build -o bw-server-finder main.go

# Run the application
./bw-server-finder
```

## Usage

1. Run the application:
   ```bash
   ./bw-server-finder
   ```

2. Use the following keyboard controls:
   - **R**: Refresh server data
   - **Esc** or **Ctrl+C**: Exit the application

3. The application will:
   - Clear the terminal screen
   - Display the header with instructions
   - Show a loading animation while fetching data
   - Display the server list in a tree format
   - Show statistics at the bottom

## Example Output

### Successful Fetch
```
Blackwake Server Finder r1.1
Author: GongSunFangYun
Press [R] To Re-Fetch Data | Press [Esc] To Exit
-----------------------------------------------------------
Commuinty Server List
├─ [RUSSIA] DreamTeam #1 6jaWheSySx
│   ├─ Player(s): 12/54
│   └─ Address: 185.162.95.8:28916
├─ [CN]Sunshine Blackwake Server
│   ├─ Player(s): 8/54
│   └─ Address: 222.129.39.253:27015
├─ [FR] Les potes Baroudeurs eeUXYzxaDJ
│   ├─ Player(s): 5/54
│   └─ Address: 137.74.178.6:27035
└─ [US East] Legitimate Enterprise - All Modes! [Ship FF]
    ├─ Player(s): 3/54
    └─ Address: 71.117.139.41:27015
-----------------------------------------------------------
Total Server Count: 4
Total Player Count: 28
Data Uptime: 14:30:22
```

### Error State
```
Blackwake Server Finder r1.1
Author: GongSunFangYun
Press [R] To Re-Fetch Data | Press [Esc] To Exit
-----------------------------------------------------------
[ERROR] Failed To Fetch Data!
-----------------------------------------------------------
```

### No Servers Found
```
Blackwake Server Finder r1.1
Author: GongSunFangYun
Press [R] To Re-Fetch Data | Press [Esc] To Exit
-----------------------------------------------------------
Commuinty Server List
-----------------------------------------------------------
No servers found
```

## Technical Details

### API Integration
The application uses the official Steam Web API to fetch server information:
- Endpoint: `IGameServersService/GetServerList/v1/`
- Game ID: 420290 (Blackwake)
- API Key: Required for authentication

### Server Name Cleaning
The tool automatically cleans server names by:
1. Removing specific tags like `|3.9|a` and `d::`
2. Removing vertical bars followed by single letters
3. Trimming whitespace

### Terminal Control
The application uses ANSI escape codes for:
- Screen clearing: `\033[2J\033[H`
- Cursor positioning: `\033[Y;XH`
- Line clearing: `\033[K`
- Screen clearing from cursor: `\033[J`

## Dependencies

- `github.com/eiannone/keyboard`: For keyboard input handling
- Standard Go libraries: `net/http`, `encoding/json`, `fmt`, `sort`, `strings`, `time`

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Test thoroughly
5. Submit a pull request

## License

This project is open source and available under the MIT License.

## Acknowledgments

- Thanks to Valve for the Steam Web API
- Blackwake game developers and community
- Contributors and testers

## Support

For issues, feature requests, or questions:
1. Check the existing issues
2. Create a new issue with detailed information
3. Include your operating system and Go version

---

*Note: This tool is for informational purposes only. Always respect server rules and community guidelines when joining game servers.*