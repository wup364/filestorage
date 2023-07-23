package repository

// GetPlaceholder 生成占位符 '?', ',', 2 => '?,?'
func GetPlaceholder(placeholder, joinstr string, len int) (res string) {
	for i := 0; i < len; i++ {
		if i == 0 {
			res = placeholder
		} else {
			res += joinstr + placeholder
		}
	}
	return res
}

// ConvertStrArray2AnyArray ConvertStrArray2AnyArray
func ConvertStrArray2AnyArray(in []string) (res []any) {
	res = make([]any, len(in))
	for i := 0; i < len(in); i++ {
		res[i] = in[i]
	}
	return
}
