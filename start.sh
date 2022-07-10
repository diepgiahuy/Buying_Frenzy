#!/bin/bash

# turn on bash's job control
set -m

# Start the primary process
./app/buying-frenzy serve &
sleep 10
./app/buying-frenzy load

fg %1