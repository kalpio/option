package option

import "testing"

func TestOption_IsNone(t *testing.T) {
	none := None[int]()
	if !none.IsNone() {
		t.Errorf("Expected `IsNone` to return true")
	}
	if none.IsSome() {
		t.Errorf("Expected `IsSome` to return false")
	}
}

func TestOption_IsSome(t *testing.T) {
	some := Some[int](42)
	if !some.IsSome() {
		t.Errorf("Expected `IsSome` to return true")
	}
	if some.IsNone() {
		t.Errorf("Expected `IsNone` to return false")
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
	none := None[int]()
	if none.UnwrapOr(21) != 21 {
		t.Errorf("Expected `UnwrapOr` to return 21")
	}
}

func TestOption_UnwrapOrElse(t *testing.T) {
	some := Some[int](42)
	if some.UnwrapOrElse(func() int { return 0 }) != 42 {
		t.Errorf("Expected `UnwrapOrElse` to return 42")
	}
	none := None[int]()
	if none.UnwrapOrElse(func() int { return 21 }) != 21 {
		t.Errorf("Expected `UnwrapOrElse` to return 21")
	}
}

type testStruct struct {
	value int
}

func TestOption_IsNone_Struct(t *testing.T) {
	none := None[testStruct]()
	if !none.IsNone() {
		t.Errorf("Expected `IsNone` to return true")
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
	if some.IsNone() {
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
	none := None[testStruct]()
	if none.UnwrapOr(testStruct{21}).value != 21 {
		t.Errorf("Expected `UnwrapOr` to return 21")
	}
}

func TestOption_UnwrapOrElse_Struct(t *testing.T) {
	some := Some[testStruct](testStruct{42})
	if some.UnwrapOrElse(func() testStruct { return testStruct{0} }).value != 42 {
		t.Errorf("Expected `UnwrapOrElse` to return 42")
	}
	none := None[testStruct]()
	if none.UnwrapOrElse(func() testStruct { return testStruct{21} }).value != 21 {
		t.Errorf("Expected `UnwrapOrElse` to return 21")
	}
}
