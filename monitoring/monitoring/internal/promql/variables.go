package promql

import (
	"fmt"
	"strings"
)

// VariableApplier converts Prometheus expressions with template variables into valid
// Prometheus expressions, and vice versa. Keys should just be the name of the variable
// (i.e. without a leading '$') and the corresponding sentinel values are assumed to be
// sufficiently unique that a reversal can be safely done.
type VariableApplier map[string]string

// ApplySentinelValues applies default sentinel variable values to the expression, such
// that the expression is a valid Prometheus query.
func (vars VariableApplier) ApplySentinelValues(expression string) string {
	for name, sentinelValue := range vars {
		varKey := "$" + name

		// If the expression uses the variable in a quoted context ("$var") then it's
		// interpreted as valid PromQL, we don't need to replace it!
		if strings.Contains(expression, fmt.Sprintf("%q", varKey)) {
			continue
		}

		// Otherwise replace all occurrences.
		expression = strings.ReplaceAll(expression, varKey, sentinelValue)
	}
	return expression
}

// RevertDefaults returns the expression that has been modified through ApplyDefaults
// and revert any defaults applied to it.
func (vars VariableApplier) RevertDefaults(originalExpression, appliedExpression string) string {
	for name, sentinelValue := range vars {
		varKey := "$" + name

		if !shouldApplyVar(originalExpression, varKey) {
			continue
		}

		appliedExpression = strings.ReplaceAll(appliedExpression, sentinelValue, varKey)
	}
	return appliedExpression
}

// If the expression uses the variable in a quoted context ("$var") then it's
// interpreted as valid PromQL, we don't need to replace it!
func shouldApplyVar(originalExpression string, varKey string) bool {
	return !strings.Contains(originalExpression, fmt.Sprintf("%q", varKey))
}
