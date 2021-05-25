package cow

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseList(pHead *ListNode) *ListNode {
	var pNewHead *ListNode
	tempNode := new(ListNode)
	for {
		if pHead == nil {
			break
		}
		tempNode = pHead
		pHead = pHead.Next
		tempNode.Next = pNewHead
		pNewHead = tempNode
	}
	return pNewHead
}
