package option

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestOption_IsNone(t *testing.T) {
	none := None[int](errors.New("some error"))
	assert.True(t, none.IsNone())
	assert.False(t, none.IsSome())
}

func TestOption_Error(t *testing.T) {
	expectedErr := errors.New("some error")
	none := None[int](expectedErr)

	assert.True(t, none.IsNone())
	assert.ErrorIs(t, none.Error(), expectedErr)
}

func TestOption_IsSome(t *testing.T) {
	some := Some[int](42)
	assert.True(t, some.IsSome())
	assert.False(t, some.IsNone())
}

func TestOption_IsSome_ErrorIsNil(t *testing.T) {
	some := Some[int](42)
	assert.True(t, some.IsSome())
	assert.Nil(t, some.Error())
}

func TestOption_Unwrap(t *testing.T) {
	some := Some[int](42)
	assert.Equal(t, 42, some.Unwrap())
}

func TestOption_UnwrapOr(t *testing.T) {
	some := Some[int](42)
	assert.Equal(t, 42, some.UnwrapOr(0))

	none := None[int](errors.New("some error"))
	assert.Equal(t, 21, none.UnwrapOr(21))
}

func TestOption_UnwrapOrElse(t *testing.T) {
	some := Some[int](42)
	assert.Equal(t, 42, some.UnwrapOrElse(func() int { return 0 }))

	none := None[int](errors.New("some error"))
	assert.Equal(t, 21, none.UnwrapOrElse(func() int { return 21 }))
}

type testStruct struct {
	value int
}

func TestOption_IsNone_Struct(t *testing.T) {
	expectedError := errors.New("testStruct: error")
	none := None[testStruct](expectedError)

	assert.True(t, none.IsNone())
	assert.ErrorIs(t, none.Error(), expectedError)
	assert.False(t, none.IsSome())
}

func TestOption_IsSome_Struct(t *testing.T) {
	some := Some[testStruct](testStruct{42})
	assert.True(t, some.IsSome())
	assert.False(t, some.IsNone())
}

func TestOption_Unwrap_Struct(t *testing.T) {
	some := Some[testStruct](testStruct{42})
	assert.Equal(t, testStruct{42}, some.Unwrap())
}

func TestOption_UnwrapOr_Struct(t *testing.T) {
	some := Some[testStruct](testStruct{42})
	assert.Equal(t, testStruct{42}, some.UnwrapOr(testStruct{0}))

	none := None[testStruct](errors.New("some error"))
	assert.Equal(t, testStruct{21}, none.UnwrapOr(testStruct{21}))
}

func TestOption_UnwrapOrElse_Struct(t *testing.T) {
	some := Some[testStruct](testStruct{42})
	assert.Equal(t, testStruct{42}, some.UnwrapOrElse(func() testStruct { return testStruct{0} }))

	none := None[testStruct](errors.New("some error"))
	assert.Equal(t, testStruct{21}, none.UnwrapOrElse(func() testStruct { return testStruct{21} }))
}

func TestOption_Nil(t *testing.T) {
	none := Some[*testStruct](nil)
	assert.True(t, none.IsNone())
	assert.False(t, none.IsSome())
}

func TestOption_Nil_Unwrap(t *testing.T) {
	none := Some[*testStruct](nil)
	assert.Panics(t, func() {
		none.Unwrap()
	})
}

func TestOption_Nil_UnwrapOr(t *testing.T) {
	none := Some[*testStruct](nil)
	assert.Equal(t, &testStruct{42}, none.UnwrapOr(&testStruct{42}))
}

func TestOption_Nil_UnwrapOrElse(t *testing.T) {
	none := Some[*testStruct](nil)
	assert.Equal(t, &testStruct{42}, none.UnwrapOrElse(func() *testStruct { return &testStruct{42} }))
}

func BenchmarkOption_IsNone(b *testing.B) {
	for i := 0; i < b.N; i++ {
		none := None[int](errors.New("some err"))
		_ = none.IsNone()
	}
}

func BenchmarkOption_IsSome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		some := Some[int](42)
		some.IsSome()
	}
}

func BenchmarkOption_Unwrap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		some := Some[int](42)
		some.Unwrap()
	}
}

func BenchmarkOption_UnwrapOr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		some := Some[int](42)
		some.UnwrapOr(0)
	}
}

func BenchmarkOption_UnwrapOrElse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		some := Some[int](42)
		some.UnwrapOrElse(func() int { return 0 })
	}
}

func BenchmarkOption_IsNone_Struct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		none := None[testStruct](errors.New("some error"))
		_ = none.IsNone()
	}
}

func BenchmarkOption_IsSome_Struct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		some := Some[testStruct](testStruct{42})
		some.IsSome()
	}
}

func BenchmarkOption_Unwrap_Struct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		some := Some[testStruct](testStruct{42})
		some.Unwrap()
	}
}

