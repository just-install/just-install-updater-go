package helpers

import "errors"

// Verbose controls verbosity.
var Verbose = false

// ErrExtractorNotImplemented should be returned if an extractor is not implemented
var ErrExtractorNotImplemented = errors.New("extractor not implemented yet")

// VersionExtractorFunc represents a function which extracts the version.
type VersionExtractorFunc func() (version string, err error)

// DownloadExtractorFunc represents a function which extracts a download link for a version.
// The version is just a hint if it is for string substitution. If not for string substitution,
// just return the latest version.
// It can return an nil string pointer for x86_64 if not available.
type DownloadExtractorFunc func(version string) (x86 string, x86_64 *string, err error)
