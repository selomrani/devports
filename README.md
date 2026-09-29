# devports 🚀

<p align="center">
  <b>A beautiful, cross-platform TUI (Terminal User Interface) for monitoring and managing local development ports.</b>
</p>

![Demo](demo.gif)

Stop typing `lsof -i :3000` and `kill -9`. **devports** gives you a sleek, real-time dashboard of every port your machine is listening on, what process is running it, and lets you kill zombie servers with a single keystroke.

---

## ✨ Features

- **Instant Port Monitoring:** View all active TCP listening ports instantly.
- **Process Details:** See the exact PID, Process Name (Node, Python, Postgres, etc.), and Memory Usage.
- **1-Click Kill:** Safely kill processes blocking your ports without memorizing PIDs (just press `k`).
- **Cross-Platform:** Works natively on macOS, Linux, and Windows.
- **Zero Dependencies:** A single, lightweight compiled binary.

---

## 📦 Installation

### macOS & Linux
The fastest way to install is via our automated bash script, which downloads the correct binary and adds it to your PATH:
```bash
curl -sL https://selomrani.github.io/devports/install.sh | bash
```

### Windows
Open **PowerShell** and run our automated installation script:
```powershell
irm https://selomrani.github.io/devports/install.ps1 | iex
```

### Using Go (All Platforms)
If you already have Go installed, you can compile and install it directly:
```bash
go install github.com/selomrani/devports@latest
```

### Manual Download
Prefer doing it yourself? Head over to the [Releases Page](https://github.com/selomrani/devports/releases) and download the `.tar.gz` or `.zip` file for your system.

---

## 💻 Usage

Simply open your terminal and type:
```bash
devports
```

### Keybindings
| Key | Action |
| :--- | :--- |
| `k` | **Kill** the currently selected process |
| `r` | **Refresh** the list of active ports |
| `↑` / `↓` | Navigate the list |
| `q` or `Ctrl+C` | Quit the application |

*Note: If a process is owned by `root` (like system services), you may need to run `sudo devports` in order to have permission to kill it.*

---

## 🛠️ Built With
- [Go](https://go.dev/)
- [Bubbletea](https://github.com/charmbracelet/bubbletea) (The powerful TUI framework)
- [Lipgloss](https://github.com/charmbracelet/lipgloss) (Style definitions)
- [gopsutil](https://github.com/shirou/gopsutil) (Cross-platform system and process monitoring)

## 🤝 Contributing
Pull requests are welcome! If you have a feature request or found a bug, please open an issue.

1. Fork the project
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 📄 License
Distributed under the MIT License. See `LICENSE` for more information.
