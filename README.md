# League Auto-Accept

Tired of waiting for the queue to pop so you can go to the bathroom? The wait is finally over! This small program automatically accepts any League of Legends queue pop until you get into your game, so you never miss a match again.

## Features

- ✅ **Automatically** accepts any queue pop
- 🚀 **Fast and lightweight** — runs efficiently in the background
- 🔧 **No extra setup required** — just run it!

## How It Works

I chose Go for this project because I like Go in many aspects, such as its speed, simplicity, and the ability to natively compile an executable. This makes the program efficient and easy to run without additional dependencies.

## How to Install and Run the Project

1. Clone the Repository

   ```
   git clone https://github.com/hirschnico/league-auto-accept
   cd league-auto-accept
   ```

2. Install Go

   Make sure you have the latest version of Go installed. You can download it from [go.dev](https://go.dev/dl/).

3. Run the Program

   To start the program without building an executable, run:

   ```
   go run cmd/main.go
   ```

4. Build a Native Executable

   If you prefer to compile it into an executable, use:

   ```
   go build cmd/main.go
   ```

   This will generate a binary that you can run directly.

## Download the Latest Release

If you just want to use the program without building it yourself, check the [latest release](https://github.com/hirschnico/league-auto-accept/releases) and download it from there.

## Challenges & Future Improvements

Some challenges still remain, such as handling cases where League of Legends is not installed under `C:\Riot Games\League of Legends`. Additionally, I'm not sure if it works on macOS yet, so future improvements could include better path detection and cross-platform support.

💡 **Want to help?** If you find issues or have ideas for improvements, feel free to open an [issue](https://github.com/hirschnico/league-auto-accept/issues) or submit a pull request!

## Preview

![](./docs/preview.gif)

## License

This project is released under a **Creative Commons Attribution-NonCommercial 4.0 International (CC BY-NC 4.0) License**.

### What This Means:

✅ You are free to use, modify, and share this project.

❌ You are **not** allowed to use it for commercial purposes.

For full license details, see: [CC BY-NC 4.0](https://creativecommons.org/licenses/by-nc/4.0/)
