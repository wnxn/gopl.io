package mysql

import (
	"fmt"
)

func SqlQuote1(x interface{}) string {
	if x == nil {
		return "NULL"
	} else if _, ok := x.(int); ok {
		return fmt.Sprintf("%d", x)
	} else if _, ok := x.(uint); ok {
		return fmt.Sprintf("%d", x)
	} else if b, ok := x.(bool); ok {
		if b {
			return "TRUE"
		}
		return "FALSE"
	} else if s, ok := x.(string); ok {
		return s
	} else {
		panic(fmt.Sprintf("unexpected type %T: %v", x, x))
	}
}

func SqlQuote2(x interface{}) string {
	fmt.Printf("%T\n", x)
	switch x := x.(type) {
	case nil:
		fmt.Printf("%T\n", x)
		return "NULL"
	case int, uint:
		fmt.Printf("%T\n", x)
		return fmt.Sprintf("%d", x)
	case bool:
		fmt.Printf("%T\n", x)
		if x {
			return "TRUE"
		}
		return "FALSE"
	case string:
		fmt.Printf("%T\n", x)
		return x
	default:
		panic(fmt.Sprintf("unexpected type %T: %v", x, x))
	}
}
