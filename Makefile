setup:
	chmod +x scripts/cli/install-linux.sh
	chmod +x scripts/cli/uninstall-linux.sh

cli:
	go build -o build/cli/url-sh cmd/cli/main.go

server:
	go build -o build/server/url-sh cmd/server/main.go