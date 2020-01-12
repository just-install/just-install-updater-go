package jiup

// RuleSet represents a set of rules for packages.
type RuleSet struct {
	rules map[string]Rule
}

// Add adds a rule to the RuleSet. It will panic if there is already a rule for
// the specified package.
func (r *RuleSet) Add(pkg string, rule Rule) {
	panic("not implemented")
}

// TODO: (*RuleSet).{Update(*registry.Registry),Test()}
