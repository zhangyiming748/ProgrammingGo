package biggerArea

import (
	"math"
	"testing"
)

/*
double calc(int a, int b, int c, int d) {
        if (a >= b + c + d) return 0;
        if (b >= a + c + d) return 0;
        if (c >= a + b + d) return 0;
        if (d >= a + b + c) return 0;
        double p = (a + b + c + d) / 2.0;
        return sqrt((p - a) * (p - b) * (p - c) * (p - d));
    }
    double solve(vector<int>& a) {
        // write code here
        double ans = 0;
        for (int i = 0; i < a.size(); i++)
            for (int j = 0; j < i; j++)
                for (int k = 0; k < j; k++)
                    for (int l = 0; l < k; l++)
                        ans = max(ans, calc(a[i], a[j], a[k], a[l]));
        return ans;
    }
*/
func TestSolve(t *testing.T) {

	//10.95445
	res := solve([]int{1, 2, 3, 4, 5})
	t.Log(res)

}
func calc(a, b, c, d int) float64 {
	if illegal(a, b, c, d) {
		return 0
	}
	p := (a + b + c + d) / 2.00
	res := math.Sqrt(float64((p - a) * (p - b) * (p - c) * (p - d)))
	return res
}
func illegal(a, b, c, d int) bool {
	if a >= b+c+d {
		return true
	}
	if b >= a+c+d {
		return true
	}
	if c >= a+b+d {
		return true
	}
	if d >= a+b+c {
		return true
	}
	return false
}
func solve(a []int) float64 {
	// write code here
	var res float64
	for i := 0; i < len(a); i++ {
		for j := 0; j < i; j++ {
			for p := 0; p < j; p++ {
				for q := 0; q < p; q++ {
					var this float64 = calc(a[i], a[j], a[p], a[q])
					if this > res {
						res = this
					}
				}
			}
		}
	}
	return res
}
