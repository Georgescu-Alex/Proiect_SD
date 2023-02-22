package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

var count = 0

func dup_first(num int) int {
	var aux []int
	for num != 0 {
		aux = append(aux, num%10)
		num = num / 10
	}

	aux = append(aux, aux[len(aux)-1])
	num = 0
	num = aux[len(aux)-1] * 10

	i := 2
	for i <= len(aux) {
		num = num + aux[len(aux)-i]
		if i != len(aux) {
			num = num * 10
		}
		i++
	}
	return num
}

func dup_numb(num []int) []int {
	i := len(num) - 1
	for i >= 0 {
		num[i] = dup_first(num[i])
		i = i - 1
	}
	return num
}

func sum_numb(num []int) int {
	sum := 0
	for _, s := range num {
		sum = sum + s
	}
	return sum
}

func sum_cmp(num int) int {
	nr := 0
	for num > 0 {
		aux := num % 10
		nr = nr + aux
		num = num / 10
	}
	return nr
}

func sum_4(numbers []int, sum int, a int, b int) int {
	i := 0
	for ind, s := range numbers {
		if sum_cmp(numbers[ind]) >= a && sum_cmp(numbers[ind]) <= b {
			sum = sum + s
			i++
		}
	}
	sum = sum / i
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

func scrmb_cuv(cuvinte []string, nr_cuv int, l_cuv int) []string {
	l_aux := make([]string, l_cuv)
	caux := make([]string, nr_cuv)
	caux = nil
	cfin := make([]string, nr_cuv)

	for i := 0; i < nr_cuv-1; i++ {
		caux = append(caux, cuvinte...)
		for j := 0; j < l_cuv; j++ {
			l_aux[j] = strings.Replace(caux[i], string(caux[i][j]), string(caux[j][i]), 1)
			caux[i] = l_aux[j]
			cfin[i] = l_aux[j]
		}
		cfin[i] = cfin[i] + string(caux[nr_cuv-1][i])
		caux = nil
	}

	for i := 0; i < nr_cuv; i++ {
		cuvinte[i] = cfin[i]
	}
	return cuvinte
}

func read_numb(f_name string, numbers []int) []int {
	file, err := os.Open(f_name)

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err == nil {
			numbers = append(numbers, i)
		}
	}
	return numbers
}

func read_cuv_1(f string, cuvinte []string) ([]string, int, int) {
	file, err := os.Open(f)

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	var aux_cuv string
	for scanner.Scan() {
		aux_cuv = scanner.Text()
		cuvinte = append(cuvinte, aux_cuv)
	}

	var nr_cuv int
	var l_cuv int
	nr_cuv = len(cuvinte)
	l_cuv = len(cuvinte[0])

	return cuvinte, nr_cuv, l_cuv
}

func prob_12(f string) string {
	var numbers []int
	sum := 0

	numbers = read_numb(f, numbers)
	sum = sum_numb(numbers)

	numbers = dup_numb(numbers)
	sum = sum_numb(numbers)

	rs := fmt.Sprint("numbers: ", numbers, ", sum: ", sum)
	return rs
}

func prob_4(f string) string {
	var numbers []int
	sum := 0
	a := 0
	b := 1

	numbers = read_numb(f, numbers)
	sum = sum_4(numbers[b:len(numbers)], sum, numbers[a], numbers[b])

	rs := fmt.Sprint("numbers: ", numbers, ", sum: ", sum)
	return rs
}

func prob_3(f string) string {
	var numbers []int
	sum := 0

	numbers = read_numb(f, numbers)
	sum = sum_numb(numbers)

	numbers = inv_numbers(numbers)
	sum = sum_numb(numbers)

	rs := fmt.Sprint("numbers: ", numbers, ", sum: ", sum)
	return rs
}

func prob_1(f string) string {
	var cuvinte []string
	var nr_cuv int
	var l_cuv int

	cuvinte, nr_cuv, l_cuv = read_cuv_1(f, cuvinte)
	cuvinte = scrmb_cuv(cuvinte, nr_cuv, l_cuv)

	cuv := strings.Join(cuvinte, " ")
	return cuv
}

func handleConnection(c net.Conn) {
	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		tmp := strings.TrimSpace(string(netData))
		if tmp == "exit" {
			break
		}
		counter := strconv.Itoa(count)
		var rsp string
		fmt.Println("Client " + counter + " Conectat")
		fmt.Println("Client " + counter + " a facut request cu datele: " + tmp)
		c.Write([]byte(string("Server a primit requestul.\t")))
		switch tmp {
		case "12.txt":
			rsp = prob_12("12.txt")
		case "3.txt":
			rsp = prob_3("3.txt")
		case "4.txt":
			rsp = prob_4("4.txt")
		case "1.txt":
			rsp = prob_1("1.txt")
		}
		c.Write([]byte(string("Server proceseaza datele.\t")))
		fmt.Println("Server trimite " + rsp + " catre client.")

		//c.Write([]byte(string(counter)))
		rsp = rsp + "\n"
		c.Write([]byte(string(rsp)))
		fmt.Println("Client " + counter + " a primit raspunsul: " + rsp)
	}
	c.Close()
}

func main() {
	args := os.Args
	if len(args) == 1 || len(args) > 2 {
		fmt.Println("Usage: go run .\\server.go (Port)")
		return
	}

	Port := ":" + args[1]
	l, err := net.Listen("tcp4", Port)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConnection(c)
		count++
	}
}
