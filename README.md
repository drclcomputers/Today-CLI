# 📅 Today CLI

A command-line tool to fetch and display historical events, births, and deaths for any date. Powered by Zenquotes' On This Day API, with a beautiful TUI and clipboard support.

---

## 🚀 Features

- 🔎 Search for events, births, and deaths by date
- 🗓️ Interactive TUI menu for event type selection
- 📋 Copy results to clipboard
- 🌐 Fetches data from 'On This Day' API
- 🖥️ Cross-platform support (Windows, macOS, Linux)
- 🧩 Modular code structure

---

## 🛠️ Installation

1. **Clone the repository:**
   ```sh
   git clone https://github.com/drclcomputers/today-cli
   cd today-cli
   ```
2. **Build the project:**
   ```sh
   go build
   ```
3. **Run the CLI:**
   ```sh
   ./today
   ```
- **Follow the prompts to select event type.**

---

## 📝 Usage

- **Show help:**
  ```sh
  ./today -h
  ```
- **Input a custom date as directed or opt for automatic event type selection.**
- **Copy results to clipboard with a single keypress.**

---

## 📂 Project Structure

```
.
├── cmd/                # CLI commands (births, deaths, events, date, root)
├── internal/
│   ├── logger/         # Logging utilities
│   ├── services/       # API, menu, spinner, event type logic
│   └── utils/          # Utility functions
├── main.go             # Entry point
└── go.mod, go.sum      # Go modules
```

---

## 🤝 Contributing

Pull requests are welcome! For major changes, please open an issue first to discuss what you would like to change.

---

## 📜 License

[MIT](LICENSE)

---

## 🙏 Acknowledgements

- [On This Day API](https://today.zenquotes.io/) for the data
- [Bubble Tea](https://github.com/charmbracelet/bubbletea) for TUI
- [golang.design/x/clipboard](https://github.com/golang-design/clipboard)
- [atotto/clipboard](https://github.com/atotto/clipboard)
