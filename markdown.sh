#!/bin/bash

# Step 1: Install the necessary npm packages
npm install remark-cli remark-preset-lint-recommended remark-preset-lint-consistent --no-bin-links

# Step 2: Check if .remarkrc exists, if not create it
if [ ! -f .remarkrc ]; then
    cat <<EOL > .remarkrc
{
  "plugins": [
    "preset-lint-recommended",
    "preset-lint-consistent"
  ]
}
EOL
fi

# Step 3: Run remark CLI tool to fix the markdown files

# Check if remark exists in the usual directory
if [ -f ./node_modules/remark-cli/cli.js ]; then
    REMARK_PATH="./node_modules/remark-cli/cli.js"
else
    echo "Error: remark binary not found in the expected location. Exiting."
    exit 1
fi

# Fix all markdown files in current directory and its sub-directories
find . -name "*.md" -exec $REMARK_PATH --use preset-lint-recommended --use preset-lint-consistent {} -o {} \;

echo "Markdown files fixed!"

