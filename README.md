# url-sh
A simple and _blazingly fast_ URL shortner and resolver tool written in go that uses MongoDB. (For teaching purposes only)

## Getting Started
1. Set the appropriate configuration variables inside `config.yaml` file
2. Run the command `make setup`

### Major Components
- Server
- CLI

#### Server
It is a simple HTTP server written in GoFiber framework.

##### How to Build
1. Make sure your have `golang` and `make` installed on your machine
2. Navigate to this project in terminal.
3. Type the command `make server`

#### CLI
It is a simple CLI app in case you wish to use this project in a terminal and not through a web app.

##### How to Install
1. Make sure your have `golang` and `make` installed on your machine.
2. Navigate to this project in terminal.
3. Change `config.yaml` file to suit your environment.
4. Type the command `make setup`.
5. Type the command `./scripts/cli/install-linux.sh`.
6. Run `url-sh` in terminal to test if app is running properly.