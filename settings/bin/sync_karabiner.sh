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
echo "START: karabiner"
echo "Current time: $(date +"%Y-%m-%d %H:%M:%S")"

cd karabiner || return
yarn run build  && \
rm -f "$HOME/.config/karabiner/karabiner.json" && \
cp "$HOME/settings/karabiner/karabiner.json" "$HOME/.config/karabiner"

log_duration_and_complete $start_time "karabiner"
echo "----------------------------------------------------------------------------------------------------"
