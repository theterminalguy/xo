package set

import (
	"errors"
	"fmt"
)

// Comparable an interface that must be implemented to compare members of the set
type Comparable interface {
	equals(a *interface{}, b *interface{}) bool
}

// Set a collection of unique elements
type Set struct {
	members    []interface{}
	memberType string
}

// Equals check if two members of a set are equal
func (s *Set) equals(a *interface{}, b *interface{}) bool {
	return *a == *b
}

// Members returns all members of the set
func (s *Set) Members() []interface{} {
	return s.members
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
		if s.equals(&member, &element) {
			return false, errors.New("element already exist in the set")
		}
	}

	s.members = append(s.members, element)

	return true, nil
}
