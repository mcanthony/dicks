package main

import (
	"math/rand"
	"regexp"
	"testing"
	"time"
)

type makeDickTestPair struct {
	opts   options
	length int
	output string
}

var makeDickTests = []makeDickTestPair{
	{options{"8", "=", "D", args{}}, 1, "8=D"},
	{options{"8", "=", "D", args{}}, 2, "8==D"},
	{options{"8", "=", "D", args{}}, 5, "8=====D"},
	{options{"8", "=", "D- - -", args{}}, 8, "8========D- - -"},
	{options{">", "- ", ">-", args{}}, 5, ">- - - - - >-"},
}

func TestMakeDick(t *testing.T) {
	for _, pair := range makeDickTests {
		v := makeDick(pair.opts, pair.length)
		if v != pair.output {
			t.Error(
				"For", pair.opts, ", ", pair.length, "\n",
				"expected", pair.output, "\n",
				"got", v,
			)
		}
	}
}

type makeDickStreamTestPair struct {
	opts  options
	count int
}

var makeDickStreamTests = []makeDickStreamTestPair{
	{options{"8", "=", "D", args{}}, 5},
	{options{"8", "=", "D", args{0}}, 5},
	{options{"8", "=", "D", args{2}}, 2},
	{options{"8", "=", "D", args{42}}, 42},
}

func TestMakeDickStream(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
	dickRegexp, _ := regexp.Compile("8[=]+D")

	for _, pair := range makeDickStreamTests {
		stream := makeDickStream(pair.opts)

		for i := 1; i <= pair.count+1; i++ {
			dick, ok := <-stream

			if ok && !dickRegexp.MatchString(dick) {
				t.Error("Invalid dick format:", dick)
			}

			if i <= pair.count && !ok {
				t.Error("channel closed early")
			} else if i == pair.count+1 && ok {
				t.Error("channel closed late")
			}
		}
	}
}
