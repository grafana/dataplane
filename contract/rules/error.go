package rules

type RuleError struct {
	Message string
}

func (r *RuleError) String() string {
	if r == nil {
		return ""
	}
	return r.Message
}
