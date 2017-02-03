package goorg

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"strings"
)

// Headline is a line preceded by a depth
type Headline struct {
	Title     string
	Text      bytes.Buffer
	Depth     int
	Headlines []*Headline
}

// FromFile parses a file into a HeadLine
func FromFile(name string) (Headline, error) {
	return Headline{}, nil
}

func parseBody(body io.Reader) (*Headline, error) {
	scanner := bufio.NewScanner(body)

	root := Headline{}
	var current *Headline
	var parsed *Headline
	var err error

	for scanner.Scan() {
		t := scanner.Text()
		parsed, err = parseHeadline(t)

		if err != nil {
			return nil, err
		}

		// first time through
		if current == nil {
			if parsed == nil {
				return nil, errors.New("parsing Body that doesn't start with a Headline")
			}
			current = parsed
			continue
		}

		// Not a new headline
		if (parsed == nil) || (parsed.Depth > current.Depth) {
			current.AddText(t)
			continue
		}

		// Is a new headline
		root.AddHeadline(current)
		current = parsed
	}

	// The last headline won't be added
	root.AddHeadline(current)

	return &root, nil
}

func (h *Headline) AddText(s string) {
	h.Text.WriteString(s)
	h.Text.WriteString("\n")
}

func (h *Headline) AddHeadline(new *Headline) {
	h.Headlines = append(h.Headlines, new)
}

// FromLine parses a line into a Headline
func parseHeadline(s string) (*Headline, error) {
	if !strings.HasPrefix(s, "*") {
		// not a headline
		return nil, nil
	}

	d := 0
	for d < len(s[d:]) && strings.HasPrefix(s[d:], "*") {
		d++
	}

	h := Headline{
		Title: strings.Trim(s[d:], " "),
		Depth: d,
	}
	return &h, nil
}