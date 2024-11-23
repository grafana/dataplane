package contract

// Compliance represents the result of a compliance check.
type Compliance bool

const (
	// Valid indicates the check passed.
	Valid Compliance = true
	// Invalid indicates the check failed.
	Invalid Compliance = false
)
