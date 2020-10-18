package set

import (
	"errors"
	"fmt"
)

// Comparable an interface that must be implemented to compare members of the set
type Comparable interface {
	Equals(a *interface{}, b *interface{}) bool
}

// Set a collection of unique elements
type Set struct {
	members    []interface{}
	memberType string
}

// Equals check if two members of a set are equal
func (s *Set) Equals(a *interface{}, b *interface{}) bool {
	return *a == *b
}

// Members returns all members of the set
func (s *Set) Members() []interface{} {
	return s.members
}

// Size returns the total number of elements in the set
func (s *Set) Size() int {
	return len(s.members)
}

// Empty returns true if there no elements in the sets
func (s *Set) Empty() bool {
	return len(s.members) == 0
}

// Add add an element to a set
func (s *Set) Add(element interface{}) (bool, error) {
	elementType := fmt.Sprintf("%T", element)
	if s.memberType == "" {
		s.memberType = elementType
	}

	if elementType != s.memberType {
		return false, fmt.Errorf("set can only be of type %s", s.memberType)
	}

	for _, member := range s.members {
		if s.Equals(&member, &element) {
			return false, errors.New("element already exist in the set")
		}
	}

	s.members = append(s.members, element)

	return true, nil
}

// Remove remove an element from a set
func (s *Set) Remove(element interface{}) bool {
	for i, member := range s.members {
		if s.Equals(&member, &element) {
			bi := i - 1
			if bi < 0 {
				bi = 0
			}

			s1 := s.members[0:i]
			s2 := s.members[i+1:]

			s.members = s1
			s.members = append(s.members, s2...)

			return true
		}
	}

	return false
}

// Contains checks if an element exists in a set
func (s *Set) Contains(element interface{}) bool {
	for _, member := range s.members {
		if s.Equals(&member, &element) {
			return true
		}
	}

	return false
}
