# url-sh
A simple and _blazingly fast_ URL shortner and resolver tool written in go that uses MongoDB. (For teaching purposes only)

## Getting Started
1. Set the appropriate environment variables inside `.env.example` file
2. Rename `.env.example` file to `.env` 

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

##### How to Build
1. Make sure your have `golang` and `make` installed on your machine
2. Navigate to this project in terminal.
3. Type the command `make cli`
