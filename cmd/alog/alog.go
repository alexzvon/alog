package alog

import (
	"github.com.alexzvon.alog/internal/helper"
	"github.com.alexzvon.alog/internal/logger"
	"strconv"
)

/*
n - количество горутин
r - записей
in - канал
*/
func Run(n, r int, in <-chan string) {
	log := logger.New()

	for i := 0; i < n; i++ {
		n := i
		go func() {
			for i := 0; i < r; i++ {
				l, ok := <-in
				if !ok {
					return
				}

				if err := log.Info(helper.ConCat("поток: ", strconv.Itoa(n), " ", l)); err != nil {
					return
				}
			}
		}()
	}
}

/*
log - список сообщений
out - канал
*/
func Log(l []string, out chan<- string) {
	for _, s := range l {
		out <- s
	}
}
