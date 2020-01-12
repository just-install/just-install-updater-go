package jiup

import "testing"

func TestRuleMix(t *testing.T) {
	var r *RuleMix
	var v Versioner = r
	var d Downloader = r
	_, _, _ = r, v, d
}
