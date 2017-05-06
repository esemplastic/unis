// Copyright 2017 Γεράσιμος Μαρόπουλος. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package unis

import (
	"strings"
)

// NewPrefixRemover accepts a "prefix" and returns a new processor
// which returns the result without that "prefix".
func NewPrefixRemover(prefix string) ProcessorFunc {
	return func(original string) (result string) {
		if strings.HasPrefix(original, prefix) {
			result = original[len(prefix)-1:]
		}
		return
	}
}

// NewPrepender accepts a "prefix" and returns a new processor
// which returns the result prepended with that "prefix".
func NewPrepender(prefix string) ProcessorFunc {
	return func(original string) (result string) {
		result = prefix + original
		return
	}
}

// NewPrependerIfNotExists accepts an "expectedIndex" as int
// and a "prefixChar" as byte and returns a new processor
// which returns the result prepended with that "prefixChar"
// if the "original" string[expectedIndex] != prefixChar.
func NewPrependerIfNotExists(expectedIndex int, prefixChar byte) ProcessorFunc {
	return func(original string) (result string) {
		if expectedIndex < len(original)-1 {
			if original[expectedIndex] != prefixChar {
				return NewPrepender(string(prefixChar)).Process(original)
			}
		}

		return original
	}
}
