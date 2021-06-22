package cow

type character struct {
	data []uint8
}

func (c *character) push(i uint8) {
	c.data = append(c.data, i)

}
func (c *character) pop() byte {
	val := c.data[len(c.data)-1]
	c.data = c.data[:len(c.data)-1]
	return val
}

func isValid(s string) bool {
	b := []byte(s)
	if b[0] == 41 || b[0] == 93 || b[0] == 125 || len(b)%2 != 0 {
		return false
	}
	var c character
	for _, v := range b {
		if v == 40 || v == 91 || v == 123 {
			c.push(v)
		}
		if v == 41 {
			val := c.pop()
			if val != 40 {
				return false
			}
		}
		if v == 93 {
			val := c.pop()
			if val != 91 {
				return false
			}
		}
		if v == 125 {
			val := c.pop()
			if val != 123 {
				return false
			}
		}
	}
	if len(c.data) != 0 {
		return false
	}
	return true
}
