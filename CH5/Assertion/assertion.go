package Assertion



func assertion(i interface{}) string {
	switch i.(type) {
	case int:
		return "is int"
	case string:
		return "is string"
	case bool:
		return "is bool"
	case float64:
		return "is float64"
	case float32:
		return "is float32"
	default:
		return "is unknown"
	}
}
