package algorithm

type Queue struct {
	l *List
}

func NewQueue() *Queue {
	return &Queue{
		l: NewList(),
	}
}

func (q *Queue) Enqueue(value interface{}) {
	q.l.Append(value)
}

func (q *Queue) Dequeue() interface{} {
	if q.l.Len() > 0 {
		e := q.l.First()
		q.l.Remove(e)
		return e.Value
	}
	return nil
}

func (q *Queue) IsEmpty() bool {
	return q.l.Len() == 0
}

func (q *Queue) Clear() {
	q.l = nil
}

func (q *Queue) Size() int {
	return q.l.Len()
}
