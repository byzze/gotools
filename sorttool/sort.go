package sorttool

import "sort"

type OrderType int

const (
	Asc OrderType = iota
	Desc
)

// 根据版本号降序排序版本列表
func SortByVersion(list []string, order OrderType) {
	sort.Slice(list, func(i, j int) bool {
		iList := processStringData(list[i])
		jList := processStringData(list[j])
		iLen := len(iList)
		jLen := len(jList)
		maxLen := iLen
		if maxLen < jLen {
			maxLen = jLen
		}
		startIndex := 0 // 处理版本号长度不一致, v1.1.1 < v1.1.1.0
		for startIndex < maxLen {
			var x, y byte
			if startIndex < iLen {
				x = iList[startIndex]
			}
			if startIndex < jLen {
				y = jList[startIndex]
			}
			if x == y {
				startIndex++
				continue
			}
			if order == Asc {
				return x < y // 升序为x < y， 降序为x > y
			}
			return x > y
		}
		return false
	})
}

// 创建一个字节数组，遍历将所有的字符放入，当字符是数字类型是，需要计算总数，再将其放入数组中，放入时机为数组末尾或遇到非数字类型的字符
func processStringData(value string) []byte {
	var byteArr []byte
	var num int
	var factor int = 1
	vLen := len(value)
	for i := 0; i < vLen; i++ {
		tmpNum := value[i] - '0'
		flag := tmpNum >= 0 && tmpNum <= 9
		if flag {
			num = int(tmpNum) + num*factor
			factor = factor * 10
			if i+1 >= vLen {
				byteArr = append(byteArr, byte(num))
				num = 0
				factor = 1
			}
			continue
		} else if num >= 0 {
			byteArr = append(byteArr, byte(num))
			num = 0
			factor = 1
		}
		byteArr = append(byteArr, value[i])
	}
	return byteArr
}
