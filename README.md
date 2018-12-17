# safetime
This lib is try to solve go issue
[time timer.Reset is not thread safe](https://github.com/golang/go/issues/26741)

Here is a sample code

```go
package main

import (
	"time"
)

func main() {
	tm := time.NewTimer(1 * time.Millisecond)
	for i := 0; i < 2; i++ {
		go func() {
			for {
				tm.Reset(1 * time.Millisecond)
			}
		}()
	}
	for {
		<-tm.C
	}
}
```

This code will be panic if you try to run.

panic in `go 1.11`

## LICENSE
[MIT](LICENSE)