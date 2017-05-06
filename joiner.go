// Copyright 2017 Γεράσιμος Μαρόπουλος. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package unis

// Joiner should be implemented by all string joiners.
type Joiner interface {
	// Join takes two pieces of strings
	// and returns a result of them, as one.
	Join(part1 string, part2 string) string
}
