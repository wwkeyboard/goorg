package goorg

import (
	"strings"
	"testing"
)

func TestParseSimpleHeading(t *testing.T) {
	l := "* Test Heading"
	h, err := parseHeadline(l)

	if err != nil {
		t.Fatal(err)
	}

	if h.Title != "Test Heading" {
		t.Error("didn't extract text from heading")
	}
}

func TestParseLineWithoutHeading(t *testing.T) {
	h, err := parseHeadline(" * blah")
	if err != nil {
		t.Error(err)
	}
	if h != nil {
		t.Error("found something for a non-headline")
	}

}

func TestReadFile(t *testing.T) {
	o, err := FromFile("./test.org")
	if err != nil {
		t.Error(err)
	}

	if len(o.Headlines) != 2 {
		t.Error("Didn't parse headings")
	}
}

var smallFile = `* line one
** subheading
* line two
something
** another`

func TestParseBody(t *testing.T) {
	h, err := parseBody(strings.NewReader(smallFile))
	if err != nil {
		t.Fatal(err)
	}

	if len(h.Headlines) != 2 {
		t.Error("Didn't count headlines correctly")
	}

	first := h.Headlines[0]
	if first.Title != "line one" {
		t.Error("Didn't parse first title")
	}

	second := h.Headlines[1]
	if second.Title != "line two" {
		t.Error("Didn't parse second title")
	}

	if !strings.Contains(second.Text.String(), "something") {
		t.Error("Didn't parse second body Text")
	}
}
