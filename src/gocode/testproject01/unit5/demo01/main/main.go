package main

import "fmt"

// 下边是测试代码，与本节无关
//func lengthOfLongestSubstring(s string) int {
//	s1 := []int{}
//	for k, v := range s {
//		for j := k + 1; j < len(s); j++ {
//			if v == s[j] {
//				s1 = append(s1, j-k)
//			}
//
//		}
//
//	}
//	return maxInSlice(s1)
//}
//func maxInSlice(s []int) int {
//	if len(s) == 0 {
//		return 0 // 或者 panic("empty slice") 根据需求处理空切片
//	}
//	max := s[0]
//	for _, v := range s {
//		if v > max {
//			max = v
//		}
//	}
//	return max
//}

func main() {
	// 定义一个数组
	var scores [5]int
	scores[0] = 95
	scores[1] = 87
	scores[2] = 79
	scores[3] = 99
	scores[4] = 89
	sum := 0
	for i := 0; i < len(scores); i++ {
		sum += scores[i]
	}
	avg := sum / len(scores)
	fmt.Println("总成绩为", sum)
	fmt.Println("平均成绩为", avg)
}
