package cmptest

import (
	"github.com/anchore/syft/syft/pkg"
	"github.com/google/go-cmp/cmp"
)

type CopyrightComparer func(x, y pkg.Copyright) bool

func DefaultCopyrightComparer(x, y pkg.Copyright) bool {
	return cmp.Equal(x, y, cmp.Comparer(
		func(x, y string) bool {
			return x == y
		},
	))
}
