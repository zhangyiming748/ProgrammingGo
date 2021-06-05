package addNode

import (
	"log"
	"testing"
)

/*
描述
假设链表中每一个节点的值都在 0 - 9 之间，那么链表整体就可以代表一个整数。
给定两个这种链表，请生成代表两个整数相加值的结果链表。
例如：链表 1 为 9->3->7，链表 2 为 6->3，最后生成新的结果链表为 1->0->0->0。
示例1
输入：
[9,3,7],[6,3]
复制
返回值：
{1,0,0,0}
复制
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func TestAddList(t *testing.T) {
	head1 := new(ListNode)
	head2 := new(ListNode)
	head1.Val = 9
	head1.Next.Val = 3
	head1.Next.Next.Val = 7
	head2.Val = 6
	head2.Next.Val = 3
	_ = addInList(head1, head2)
}

func addInList(head1 *ListNode, head2 *ListNode) *ListNode {
	// write code here
	val1 := head1.Val
	val2 := head2.Val
	next1 := head1.Next
	next2 := head2.Next
	for {

		if next1.Next != nil {
			next1 = next1.Next
			val1 = 10*val1 + next1.Val
		} else {
			log.Println(val1)
			break
		}
	}
	for {
		if next2.Next != nil {
			next2 = next2.Next
			val2 = 10*val2 + next2.Val
		} else {
			log.Println(val2)
			break
		}
	}
	return head2
}
