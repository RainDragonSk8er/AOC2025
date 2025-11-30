#!/bin/sh

# 0. Update Repo
git pull

# 1. Run Scaffolder (Auto-detects date)
# It will only create files if today is Nov 30 or Dec 1-12
scaffold

# 2. Run Tracker (Updates README)
tracker

# 3. Git Operations
# Check if there are changes
if [ -n "$(git status --porcelain)" ]; then
    echo "Changes detected. Committing..."
    git add .
    git commit -m "Auto-update: Leaderboard & Scaffolding $(date)"
    git push
else
    echo "No changes to commit."
fi
