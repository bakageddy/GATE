#!/usr/bin/env sh

ls templates/*.templ | entr -rcs 'templ generate'
