// Package group parsed content and files in form of /etc/group
package group

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strings"
)

// An Entry contains all the fields for a specific user
type Entry struct {
	Pass    string
	GID     string
	Members []string
}

// Parse opens the '/etc/group' file and parses it into a map from usernames
// to Entries
func Parse() (map[string]Entry, error) {
	return ParseFile("/etc/group")
}

// ParseFile opens the file and parses it into a map from usernames to Entries
func ParseFile(path string) (map[string]Entry, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	return ParseReader(file)
}

// ParseReader consumes the contents of r and parses it into a map from
// usernames to Entries
func ParseReader(r io.Reader) (map[string]Entry, error) {
	lines := bufio.NewReader(r)
	entries := make(map[string]Entry)
	for {
		line, _, err := lines.ReadLine()
		if err != nil {
			break
		}
		name, entry, err := parseLine(string(cloneBytes(line)))
		if err != nil {
			return nil, err
		}
		entries[name] = entry
	}
	return entries, nil
}

func parseLine(line string) (string, Entry, error) {
	fs := strings.Split(line, ":")
	if len(fs) != 4 {
		return "", Entry{}, errors.New("Unexpected number of fields in /etc/group")
	}
	return fs[0], Entry{
		Pass:    fs[1],
		GID:     fs[2],
		Members: strings.Split(fs[3], ","),
	}, nil
}

func cloneBytes(x []byte) []byte {
	y := make([]byte, len(x))
	copy(y, x)
	return y
}
