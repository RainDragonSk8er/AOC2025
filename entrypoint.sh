#!/bin/sh

# Ensure git is configured (if mounted)
# We assume the user mounts their .gitconfig or we set basic config here
git config --global user.email "aoc-bot@example.com"
git config --global user.name "AOC Bot"

# Add safe directory
git config --global --add safe.directory /home/jho/Code/AOC2025
# Also add the current directory just in case we mount it differently
git config --global --add safe.directory /app

echo "Starting AOC Tracker..."

# Run immediately on start
./tracker

# Loop to run every hour
while true; do
    echo "Sleeping for 1 hour..."
    sleep 3600
    echo "Running tracker..."
    ./tracker
done
