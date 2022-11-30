package gog

import "golang.org/x/exp/constraints"

//https://github.com/icza/gog

// If returns vtrue is cond is true, vfalse otherwise.
//
// Useful to avoid an if statement when initializing variables, for example:
//
//	min := If(i > 0, i, 0)
func If[T any](cond bool, vtrue, vfalse T) T {
	if cond {
		return vtrue
	}
	return vfalse
}

// Ptr returns a pointer to the passed value.
//
// Useful when you have a value and need a pointer, e.g.:
//
//	func f() string { return "foo" }
//
//	foo := struct{
//	    Bar *string
//	}{
//	    Bar: Ptr(f()),
//	}
func Ptr[T any](v T) *T {
	return &v
}

// DePtr returns a dereference variable value.
// For example:
//
//	foo := struct{
//	    Bar *string
//	}{}
//
//	bar := DePtr[string](foo.Bar)
func DePtr[T any](v *T) (res T) {
	if v != nil {
		res = *v
	}
	return
}

// Must takes 2 arguments, the second being an error.
// If err is not nil, Must panics. Else the first argument is returned.
//
// Useful when inputs to some function are provided in the source code,
// and you are sure they are valid (if not, it's OK to panic).
// For example:
//
//	t := Must(time.Parse("2006-01-02", "2022-04-20"))
func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

// First returns the first argument.
// Useful when you want to use the first result of a function call that has more than one return values
// (e.g. in a composite literal or in a condition).
//
// For example:
//
//	func f() (i, j, k int, s string, f float64) { return }
//
//	p := image.Point{
//	    X: First(f()),
//	}
func First[T any](first T, _ ...any) T {
	return first
}

// Second returns the second argument.
// Useful when you want to use the second result of a function call that has more than one return values
// (e.g. in a composite literal or in a condition).
//
// For example:
//
//	func f() (i, j, k int, s string, f float64) { return }
//
//	p := image.Point{
//	    X: Second(f()),
//	}
func Second[T any](_ any, second T, _ ...any) T {
	return second
}

// Third returns the third argument.
// Useful when you want to use the third result of a function call that has more than one return values
// (e.g. in a composite literal or in a condition).
//
// For example:
//
//	func f() (i, j, k int, s string, f float64) { return }
//
//	p := image.Point{
//	    X: Third(f()),
//	}
func Third[T any](_, _ any, third T, _ ...any) T {
	return third
}

// ByDefault return default argument
// For example:
//
//	ByDefault(default, optional...)
func ByDefault[T any](first T, optional ...T) T {
	defFirst := first

	if len(optional) >= 1 {
		defFirst = optional[0]
	}

	return defFirst
}

func Min[T constraints.Ordered](args ...T) T {
	min := args[0]
	for _, x := range args {
		if x < min {
			min = x
		}
	}
	return min
}

func Max[T constraints.Ordered](args ...T) T {
	max := args[0]
	for _, x := range args {
		if x > max {
			max = x
		}
	}
	return max
}

func Reverse[S ~[]E, E any](slice S) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}
