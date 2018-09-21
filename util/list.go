package util

// List 参数列表
type List []string

func (list List) Len() int {
	return len(list)
}

func (list List) Less(i, j int) bool {
	if list[i] < list[j] {
		return true
	} else if list[i] > list[j] {
		return false
	} else {
		return true
	}
}

func (list List) Swap(i, j int) {
	temp := list[i]
	list[i] = list[j]
	list[j] = temp
}

func (list List) String() string {
	echoStr := ""
	for _, obj := range list {
		echoStr = echoStr + obj
	}
	return echoStr
}
