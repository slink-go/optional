package optional

import (
	"fmt"
	"time"
)

// region - value

type Value[T any] struct {
	value   T
	isEmpty bool
}

func (o *Value[T]) Get() T {
	return o.value
}
func (o *Value[T]) Set(value T) {
	o.value = value
	o.isEmpty = false
}
func (o *Value[T]) Empty() bool {
	return o.isEmpty
}
func (o *Value[T]) OrElse(value T) T {
	if o.isEmpty {
		return value
	}
	return o.value
}
func (o *Value[T]) OrElseString(value string) string {
	if o.isEmpty {
		return value
	}
	return fmt.Sprintf("%v", o.value)
}
func (o *Value[T]) OrElseFormatted(format, value string) string {
	if o.isEmpty {
		return value
	}
	var v interface{}
	v = o.value
	switch x := v.(type) {
	case time.Time:
		return x.Format(format)
	default:
		return fmt.Sprintf(format, o.value)
	}

}

func NewFloat() Value[float64] {
	return Value[float64]{
		value:   0,
		isEmpty: true,
	}
}
func NewFloatWithValue(value float64) Value[float64] {
	return Value[float64]{
		value:   value,
		isEmpty: false,
	}
}

func NewInt() Value[int] {
	return Value[int]{
		value:   0,
		isEmpty: true,
	}
}
func NewIntWithValue(value int) Value[int] {
	return Value[int]{
		value:   value,
		isEmpty: false,
	}
}

func NewDate() Value[time.Time] {
	return Value[time.Time]{
		value:   time.Time{},
		isEmpty: true,
	}
}
func NewDateWithValue(dt time.Time) Value[time.Time] {
	return Value[time.Time]{
		value:   dt,
		isEmpty: false,
	}
}

// endregion
