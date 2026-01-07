# Project Hopeline
Disaster Relief and Family Reconnection System - This project focuses on reuniting families and restoring dignity and hope, not just technical complexity.


> Your code may be the difference between someone feeling lost and someone feeling found.


## Prerequisites
Make sure you have:
- WSL installed (WSL 2 recommended)
- A Linux distro installed (e.g. Ubuntu)
- Go installed inside WSL
- Git installed inside WSL

Source: https://learn.microsoft.com/en-us/windows/wsl/install

Build essentials
- sudo apt-get update
- sudo apt-get upgrade
- sudo apt-get install build-essential -y
- sudo snap install go --classic
- git config --global user.email "your_email"
- git config --global user.name "your_username"
- ssh-keygen -t ed25519 -C "your_email"
- eval "$(ssh-agent -s)"
- ssh-add /home/username/.ssh/id_ed25519
- git clone https://github.com/illusioniststg/hopeline.git
- cd hopeline
- make run

source: 
- [Generate SSH Key](https://docs.github.com/en/authentication/connecting-to-github-with-ssh/generating-a-new-ssh-key-and-adding-it-to-the-ssh-agent)

- [Add SSH key to Github Account](https://docs.github.com/en/authentication/connecting-to-github-with-ssh/adding-a-new-ssh-key-to-your-github-account)

## Make Commands

| Command      | Description             |
| ------------ | ----------------------- |
| `make run`   | Run the server          |
| `make build` | Build the binary        |
| `make clean` | Remove the built binary |

