#!/usr/bin/env sh

ls types/*.go templates/*.go main.go | entr -rcs 'go run main.go'
