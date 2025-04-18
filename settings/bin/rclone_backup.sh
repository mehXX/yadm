#!/bin/bash

LOCK_FILE="/tmp/rclone_backup.lock"
[ -e "$LOCK_FILE" ] && exit 0
touch "$LOCK_FILE"

ERROR_LOG=$(mktemp)
exec 3>&2
exec 2> "$ERROR_LOG"

export TELEGRAM_BOT_API_TOKEN=$(security find-generic-password -a rclone_backup -s TELEGRAM_BOT_API -w)
export RCLONE_CONFIG_PASS=$(security find-generic-password -a rclone_backup -s RCLONE_CONFIG_PASS -w)

on_error() {
    local code=$?
    local line=${BASH_LINENO[0]}
    local cmd=$BASH_COMMAND
    local msg=$(cat "$ERROR_LOG")
    curl -s -X POST "https://api.telegram.org/$TELEGRAM_BOT_API_TOKEN/sendMessage" \
         -d chat_id="253872226" \
         --data-urlencode text="rclone‑backup failed: exit $code, line $line, cmd: $cmd\n$msg"
    exit $code
}

set -Ee                       # -e: exit on error, -E: inherit ERR in functions :contentReference[oaicite:3]{index=3}
trap on_error ERR

cleanup() {
    rm -f "$LOCK_FILE"
    exec 2>&3
    rm -f "$ERROR_LOG"
}
trap cleanup EXIT

date

[[ "$USER" == "ymka" ]] || { echo "USER is not ymka" >&2; exit 1; }

get_current_time() { /opt/homebrew/bin/gdate +%s%N; }

log_duration_and_complete() {
    local start=$1 task=$2 end=$(get_current_time)
    printf 'DONE: %s in %d ms\n' "$task" $(( (end - start) / 1000000 ))
}

echo "BEGINNING /Users/ymka/settings/bin/rclone_backup.sh"
echo "Current time: $(date +"%Y-%m-%d %H:%M:%S")"

start=$(get_current_time)
echo "START: backing up icloud to backblaze-b2"
/opt/homebrew/bin/rclone sync "/Users/ymka/Library/Mobile Documents/com~apple~CloudDocs" encrypted-backblaze-b2: --exclude '**/.DS_Store' --exclude '.DS_Store'
log_duration_and_complete "$start" "backing up icloud to backblaze-b2"

start=$(get_current_time)
echo "START: backing up icloud to google-drive"
/opt/homebrew/bin/rclone sync "/Users/ymka/Library/Mobile Documents/com~apple~CloudDocs" google-drive:iCloud --exclude '**/.DS_Store' --exclude '.DS_Store'
log_duration_and_complete "$start" "backing up icloud to google-drive"

