package rules

import "github.com/grafana/grafana-plugin-sdk-go/data"

// Rule defines a rule for checking bool.
type Rule struct {
	Name         string
	IsCompatible func(frames data.Frames) bool
}

// Rules is a collection of Rules.
type Rules []Rule

func (rules *Rules) Compliance(frames data.Frames) []RuleError {
	errors := []RuleError{}
	for _, rule := range *rules {
		if !rule.IsCompatible(frames) {
			errors = append(errors, RuleError{Message: rule.Name})
		}
	}
	return errors
}
