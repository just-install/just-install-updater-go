package rules

import (
	// dot imports are used here, as the only exposed variable in this package
	// is Rules, and this makes it possible for the rule definitions to be way
	// cleaner.
	. "github.com/just-install/just-install-updater-go/jiup-ng"
	. "github.com/just-install/just-install-updater-go/jiup-ng/sources"
	. "github.com/just-install/just-install-updater-go/jiup-ng/util"
)

// Rules contains rules for jiup-go.
var Rules = new(RuleSet)

func init() {
	Rules.Add("bootnext", RuleMix(
		AppVeyorBranchVersion{
			Repo:   `geek1011/bootnext`,
			Branch: `master`,
		},
		AppVeyorArtifact{
			Files: RegexpMap{
				Arch64: Literal(`bootnext.msi`),
			},
		},
	))
}
