// intsToString类似于fmt.Sprintf()，但添加了逗号
package main

import (
	"bytes"
	"fmt"
)

func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')

	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)

		// TODO 这里处理比我写得好，我这样写还要处理最后一个逗号
		// buf.WriteString(strconv.Itoa(v))
		// buf.WriteString(",")
	}

	buf.WriteByte(']')
	return buf.String()
}

func main() {
	v := []int{1, 2, 3, 4, 5}
	fmt.Println(intsToString(v)) // [1, 2, 3, 4, 5]
	fmt.Println(v)               // [1 2 3 4 5]
}
