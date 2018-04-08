package defs

import (
	"fmt"
	"github.com/Noah-Huppert/kube-bot/chat"
)

// AllLoader loads all other RegLoaderables in the defs package
type AllLoader struct{}

// NewAllLoader creates and returns a new AllLoader
func NewAllLoader() *AllLoader {
	return &AllLoader{}
}

// Load implements RegLoadable for AllLoader by loading all RegLoadables in
// defs package.
func (l AllLoader) Load(registry chat.Registry) error {
	// Augments
	augLdr := NewAugmentLoader()
	if err := augLdr.Load(registry); err != nil {
		return fmt.Errorf("error loading augments: %s", err.Error())
	}

	// Done
	return nil
}
