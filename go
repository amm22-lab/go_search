#!/usr/bin/env bash

# Safer script behaviour: exit on errors, on unset vars, and failed pipelines
set -euo pipefail

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

# Builds the final search URL
url="https://www.google.com/search?q=$encoded"

# Opens in your default browser
open "$url"
