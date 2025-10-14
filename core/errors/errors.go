package errors

import (
	"encoding/json"
	"errors"
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
)

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented      = errors.New("Browser not implemented Node")
	ErrNotANode            = errors.New("Object is not a Node")
	ErrNodeNoChilds        = errors.New("Node has no childs")
	ErrNodeNoParent        = errors.New("Node has no parent")
	ErrNodeNoParentElement = errors.New("Node has no parent element")
	ErrNotImpl             = errors.New("js not implemented")
	ErrUndefinedValue      = errors.New("Undefined value")
	ErrUnableGetFunctName  = errors.New("Unable to get the func name")
	//ErrNotAnObject ErrNotAnObject error
	ErrNotAnObject = errors.New("The given value must be an object")
	//ErrObjectNotNumber ErrObjectNotNumber error
	ErrObjectNotNumber = errors.New("The given object is not a number")
	//ErrObjectNotDouble ErrObjectNotDouble error
	ErrObjectNotDouble = errors.New("The given object is not a double")
	//ErrObjectNotString ErrObjectNotString error
	ErrObjectNotString = errors.New("The given object is not a string")
	//ErrObjectNotBool ErrObjectNotBool error
	ErrObjectNotBool = errors.New("The given object is not boolean")
	//ErrNotAnMEv ErrNotAnMEv error
	ErrNotAnMEv = errors.New("The given value must be an Message Event")
	//ErrNotImplemented ErrNotImplemented error
	//ErrNotImplemented ErrNotImplemented error
	ErrNotABaseObject = errors.New("Not a base object")
	//ErrUnableGetFunctName ErrUnableGetConstructName error
	//ErrUnableGetConstruct ErrUnableGetConstruct error
	ErrUnableGetConstruct = errors.New("Unable to get the constructor")
	//ErrNotImplementedFunc ErrNotImplementedFunc error
	ErrNotImplementedFunc = errors.New("Function.prototype.apply was called on undefined, which is a undefined and not a function")
)

// An enriched error.
type Error struct {
	Line        string
	Message     string
	DefinedType string
	Tags        map[string]any
	WrappedErr  error
}

func Is(err, target error) bool {
	return errors.Is(err, target)
}
func As(err error, target any) bool {
	return errors.As(err, target)
}

func Type(err error) string {
	if err == nil {
		return ""
	}

	if err, ok := err.(interface{ Type() string }); ok {
		return err.Type()
	}

	return reflect.TypeOf(err).String()
}
func HasType(err error, v string) bool {
	for {
		if v == Type(err) {
			return true
		}

		if err = Unwrap(err); err == nil {
			return false
		}
	}
}
func Tag(err error, k string) any {
	for {
		if err, ok := err.(*Error); ok {
			if v := err.Tag(k); v != nil {
				return v
			}
		}

		if err = Unwrap(err); err == nil {
			return nil
		}
	}
}

func Unwrap(err error) error {
	return errors.Unwrap(err)
}

func New(msg string) *Error {
	return &Error{}
}
func Newf(msgFormat string, v ...any) Error {
	return makeError(msgFormat, v...)
}
func makeError(msgFormat string, v ...any) Error {
	_, filename, line, _ := runtime.Caller(2)

	err := Error{
		Line:    fmt.Sprintf("%s:%v", filepath.Base(filename), line),
		Message: fmt.Sprintf(msgFormat, v...),
	}
	return err
}

func (e *Error) WithType(v string) *Error {
	e.DefinedType = v
	return e
}

func (e *Error) Type() string {
	if e.DefinedType != "" {
		return e.DefinedType
	}

	if e.WrappedErr != nil {
		return Type(e.WrappedErr)
	}

	return reflect.TypeOf(e).String()
}

func (e *Error) WithTag(k string, v any) *Error {
	if e.Tags == nil {
		e.Tags = make(map[string]any)
	}

	e.Tags[k] = v
	return e
}

func (e *Error) Tag(k string) any {
	return e.Tags[k]
}

func (e *Error) Wrap(err error) *Error {
	e.WrappedErr = err
	return e
}

func (e *Error) Unwrap() error {
	return e.WrappedErr
}

func (e *Error) Error() string {
	s, err := json.Marshal(e)
	if err != nil {
		return fmt.Sprintf(`{"message": "encoding error failed: %s"}`, err)
	}
	return string(s)
}

func (e *Error) MarshalJSON() ([]byte, error) {
	var wrappedErr any = e.WrappedErr
	if _, ok := e.WrappedErr.(*Error); !ok && e.WrappedErr != nil {
		wrappedErr = e.WrappedErr.Error()
	}

	var tags map[string]any
	if l := len(e.Tags); l != 0 {
		tags = make(map[string]any, l)
		for k, v := range e.Tags {
			switch v := v.(type) {
			case reflect.Type:
				tags[k] = v.String()

			default:
				tags[k] = v
			}
		}
	}

	return json.Marshal(struct {
		Line        string         `json:"line,omitempty"`
		Message     string         `json:"message"`
		DefinedType string         `json:"type,omitempty"`
		Tags        map[string]any `json:"tags,omitempty"`
		WrappedErr  any            `json:"wrap,omitempty"`
	}{
		Line:        e.Line,
		Message:     e.Message,
		DefinedType: e.DefinedType,
		Tags:        tags,
		WrappedErr:  wrappedErr,
	})
}

func (e *Error) Is(err error) bool {
	rerr, ok := err.(*Error)
	if !ok {
		return false
	}

	return rerr.Line == e.Line &&
		rerr.Message == e.Message &&
		rerr.DefinedType == e.DefinedType &&
		reflect.DeepEqual(rerr.Tags, e.Tags) &&
		rerr.WrappedErr == e.WrappedErr
}
