#!/bin/bash

for file in *; do
  if [ -f "$file" ]; then
    sed -i 's/photosynthesisv5/photosynthesisv5/g' "$file"
  fi
done

