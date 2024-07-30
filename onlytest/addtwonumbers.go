package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	} else if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}

	pb1 := l1
	pb2 := l2
	var res, tail *ListNode

	nextadd := 0
	for pb1 != nil && pb2 != nil {
		icur := pb1.Val + pb2.Val + nextadd
		nextadd = icur / 10
		down := icur % 10
		cnode := &ListNode{
			Val:  down,
			Next: nil,
		}

		if res == nil {
			res = cnode
			tail = cnode
		} else {
			tail.Next = cnode
			tail = tail.Next
		}
		pb1 = pb1.Next
		pb2 = pb2.Next
	}
	for pb1 != nil {
		midsum := pb1.Val + nextadd
		nextadd = midsum / 10
		cadd := midsum % 10
		cnode := &ListNode{
			Val:  cadd,
			Next: nil,
		}

		tail.Next = cnode
		tail = tail.Next
		pb1 = pb1.Next
	}
	for pb2 != nil {
		midsum := pb2.Val + nextadd
		nextadd = midsum / 10
		cadd := midsum % 10
		cnode := &ListNode{
			Val:  cadd,
			Next: nil,
		}

		tail.Next = cnode
		tail = tail.Next
		pb2 = pb2.Next
	}
	if nextadd != 0 {
		cnode := &ListNode{
			Val:  nextadd,
			Next: nil,
		}
		tail.Next = cnode
	}
	return res
}
