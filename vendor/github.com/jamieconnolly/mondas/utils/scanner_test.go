package utils_test

import (
	"bufio"
	"strings"
	"testing"

	"github.com/jamieconnolly/mondas/utils"
	"github.com/stretchr/testify/assert"
)

func testScanMetadata(t *testing.T, text string, values []string) {
	s := bufio.NewScanner(strings.NewReader(text))
	s.Split(utils.ScanMetadata)
	var i int
	for i = 0; s.Scan(); i++ {
		if i >= len(values) {
			break
		}
		assert.Equal(t, values[i], s.Text())
	}
	assert.Equal(t, len(values), i)
	assert.NoError(t, s.Err())
}

func TestScanMetadata_WithEmptyInput(t *testing.T) {
	testScanMetadata(t, "", []string{})
}

func TestScanMetadata_WithMultiline(t *testing.T) {
	testScanMetadata(t,
		"# foo:\n#   bar\n#\n#   baz\n\n#   qux\n# qux:\n#   baz\n#   bar\n\n#  foo\n",
		[]string{"foo:\n  bar\n\n  baz", "qux:\n  baz\n  bar"},
	)
}

func TestScanMetadata_WithNoMatchingTokens(t *testing.T) {
	testScanMetadata(t, "foo: bar\nbaz: qux\n", []string{})
}

func TestScanMetadata_WithSingleLine(t *testing.T) {
	testScanMetadata(t, "# foo: bar\n\n# baz: qux\n", []string{"foo: bar", "baz: qux"})
}

func TestScanMetadata_WithNoNewlineToEnd(t *testing.T) {
	testScanMetadata(t, "# foo: bar", []string{"foo: bar"})
}
