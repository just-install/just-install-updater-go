#!/bin/bash

err() {
    echo "$1" >&2
    exit 1
}

echo "Cloning registry"
rm -rf registry
git clone https://$GITHUB_TOKEN:x-oauth-basic@github.com/just-install/registry >/dev/null 2>/dev/null || err "Could not clone repo"
cd registry

echo "Configuring git user and email"
git config user.name "just-install-bot"
git config user.name
git config user.email "just-install-bot@outlook.com"
git config user.email

echo "Updating registry"
go run ../main.go just-install.json || err "Could not update registry"

echo "Committing changes"
git add -A
git commit -S -m "jiup-go automatic commit"

echo "Pushing changes"
git push >/dev/null 2>/dev/null || err "Could not push changes"

echo "Cleaning up"
cd ..
rm -rf registry