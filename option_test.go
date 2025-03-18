package option

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
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
