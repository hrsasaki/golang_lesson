package ex16

func Join(sep string, strs ...string) string {
	if len(strs) == 0 {
		return ""
	} else if len(strs) == 1 {
		return strs[0]
	}
	return strs[0] + sep + Join(sep, strs[1:]...)
}
