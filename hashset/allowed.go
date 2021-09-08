package hashset

import (
	"errors"
)

type Hashset struct {
	set map[string]bool
}

func (h *Hashset) Add(value string) error {
	if _, ok := h.set[value]; ok {
		return errors.New("Cannot add duplicate values.")
	}

	h.set[value] = true
	return nil
}

func New() *Hashset {
	hashset := Hashset{make(map[string]bool)}
	return &hashset
}

func (h *Hashset) Contains(value string) bool {
	return h.set[value]
}