func BenchmarkOption_UnwrapOr_Struct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		some := Some[testStruct](testStruct{42})
		some.UnwrapOr(testStruct{0})
	}
}

func BenchmarkOption_UnwrapOrElse_Struct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		some := Some[testStruct](testStruct{42})
		some.UnwrapOrElse(func() testStruct { return testStruct{0} })
	}
}

func BenchmarkOption_Nil(b *testing.B) {
	for i := 0; i < b.N; i++ {
		none := Some[*testStruct](nil)
		_ = none.IsNone()
	}
}

func BenchmarkOption_Nil_Unwrap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		none := Some[*testStruct](nil)
		assert.Panics(b, func() {
			none.Unwrap()
		})
	}
}

func BenchmarkOption_Nil_UnwrapOr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		none := Some[*testStruct](nil)
		none.UnwrapOr(&testStruct{42})
	}
}

func BenchmarkOption_Nil_UnwrapOrElse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		none := Some[*testStruct](nil)
		none.UnwrapOrElse(func() *testStruct { return &testStruct{42} })
	}
}

func BenchmarkOption_IsNone_Parallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		none := None[int](errors.New("some error"))
		for pb.Next() {
			_ = none.IsNone()
		}
	})
}

func BenchmarkOption_IsSome_Parallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		some := Some[int](42)
		for pb.Next() {
			some.IsSome()
		}
	})
}

func BenchmarkOption_Unwrap_Parallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		some := Some[int](42)
		for pb.Next() {
			some.Unwrap()
		}
	})
}

func BenchmarkOption_UnwrapOr_Parallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		some := Some[int](42)
		for pb.Next() {
			some.UnwrapOr(0)
		}
	})
}

func BenchmarkOption_UnwrapOrElse_Parallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		some := Some[int](42)
		for pb.Next() {
			some.UnwrapOrElse(func() int { return 0 })
		}
	})
}

func BenchmarkOption_IsNone_Struct_Parallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		none := None[testStruct](errors.New("some error"))
		for pb.Next() {
			_ = none.IsNone()
		}
	})
}

func BenchmarkOption_IsSome_Struct_Parallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		some := Some[testStruct](testStruct{42})
		for pb.Next() {
			some.IsSome()
		}
	})
}

func BenchmarkOption_Unwrap_Struct_Parallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		some := Some[testStruct](testStruct{42})
		for pb.Next() {
			some.Unwrap()
		}
	})
}

func BenchmarkOption_UnwrapOr_Struct_Parallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		some := Some[testStruct](testStruct{42})
		for pb.Next() {
			some.UnwrapOr(testStruct{0})
		}
	})
}

func BenchmarkOption_UnwrapOrElse_Struct_Parallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		some := Some[testStruct](testStruct{42})
		for pb.Next() {
			some.UnwrapOrElse(func() testStruct { return testStruct{0} })
		}
	})
}

func BenchmarkOption_Nil_Parallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		none := Some[*testStruct](nil)
		for pb.Next() {
			_ = none.IsNone()
		}
	})
}

func BenchmarkOption_Nil_Unwrap_Parallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		none := Some[*testStruct](nil)
		for pb.Next() {
			assert.Panics(b, func() {
				none.Unwrap()
			})
		}
	})
}

func BenchmarkOption_Nil_UnwrapOr_Parallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		none := Some[*testStruct](nil)
		for pb.Next() {
			none.UnwrapOr(&testStruct{42})
		}
	})
}

func BenchmarkOption_Nil_UnwrapOrElse_Parallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		none := Some[*testStruct](nil)
		for pb.Next() {
			none.UnwrapOrElse(func() *testStruct { return &testStruct{42} })
		}
	})
}

func ExampleOption_IsNone() {
	none := None[int](errors.New("some error"))
	isNone := none.IsNone()
	err := none.Error()
	fmt.Printf("IsNone: %t, Error: %q", isNone, err)
	// Output: IsNone: true, Error: "some error"
}

func ExampleOption_IsSome() {
	some := Some[int](42)
	fmt.Println(some.IsSome())
	// Output: true
}

func TestOption_Filter(t *testing.T) {
	// Test with Some value that passes the filter
	some := Some(42)
	filtered := some.Filter(func(n int) bool { return n > 0 })
	assert.True(t, filtered.IsSome())
	assert.Equal(t, 42, filtered.Unwrap())
	assert.Nil(t, filtered.Error())

	// Test with Some value that doesn't pass the filter
	some = Some(-5)
	filtered = some.Filter(func(n int) bool { return n > 0 })
	assert.True(t, filtered.IsNone())
	assert.False(t, filtered.IsSome())
	assert.NotNil(t, filtered.Error())
	assert.Equal(t, "option: value did not satisfy predicate", filtered.Error().Error())

	// Test with None value
	expectedErr := errors.New("original error")
	none := None[int](expectedErr)
	filtered = none.Filter(func(n int) bool { return n > 0 })
	assert.True(t, filtered.IsNone())
	assert.False(t, filtered.IsSome())
	assert.ErrorIs(t, filtered.Error(), expectedErr)
}

