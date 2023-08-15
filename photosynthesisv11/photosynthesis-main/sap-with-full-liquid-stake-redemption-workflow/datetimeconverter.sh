#!/bin/bash

timestamp_nanoseconds=1690287394738055641
timestamp_seconds=$((timestamp_nanoseconds / 1000000000))

date -d @$timestamp_seconds

