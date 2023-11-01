#!/bin/bash

# Number of runs
total_runs=20

for ((i=1; i<=$total_runs; i++))
do
    # Client app starts
    ./test_client.sh

    sleep 1
done
