#!/bin/bash

err() {
    echo "$1" >&2
    exit 1
}

echo '$ cd registry'
cd registry || err "Cannot find registry"
echo

echo '$ git config user.name "just-install-bot"'
git config user.name "just-install-bot"
echo

echo '$ git config user.email "just-install-bot@outlook.com"'
git config user.email "just-install-bot@outlook.com"
echo

echo '$ go run ../main.go -c message.txt just-install.json'
go run ../main.go -c message.txt just-install.json || err "Could not update registry"
echo

echo '$ git add just-install.json'
git add just-install.json
echo

echo '$ cat message.txt'
cat message.txt
echo

echo '$ git commit -F message.txt'
git commit -F message.txt
echo

echo '$ git push https://github.com/just-install/registry.git master'
[[ -z $GITHUB_TOKEN ]] && die "No github token"
git push "https://$GITHUB_TOKEN@github.com/just-install/registry.git" master >/dev/null 2>/dev/null || err "Could not push changes"
echo

echo '$ cd ..'
cd ..
echo

echo '$ rm -rf registry'
rm -rf registry
echo