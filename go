#!/usr/bin/env bash

# Safer script behaviour: exit on errors, on unset vars, and failed pipelines
set -euo pipefail

usage() {
	echo "Usage: go [-n] [-d] <search terms>"
	exit 1
}

incognito="no"
engine="google"

while getopts ":nd" opt; do
	case "$opt" in
		n) incognito="yes" ;;
		d) engine="duckduckgo" ;;
		*) usage ;;
	esac
done
shift $((OPTIND -1))

# Checks if you have passed at least one word
if [ "$#" -eq 0 ]; then
	echo "Usage: go <search terms>"
	exit 1
fi

# Stores all arguments as one string
query="$*"

# Uses python to URL-encode the query so spaces/special chars are safe
encoded=$(python3 - "$query" <<'PY'
import urllib.parse, sys
print(urllib.parse.quote(' '.join(sys.argv[1:])))
PY
)

if [ "$engine" = "duckduckgo" ]; then
	url="https://duckduckgo.com/?q=$encoded"
else
	url="https://www.google.com/search?q=$encoded"
fi

if [ "$incognito" = "yes" ]; then
	open -na "Google Chrome" --args --incognito "$url"
else
	open "$url"
fi
