// Package radamsa implements a Go interface to libradamsa.
package radamsa

// Copyright © 2019 Ramsey Dow. All rights reserved.
// SPDX-License-Identifier: MIT

// #cgo CFLAGS: -Icache
// #cgo LDFLAGS: -Lcache -lradamsa
// #include <radamsa.h>
import "C"

import (
	"errors"
	"unsafe"
)

// Radamsa contains the state for a given fuzzer instance.
type Radamsa struct {
	seed      int64
	inPlace   bool
	iteration uint64
}

// New creates a new Radamsa instance. It supports configuration
// by passing in any number of Option types.
func New(opts ...Option) *Radamsa {
	r := &Radamsa{}
	for _, opt := range opts {
		opt(r)
	}
	C.radamsa_init()
	return r
}

// Option is a type that operates against a receiver that is a Rasamsa pointer.
type Option func(*Radamsa)

// WithSeed allows you to specify the PRNG seed value used by Radamsa.
func WithSeed(seed int64) Option {
	return func(r *Radamsa) {
		r.seed = seed
	}
}

// InPlace allows you to specify that Radamsa should fuzz its input in place.
func InPlace(inPlace bool) Option {
	return func(r *Radamsa) {
		r.inPlace = inPlace
	}
}

// Seed returns the currently configured seed value.
func (r *Radamsa) Seed() int64 {
	return r.seed
}

// Iteration returns the current iteration number.
func (r *Radamsa) Iteration() uint64 {
	return r.iteration
}

// Fuzz will fuzz the input byte slice, which has the specified length.
// The result will he written to the output buffer, which has the specified
// capacity.
func (r *Radamsa) Fuzz(input []byte, length int, output []byte, capacity int) (uint, error) {
	if length == 0 {
		return 0, errors.New("input buffer has no data")
	}
	if capacity == 0 {
		return 0, errors.New("output buffer has 0 capacity")
	}

	r.iteration++
	// FIXME: The inplace cgo interop is currently broken.
	// if r.inPlace {
	// 	return fuzz(input, length, input, capacity, r.seed), nil
	// }
	return fuzz(input, length, output, capacity, r.seed), nil
}

// If the below code looks uncomfortably weird, don't worry. We are attempting
// to bridge the gap between the C types that the libradamsa API expects and
// the types employed by idiomatic Go. len(xb) returns an int, not the Go 
// equivalent of a size_t. Likewise, the default seed value acquired from
// UnixNano() is an int64, not an unsigned int as expected by libradamsa. By
// performing these type conversions interstitially, the caller is free to
// deal in native Go types without having to give consideration to the
// underlying C types.
func fuzz(input []byte, length int, output []byte, capacity int, seed int64) uint {
	n := C.radamsa(
		(*C.uint8_t)(unsafe.Pointer(&input[0])),
		C.size_t(uint(length)),
		(*C.uint8_t)(unsafe.Pointer(&output[0])),
		C.size_t(uint(capacity)),
		C.uint(uint(seed)))
	return uint(n)
}
