package d

import (
	"net/http"
	"net/url"
	"regexp"

	c "github.com/just-install/just-install-updater-go/jiup/rules/common"
	h "github.com/just-install/just-install-updater-go/jiup/rules/helper"
)

// AppVeyorArtifacts returns a download extractor for a AppVeyor artifacts
// with a deployment name matching the provided regexps. Either regexp can be nil.
//
// The version passed to the extractor must be a valid AppVeyor build version.
func AppVeyorArtifacts(repo string, x86FileRe, x64FileRe *regexp.Regexp) c.DownloadExtractorFunc {
	return func(version string) (*string, *string, error) {
		var build struct {
			Build struct {
				Jobs []struct {
					JobID string `json:"jobId"`
				}
			}
		}

		if err := h.GetJSON(
			nil,
			"https://ci.appveyor.com/api/projects/"+repo+"/build/"+url.PathEscape(version),
			map[string]string{"Accept": "application/json"},
			[]int{http.StatusOK},
			&build,
		); err != nil {
			return nil, nil, err
		}

		var x86, x64 *string
		for _, job := range build.Build.Jobs {
			var artifacts []struct {
				FileName string `json:"fileName"`
				Name     string
			}

			if job.JobID == "" {
				continue
			} else if err := h.GetJSON(
				nil,
				"https://ci.appveyor.com/api/buildjobs/"+url.PathEscape(job.JobID)+"/artifacts",
				map[string]string{"Accept": "application/json"},
				[]int{http.StatusOK},
				&artifacts,
			); err != nil {
				return nil, nil, err
			}

			for _, artifact := range artifacts {
				if x86FileRe != nil && x86 == nil && x86FileRe.MatchString(artifact.Name) {
					u := "https://ci.appveyor.com/api/buildjobs/" + url.PathEscape(job.JobID) + "/artifacts/" + url.PathEscape(artifact.FileName)
					x86 = &u
				}
				if x64FileRe != nil && x64 == nil && x64FileRe.MatchString(artifact.Name) {
					u := "https://ci.appveyor.com/api/buildjobs/" + url.PathEscape(job.JobID) + "/artifacts/" + url.PathEscape(artifact.FileName)
					x64 = &u
				}
			}

			if (x86 != nil || x86FileRe == nil) && (x64 != nil || x64FileRe == nil) {
				break
			}
		}
		return x86, x64, nil
	}
}
