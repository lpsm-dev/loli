package cli

import "strings"

// Version is the CLI version that is built for the release.
func (r *Release) Version() string {
	return strings.TrimPrefix(r.TagName, "v")
}
