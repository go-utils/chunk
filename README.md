# What is this?
Obtains an index of specified size length and returns it.  

# Usage

https://go.dev/play/p/KzSjIei9_vL

```go
package main

import (
	"fmt"

	"github.com/go-utils/chunk"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	chunkedArr := make([][]int, 0)
	// chunk size is 3
	for idx := range chunk.GetIndex(len(arr), 3) {
		chunkedArr = append(chunkedArr, arr[idx.From:idx.To])
	}

	fmt.Printf("%#v\n", chunkedArr)
	// Output: [][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}, []int{10}}
}
```

## License
- Under the [MIT License](./LICENSE)
- Copyright (C) 2022 go-utils
