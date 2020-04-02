package hpeg

import "math/rand"

func PickAndDelete(list *[]string) string {
	if len(*list) < 1 {
		return "ERROR"
	}
	lst := *list
	idx := rand.Intn(len(lst))
	picked := lst[idx]
	*list = append(lst[:idx], lst[idx+1:]...)
	return picked
}
