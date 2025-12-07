package utils

type Fifo struct {
	fifo []interface{}
	tail int
	head int
	size int
}

func NewFifo(s int) Fifo {
	f := Fifo{
		fifo: make([]interface{}, s),
		tail: 0,
		head: 0,
		size: s,
	}
	return f
}

func (f *Fifo) Push(i interface{}) {
	f.fifo[f.head] = i
	f.head = (f.head + 1) % f.size
	if f.head == f.tail {
		f.tail = (f.tail + 1) % f.size
	}
}

func (f *Fifo) Get() *interface{} {
	f.head = (f.head - 1 + f.size) % f.size
	if f.head == f.tail {
		return nil
	} else {
		return &f.fifo[f.head]
	}
}
