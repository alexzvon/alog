package alog

import (
	"flag"
	"github.com.alexzvon.alog/internal/helper"
	"strconv"
	"testing"
)

var cn = flag.Int("cn", 5, "Количество потоков")
var cr = flag.Int("cr", 50, "Количество записей")

func TestLog(t *testing.T) {
	n := *cn
	r := *cr

	c := n * r

	log := make([]string, c)

	for i := 0; i < c; i++ {
		log[i] = helper.ConCat("Message ", strconv.Itoa(i))
	}

	t.Run("Test Log", func(t *testing.T) {
		inout := make(chan string)

		Run(n, r, inout)
		Log(log, inout)

		close(inout)
	})
}
