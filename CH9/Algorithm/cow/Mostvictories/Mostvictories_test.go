package Mostvictories

import "testing"

/*
描述
Alice和Bob打牌，每人都有n张牌
Alice的牌里有p1张石头牌，q1张剪刀牌，m1张布牌。
Bob的牌里有p2张石头牌，q2张剪刀牌，m2张布牌。
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
备注：
对于100\%的数据，1\leq n\leq 1e9,0\leq p1,p2,q1,q2,m1,m2\leq 1e9,p1+q1+m1=n,p2+q2+m2=n对于100%的数据，1≤n≤1e9,0≤p1,p2,q1,q2,m1,m2≤1e9,p1+q1+m1=n,p2+q2+m2=n
*/
func Mostvictories(n int, p1 int, q1 int, m1 int, p2 int, q2 int, m2 int) int {
	// write code here
}
func TestMostvictories(t *testing.T) {

}
