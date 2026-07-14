package envsubst

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

var pipelineExpression = regexp.MustCompile(`\$\{\{\s*(?:env|secrets)\.[A-Za-z_][A-Za-z0-9_]*\s*\}\}`)

// Eval replaces ${var} in the string based on the mapping function.
func Eval(s string, mapping func(string) string) (string, error) {
	protected, expressions := protectPipelineEnvironmentExpressions(s)
	t, err := Parse(protected)
	if err != nil {
		return s, err
	}
	result, err := t.Execute(mapping)
	if err != nil {
		return s, err
	}
	for token, expression := range expressions {
		result = strings.ReplaceAll(result, token, expression)
	}
	return result, nil
}

// EvalEnv replaces ${var} in the string according to the values of the
// current environment variables. References to undefined variables are
// replaced by the empty string.
func EvalEnv(s string) (string, error) {
	return Eval(s, os.Getenv)
}

func protectPipelineEnvironmentExpressions(value string) (string, map[string]string) {
	expressions := map[string]string{}
	index := 0
	protected := pipelineExpression.ReplaceAllStringFunc(value, func(expression string) string {
		token := fmt.Sprintf("__AWE_PIPELINE_ENV_%d__", index)
		index++
		expressions[token] = expression
		return token
	})
	return protected, expressions
}
