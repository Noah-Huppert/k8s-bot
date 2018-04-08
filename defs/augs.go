package defs

import (
	"fmt"
	"github.com/Noah-Huppert/kube-bot/chat"
)

// AugmentLoader implements the RegLoadable interface for the "reply" augment.
// This specifies which method of messaging the bot should use to reply
type AugmentLoader struct{}

// NewAugmentLoader creates and returns a new AugmentLoader
func NewAugmentLoader() *AugmentLoader {
	return &AugmentLoader{}
}

// Load implements RegLoadable
func (l AugmentLoader) Load(registry chat.Registry) error {
	// Make keyword struct
	reply, err := chat.NewKeyword("reply",
		[]string{"thread", "channel", "pm"},
		true)
	if err != nil {
		return fmt.Errorf("error creating keyword: %s", err.Error())
	}

	// Add
	if err = registry.AddAugment(reply); err != nil {
		return fmt.Errorf("error adding reply keyword to registry: %s", err.Error())
	}

	return nil
}
