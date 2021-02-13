package hw04_lru_cache //nolint:golint,stylecheck

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

type cacheItem struct {
	Key   Key
	Value interface{}
}

func (l *lruCache) Set(key Key, value interface{}) bool {
	val, ok := l.items[key]
	newCashItem := cacheItem{Key: key, Value: value}
	if !ok {
		if l.queue.Len() == l.capacity {
			backQueueElement := l.queue.Back()
			delete(l.items, backQueueElement.Value.(cacheItem).Key)
			l.queue.Remove(backQueueElement)
		}
		newElement := l.queue.PushFront(newCashItem)
		l.items[key] = newElement
		return false
	}
	val.Value = newCashItem
	l.queue.MoveToFront(val)
	l.items[key] = l.queue.Front()
	return true
}

func (l *lruCache) Get(key Key) (interface{}, bool) {
	val, ok := l.items[key]
	if !ok {
		return nil, false
	}

	l.queue.MoveToFront(val)
	l.items[key] = l.queue.Front()
	return val.Value.(cacheItem).Value, true
}

func (l *lruCache) Clear() {
	l.queue = NewList()
	l.items = make(map[Key]*ListItem, l.capacity)
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
