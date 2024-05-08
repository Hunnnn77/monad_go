package src

type Result[T any] struct {
	*Ok[T]
	*Err
}
type Ok[T any] struct {
	Value T
}
type Err struct {
	E error
}

func (r Result[T]) Unwrap() T {
	if r.Ok == nil || r.IsErr() {
		panic(r.E)
	}
	return r.Ok.Value
}
func (r Result[T]) UnwrapOr(f func(e error) T) T {
	if r.IsOk() {
		return r.Value
	}
	return f(r.Err.E)
}
func (r Result[T]) IsOk() bool {
	return r.Ok != nil
}
func (r Result[T]) IsErr() bool {
	return r.Err != nil
}
func (r Result[T]) Match(f func(value T) *T, ef func(e error) *T) *T {
	if r.IsOk() {
		return f(r.Value)
	}
	return ef(r.E)
}

type Option[T any] struct {
	*Some[T]
	*None
}
type Some[T any] struct {
	Value T
}
type None struct{}

func (o Option[T]) Unwrap() T {
	if o.Some == nil || o.IsNone() {
		panic("none")
	}
	return o.Some.Value
}
func (o Option[T]) UnwrapOr(f func() T) T {
	if o.IsSome() {
		return o.Value
	}
	return f()
}
func (o Option[T]) IsSome() bool {
	return o.Some != nil
}
func (o Option[T]) IsNone() bool {
	return o.None != nil
}
func (o Option[T]) Match(f func(value T) *T, ef func() *T) *T {
	if o.IsSome() {
		return f(o.Value)
	}
	return ef()
}
