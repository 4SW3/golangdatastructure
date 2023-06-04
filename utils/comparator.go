package utils

// Comparator will make type assertion and panic if it fails.
// It will return -1 if a < b, 0 if a == b, 1 if a > b
type Comparator func(a, b interface{}) int

func StringComparator(a, b interface{}) int {
	aStr, ok := a.(string)
	if !ok {
		panic("Type error: expected a string")
	}
	bStr, ok := b.(string)
	if !ok {
		panic("Type error: expected a string")
	}

	min := len(bStr)
	if len(aStr) < min {
		min = len(aStr)
	}

	diff := 0
	for i := 0; i < min && diff == 0; i++ {
		diff = int(aStr[i]) - int(bStr[i])
	}

	if diff == 0 {
		diff = len(aStr) - len(bStr)
	}

	if diff < 0 {
		return -1
	}

	if diff > 0 {
		return 1
	}

	return 0
}

func IntComparator(a, b interface{}) int {
	aInt, ok := a.(int)
	if !ok {
		panic("Type error: expected an integer")
	}
	bInt, ok := b.(int)
	if !ok {
		panic("Type error: expected an integer")
	}
	switch {
	case aInt < bInt:
		return -1
	case aInt > bInt:
		return 1
	default:
		return 0
	}
}
