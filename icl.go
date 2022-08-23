package icl

import "fmt"

type Integers []int

func (i Integers) GetCollection() []int {
	return i
}

type Collection interface {
	GetCollection() []int
}

type CollectionAdapter func() []int

func (ca CollectionAdapter) GetCollection() []int {
	return ca()
}

type Decorator func(Collection) Collection

func Decorate(cl Collection, ds ...Decorator) Collection {
	decorated := cl
	for _, decorator := range ds {
		decorated = decorator(decorated)
	}
	return decorated
}

func PrintIntegers(cl Collection) {
	fmt.Printf("%v\n", cl.GetCollection())
}
