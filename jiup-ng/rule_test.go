package jiup

import "testing"

func TestRuleMix(t *testing.T) {
	r := RuleMix(nil, nil)
	var v Versioner = r
	var d Downloader = r
	_, _, _ = r, v, d
}
