package ex19

func ex19() int {
	defer func() {
		recover()
	}()
	panic(1)
}
