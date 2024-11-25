#!/bin/bash

# create 100 directories and 100 files in each of them
for i in {000..099}; do
  mkdir -p "dir_$i"
  for j in {00..99}; do
    touch "dir_$i/file_$j.pdf"
  done
done
