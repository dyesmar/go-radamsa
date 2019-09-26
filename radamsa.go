// Package radamsa implements a Go interface to libradamsa.
package radamsa

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

// New creates a new Radamsa instance. it supports configuration
// by passing in any number of Option types.
func New(opts ...Option) *Radamsa {
	r := &Radamsa{}
	for _, opt := range opts {
		opt(r)
	}
	C.init()
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

func fuzz(input []byte, length int, output []byte, capacity int, seed int64) uint {
	n := C.radamsa(
		(*C.uchar)(unsafe.Pointer(&input[0])),  // ptr
		C.ulong(uint(length)),                  // len
		(*C.uchar)(unsafe.Pointer(&output[0])), // target
		C.ulong(uint(capacity)),                // max
		C.uint(uint(seed)))                     // seed
	return uint(n)
}
