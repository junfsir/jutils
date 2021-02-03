package main

import (
	"context"
	"fmt"
	"github.com/jfsir/jutils/client"
)

func main()  {
	ctx := context.Background()
	client.Cli.ContainerList()
}
