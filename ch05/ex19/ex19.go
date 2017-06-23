package ex19

func ex19() (val int) {
	defer func() {
		recover()
		val = 1
	}()
	panic(1)
}
