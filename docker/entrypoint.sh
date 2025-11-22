#!/bin/sh

# Ensure git is configured (if mounted)
# We assume the user mounts their .gitconfig or we set basic config here
git config --global user.email "aoc-bot@example.com"
git config --global user.name "AOC Bot"
git config --global credential.helper 'store --file=/root/.git-credentials'

if [ ! -f /root/.git-credentials ]; then
    echo "WARNING: /root/.git-credentials not found!"
else
    echo "Found /root/.git-credentials"
fi
# Add safe directory
git config --global --add safe.directory /home/jho/Code/AOC2025
# Also add the current directory just in case we mount it differently
git config --global --add safe.directory /app

echo "Starting AOC Tracker..."

# Run immediately on start
run_update.sh

# Loop to run every hour
# We use crond logic or just a simple sleep loop.
# Since we want it to run at midnight, a sleep loop might drift.
# But for simplicity in this container, a sleep loop checking the time or just running hourly is okay.
# Running hourly ensures we hit the midnight window (00:xx) eventually.
while true; do
    echo "Sleeping for 1 hour..."
    sleep 3600
    echo "Running update..."
    run_update.sh
done
