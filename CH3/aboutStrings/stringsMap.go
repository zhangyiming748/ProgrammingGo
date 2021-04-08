package aboutStrings

import "strings"

func replace(s string)string  {
	transformed:= func(r rune)rune {
		if r==' '{
			return 0
		}
		return r
	}
	return strings.Map(transformed,s)
}
