#!/bin/sh

set -e

# count all user-agents in the text file, remove all that have value "-", and print the top 10

less data.txt | awk -F '"' '{print $6}' | grep -wv "^-$" | sort | uniq -c | sort | tail -n 10
