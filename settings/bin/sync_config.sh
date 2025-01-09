#!/bin/bash

get_current_time() {
    echo "$(/opt/homebrew/bin/gdate +%s%N)"
}

log_duration_and_complete() {
    local start_time=$1
    local task_name=$2
    local end_time=$(get_current_time)
    local elapsed_time=$(( (end_time - start_time) / 1000000 )) # Convert to milliseconds
    echo "DONE: $task_name in ${elapsed_time} ms"
}

start_time=$(get_current_time)
echo "START: yadm"
echo "Current time: $(date +"%Y-%m-%d %H:%M:%S")"

#  /opt/homebrew/bin/git --git-dir="$HOME/.git_settings/" --work-tree="$HOME" commit -m "karabiner.json updated" && \
#  /opt/homebrew/bin/git --git-dir="$HOME/.git_settings/" --work-tree="$HOME" push origin master
/opt/homebrew/bin/yadm add -u
/opt/homebrew/bin/yadm add "$HOME/settings"
/opt/homebrew/bin/yadm commit -m "qwe"
/opt/homebrew/bin/yadm push origin master

log_duration_and_complete $start_time "yadm"
echo "----------------------------------------------------------------------------------------------------"
