package tasks

import "reflect"

func EqualSlices(sl1, sl2 []int) bool {
	isEqual := reflect.DeepEqual(sl1, sl2)
	return isEqual
}
