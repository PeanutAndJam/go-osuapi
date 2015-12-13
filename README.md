# go-osuapi

[![docs](https://godoc.org/github.com/thehowl/go-osuapi?status.svg)](https://godoc.org/github.com/thehowl/go-osuapi)

go-osuapi is an osu! API library for Golang.

## Get started

```go
package main

import (
	"fmt"
	"github.com/thehowl/go-osuapi"
)

func main() {
	c := osuapi.NewClient("Your API key")
	user, err := c.GetUser("peppy", osuapi.Standard)
	if err != nil {
		panic(err)
	}
	fmt.Printf("User %s is rank #%d on osu! standard.", user.Username, user.Rank)
}
```
