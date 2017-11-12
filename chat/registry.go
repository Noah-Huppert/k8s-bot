package chat

import (
	"errors"
	"fmt"
	"log"
	"os"
)

// Registry provides an interface for storing and querying command and augment
// definitions.
type Registry interface {
	// AddCmd stores the provided command definition. Returns nil on success,
	// or an error on failure.
	AddCmd(cmd *Cmd) error

	// AddAugment stores the provided augment definition. Returns nil on
	// success, or an error on failure.
	AddAugment(aug *Keyword) error

	// Cmd returns the command definition with the provided name. Nil if not
	// found
	Cmd(name string) *Cmd

	// Augment returns the augment definition containing the provided value.
	// Nil if not found
	Augment(name string) *Keyword
}

// DefaultRegistry implements Registry to store command and augment definitions.
type DefaultRegistry struct {
	// cmds holds all the registered commands
	cmds map[string]*Cmd

	// augments holds all valid argumentation that can be understood by
	// the bot. Keys are augment names. Values are Keywords.
	augments map[string]*Keyword

	// logger is used to display messages
	logger *log.Logger
}

// NewDefaultRegistry returns a new empty DefaultRegistry
func NewDefaultRegistry() DefaultRegistry {
	var reg DefaultRegistry

	reg.cmds = make(map[string]*Cmd)
	reg.augments = make(map[string]*Keyword)

	reg.logger = log.New(os.Stdout, "registry: ", 0)

	return reg
}

// AddCmd registers a command definition with the registry. Returning an error
// if one occurs, nil otherwise.
func (r DefaultRegistry) AddCmd(cmd *Cmd) error {
	// Check name not empty
	if len(cmd.Name) == 0 {
		return errors.New("command name empty")
	}

	// Check if exists
	if _, ok := r.cmds[cmd.Name]; ok {
		return fmt.Errorf("command with name \"%s\" already registered", cmd.Name)
	}

	// Add
	r.cmds[cmd.Name] = cmd

	return nil
}

// AddAugment registers an augmentation with the registry. Returning an error
// if one occurs, nil otherwise
func (r DefaultRegistry) AddAugment(aug *Keyword) error {
	// Check augment values are not empty, and not stomping on any existing
	// augments
	for i, val := range aug.Values {
		// Check not empty
		if len(val) == 0 {
			return fmt.Errorf("augment value at index %d is empty, can not be", i)
		}

		// Check value not already registered
		if keyword, ok := r.augments[val]; ok {
			return fmt.Errorf("augment with value \"%s\" already registered (.Values index = %d) registered to \"%s\" augment", val, i, keyword.Name)
		}
	}

	// If all values pass checks, add
	for _, val := range aug.Values {
		r.augments[val] = aug
	}

	return nil
}

// Augment implements the Augment Registry method
func (r DefaultRegistry) Augment(value string) *Keyword {
	if val, ok := r.augments[value]; ok {
		return val
	}

	return nil
}

// Cmd implements the Cmd Registry method
func (r DefaultRegistry) Cmd(name string) *Cmd {
	if val, ok := r.cmds[name]; ok {
		return val
	}

	return nil
}

// parseAugments parses augments out of a string of text, for internal use only.
// An error is returned if one occurs, nil if successful.
/*
func (r Registry) parseAugments(cmdReq CmdReq, words []string) (map[string]string, error) {
	augments := make(map[string]string)

	for len(words) > 1 {
		first := words[0]
		keyword, aOk := r.Augments[first]

		// Check augment was found
		if !aOk {
			// Skip item if not command
			if _, cmdOk := r.Cmds[first]; !cmdOk {
				words = words[1:]
				continue
			} else {
				break
			}
		}

		// Check if augment already provided
		if val, ok := cmdReq.Augments[keyword.Name]; ok {
			return augments, fmt.Errorf("augment already provided with value \"%s\"", val)
		}

		// Add to augmentations list
		cmdReq.Augments[keyword.Name] = first

		// Pop first item off words
		words = words[1:]
		r.logger.Println(len(words))
	}

	return augments, nil
}
*/
