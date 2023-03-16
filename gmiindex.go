package main

import (
	"bufio"
	"fmt"
	"io"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

type entry struct {
	title string
	path  string
	date  string
}

type ByTitle []entry

func (t ByTitle) Len() int      { return len(t) }
func (t ByTitle) Swap(i, j int) { t[i], t[j] = t[j], t[i] }
func (t ByTitle) Less(i, j int) bool {
	a := t[i].date
	b := t[j].date
	reverse := true
	if a == "" {
		a = t[i].title
		reverse = false
	}
	if b == "" {
		b = t[j].title
		reverse = false
	}
	if reverse {
		return a > b
	} else {
		return a < b
	}
}

func main() {
	files := os.Args[1:]
	dateMatch := regexp.MustCompile(`(\d{4}-\d{2}-\d{2})?.*`)
	entries := make([]entry, len(files))
	for i, path := range files {
		entry, err := getEntry(dateMatch, path)
		fatal("getting entry information", err)
		entries[i] = entry
	}

	sort.Sort(ByTitle(entries))

	for _, e := range entries {
		fmt.Printf("=> %s %s\n", e.path, e.title)
	}
}

func getEntry(dateMatch *regexp.Regexp, path string) (entry, error) {
	title := filepath.Base(path)
	datem := dateMatch.FindStringSubmatch(title)

	date := ""
	if datem != nil && datem[1] != "" {
		date = datem[1]
	}

	if strings.HasSuffix(path, ".gmi") || strings.HasSuffix(path, ".gemini") {
		innerTitle, err := getTitle(path)
		if err != nil {
			return entry{}, err
		}
		if innerTitle != "" {
			title = innerTitle
			if date != "" {
				title = date + " - " + title
			}
		}
	}

	u, err := url.Parse(path)
	if err != nil {
		return entry{}, err
	}
	return entry{
		title: title,
		date:  date,
		path:  u.EscapedPath(),
	}, nil
}

func getTitle(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	r := bufio.NewReader(f)
	line, _, err := r.ReadLine()
	if err != nil {
		if err != io.EOF {
			return "", err
		}
	}
	strLine := string(line)

	if strings.HasPrefix(strLine, "#") {
		return strings.TrimRight(strings.TrimLeft(strLine, "# "), " \t"), nil
	}
	return "", nil
}

func fatal(prefix string, err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, fmt.Sprintf("%s: %v", prefix, err))
		os.Exit(1)
	}
}
