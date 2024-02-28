// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !windows

package url

// Code copied from https://github.com/golang/go/blob/7c2b69080a0b9e35174cc9c93497b6e7176f8275/src/cmd/go/internal/web/url_other_test.go

var urlTests = []struct {
	url          string
	filePath     string
	canonicalURL string // If empty, assume equal to url.
	wantErr      string
}{
	// Examples from RFC 8089:
	{
		url:      `file:///path/to/file`,
		filePath: `/path/to/file`,
	},
	{
		url:          `file:/path/to/file`,
		filePath:     `/path/to/file`,
		canonicalURL: `file:///path/to/file`,
	},
	{
		url:          `file://localhost/path/to/file`,
		filePath:     `/path/to/file`,
		canonicalURL: `file:///path/to/file`,
	},

	// We reject non-local files.
	{
		url:     `file://host.example.com/path/to/file`,
		wantErr: "file URL specifies non-local host",
	},
}
