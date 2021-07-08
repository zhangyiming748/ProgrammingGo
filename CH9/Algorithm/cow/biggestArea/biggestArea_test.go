package biggestArea

import (
	"math"
	"sort"
	"testing"
)

func TestSolve(t *testing.T) {
	array := []int{1, 2, 3, 4, 5}
	ret := solve(array)
	t.Log(ret)
}

/*
double solve(vector<int>& a) {
        double res = 0;
        for(int i=0;i<a.size()-3;i++){
            for(int j=i+1;j<a.size()-2;j++){
                for(int k=j+1;k<a.size()-1;k++){
                    for(int l=k+1;l<a.size();l++){
                        int a1=a[i],a2=a[j],a3=a[k],a4=a[l];
                        if(a1<a2+a3+a4&&a2<a1+a3+a4&&a3<a1+a2+a4&&a4<a1+a2+a3){
                            double z = (double)(a1+a2+a3+a4)/2;
                            res = max(res,sqrt((z-a1)*(z-a2)*(z-a3)*(z-a4)));
                        }
                    }
                }
            }
        }
        return res;
    }
*/
func solve(a []int) float64 {
	var res float64
	sort.Slice(a, func(i, j int) bool {
		return a[i] > a[j]
	})
	for i := 0; i < len(a)-3; i++ {
		for j := i + 1; j < len(a)-2; j++ {
			for k := j + 1; k < len(a)-1; k++ {
				for l := k + 1; l < len(a); l++ {
					a1, a2, a3, a4 := a[i], a[j], a[k], a[l]
					if a1 < a2+a3+a4 && a2 < a1+a3+a4 && a3 < a1+a2+a4 && a4 < a1+a2+a3 {
						var x float64 = float64(a1+a2+a3+a4) / 2
						res = math.Sqrt((x - float64(a1)) * (x - float64(a2)) * (x - float64(a3)) * (x - float64(a4)))
					}
				}
			}
		}
	}
	return res
}
