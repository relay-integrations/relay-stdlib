#!/bin/bash

# Copyright (C) 2021 Noah Van Tiggel
# Licensed under Apache 2.0

JQ="${JQ:-jq}"
NI="${NI:-ni}"

# Change IFS so we can have spaces in values
OLDIFS=$IFS
IFS=$'\n'

METHOD=$(ni get -p {.method})
BODY=$(ni get | jq -rc 'try .body // empty')
URL=$(ni get -p {.url})
HEADERS=($(ni get | jq -r '.headers // empty | to_entries[] | "\(.key): \(.value)" | @sh'))
QUERY=($(ni get | jq -r '.query // empty | to_entries[] | "\(.key)=\(.value)"'))

ARGS=""
QUERYPARAMS=""

for header in "${HEADERS[@]}"
do
	ARGS+="-H ${header} "
done

[ -n "${BODY}" ] && ARGS+="-d '${BODY}'"

INDEX=0
for queryparam in "${QUERY[@]}"
do
    [ $INDEX -eq 0 ] && QUERYPARAMS+="?"
    QUERYPARAMS+="$queryparam"
    [ "$queryparam" != "${QUERY[-1]}" ] && QUERYPARAMS+="\&"
    INDEX+=1
done

FULL="curl -s -L -o response.txt -w "%{http_code}" -X ${METHOD} "${ARGS}" "${URL}${QUERYPARAMS}""
CODE="$(eval $FULL)"

ni output set -k code -v "$CODE"

IFS=$OLDIFS

EXPECTS=($(ni get | jq -r '.expects // empty | @sh'))
FAILON=($(ni get | jq -r '.failon // empty | @sh'))

# Check if the http code matches expects or failon
# First check failon

for failcode in "${FAILON[@]}"
do
    if [ "$CODE" == "$failcode" ] ; then
        echo "Request failed! Call returned ${CODE}"
        exit 1
    fi
done

# Then check expects
found=false
for expectedcode in "${EXPECTS[@]}"
do
    [ "$CODE" == "$expectedcode" ] && found=true 
done

if [ "$found" == false ] && [ "${#EXPECTS[@]}" -gt 0 ] ; then
    echo "Request failed! Call returned ${CODE} but expected ${EXPECTS[@]}"
    exit 1
fi

# No errors, continue

RESPONSE=$(cat response.txt)

ni output set -k response -v "$RESPONSE"