func TestMap(t *testing.T) {
	// Test with Some value
	some := Some(42)
	mapped := Map(some, func(n int) string { return fmt.Sprintf("value: %d", n) })
	assert.True(t, mapped.IsSome())
	assert.Equal(t, "value: 42", mapped.Unwrap())
	assert.Nil(t, mapped.Error())

	// Test with None value
	expectedErr := errors.New("original error")
	none := None[int](expectedErr)
	mapped = Map(none, func(n int) string { return fmt.Sprintf("value: %d", n) })
	assert.True(t, mapped.IsNone())
	assert.False(t, mapped.IsSome())
	assert.ErrorIs(t, mapped.Error(), expectedErr)

	// Test with different types
	strOpt := Some("hello")
	lenOpt := Map(strOpt, func(s string) int { return len(s) })
	assert.True(t, lenOpt.IsSome())
	assert.Equal(t, 5, lenOpt.Unwrap())
}

func TestFlatMap(t *testing.T) {
	// Test with Some value that maps to Some
	some := Some("42")
	result := FlatMap(some, func(s string) Option[int] {
		n, err := strconv.Atoi(s)
		if err != nil {
			return None[int](err)
		}
		return Some(n)
	})
	assert.True(t, result.IsSome())
	assert.Equal(t, 42, result.Unwrap())
	assert.Nil(t, result.Error())

	// Test with Some value that maps to None
	some = Some("not a number")
	result = FlatMap(some, func(s string) Option[int] {
		n, err := strconv.Atoi(s)
		if err != nil {
			return None[int](err)
		}
		return Some(n)
	})
	assert.True(t, result.IsNone())
	assert.False(t, result.IsSome())
	assert.NotNil(t, result.Error())
	assert.Contains(t, result.Error().Error(), "strconv.Atoi")

	// Test with None value
	expectedErr := errors.New("original error")
	none := None[string](expectedErr)
	result = FlatMap(none, func(s string) Option[int] {
		n, err := strconv.Atoi(s)
		if err != nil {
			return None[int](err)
		}
		return Some(n)
	})
	assert.True(t, result.IsNone())
	assert.False(t, result.IsSome())
	assert.ErrorIs(t, result.Error(), expectedErr)
}

func TestOption_Filter_Struct(t *testing.T) {
	// Test with Some value that passes the filter
	some := Some(testStruct{42})
	filtered := some.Filter(func(ts testStruct) bool { return ts.value > 0 })
	assert.True(t, filtered.IsSome())
	assert.Equal(t, testStruct{42}, filtered.Unwrap())
	assert.Nil(t, filtered.Error())

	// Test with Some value that doesn't pass the filter
	some = Some(testStruct{-5})
	filtered = some.Filter(func(ts testStruct) bool { return ts.value > 0 })
	assert.True(t, filtered.IsNone())
	assert.False(t, filtered.IsSome())
	assert.NotNil(t, filtered.Error())
}

func TestMap_Struct(t *testing.T) {
	// Test with Some value
	some := Some(testStruct{42})
	mapped := Map(some, func(ts testStruct) string {
		return fmt.Sprintf("value: %d", ts.value)
	})
	assert.True(t, mapped.IsSome())
	assert.Equal(t, "value: 42", mapped.Unwrap())
	assert.Nil(t, mapped.Error())

	// Test mapping to a different struct
	type otherStruct struct {
		text string
		num  int
	}

	mapped2 := Map(some, func(ts testStruct) otherStruct {
		return otherStruct{
			text: fmt.Sprintf("value: %d", ts.value),
			num:  ts.value * 2,
		}
	})
	assert.True(t, mapped2.IsSome())
	assert.Equal(t, otherStruct{"value: 42", 84}, mapped2.Unwrap())
}

func TestFlatMap_Struct(t *testing.T) {
	// Test with Some value that maps to Some
	some := Some(testStruct{42})
	result := FlatMap(some, func(ts testStruct) Option[string] {
		if ts.value > 0 {
			return Some(fmt.Sprintf("positive: %d", ts.value))
		}
		return None[string](errors.New("negative value"))
	})
	assert.True(t, result.IsSome())
	assert.Equal(t, "positive: 42", result.Unwrap())
	assert.Nil(t, result.Error())

	// Test with Some value that maps to None
	some = Some(testStruct{-5})
	result = FlatMap(some, func(ts testStruct) Option[string] {
		if ts.value > 0 {
			return Some(fmt.Sprintf("positive: %d", ts.value))
		}
		return None[string](errors.New("negative value"))
	})
	assert.True(t, result.IsNone())
	assert.False(t, result.IsSome())
	assert.NotNil(t, result.Error())
	assert.Equal(t, "negative value", result.Error().Error())
}
