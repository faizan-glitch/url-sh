#!/usr/bin/bash

go build -o build/cli/url-sh cmd/cli/main.go

mv build/cli/url-sh /usr/local/bin/url-sh

if [ ! -d "/usr/local/etc/url-sh" ]; then
    mkdir /usr/local/etc/url-sh
    cp config.yaml /usr/local/etc/url-sh
    echo "url-sh has been installed successfully."
else
    echo "url-sh is already installed."
    exit 1
fi
