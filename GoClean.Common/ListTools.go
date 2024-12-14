package GoClean_Common

import "sort"

type Slice[T any] []T

func (s *Slice[T]) Add(value T) {
	*s = append(*s, value)
}

func (s *Slice[T]) AddRange(values []T) {
	*s = append(*s, values...)
}

func (s Slice[T]) Where(test func(T) bool) Slice[T] {
	var result Slice[T]
	for _, v := range s {
		if test(v) {
			result = append(result, v)
		}
	}
	return result
}
func (s Slice[T]) Delete(test func(T) bool) Slice[T] {
	var result Slice[T]
	for index, v := range s {
		if test(v) {
			result = append(s[:index], s[index+1:]...)
		}
	}
	return result
}

func (s Slice[T]) FirstOrDefault(test ...func(T) bool) *T {

	if len(test) == 0 {
		if len(s) > 0 {
			return &s[0]
		}
		return nil
	}
	for _, v := range s {
		if test[0](v) {
			return &v
		}
	}
	return nil
}
func (s Slice[T]) Any(test ...func(T) bool) bool {

	if len(test) == 0 {
		if len(s) > 0 {
			return false
		}
		return false
	}
	for _, v := range s {
		if test[0](v) {
			return true
		}
	}
	return false
}
func (s Slice[T]) LastOrDefault(test ...func(T) bool) T {
	var zero T // مقدار پیش‌فرض
	if len(test) == 0 {
		if len(s) > 0 {
			return s[len(s)-1]
		}
		return zero
	}
	for i := len(s) - 1; i >= 0; i-- {
		if test[0](s[i]) {
			return s[i]
		}
	}
	return zero
}

func (s *Slice[T]) Sort(less func(a, b T) bool) {
	sort.Slice(*s, func(i, j int) bool {
		return less((*s)[i], (*s)[j])
	})
}

func (s Slice[T]) Contains(value T, equals func(a, b T) bool) bool {
	for _, v := range s {
		if equals(v, value) {
			return true
		}
	}
	return false
}

func (s Slice[T]) Count() int {
	return len(s)
}

// تبدیل لیست به آرایه
func (s Slice[T]) ToArray() []T {
	return s
}

func SelectMany[T, U any](input Slice[T], selector func(T) Slice[U]) Slice[U] {
	var result Slice[U]
	for _, item := range input {
		innerList := selector(item)
		result = append(result, innerList...)
	}
	if result == nil {
		result = Slice[U]{}
	}
	return result
}

func ListSelect[T any, U any](input Slice[T], selector func(T) U) Slice[U] {
	var result Slice[U]
	for _, item := range input {
		mappedItem := selector(item)
		result = append(result, mappedItem)
	}
	return result
}
