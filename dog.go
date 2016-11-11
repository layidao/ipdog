package ipdog

func GetProvinceCode(ipint int) string {
	index := searchIndex(ipint)
	sn := snSlice[index]
	return codeSlice[sn]
}

// 查找给定值所在数组中的位置
func searchIndex(value int) int {
	left, right := 0, len(ips)-1

	for left <= right {
		mid := left + (right-left)/2

		if ips[mid] < value {
			left = mid + 1
		} else if ips[mid] > value {
			right = mid - 1
		} else {
			return mid
		}
	}
	return left - 1
}
