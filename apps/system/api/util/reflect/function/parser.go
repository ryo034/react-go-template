package function

import "reflect"

type Returns interface {
	Error() error
	// Check Error() returns nil before invoke.
	// When wrapped function returns (int, string, error), value(0) returns int value,
	// value(1) returns string and value(2) returns nil. And their types are interface{}.
	// Casting the return value to the original type will be always successful.
	Value(index int) interface{}
}

type returns struct {
	values   []interface{}
	errIndex int
}

func ErrReturns(err error) Returns {
	return &returns{
		[]interface{}{err},
		0,
	}
}

func (r *returns) Error() error {
	result := r.values[r.errIndex]
	if result == nil {
		return nil
	}
	return result.(error)
}

func (r *returns) Value(index int) interface{} {
	return r.values[index]
}

type AnyFunc func() Returns

func Parse(fn interface{}) AnyFunc {
	t := reflect.TypeOf(fn)
	if t.Kind() != reflect.Func {
		panic("param should be a function")
	}
	errType := reflect.TypeOf((*error)(nil)).Elem()
	errIdx := -1
	for i := 0; i < t.NumOut(); i++ {
		if t.Out(i).Implements(errType) {
			errIdx = i
			break
		}
	}
	if errIdx == -1 {
		panic("function should return an error.")
	}
	return func() Returns {
		in := make([]reflect.Value, t.NumIn())
		outputs := reflect.ValueOf(fn).Call(in)
		result := make([]interface{}, t.NumOut())
		for i, output := range outputs {
			result[i] = output.Interface()
		}
		return &returns{result, errIdx}
	}
}
