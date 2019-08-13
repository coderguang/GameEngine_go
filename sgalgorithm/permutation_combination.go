package sgalgorithm

//resultList must be a slice
//从srcList中,选出num个做全排列
func GenPermutation(srcList []string, num int, resultList *[]string) {
	if num <= 0 {
		return
	}
	if num > len(srcList) {
		num = len(srcList)
	}
	flags := make([]int, num, num*2)
	for _, n := range flags {
		flags[n] = 0
	}
	genPermutationItem(srcList, flags, len(srcList), resultList)
	return
}

func genPermutationItem(srcList []string, flags []int, srcLength int, resultList *[]string) {
	str := ""
	for _, n := range flags {
		str += srcList[n]
	}
	*resultList = append(*resultList, str)
	if flags[len(flags)-1] < (srcLength - 1) { //变动最后一位
		flags[len(flags)-1]++
		genPermutationItem(srcList, flags, srcLength, resultList)
	} else { //最后一位不能变动时，往前查找可以变动的位置
		flags[len(flags)-1] = 0
		if len(flags)-2 < 0 {
			return
		}
		for i := len(flags) - 2; i >= 0; i-- {
			if flags[i] < srcLength-1 {
				flags[i]++
				genPermutationItem(srcList, flags, srcLength, resultList)
				return
			} else {
				flags[i] = 0
			}
		}
		return
	}
}
