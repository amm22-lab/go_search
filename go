#!/usr/bin/env bash

# Safer script behaviour: exit on errors, on unset vars, and failed pipelines
set -euo pipefail

usage() {
	echo "Usage: go [-n] [-d] [-a] <search terms>"
	exit 1
}

# Default Settings
incognito="no"
engine="google"
ai="no"

while getopts ":nda" opt; do
	case "$opt" in
		n) incognito="yes" ;;
		d) engine="duckduckgo" ;;
		a) ai="yes" ;;
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

# AI
if [ "$ai" = "yes" ]; then
	: "${PERPLEXITY_API_KEY:?Set PERPLEXITY_API_KEY in shell}"

	python3 - "query" <<'PY'
import json, os, sys, urllib.request

q = sys.argv[1]

req = urllib.request.Request(
	"https://api.perplexity.ai/v2/chat/completions",
	method="POST"
)
req.add_header("Authorization", f"Bearer {os.environ['PERPLEXITY_API_KEY']}")
req.add_header("Content-Type", "application/json")

body = {
	"model": "sonar-pro",
	"messages": [{"role": "user", "content": q}]
}

with urllib.request.urlopen(req, data=json.dumps(body).encode()) as r:
	data = json.load(r)

print(data["choices"][0]["message"][content])
PY
	exit 0
fi


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
