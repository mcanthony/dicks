package main

import (
	"fmt"
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
	{options{"8", "=", "D", 10, 1, args{}}, 1, "8=D"},
	{options{"8", "=", "D", 10, 1, args{}}, 2, "8==D"},
	{options{"8", "=", "D", 10, 1, args{}}, 5, "8=====D"},
	{options{"8", "=", "D- - -", 10, 1, args{}}, 8, "8========D- - -"},
	{options{">", "- ", ">-", 10, 1, args{}}, 5, ">- - - - - >-"},
}

func TestMakeDick(t *testing.T) {
	for _, pair := range makeDickTests {
		v := makeDick(pair.opts, pair.length)
		if v != pair.output {
			t.Error("\n",
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
	{options{"8", "=", "D", 10, 1, args{}}, 5},
	{options{"8", "=", "D", 10, 1, args{0}}, 5},
	{options{"8", "=", "D", 10, 1, args{2}}, 2},
	{options{"8", "=", "D", 10, 1, args{42}}, 42},
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

type randomIntWithinTestPair struct {
	input  []int
	output []int
}

var randomIntWithinTests = []randomIntWithinTestPair{
	{[]int{1, 10}, []int{1, 10}},
	{[]int{10, 100}, []int{10, 100}},
	{[]int{1, 1}, []int{1, 1}},
	{[]int{20, 10}, []int{10, 10}},
	{[]int{20, 0}, []int{0, 0}},
}

func TestRandomIntWithin(t *testing.T) {
	for _, pair := range randomIntWithinTests {
		for i := 0; i < 100; i++ {
			v := randomIntWithin(pair.input[0], pair.input[1])
			if v < pair.output[0] || v > pair.output[1] {
				t.Error("\n",
					"For:",
					fmt.Sprintf("min=%d, max=%d",
						pair.input[0], pair.input[1],
					), "\n",
					fmt.Sprintf("Expected int within: %d-%d",
						pair.output[0], pair.output[1],
					), "\n",
					"Got:", v,
				)
			}
		}
	}
}
