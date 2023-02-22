package main

import (
	"fmt"
	"strings"
	"sync"
)

func Input(inp [3]chan<- string) {
	input := [][]string{
		{"aabbb", "ebep", "blablablaa", "hijk", "wswww"},
		{"abba", "eeeppp", "cocor", "ppppppaa", "qwerty", "acasq"},
		{"lalala", "lalal", "papapa", "papap"},
	}

	for i := range inp {
		go func(in chan<- string, word []string) {
			for _, wrd := range word {
				in <- wrd
			}
			close(in)
		}(inp[i], input[i])
	}
}

func Map(in <-chan string, out chan<- map[string]int) {
	count := map[string]int{}
	vow := []string{"a", "e", "i", "o", "u"}
	for word := range in {
		tmp := word
		nr_vow := 0
		nr_con := 0
		for _, i := range tmp {
			for _, j := range vow {
				if strings.Contains(string(i), j) {
					nr_vow++
				}
			}
			nr_con = len(tmp) - nr_vow
		}
		if nr_vow%2 == 0 {
			if nr_con%3 == 0 {
				count["vow_con"] = count["vow_con"] + 1
			}
		}
	}
	out <- count
	close(out)
}

func Shuffle(in []<-chan map[string]int, out chan<- int) {
	var wg sync.WaitGroup
	wg.Add(len(in))
	for _, ch := range in {
		go func(c <-chan map[string]int) {
			for m := range c {
				nvc, vc := m["vow_con"]
				if vc {
					out <- nvc
				}
			}
			wg.Done()
		}(ch)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
}

func Reduce(in <-chan int, out chan<- float64) {
	sum, count := 0, 0
	for n := range in {
		sum = sum + n
		count++
	}
	out <- float64(sum) / float64(count)
	close(out)
}

func main() {
	input1 := make(chan string)
	input2 := make(chan string)
	input3 := make(chan string)
	map1 := make(chan map[string]int)
	map2 := make(chan map[string]int)
	map3 := make(chan map[string]int)
	reduce := make(chan int)
	avg := make(chan float64)
	go Input([3]chan<- string{input1, input2, input3})
	go Map(input1, map1)
	go Map(input2, map2)
	go Map(input3, map3)
	go Shuffle([]<-chan map[string]int{map1, map2, map3}, chan<- int(reduce))
	go Reduce(reduce, avg)
	fmt.Printf("Numar mediu: %f\n", <-avg)
}
