package nn

import "fmt"

type TTT struct {
}

func (*TTT) T(args ...interface{}) {
	fmt.Println(args)
}
