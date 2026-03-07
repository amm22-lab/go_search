# go

A lightweight Bash CLI tool that lets you search the web or query an AI directly from your terminal — without ever leaving your coding environment.

Built for developers who want to stay in flow.

---

## Features

- 🔍 **Instant web search** — launch Google searches straight from the terminal
- 🦆 **DuckDuckGo support** — switch search engines with a single flag
- 🕵️ **Incognito mode** — open searches in a private Chrome window
- 🤖 **AI mode** — ask questions and get answers inline in the terminal via the Perplexity API
- 🔐 **Safe scripting** — built with `set -euo pipefail` for robust error handling
- ⚡ **Zero dependencies** — pure Bash + Python 3 (already on your Mac)

---

## Requirements

- macOS
- Python 3
- Google Chrome (for incognito mode)
- A [Perplexity API key](https://www.perplexity.ai/) (only required for AI mode)

---

## Installation

**1. Clone the repository**

```bash
git clone https://github.com/amm22-lab/go.git
cd go
```

**2. Make the script executable**

```bash
chmod +x go.sh
```

**3. Move it to your PATH so you can run it from anywhere**

```bash
sudo mv go.sh /usr/local/bin/go
```

**4. Set your Perplexity API key (required for AI mode only)**

Add this to your `~/.zshrc` or `~/.bash_profile`:

```bash
export PERPLEXITY_API_KEY="your_api_key_here"
```

Then reload your shell:

```bash
source ~/.zshrc
```

---

## Usage

```
go [-n] [-d] [-a] <query>
```

| Flag | Description |
|------|-------------|
| `-n` | Open search in incognito mode |
| `-d` | Use DuckDuckGo instead of Google |
| `-a` | Use AI mode (Perplexity API) |

---

## Examples

**Basic Google search**
```bash
go how to reverse a linked list in python
```

**Search with DuckDuckGo**
```bash
go -d best vim shortcuts
```

**Open search in incognito**
```bash
go -n javascript promises explained
```

**Ask the AI a question inline**
```bash
go -a what is the difference between REST and GraphQL
```

```
GraphQL gives clients the ability to request exactly the data they need,
while REST returns fixed data structures defined by the server...
```

**Combine flags — DuckDuckGo in incognito**
```bash
go -n -d privacy focused search engines
```

---

## How It Works

1. Your query is passed as a string argument
2. Python's `urllib.parse.quote` URL-encodes it so spaces and special characters are handled safely
3. In standard mode, the encoded URL is opened in your default browser via `open`
4. In incognito mode (`-n`), Chrome is launched with the `--incognito` flag
5. In AI mode (`-a`), the query is sent to the Perplexity API (`sonar-pro` model) and the response is printed directly in the terminal

---

## Project Structure

```
go/
└── go.sh       # Main CLI script
```

---

## Author

**Alex Mutua**
[github.com/amm22-lab](https://github.com/amm22-lab)
