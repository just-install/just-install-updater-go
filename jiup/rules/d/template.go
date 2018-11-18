package d

import (
	"errors"
	"strings"

	"github.com/just-install/just-install-updater-go/jiup/rules/h"

	"github.com/just-install/just-install-updater-go/jiup/rules/c"
)

// Template creates a download link based on substituting {{.Version}}. Leave a template empty if no link for that version.
func Template(x86Tmpl, x64Tmpl string) c.DownloadExtractorFunc {
	return func(version string) (*string, *string, error) {
		if x86Tmpl == "" && x64Tmpl == "" {
			return nil, nil, errors.New("at least one of x86 and x64 templates must be defined")
		}

		r := func(i string) string {
			o := i
			o = strings.Replace(o, "{{.Version}}", version, -1)
			o = strings.Replace(o, "{{.VersionU}}", strings.Replace(version, ".", "_", -1), -1)
			o = strings.Replace(o, "{{.VersionD}}", strings.Replace(version, ".", "-", -1), -1)
			o = strings.Replace(o, "{{.Version0}}", strings.Split(version, ".")[0], -1)
			o = strings.Replace(o, "{{.VersionN}}", strings.Replace(version, ".", "", -1), -1)
			return o
		}

		var x86dl, x64dl *string
		if x86Tmpl != "" {
			x86dl = h.StrPtr(r(x86Tmpl))
		}
		if x64Tmpl != "" {
			x64dl = h.StrPtr(r(x64Tmpl))
		}
		return x86dl, x64dl, nil
	}
}
