#!/bin/bash

LOCK_FILE="/tmp/restic_backup.lock"

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

date

if [[ "$USER" != "ymka" ]]; then
  echo "USER is not ymka"
  rm "$LOCK_FILE"
  exit
fi

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

echo "BEGINNING"
echo "Current time: $(date +"%Y-%m-%d %H:%M:%S")"

start_time=$(get_current_time)
echo "START: backing up icloud"

export B2_ACCOUNT_ID=$(security find-generic-password -a restic_backup -s b2-account-id -w)
export B2_ACCOUNT_KEY=$(security find-generic-password -a restic_backup -s b2-account-key -w)
export RESTIC_REPOSITORY=$(security find-generic-password -a restic_backup -s restic-repository -w)
export RESTIC_PASSWORD=$(security find-generic-password -a restic_backup -s restic-password -w)

/opt/homebrew/bin/restic backup "/Users/ymka/Library/Mobile Documents/com~apple~CloudDocs"
log_duration_and_complete $start_time "backing up icloud"

