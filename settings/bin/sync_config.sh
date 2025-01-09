#!/bin/bash

#  /opt/homebrew/bin/git --git-dir="$HOME/.git_settings/" --work-tree="$HOME" commit -m "karabiner.json updated" && \
#  /opt/homebrew/bin/git --git-dir="$HOME/.git_settings/" --work-tree="$HOME" push origin master
/opt/homebrew/bin/yadm add -u
/opt/homebrew/bin/yadm add "$HOME/settings"
/opt/homebrew/bin/yadm commit -m "qwe"
/opt/homebrew/bin/yadm push origin master
