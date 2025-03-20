#!/bin/bash

curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.5/install.sh | bash
source ~/.bashrc
nvm install node
npm install -g nodemon
go install github.com/swaggo/swag/cmd/swag@latest
cd /workspaces/Mono-repo-template/Backend-Go-template/app && go get .