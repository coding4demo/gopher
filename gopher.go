package main

import (
	"fmt"

	"github.com/coding4demo/private-repo/private"
	"github.com/coding4demo/public-repo/public"
)

func main() {

	fmt.Println(public.Message())

	fmt.Println(private.Message())
}
