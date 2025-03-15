package option

import (
	"errors"
	"fmt"
	"testing"
)

func TestOption_IsNone(t *testing.T) {
	none := None[int](errors.New("some error"))
	if isNone := none.IsNone(); !isNone {
		t.Errorf("Expected `IsNone` to return true")
	}
	if none.IsSome() {
		t.Errorf("Expected `IsSome` to return false")
	}
}

func TestOption_Error(t *testing.T) {
	expectedErr := errors.New("some error")
	none := None[int](expectedErr)
	if isNone := none.IsNone(); !isNone {
		t.Errorf("Expected `IsNone` to return true")
	}

	if errors.Is(none.Error(), expectedErr) != true {
		t.Errorf("Expected `Error()` to be `%v`", expectedErr)
	}
}

func TestOption_IsSome(t *testing.T) {
	some := Some[int](42)
	if !some.IsSome() {
		t.Errorf("Expected `IsSome` to return true")
	}
	if isNone := some.IsNone(); isNone {
		t.Errorf("Expected `IsNone` to return false")
	}
}

func TestOption_IsSome_ErrorIsNil(t *testing.T) {
	some := Some[int](42)
	if !some.IsSome() {
		t.Errorf("Expected `IsSome` to return true")
	}
	if some.Error() != nil {
		t.Errorf("Expected `Error()` to return nil")
	}
}

func TestOption_Unwrap(t *testing.T) {
	some := Some[int](42)
	if some.Unwrap() != 42 {
		t.Errorf("Expected `Unwrap` to return 42")
	}
}

func TestOption_UnwrapOr(t *testing.T) {
	some := Some[int](42)
	if some.UnwrapOr(0) != 42 {
		t.Errorf("Expected `UnwrapOr` to return 42")
	}
	none := None[int](errors.New("some error"))
	if none.UnwrapOr(21) != 21 {
		t.Errorf("Expected `UnwrapOr` to return 21")
	}
}

func TestOption_UnwrapOrElse(t *testing.T) {
	some := Some[int](42)
	if some.UnwrapOrElse(func() int { return 0 }) != 42 {
		t.Errorf("Expected `UnwrapOrElse` to return 42")
	}
	none := None[int](errors.New("some error"))
	if none.UnwrapOrElse(func() int { return 21 }) != 21 {
		t.Errorf("Expected `UnwrapOrElse` to return 21")
	}
}

type testStruct struct {
	value int
}

func TestOption_IsNone_Struct(t *testing.T) {
	expectedError := errors.New("testStruct: error")
	var isNone bool
	none := None[testStruct](expectedError)
	if isNone = none.IsNone(); !isNone {
		t.Errorf("Expected `IsNone` to return true")
	}
	if !errors.Is(none.Error(), expectedError) {
		t.Errorf("Expected `err` to be `%q`", expectedError)
	}
	if none.IsSome() {
		t.Errorf("Expected `IsSome` to return false")
	}
}

func TestOption_IsSome_Struct(t *testing.T) {
	some := Some[testStruct](testStruct{42})
	if !some.IsSome() {
		t.Errorf("Expected `IsSome` to return true")
	}
	if isNone := some.IsNone(); isNone {
		t.Errorf("Expected `IsNone` to return false")
	}
}

func TestOption_Unwrap_Struct(t *testing.T) {
	some := Some[testStruct](testStruct{42})
	if some.Unwrap().value != 42 {
		t.Errorf("Expected `Unwrap` to return 42")
	}
}

func TestOption_UnwrapOr_Struct(t *testing.T) {
	some := Some[testStruct](testStruct{42})
	if some.UnwrapOr(testStruct{0}).value != 42 {
		t.Errorf("Expected `UnwrapOr` to return 42")
	}
	none := None[testStruct](errors.New("some error"))
	if none.UnwrapOr(testStruct{21}).value != 21 {
		t.Errorf("Expected `UnwrapOr` to return 21")
	}
}

func TestOption_UnwrapOrElse_Struct(t *testing.T) {
	some := Some[testStruct](testStruct{42})
	if some.UnwrapOrElse(func() testStruct { return testStruct{0} }).value != 42 {
		t.Errorf("Expected `UnwrapOrElse` to return `testStruct { value: 42 }`")
	}
	none := None[testStruct](errors.New("some error"))
	if none.UnwrapOrElse(func() testStruct { return testStruct{21} }).value != 21 {
		t.Errorf("Expected `UnwrapOrElse` to return `testStruct { value: 21 }`")
	}
}

func TestOption_Nil(t *testing.T) {
	none := Some[*testStruct](nil)
	if isNone := none.IsNone(); !isNone {
		t.Errorf("Expected `IsNone` to return true")
	}
	if none.IsSome() {
		t.Errorf("Expected `IsSome` to return false")
	}
}

func TestOption_Nil_Unwrap(t *testing.T) {
	none := Some[*testStruct](nil)
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected `Unwrap` to panic")
		}
	}()
	none.Unwrap()
}

func TestOption_Nil_UnwrapOr(t *testing.T) {
	none := Some[*testStruct](nil)
	if none.UnwrapOr(&testStruct{42}).value != 42 {
		t.Errorf("Expected `UnwrapOr` to return `&testStruct { value: 42 }`")
	}
}

func TestOption_Nil_UnwrapOrElse(t *testing.T) {
	none := Some[*testStruct](nil)
	if none.UnwrapOrElse(func() *testStruct { return &testStruct{42} }).value != 42 {
		t.Errorf("Expected `UnwrapOrElse` to return `&testStruct { value: 42 }`")
	}
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
		func() {
			defer func() {
				if r := recover(); r == nil {
					b.Errorf("Expected `Unwrap` to panic")
				}
			}()
			none.Unwrap()
		}()
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
			func() {
				defer func() {
					if r := recover(); r == nil {
						b.Errorf("Expected `Unwrap` to panic")
					}
				}()
				none.Unwrap()
			}()
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
