package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"math/rand"
	"strings"
	"time"
)

type args struct {
	Count int `description:"Number of dicks to print"`
}

type options struct {
	Balls     string `short:"b" long:"balls" default:"8" description:"Character to used to represent the balls"`
	Shaft     string `short:"s" long:"shaft" default:"=" description:"Character to used to represent the shaft"`
	Head      string `short:"H" long:"head" default:"D" description:"Character to used to represent the head"`
	Args      args   `positional-args:"yes"`
}

func makeDick(opts options, length int) string {
	return fmt.Sprintf("%s%s%s",
		opts.Balls,
		strings.Repeat(opts.Shaft, length),
		opts.Head,
	)
}

func makeDickStream(opts options) chan string {
	rand.Seed(time.Now().UTC().UnixNano())
	stream := make(chan string)

	count := 5
	if opts.Args.Count != 0 {
		count = opts.Args.Count
	}

	go func() {
		for i := 0; i < count; i++ {
			stream <- makeDick(opts, rand.Intn(10)+1)
		}
		close(stream)
	}()

	return stream
}

func main() {
	var opts options
	flags.Parse(&opts)

	dickStream := makeDickStream(opts)
	for dick := range dickStream {
		fmt.Println(dick)
	}
}
