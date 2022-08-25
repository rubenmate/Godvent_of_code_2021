package helpers

import (
	"bufio"
	"log"
	"os"
	// "strings"
)

func Getlines(path string) []string {
    file, err := os.Open(path)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    sc := bufio.NewScanner(file)
    lines := make([]string, 0)

    // Read through 'tokens' until an EOF is encountered.
    for sc.Scan() {
        lines = append(lines, sc.Text())
    }

    if err := sc.Err(); err != nil {
        log.Fatal(err)
    }
    return lines
}
