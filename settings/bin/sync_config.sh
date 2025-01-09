#!/bin/bash

start_time=$(get_current_time)
echo "START: yadm"

#  /opt/homebrew/bin/git --git-dir="$HOME/.git_settings/" --work-tree="$HOME" commit -m "karabiner.json updated" && \
#  /opt/homebrew/bin/git --git-dir="$HOME/.git_settings/" --work-tree="$HOME" push origin master
/opt/homebrew/bin/yadm add -u
/opt/homebrew/bin/yadm add "$HOME/settings"
/opt/homebrew/bin/yadm commit -m "qwe"
/opt/homebrew/bin/yadm push origin master

log_duration_and_complete $start_time "END: yadm done"