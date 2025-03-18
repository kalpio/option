// Package option provides a generic Option type implementation for Go,
// similar to Rust's Option enum, allowing for explicit handling of optional values.
package option

import (
  "errors"
  "reflect"
)

var (
  ErrNilValue = errors.New("option: value cannot be nil")
)

// Option represents a value that may or may not be present.
type Option[T any] struct {
  none bool
  some T
  err  error
}

func Some[T any](value T) Option[T] {
  if isNil(value) {
    return Option[T]{none: true, err: ErrNilValue}
  }
  return Option[T]{some: value}
}

// None creates a new Option in the None state with the provided error.
// This represents the absence of a value, with an optional error explaining why.
//
// Example:
//
//	// Create a None option with an error
//	opt := option.None[string](errors.New("value not available"))
//
//	// Create a None option without an error
// None creates a new Option in the None state with the provided error.
// This represents the absence of a value, with an optional error explaining why.
//
// Example:
//
//	// Create a None option with an error
//	opt := option.None[string](errors.New("value not available"))
//
//	// Create a None option without an error
//	opt := option.None[int](nil)
//
// The type parameter T specifies what type the Option would contain if it were Some.
// When using None, you must explicitly specify the type parameter since it cannot be inferred.
func None[T any](err error) Option[T] {
  return Option[T]{none: true, err: err}
}

// IsSome returns true if the Option is in the Some state (contains a value),
// and false if it is in the None state.
//
// Example:
//
//	opt := option.Some(42)
//	if opt.IsSome() {
//		// Handle the case where a value is present
//	}
func (o Option[T]) IsSome() bool {
  return !o.none
}

// IsNone returns true if the Option is in the None state (contains no value),
// and false if it is in the Some state.
//
// Example:
//
//	opt := option.None[int](errors.New("no value"))
//	if opt.IsNone() {
//		// Handle the case where no value is present
//		err := opt.Error()
//		// Process the error...
//	}
func (o Option[T]) IsNone() bool {
  return o.none
}

// Error returns the error associated with a None option, or nil if the option is Some.
// This method is useful for checking why a None value was created.
//
// Example:
//
//	opt := option.None[string](errors.New("data not found"))
//	if err := opt.Error(); err != nil {
//		fmt.Printf("Error: %v\n", err)
//	}
//
// For Some options, this method always returns nil.
func (o Option[T]) Error() error {
  if o.none {
    return o.err
  }

  return nil
}

// Unwrap extracts and returns the contained value if the Option is Some.
// Panics if the Option is None.
//
// Example:
//
//	opt := option.Some(42)
//	value := opt.Unwrap() // Returns 42
//
//	// This will panic:
//	opt := option.None[int](errors.New("no value"))
//	value := opt.Unwrap()
//
// It's recommended to check IsSome() before calling Unwrap to avoid panics.
func (o Option[T]) Unwrap() T {
  if o.none {
    panic("`Unwrap` called on `None` value")
  }
  return o.some
}

// UnwrapOr returns the contained value if the Option is Some,
// otherwise returns the provided default value.
//
// Example:
//
//	opt := option.Some(42)
//	value := opt.UnwrapOr(0) // Returns 42
//
//	opt := option.None[int](errors.New("no value"))
//	value := opt.UnwrapOr(0) // Returns 0
//
// This method provides a safe way to extract a value without risking a panic.
func (o Option[T]) UnwrapOr(def T) T {
  if o.none {
    return def
  }
  return o.some
}

// UnwrapOrElse returns the contained value if the Option is Some,
// otherwise calls the provided function and returns its result.
//
// Example:
//
//	opt := option.Some(42)
//	value := opt.UnwrapOrElse(func() int { return computeDefault() }) // Returns 42
//
//	opt := option.None[int](errors.New("no value"))
//	value := opt.UnwrapOrElse(func() int { return computeDefault() }) // Returns result of computeDefault()
//
// This is useful when the default value is expensive to compute or needs to be determined dynamically.
func (o Option[T]) UnwrapOrElse(f func() T) T {
  if o.none {
    return f()
  }
  return o.some
}

// Filter returns None if the option is None or if the predicate returns false when applied to the contained value.
// Otherwise returns the original option.
//
// Example:
//
//	// Keep only positive numbers
//	opt := option.Some(42)
//	filtered := opt.Filter(func(n int) bool { return n > 0 }) // Still Some(42)
//
//	opt := option.Some(-3)
//	filtered := opt.Filter(func(n int) bool { return n > 0 }) // None with error
//
//	// None values remain None
//	opt := option.None[int](errors.New("no value"))
//	filtered := opt.Filter(func(n int) bool { return n > 0 }) // Still None with original error
//
// This method is useful for conditionally processing values based on their properties.
func (o Option[T]) Filter(predicate func(T) bool) Option[T] {
  if o.none || !predicate(o.some) {
    var err error
    if o.none {
      err = o.err
    } else {
      err = errors.New("option: value did not satisfy predicate")
    }
    return None[T](err)
  }
  return o
}

// Map transforms the contained value using the provided function if the option is Some.
// Returns None if the option is None.
//
// Example:
//
//	// Convert an integer to a string
//	opt := option.Some(42)
//	mapped := Map(opt, func(n int) string { return strconv.Itoa(n) }) // Some("42")
//
//	// None values pass through unchanged
//	opt := option.None[int](errors.New("no value"))
//	mapped := Map(opt, func(n int) string { return strconv.Itoa(n) }) // None with original error
//
// This function is useful for transforming values without having to manually check if they exist.
// The type parameters T and U represent the input and output types of the transformation.
func Map[T, U any](o Option[T], f func(T) U) Option[U] {
  if o.none {
    return None[U](o.err)
  }
  return Some(f(o.some))
}

// FlatMap transforms the contained value using the provided function if the option is Some.
// The function must return an Option. Returns None if the original option is None.
//
// Example:
//
//	// Parse a string to an integer, which might fail
//	opt := option.Some("42")
//	result := FlatMap(opt, func(s string) Option[int] {
//		n, err := strconv.Atoi(s)
//		if err != nil {
//			return None[int](err)
//		}
//		return Some(n)
//	}) // Some(42)
//
//	// With invalid input
//	opt := option.Some("not a number")
//	result := FlatMap(opt, func(s string) Option[int] {
//		n, err := strconv.Atoi(s)
//		if err != nil {
//			return None[int](err)
//		}
//		return Some(n)
//	}) // None with parsing error
//
//	// None values pass through unchanged
//	opt := option.None[string](errors.New("no value"))
//	result := FlatMap(opt, func(s string) Option[int] {
//		// This function is never called
//		return Some(0)
//	}) // None with original error
//
// This function is useful for chaining operations that might fail, similar to monadic bind operations.
// The type parameters T and U represent the input and output types of the transformation.
func FlatMap[T, U any](o Option[T], f func(T) Option[U]) Option[U] {
  if o.none {
    return None[U](o.err)
  }
  return f(o.some)
}

func isNil[T any](value T) bool {
  v := reflect.ValueOf(value)
  switch v.Kind() {
  case reflect.Chan,
    reflect.Func,
    reflect.Map,
    reflect.Ptr,
    reflect.UnsafePointer,
    reflect.Interface,
    reflect.Slice:
    return v.IsNil()
  default:
    return false
  }
}
