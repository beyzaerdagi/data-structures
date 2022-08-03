package queue

func (q *Queue) Dequeue() {

	if q.First == nil {
		return
	}
	if q.Size == 1 {
		q.First = nil
		q.Last = nil
	} else {
		q.First = q.First.Next
	}
	q.Size--
}
