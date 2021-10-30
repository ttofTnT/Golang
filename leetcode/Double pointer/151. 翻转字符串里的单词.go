package Double_pointer
import "strings"

// Fields 以连续的空白字符为分隔符，将 s 切分成多个子串，结果中不包含空白字符本身
// 空白字符有：\t, \n, \v, \f, \r, ' ', U+0085 (NEL), U+00A0 (NBSP)
// 如果 s 中只包含空白字符，则返回一个空列表

func reverseWords(s string) string {
	ss := strings.Fields(s)
	reverse151(&ss, 0, len(ss)-1)
	return strings.Join(ss, " ")
}
func reverse151(m *[]string, i int, j int) {
	for i <= j {
		(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
		i++
		j--
	}
}

