package problem0258

// 这道题的标签是数学，题目要求在O(1)时间复杂度内解决，显然就是找规律
// 数字	各位相加
// 1	1
// 2	2
// 3	3
// 4	4
// 5	5
// 6	6
// 7	7
// 8	8
// 9	9
// 10	1
// 11	2
// 12	3
// 13	4
// 14	5
// 15	6
// 16	7
// 17	8
// 18	9
// 19	1
// 20	2
// 21	3
// 22	4
// 23	5
// 24	6
// 可读性最好的答案
func addDigits(num int) int {
	if num == 0 {
		return 0
	} else if num%9 == 0 {
		return 9
	} else {
		return num % 9
	}
}
