package tasks

// Отсортируйте элементы слайса целых чисел. Так чтобы первое число заняло последнее место и наоборот, второе число – предпоследнее и наоборот. (3,2,1,6,7) –&gt; (7,6,1,2,3)

func SortSlice(s []int) []int {

	newSlice := []int{}

	for i := len(s) - 1; i > -1; i-- {
		newSlice = append(newSlice, s[i])
	}
	return newSlice
}
