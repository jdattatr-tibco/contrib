package json

import (
	"fmt"
	"strings"

	"github.com/oliveagle/jsonpath"
	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expression/function"
)

func init() {
	_ = function.Register(&fnCheckExists{})
}

type fnCheckExists struct {
}

// Name returns the name of the function
func (fnCheckExists) Name() string {
	return "checkExists"
}

// Sig returns the function signature
func (fnCheckExists) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeAny, data.TypeString}, false
}

// Eval executes the function
func (fnCheckExists) Eval(params ...interface{}) (interface{}, error) {
	expression, ok := params[1].(string)
	if !ok {
		return false, fmt.Errorf("The JSON key/path must be a string")
	}
	//tmp fix to take $loop as $. for now
	if strings.HasPrefix(strings.TrimSpace(expression), "$loop.") {
		expression = strings.Replace(expression, "$loop", "$", -1)
	}
	_, err := jsonpath.JsonPathLookup(params[0], expression)
	if err != nil {
		return false, err
	}
	return true, nil
}