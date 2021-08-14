package either

import "github.com/dynastywind/go-commons/optional"

// A type containing either left or right element
// Note that only left or right should appear in this struct and the other should be nil
// When trying to access the wrapped data, one MUST call HasLeft or HasRight first to test if desired data exists or not
type Either struct {
	left  interface{}
	right interface{}
}

func OfLeft(data interface{}) Either {
	if data == nil {
		panic("Data should not be nil")
	}
	return Either{
		left:  data,
		right: nil,
	}
}

func OfRight(data interface{}) Either {
	if data == nil {
		panic("Data should not be nil")
	}
	return Either{
		left:  nil,
		right: data,
	}
}

func (e Either) HasLeft() bool {
	return e.left != nil
}

func (e Either) HasRight() bool {
	return e.right != nil
}

func (e Either) GetLeft() interface{} {
	if e.HasLeft() {
		return e.left
	}
	panic("Left should not be nil")
}

func (e Either) GetRight() interface{} {
	if e.HasRight() {
		return e.right
	}
	panic("Right should not be nil")
}

func (e Either) MapLeft(mapper func(interface{}) interface{}) optional.Optional {
	return optional.OfNillable(e.left).Map(mapper)
}

func (e Either) MapRight(mapper func(interface{}) interface{}) optional.Optional {
	return optional.OfNillable(e.right).Map(mapper)
}

func (e Either) IfLeftPresent(consumer func(interface{})) {
	if e.HasLeft() {
		consumer(e.GetLeft())
	}
}

func (e Either) IfRightPresent(consumer func(interface{})) {
	if e.HasRight() {
		consumer(e.GetRight())
	}
}
