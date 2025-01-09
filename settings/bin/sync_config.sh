#!/bin/bash

/opt/homebrew/bin/git --git-dir="$HOME/.git_settings/" --work-tree="$HOME" add -u
/opt/homebrew/bin/git --git-dir="$HOME/.git_settings/" --work-tree="$HOME" add "$HOME/settings"
CHANGES=$(/opt/homebrew/bin/git --git-dir="$HOME/.git_settings/" --work-tree="$HOME" status --porcelain)

if [ -n "$CHANGES" ]; then
  /opt/homebrew/bin/git --git-dir="$HOME/.git_settings/" --work-tree="$HOME" commit -m "karabiner.json updated" && \
  /opt/homebrew/bin/git --git-dir="$HOME/.git_settings/" --work-tree="$HOME" push origin master
fi
