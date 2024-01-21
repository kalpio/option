package option

import "reflect"

type Option[T any] struct {
	none bool
	some T
}

func Some[T any](value T) Option[T] {
	if isNil(value) {
		return Option[T]{none: true}
	}
	return Option[T]{some: value}
}

func None[T any]() Option[T] {
	return Option[T]{none: true}
}

func (o Option[T]) IsSome() bool {
	return !o.none
}

func (o Option[T]) IsNone() bool {
	return o.none
}

func (o Option[T]) Unwrap() T {
	if o.none {
		panic("`Unwrap` called on `None` value")
	}
	return o.some
}

func (o Option[T]) UnwrapOr(def T) T {
	if o.none {
		return def
	}
	return o.some
}

func (o Option[T]) UnwrapOrElse(f func() T) T {
	if o.none {
		return f()
	}
	return o.some
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
