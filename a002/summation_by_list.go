package a002

type DigitList struct {
	Val  int
	Next *DigitList
}

func sumInOrder(l1 *DigitList, l2 *DigitList) *DigitList {
	head := &DigitList{Val: 0, Next: nil}
	if l1 != nil || l2 != nil {
		var carry int
		next := head
		for l1 != nil || l2 != nil {
			var x, y int
			if l1 != nil {
				x = l1.Val
				l1 = l1.Next
			}
			if l2 != nil {
				y = l2.Val
				l2 = l2.Next
			}
			next.Val = (x + y + carry) % 10
			carry = (x + y + carry) / 10
			next.Next = &DigitList{Val: 0, Next: nil}
			next = next.Next
		}
		if carry > 0 {
			next.Val = carry
		}
	}
	return head
}
