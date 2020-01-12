package sources

import (
	"errors"

	"github.com/just-install/just-install-updater-go/jiup-ng"
)

// TODO: maybe consider moving each source into it's own package?

type appveyorData string

const (
	appVeyorDataJobIDs appveyorData = "appVeyorDataJobIDs" // []string
)

// AppVeyorBranchVersion implements a Versioner for retrieving the latest
// version from an AppVeyor project branch.
type AppVeyorBranchVersion struct {
	Repo   string
	Branch string
}

// Version implements jiup.Versioner.
func (a AppVeyorBranchVersion) Version(data *jiup.RuleData) (string, error) {
	panic("not implemented")
}

// AppVeyorArtifact implements a Downloader for retrieving artifacts from an
// AppVeyor project.
type AppVeyorArtifact struct {
	Repo string // optional if the Versioner is an AppVeyorBranchVersion

	// Files contains regexps to match the files for each supported architecture.
	Files jiup.RegexpMap
}

// Download implements jiup.Downloader.
func (a AppVeyorArtifact) Download(version string, data *jiup.RuleData) (jiup.LinkMap, error) {
	if _, ok := data.Get(appVeyorDataJobIDs); !ok {
		if a.Repo == "" {
			return nil, errors.New("AppVeyorArtifact.Repo is required if AppVeyorBranchVersion was not the Versioner")
		}
		// TODO: get job ids for version build
	}

	for _, jobID := range data.MustGet(appVeyorDataJobIDs).([]string) {
		// TODO: get artifacts for job, match
		_ = jobID
	}

	panic("not implemented")
}
