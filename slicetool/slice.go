package slicetool

// Del 切片删除指元素,使用移位法删除
func Del[T comparable](s *[]T, val T) {
	temp := *s
	if len(temp) == 0 {
		return
	}
	var i int
	for key, v := range temp {
		if v != val {
			temp[i] = temp[key]
			i++
		}
	}
	*s = temp[:i]
}

// DelByIndex 切片删除指定索引元素,使用移位法删除
func DelByIndex[T any](s *[]T, index int) {
	temp := *s
	count := len(temp)
	if count == 0 || index > count {
		return
	}
	count--
	var i int = index
	for index < count {
		index++
		temp[i] = temp[index]
		i++
	}
	*s = temp[:i]
}
