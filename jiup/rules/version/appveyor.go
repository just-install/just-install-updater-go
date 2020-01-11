package v

import (
	"errors"
	"net/http"
	"net/url"

	c "github.com/just-install/just-install-updater-go/jiup/rules/common"
	h "github.com/just-install/just-install-updater-go/jiup/rules/helper"
)

// AppVeyorBranch returns a version extractor for an AppVeyor branch.
func AppVeyorBranch(repo, branch string) c.VersionExtractorFunc {
	return func() (string, error) {
		var build struct {
			Build struct {
				Version string
			}
		}

		if err := h.GetJSON(
			nil,
			"https://ci.appveyor.com/api/projects/"+repo+"/branch/"+url.PathEscape(branch),
			map[string]string{"Accept": "application/json"},
			[]int{http.StatusOK},
			&build,
		); err != nil {
			return "", err
		} else if build.Build.Version == "" {
			return "", errors.New("no version in response")
		}

		return build.Build.Version, nil
	}
}
