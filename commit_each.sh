#!/bin/bash

git status --porcelain | while read line; do

file="${line:3}"

if [ -f "$file" ]; then
    echo "Committing $file..."
    git add "$file"
    git commit -m "Add $file"
fi
done

git push origin main