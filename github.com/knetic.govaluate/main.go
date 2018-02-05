package main

import (
    "github.com/Knetic/govaluate"
    "github.com/bmizerany/assert"
)

func main() {
    expression, _ := govaluate.NewEvaluableExpression("foo > 0")

    parameters := make(map[string]interface{}, 8)

    parameters["foo"] = -1
    result, _ := expression.Evaluate(parameters)
    assert.Equal(nil, result, false)

    parameters["foo"] = 0
    result, _ = expression.Evaluate(parameters)
    assert.Equal(nil, result, false)

    parameters["foo"] = 1
    result, _ = expression.Evaluate(parameters)
    assert.Equal(nil, result, true)
}
