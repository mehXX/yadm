#!/bin/bash

LOCK_FILE="/tmp/proton_sync.lock"

if [ -e "$LOCK_FILE" ]; then
    exit 0
fi

touch "$LOCK_FILE"

# Function to remove the lock file
cleanup() {
    rm -f "$LOCK_FILE"
    echo "Lock file removed"
    echo "----------------------------------------------------------------------------------------------------"
}

# Set the trap to remove the lock file on script exit (EXIT signal)
trap cleanup EXIT

# Function to get the current time in nanoseconds
get_current_time() {
    echo "$(/opt/homebrew/bin/gdate +%s%N)"
}

# Function to log the duration of a task
log_duration_and_complete() {
    local start_time=$1
    local task_name=$2
    local end_time=$(get_current_time)
    local elapsed_time=$(( (end_time - start_time) / 1000000 )) # Convert to milliseconds
    echo "DONE: $task_name in ${elapsed_time} ms"
}

echo "BEGINNING"
echo "Current time: $(date +"%Y-%m-%d %H:%M:%S")"

# Sync Settings
start_time=$(get_current_time)
echo "START: syncing Settings"
/opt/homebrew/bin/rclone sync ~/Library/Mobile\ Documents/com\~apple\~CloudDocs/Settings proton:Settings
log_duration_and_complete $start_time "syncing Settings"

# Sync Documents
start_time=$(get_current_time)
echo "START: syncing Documents"
/opt/homebrew/bin/rclone sync ~/Library/Mobile\ Documents/com\~apple\~CloudDocs/Documents proton:Documents
log_duration_and_complete $start_time "syncing Documents"

# Sync Health
start_time=$(get_current_time)
echo "START: syncing Health"
/opt/homebrew/bin/rclone sync ~/Library/Mobile\ Documents/com\~apple\~CloudDocs/Health proton:Health
log_duration_and_complete $start_time "syncing Health"

# Sync Passwords
start_time=$(get_current_time)
echo "START: syncing Passwords"
/opt/homebrew/bin/rclone sync ~/Library/Mobile\ Documents/com\~apple\~CloudDocs/Passwords proton:Passwords
log_duration_and_complete $start_time "syncing Passwords"

