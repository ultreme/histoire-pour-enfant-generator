package hpeg

import "math/rand"

func PickAndDelete(list *[]string) string {
	lst := *list
	idx := rand.Intn(len(lst))
	picked := lst[idx]
	*list = append(lst[:idx], lst[idx+1:]...)
	return picked
}
