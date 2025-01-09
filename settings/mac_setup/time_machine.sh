#!/bin/bash

set -euo pipefail
set -vx

sudo tmutil addexclusion -p ~/.cache
sudo tmutil addexclusion -p ~/Downloads
sudo tmutil addexclusion -p "$HOME/Library/Application Support"
sudo tmutil addexclusion -p ~/Library/Caches
sudo tmutil addexclusion -p ~/Library/Containers
sudo tmutil addexclusion -p ~/Music
sudo tmutil addexclusion -p ~/go
