package set

type Set struct {
	dense []uint
	sparse []uint
	n uint
}

func New(max uint) *Set {
	return &Set{
		dense: make([]uint, max, max),
		sparse: make([]uint, max, max),
	}
}

func (s *Set) Add(i uint) {
	if s.Contains(i) {
		return
	}
	s.dense[s.n] = i
	s.sparse[i] = s.n
	s.n++
}

func (s *Set) Contains(i uint) bool {
	return s.sparse[i] < s.n && s.dense[s.sparse[i]] == i
}

func (s *Set) Clear() {
	s.n = 0
}
