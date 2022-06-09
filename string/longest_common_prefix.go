package main

func longestCommonPrefix(strs []string) string {
	i := 0
	common := ""
	for {
		cl := ""
		for _, v := range strs {
			if len(v) <= i {
				return common
			} else if cl == "" {
				cl = string(v[i])
			} else if cl != string(v[i]) {
				return common
			}
		}
		common += cl
		i++
	}
}
