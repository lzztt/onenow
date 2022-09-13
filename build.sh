#!/bin/bash

# Build the homepage for GitHub Pages
for i in `ls note | sort -n`; do
    echo "- [`head -n 1 note/$i | sed 's/^# //'`](note/`echo $i | sed 's/.md$//'`)";
done > README.md

# Build the site in Jekyll container
if [ ! -z "$JEKYLL_DOCKER_NAME" ]; then
    for i in note/*.md; do
        sed -i '1 i\---\nlayout: default\n---' $i;
    done

    cp README.md index.md
    sed -i '1 i\---\nlayout: default\n---' index.md

    jekyll build
fi
