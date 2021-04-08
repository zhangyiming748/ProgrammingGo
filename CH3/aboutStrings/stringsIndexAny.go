package aboutStrings

import "strings"

func returnFirst(s,chars string)int  {
	return strings.IndexAny(s,chars)
}
