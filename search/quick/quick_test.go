// Copyright 2017 Koichi Shiraishi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package quick

import (
	"testing"
)

func TestSearch(t *testing.T) {
	type args struct {
		P string
		m int
		T string
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example",
			args: args{
				P: "GCAGAGAG",
				m: 8,
				T: "GCATCGCAGAGAGTATACAGTACG",
				n: 24,
			},
			want: 5,
		},
		{
			name: "same",
			args: args{
				P: "GCAGAGAG",
				m: 8,
				T: "GCAGAGAG",
				n: 8,
			},
			want: 0,
		},
		{
			name: "before",
			args: args{
				P: "GCAGAGAG",
				m: 8,
				T: "GCAGAGAGABC",
				n: 11,
			},
			want: 0,
		},
		{
			name: "after",
			args: args{
				P: "GCAGAGAG",
				m: 8,
				T: "ABCGCAGAGAG",
				n: 11,
			},
			want: 3,
		},
		{
			name: "not match",
			args: args{
				P: "GCAGAGAG",
				m: 8,
				T: "test",
				n: 4,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Search(tt.args.P, tt.args.m, tt.args.T, tt.args.n); got != tt.want {
				t.Errorf("Search(%v, %v, %v, %v) = %v, want %v", tt.args.P, tt.args.m, tt.args.T, tt.args.n, got, tt.want)
			}
		})
	}
}
