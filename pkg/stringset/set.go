package stringset

import (
	"sort"
	"strings"
)

type Set struct {
	elements map[string]bool
}

func NewSet(data ...string) *Set {
	s := &Set{}
	s.elements = make(map[string]bool)
	for _, v := range data {
		s.elements[v] = true
	}
	return s
}

func (s *Set) Add(data string) {
	s.elements[data] = true
}

func (s *Set) Remove(data string) {
	delete(s.elements, data)
}

func (s *Set) Contains(data string) bool {
	return s.elements[data]
}

func (s *Set) Size() int {
	return len(s.elements)
}

func (s *Set) Elements() []string {
	var elements []string
	for k := range s.elements {
		elements = append(elements, k)
	}
	sort.Strings(elements)
	return elements
}

func (s *Set) ToString() string {
	e := s.Elements()
	return strings.Join(e, ",")
}
