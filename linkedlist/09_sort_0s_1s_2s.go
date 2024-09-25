package linkedlist

func sort012(head *Node) *Node {

	if head == nil {
		return nil
	}

	var zeroHead *Node = nil
	var zeroTail *Node = nil

	var oneHead *Node = nil
	var oneTail *Node = nil

	var twoHead *Node = nil
	var twoTail *Node = nil

	curr := head

	for curr != nil {

		if curr.Data == 0 {
			if zeroHead == nil {
				zeroHead = curr
				zeroTail = curr
			} else {
				zeroTail.Next = curr
				zeroTail = curr
			}
		} else if curr.Data == 1 {
			if oneHead == nil {
				oneHead = curr
				oneTail = curr
			} else {
				oneTail.Next = curr
				oneTail = curr
			}
		} else if curr.Data == 2 {
			if twoHead == nil {
				twoHead = curr
				twoTail = curr
			} else {
				twoTail.Next = curr
				twoTail = curr
			}
		}

		curr = curr.Next

	}

	if zeroTail != nil {
		zeroTail.Next = oneHead
	}

	if oneTail != nil {
		oneTail.Next = twoHead
	}

	if twoTail != nil {
		twoTail.Next = nil
	}

	if zeroHead != nil {
		return zeroHead
	} else if oneHead != nil {
		return oneHead
	} else {
		return twoHead
	}

}
