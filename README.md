# palcheck

**palcheck** – a ultra‑small, zero‑dependency command‑line tool to test if strings are palindromes.

## Features
- Reads from STDIN or a file passed as the first argument.
- Ignores case, whitespace, punctuation, and any non‑alphanumeric characters.
- Outputs `✅ <original line>` for palindromes and `❌ <original line>` for non‑palindromes.
- Single‑file Go program – just `go run main.go` and you’re good to go.

## Installation
```bash
# Clone (or just copy the single file)
git clone https://github.com/yourname/palcheck.git
cd palcheck
# Build (optional)
go build -o palcheck main.go
```

## Usage
```bash
# From a pipe
echo -e "Madam\nHello" | ./palcheck
# From a file
./palcheck sentences.txt
```

## License
MIT – see the `LICENSE` file (included in the repository).