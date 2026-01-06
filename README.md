# Project Hopeline
Disaster Relief and Family Reconnection System

## Prerequisites
Make sure you have:
- WSL installed (WSL 2 recommended)
- A Linux distro installed (e.g. Ubuntu)
- Go installed inside WSL
- Make installed inside WSL

Source: https://learn.microsoft.com/en-us/windows/wsl/install

## Github Setup

- git config --global user.email "your_email"

- git config --global user.name "your_username"

- ssh-keygen -t ed25519 -C "your_email"

- eval "$(ssh-agent -s)"

- ssh-add /home/username/.ssh/id_ed25519

source: 
- [Generate SSH Key](https://docs.github.com/en/authentication/connecting-to-github-with-ssh/generating-a-new-ssh-key-and-adding-it-to-the-ssh-agent)

- [Add SSH key to Github Account](https://docs.github.com/en/authentication/connecting-to-github-with-ssh/adding-a-new-ssh-key-to-your-github-account)

## Make Commands

| Command      | Description             |
| ------------ | ----------------------- |
| `make run`   | Run the server          |
| `make build` | Build the binary        |
| `make clean` | Remove the built binary |
