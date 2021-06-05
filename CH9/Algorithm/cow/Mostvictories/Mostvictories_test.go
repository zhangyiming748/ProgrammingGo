package Mostvictories

import "testing"

/*
描述
Alice和Bob打牌，每人都有n张牌

Alice知道Bob每次要出什么牌，请你安排策略，使Alice获胜次数最多。
输出获胜次数。
示例1
输入：
3,3,0,0,0,0,3
复制
返回值：
0
复制
说明：
Alice只有石头，Bob只有布，每一场Alice都必败，所以Alice只能赢0局
示例2
输入：
6,2,2,2,2,2,2
复制
返回值：
6
复制
说明：
Alice可以在Bob出石头的时候出布，在Bob出布的时候出剪刀，在Bob出剪刀的时候出石头，按照这个策略Alice最多能赢下所有的比赛，所以最多能赢6局
Alice的牌里有p1张石头牌，q1张剪刀牌，m1张布牌。
Bob的牌里有p2张石头牌，q2张剪刀牌，m2张布牌。  def Mostvictories(self , n , p1 , q1 , m1 , p2 , q2 , m2 ):
        k=0
        if p1>=q2:
            k=k+q2
        else:
            k=k+p1
        if q1>=m2:
            k=k+m2
        else:
            k=k+q1
        if m1>=p2:
            k=k+p2
        else:
            k=k+m1
        return k
*/

func Mostvictories(n int, p1 int, q1 int, m1 int, p2 int, q2 int, m2 int) int {
	count := 0
	if p1 >= q2 {
		count += q2
	} else {
		count += p1
	}
	if q1 >= m2 {
		count += m2
	} else {
		count += q1
	}
	if m1 >= p2 {
		count += p2
	} else {
		count += m1
	}

	return count
}
func TestMostvictories(t *testing.T) {

	r1 := Mostvictories(3, 3, 0, 0, 0, 0, 3)
	r2 := Mostvictories(1000000000, 110000000, 320000000, 570000000, 110000000, 320000000, 570000000)
	t.Logf("r1=%d\tr2=%d\n", r1, r2)
}
