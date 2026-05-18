// palcheck – tiny palindrome checker
//
// Run with: go run main.go [optional-file]
// If a file path is supplied, lines are read from the file; otherwise, STDIN is used.

package main

import (
    "bufio"
    "fmt"
    "os"
    "regexp"
    "strings"
)

// isPalindrome returns true if s is a palindrome after stripping
// non‑alphanumeric characters and normalising case.
func isPalindrome(s string) bool {
    // Keep only letters and digits.
    re := regexp.MustCompile(`[A-Za-z0-9]`)
    cleaned := strings.ToLower(strings.Join(re.FindAllString(s, -1), ""))
    // Compare mirrored characters.
    for i, j := 0, len(cleaned)-1; i < j; i, j = i+1, j-1 {
        if cleaned[i] != cleaned[j] {
            return false
        }
    }
    return true
}

func process(scanner *bufio.Scanner) {
    for scanner.Scan() {
        line := scanner.Text()
        if isPalindrome(line) {
            fmt.Printf("✅ %s\n", line)
        } else {
            fmt.Printf("❌ %s\n", line)
        }
    }
    if err := scanner.Err(); err != nil {
        fmt.Fprintf(os.Stderr, "error reading input: %v\n", err)
        os.Exit(1)
    }
}

func main() {
    var scanner *bufio.Scanner
    if len(os.Args) > 1 {
        f, err := os.Open(os.Args[1])
        if err != nil {
            fmt.Fprintf(os.Stderr, "cannot open file %s: %v\n", os.Args[1], err)
            os.Exit(1)
        }
        defer f.Close()
        scanner = bufio.NewScanner(f)
    } else {
        scanner = bufio.NewScanner(os.Stdin)
    }
    process(scanner)
}
