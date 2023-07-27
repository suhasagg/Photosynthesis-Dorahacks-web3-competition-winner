#!/bin/bash

find . -type f -exec sh -c '
  for file do
    echo "Processing: $file"
    sed -i 's/photosynthesisv5/photosynthesisv5/g' "$file"
  done
' sh {} +

