package iformula

import (
	"fmt"
	"github.com/expr-lang/expr"
)

// ValidateExpression 检查给定的表达式是否合法，并可选地验证变量是否存在
func ValidateExpression(expression string, allowedVariables map[string]interface{}) error {
	// 编译表达式
	_, err := expr.Compile(expression, expr.Env(allowedVariables))
	if err != nil {
		return fmt.Errorf("表达式无效: %w", err)
	}
	return nil
}
