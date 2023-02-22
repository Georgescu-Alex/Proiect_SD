package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var num_tmp = [10][]int{
	{11, 12, 13},
	{12, 13, 14},
	{13, 14, 15},
	{14, 15, 16},
	{15, 16, 17},
	{16, 17, 18},
	{17, 18, 19},
	{18, 19, 20},
	{19, 20, 21},
	{20, 21, 22},
}

var rez [numProcesses]string

func sum_numb(num []int) int {
	sum := 0
	for _, s := range num {
		sum = sum + s
	}
	return sum
}

func inv_nr(num int) int {
	nr := 0
	for num > 0 {
		aux := num % 10
		nr = (nr * 10) + aux
		num = num / 10
	}
	return nr
}

func inv_numbers(numbers []int) []int {
	i := len(numbers) - 1
	for i >= 0 {
		numbers[i] = inv_nr(numbers[i])
		i = i - 1
	}
	return numbers
}

const numProcesses = 10
const numNames = 10

var names = [numNames]bool{}
var mux sync.Mutex

func chooseName(i int, initialName int) int {
	mux.Lock()
	defer mux.Unlock()

	if !names[initialName] {
		names[initialName] = true
		return initialName
	}

	for {
		name := rand.Intn(numNames)
		if !names[name] {
			names[name] = true
			return name
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var wg sync.WaitGroup
	wg.Add(numProcesses)

	for i := 0; i < numProcesses; i++ {
		initialName := rand.Intn(numNames)
		go func(i int, initialName int) {
			defer wg.Done()

			name := chooseName(i, initialName)
			fmt.Printf("Procesul %d cu numele %d alege un nou nume %d | date - %d\n", i, initialName, name, num_tmp[i])

			sum := sum_numb(inv_numbers(num_tmp[i]))
			rez[i] = fmt.Sprintf("Rezultatul lui %d: numere - %d, suma - %d\n", name, num_tmp[i], sum)
		}(i, initialName)
	}

	wg.Wait()

	fmt.Println("\n**********************\n")
	for k := 0; k < len(rez); k++ {
		fmt.Print(rez[k])
	}
}
