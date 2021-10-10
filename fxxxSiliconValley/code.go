package fxxxSiliconValley

func ptr(i *int) {
	*i = 32
}
func getSumAndSub(i, j int) (int, int) {

	if sub := i - j; sub > 0 {
		return i + j, sub
	} else {
		return i + j, -sub
	}
}
