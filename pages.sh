#!/usr/bin/env bash

NAMES=$(find . -type f | grep '.pdf')
TOTAL=0
for i in $NAMES 
do
	PAGES=$(pdfinfo $i | grep 'Pages' | sed 's/Pages\://')
	TOTAL=$(expr $PAGES + $TOTAL)
done
echo $TOTAL
