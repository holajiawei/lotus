package validation

// This is incorrect. Its purpose is to illustrate my intent to see if there's a
// clean and simple way to do this (and even if it's worth doing it). `Errorf`,
//  * https://pkg.go.dev/golang.org/x/xerrors?tab=doc#Errorf
// allows to wrap a child error into a new parent error (using OOP language here
// is already wrong). In this way one error can have many parents, what I would
// like is a hierarchy of errors where one parent has many possible child errors
// according to its category. The objective is to programmatically  preserve in
// the language information liking errors from different contexts.
// To achieve this, instead of using `xerrors.wrapError` and keep a reference of
// the child, I define a new struct that references the parent. To keep `xerrors.Is()`
// semantics valid `Unwrap()` now returns the parent error navigating the chain upwards.
// FIXME: This is general enough to be outside of this package.
type HierarchicalError struct {
	parent error
	msg    string
}

func (e *HierarchicalError) Error() string {
	// return fmt.Sprint(e)
	return e.msg
	// FIXME: Using `fmt.Sprint()` overflows the stack, probably because
	//  I'm using references in the opposite direction.
}

func (e *HierarchicalError) Unwrap() error {
	return e.parent
}

func ErrorWrapString(parent error, msg string) error {
	return &HierarchicalError{
		parent: parent,
		msg:    msg,
	}
}

func ErrorWrapError(parent error, err error) error {
	return &HierarchicalError{
		parent: parent,
		msg:    err.Error(),
	}
}
