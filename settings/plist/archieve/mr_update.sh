#!/bin/bash

REPOS_DIR="/Users/n.naumov/gitlab"

exec &>> /tmp/mr.log

date

/opt/homebrew/bin/mr --directory /Users/n.naumov/gitlab --jobs 4 update

# Loop through the repositories and index them
for repo in "$REPOS_DIR"/*; do
  if [ -d "$repo/.git" ]; then
    echo "Indexing $repo..."
    /Users/n.naumov/go/bin/zoekt-git-index -index "$REPOS_DIR/index" "$repo"
  else
    echo "$repo is not a git repository."
  fi
done
