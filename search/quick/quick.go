// Copyright 2017 Koichi Shiraishi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package quick implements Quick Search algorithm in Go.
package quick

// ALPHABETLEN constant alphabet length
const ALPHABETLEN = 256

func preBadCharacter(P string, m int) []int {
	qbc := make([]int, ALPHABETLEN)

	for i := 0; i < ALPHABETLEN; i++ {
		qbc[i] = m + 1
	}
	for i := 0; i < m; i++ {
		qbc[P[i]] = m - i
	}

	return qbc
}

// Search search index of P in T using Quick Search algorithm.
func Search(P string, m int, T string, n int) int {
	// Preprocessing
	qsbc := preBadCharacter(P, m)

	// Searching
	for j := 0; j <= n-m; j += qsbc[T[j+m]] {
		if P[:m] == T[j:][:m] {
			return j
		}
	}

	return 0 // not match
}
