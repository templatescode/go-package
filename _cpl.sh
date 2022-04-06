#!/bin/bash
clear

nameAPP="app"

# Test
echo "** Testing..."
go test -v -cover ./..
echo "** Finish!" 

# Compile
echo
echo "** Compiling with race..."
go build -race -o $nameAPP
echo "** Finish!" 

# Check/Lint
echo
echo "** Checking..."
staticcheck ./...
echo "** Finish!"

# Environment
export DB_USER="devuser"
export DB_PWD="devpwd"
export DB_HOST="localhost"
export DB_PORT="5432"
export DB_NAME="db-test"

# Run
echo
echo "** Starting..."
echo
./$nameAPP
