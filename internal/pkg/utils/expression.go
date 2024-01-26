package commUtils

import "strings"

func IsJsonPathExpression(str string) (variName, jsonPathExpression string, yes bool) {
	index := strings.IndexAny(str, ".[")

	if index > -1 {
		yes = true

		variName = str[:index]
		jsonPathExpression = "$" + str[index:]
	} else {
		variName = str
	}

	return
}
