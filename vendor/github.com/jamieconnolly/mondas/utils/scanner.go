package utils

import "regexp"

var metadataRegexp = regexp.MustCompile(`(?im)^(?:[^\n]*#[ \w]*:[^\n]*\n?(?:#(?:[ ]{3,}.*?)?(?:\n|$))*)`)

// ScanMetadata is a split function for a Scanner that returns each
// section of metadata as a token. It will never return an empty string.
func ScanMetadata(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	if m := metadataRegexp.FindIndex(data); m != nil {
		data = data[m[0]:m[1]]

		if data[len(data)-1] == '\n' || data[len(data)-1] == '\r' {
			data = data[:len(data)-1]
		}

		metadata := make([]byte, 0, len(data))
		for i, b := range data {
			if b == '#' && i == 0 {
				continue
			}

			if b == '#' && i > 0 && (data[i-1] == '\n' || data[i-1] == '\r') {
				continue
			}

			if b == ' ' && i > 0 && (data[i-1] == '#') {
				continue
			}

			metadata = append(metadata, b)
		}

		return m[1], metadata, nil
	}

	return len(data), nil, nil
}
