package bimap

type BiMap[T, V comparable] struct {
	forward map[T]V
	backward map[V]T
}

func NewBiMap[T, V comparable]() *BiMap[T, V] {
	return &BiMap[T, V]{
		forward: make(map[T]V),
		backward: make(map[V]T),
	}
}

func (m *BiMap[T, V]) Get(key T) (V, bool) {
	v, exists := m.forward[key]
	return v, exists
}

func (m *BiMap[T, V]) GetInverse(key V) (T, bool) {
	v, exists := m.backward[key]
	return v, exists
}

func (m *BiMap[T, V]) Insert(key T, value V) {
	m.forward[key] = value
	m.backward[value] = key
}
