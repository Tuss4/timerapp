#!/bin/bash
cat requirements.txt | while read line
do
    exec go get $line
done
