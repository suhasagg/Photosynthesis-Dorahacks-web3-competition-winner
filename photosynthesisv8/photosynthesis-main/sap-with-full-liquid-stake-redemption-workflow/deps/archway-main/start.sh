#!/bin/sh

# Start crond in background
crond -b -d 8

# Start archwayd in foreground
archwayd --home /home/photo/.photo start
