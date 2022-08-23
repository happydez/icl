package icl

func OnlyEvens() Decorator {
	return func(c Collection) Collection {
		return CollectionAdapter(func() []int {
			integers := c.GetCollection()
			evenIntegers := make([]int, 0, len(integers))
			for _, value := range integers {
				if value%2 == 0 {
					evenIntegers = append(evenIntegers, value)
				}
			}
			return evenIntegers
		})
	}
}

func OnlyOdds() Decorator {
	return func(c Collection) Collection {
		return CollectionAdapter(func() []int {
			integers := c.GetCollection()
			evenIntegers := make([]int, 0, len(integers))
			for _, value := range integers {
				if value%2 != 0 {
					evenIntegers = append(evenIntegers, value)
				}
			}
			return evenIntegers
		})
	}
}

func OnlyPositives(withZero bool) Decorator {
	return func(c Collection) Collection {
		return CollectionAdapter(func() []int {
			integers := c.GetCollection()
			posIntegers := make([]int, 0, len(integers))
			for _, vl := range integers {
				if vl > 0 {
					posIntegers = append(posIntegers, vl)
				} else if vl == 0 && withZero {
					posIntegers = append(posIntegers, 0)
				}
			}
			return posIntegers
		})
	}
}

func Double() Decorator {
	return func(c Collection) Collection {
		return CollectionAdapter(func() []int {
			integers := c.GetCollection()
			for i := range integers {
				integers[i] *= 2
			}
			return integers
		})
	}
}
