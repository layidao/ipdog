package ipdog

import (
	// "fmt"
	"strconv"
	"strings"
)

// 获取IP所属省份的sn
func GetProvinceSn(ip string) int64 {
	ipint := ip2Int(ip)
	index := searchIndex(ipint)
	return snSlice[index]
}

// 获取IP所属省份 code
func GetProvinceCode(ip string) string {
	ipint := ip2Int(ip)
	index := searchIndex(ipint)
	sn := snSlice[index]
	return codeSlice[sn]
}

// 查找给定值所在数组中的位置
func searchIndex(value int64) int {
	if value <= 0 {
		return 0
	}
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

// IP 地址转为正数
func ip2Int(ip string) int64 {
	bits := strings.Split(ip, ".")
	if len(bits) != 4 {
		return 0
	}

	b0, _ := strconv.Atoi(bits[0])
	b1, _ := strconv.Atoi(bits[1])
	b2, _ := strconv.Atoi(bits[2])
	b3, _ := strconv.Atoi(bits[3])

	var sum int64

	sum += int64(b0) << 24
	sum += int64(b1) << 16
	sum += int64(b2) << 8
	sum += int64(b3)

	return sum
}
