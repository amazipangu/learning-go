# 📝 Go Learning Journal – [2025-07-31]

## 📌 Topic / Concept
> [Tutorial: get started with Go](https://go.dev/doc/tutorial/getting-started)

---

## 🧠 What I Learned Today

- [x] install Go (Referred to [Download and install](https://go.dev/doc/install))
- [x] enable dependencies tracks executing [go mod init](https://go.dev/ref/mod#go-mod-init)
- [x] declare a main package and why it's required
- [x] import an external module using [go mod tidy](https://go.dev/ref/mod#authenticating)

### 🧭 Detail: declare a main package and why it's required

Because Go compiles code into either:

1. Executable programs (with package main + func main())
1. Reusable libraries (with any other package name)

When you run go run or go build, Go looks for:

* A file with package main
* A function named func main() inside it

```go
package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go())
}

```