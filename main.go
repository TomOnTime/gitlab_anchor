// Various functions helpful when dealing with MarkDown files.
package markdownutils

import (
	"strings"
	"unicode"
)

// CreateGitHubAnchor produces GitHub-compatible HTML anchors.
func CreateGitHubAnchor(text string) string {
	// The algorithm was created via blackbox testing and may not be a
	// perfect emulation.  So far it is valid for all tests I can
	// find.  Please report any cases where the emulation is not perfect.
	var anchorName []rune

	for _, r := range []rune(strings.TrimSpace(text)) {
		switch {
		case r == ' ' || r == '-':
			anchorName = append(anchorName, '-')
		case unicode.IsLetter(r) || unicode.IsNumber(r):
			anchorName = append(anchorName, unicode.ToLower(r))
		default:
		}
	}

	return string(anchorName)
}

// CreateGitLabAnchor produces GitLab-compatible HTML anchors.
func CreateGitLabAnchor(text string) string {
	// The algorithm was created by translating GitLab's algorithm
	// from the original Ruby.  No known incompatibilities at this
	// time.  Please report any cases where the emulation is not perfect.
	var anchorName []rune
	var lastWasDash = false

	for _, r := range []rune(strings.TrimSpace(text)) {
		switch {
		case r == ' ' || r == '-':
			if !lastWasDash {
				anchorName = append(anchorName, '-')
				lastWasDash = true
			}
		case unicode.IsLetter(r) || unicode.IsNumber(r):
			anchorName = append(anchorName, unicode.ToLower(r))
			lastWasDash = false
		default:
		}
	}

	return string(anchorName)
}
