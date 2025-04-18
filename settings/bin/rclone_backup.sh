#!/bin/bash

set -o errexit -o pipefail -o nounset

LOCK_FILE="/tmp/rclone_backup.lock"
NOTIFY_FLAG="/tmp/rclone_backup_error_notified"
LAST_BACKUP_FILE="/tmp/rclone_backup.last"

# load Telegram token early so we can alert if backups are stale
export TELEGRAM_BOT_API_TOKEN=$(security find-generic-password -a rclone_backup -s TELEGRAM_BOT_API -w)

# --- check for stale backup (>24h) and notify once per script run ---
if [ -e "$LAST_BACKUP_FILE" ]; then
    last_ts=$(cat "$LAST_BACKUP_FILE")
    current_ts=$(date +%s)
    # if more than 86400 seconds (1 day) have passed since last backup
    if [ $((current_ts - last_ts)) -gt 86400 ]; then
        last_human=$(date -r "$last_ts" +"%Y-%m-%d %H:%M:%S")
        telegram_text=$(printf "❗️rclone-backup warning: no backup in the last 24h. Last backup was at %s" "$last_human")
        curl -s -X POST "https://api.telegram.org/$TELEGRAM_BOT_API_TOKEN/sendMessage" \
             -d chat_id="253872226" \
             --data-urlencode text="$telegram_text"
    fi
fi

# prevent overlapping runs
[ -e "$LOCK_FILE" ] && exit 0
touch "$LOCK_FILE"

# capture stderr to a temp file
ERROR_LOG=$(mktemp)
exec 3>&2
exec 2> "$ERROR_LOG"

cleanup() {
    rm -f "$LOCK_FILE"
    exec 2>&3
    rm -f "$ERROR_LOG"
}
trap cleanup EXIT

export RCLONE_CONFIG_PASS=$(security find-generic-password -a rclone_backup -s RCLONE_CONFIG_PASS -w)

on_error() {
    local code=$?
    local line=${BASH_LINENO[0]}
    local cmd=$BASH_COMMAND
    local msg
    msg=$(<"$ERROR_LOG")

    # always echo full error to stdout
    printf '[%s] rclone-backup failed: exit %d, line %d, cmd: %s\n' \
        "$(date +"%Y-%m-%d %H:%M:%S")" \
        "$code" "$line" "$cmd"
    printf '%s\n' "$msg"

    # Telegram alert only once
    if [ ! -e "$NOTIFY_FLAG" ]; then
        local telegram_text
        telegram_text=$(printf "rclone-backup failed: exit %d, line %d, cmd: %s\n%s" \
            "$code" "$line" "$cmd" "$msg")

        curl -s -X POST "https://api.telegram.org/$TELEGRAM_BOT_API_TOKEN/sendMessage" \
             -d chat_id="253872226" \
             --data-urlencode text="$telegram_text"

        touch "$NOTIFY_FLAG"
    fi

    exit $code
}
trap on_error ERR

# Start logging to stdout
echo "----------------------------------------------------------------------------------------------------"
date
echo "BEGINNING /Users/ymka/settings/bin/rclone_backup.sh"
echo "Current time: $(date +"%Y-%m-%d %H:%M:%S")"

get_current_time() { /opt/homebrew/bin/gdate +%s%N; }
log_duration_and_complete() {
    local start=$1 task=$2 end=$(get_current_time)
    printf 'DONE: %s in %d ms\n' "$task" $(( (end - start) / 1000000 ))
}

[[ "$USER" == "ymka" ]] || { echo "USER is not ymka"; exit 1; }

start=$(get_current_time)
echo "START: backing up iCloud to Backblaze B2"
/opt/homebrew/bin/rclone sync \
    "/Users/ymka/Library/Mobile Documents/com~apple~CloudDocs" \
    encrypted-backblaze-b2: \
    --exclude '**/.DS_Store' --exclude '.DS_Store'
log_duration_and_complete "$start" "backing up iCloud to Backblaze B2"

start=$(get_current_time)
echo "START: backing up iCloud to Google Drive"
/opt/homebrew/bin/rclone sync \
    "/Users/ymka/Library/Mobile Documents/com~apple~CloudDocs" \
    google-drive:iCloud \
    --exclude '**/.DS_Store' --exclude '.DS_Store'
log_duration_and_complete "$start" "backing up iCloud to Google Drive"

# if we had previously failed, now report recovery
if [ -e "$NOTIFY_FLAG" ]; then
    printf '[%s] rclone-backup succeeded after previous failures\n' \
        "$(date +"%Y-%m-%d %H:%M:%S")"
    curl -s -X POST "https://api.telegram.org/$TELEGRAM_BOT_API_TOKEN/sendMessage" \
         -d chat_id="253872226" \
         --data-urlencode text="rclone-backup succeeded at $(date +"%Y-%m-%d %H:%M:%S") after previous failures."
    rm -f "$NOTIFY_FLAG"
fi

# --- record the time of this successful backup ---
date +%s > "$LAST_BACKUP_FILE"
