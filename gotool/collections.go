package gotool

// 移除重复元素
func RemoveDuplicates(slice []string) []string {
	seen := make(map[string]struct{}, len(slice))
	index := 0
	for _, v := range slice {
		if _, ok := seen[v]; ok {
			continue
		}
		seen[v] = struct{}{}
		slice[index] = v
		index++
	}
	return slice[:index]
}
