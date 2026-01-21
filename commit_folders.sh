#!/bin/bash

# Iterate over untracked files and directories (lines starting with ??)
git status --porcelain | grep "^??" | while read line; do
    # Extract the path (remove the first 3 characters "?? ")
    path="${line:3}"
    
    # Remove trailing slash if it's a directory to make the commit message cleaner
    clean_name="${path%/}"
    
    echo "Committing $clean_name..."
    git add "$path"
    git commit -m "Add $clean_name"
done

echo "Pushing to origin main..."
git push origin main
