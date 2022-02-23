package package_interface_type

type any interface{}

func CastToString(arg any) string {
	castArg, castCheck := arg.(string)

	if castCheck {
		return castArg
	} else {
		return "cast fail"
	}
}

func TypeCheck(arg any) string {
	switch arg.(type) {
		case bool:
			return "bool!"
		case int:
			return "int!"
		case string:
			return "string!"
		default:
			return "other!"
	}
}

type listType = map[string]any
var list = map[string]any{}
func AddListItem(key string, value any) listType {
	list[key] = value

	return list
}
