package queue

func (q *Queue) Enqueue(node *Node) {

	if q.First == nil {
		q.First = node
		q.Last = node
	} else {
		q.Last.Next = node
		q.Last = node
	}
	q.Size++
}
