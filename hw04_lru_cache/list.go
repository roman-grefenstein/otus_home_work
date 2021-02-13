package hw04_lru_cache //nolint:golint,stylecheck

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{} // значение
	Next  *ListItem   // следующий элемент
	Prev  *ListItem   // предыдущий элемент
}

type list struct {
	len   int
	front *ListItem
	back  *ListItem
}

func (l list) Len() int {
	return l.len
}

func (l list) Front() *ListItem {
	return l.front
}

func (l list) Back() *ListItem {
	return l.back
}

func (l *list) PushFront(v interface{}) *ListItem {
	newListItem := &ListItem{
		Value: v,
		Prev:  nil,
		Next:  l.front,
	}
	if l.front != nil {
		l.front.Prev = newListItem
	}
	l.front = newListItem
	if l.len == 0 {
		l.back = newListItem
	}
	l.len++
	return newListItem
}

func (l *list) PushBack(v interface{}) *ListItem {
	newListItem := &ListItem{
		Value: v,
		Prev:  l.back,
		Next:  nil,
	}
	if l.back != nil {
		l.back.Next = newListItem
	}
	l.back = newListItem
	if l.len == 0 {
		l.front = newListItem
	}
	l.len++
	return newListItem
}

func (l *list) Remove(i *ListItem) {
	if i.Next == nil {
		i.Prev.Next = nil
		l.back = i.Prev
	} else {
		i.Next.Prev = i.Prev
	}

	if i.Prev == nil {
		i.Next.Prev = nil
		l.front = i.Next
	} else {
		i.Prev.Next = i.Next
	}
	l.len--
}

func (l *list) MoveToFront(i *ListItem) {
	if i == l.front {
		return
	}
	l.PushFront(i.Value)
	l.Remove(i)
}

func NewList() List {
	return new(list)
}
