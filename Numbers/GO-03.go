package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	//Constant
	const r = 100
	const g = 70e5 / r

	fmt.Println("#Constant---------")
	fmt.Println(g)
	fmt.Println(int64(g))
	fmt.Println(math.Sin(g))
	fmt.Println(math.Cos(g))
	fmt.Println("-------------")

	//Random Numbers
	fmt.Println("#Random Numbers")
	fmt.Print(rand.Intn(100), ",")
	fmt.Print(rand.Intn(100))
	fmt.Println()

	fmt.Println(rand.Float64())

	fmt.Print((rand.Float64()*3)+7, ",")
	fmt.Print((rand.Float64() * 7) + 5)
	fmt.Println()

	g1 := rand.NewSource(time.Now().UnixNano())
	x1 := rand.New(g1)

	fmt.Print(x1.Intn(100), ",")
	fmt.Print(x1.Intn(100))
	fmt.Println()

	g2 := rand.NewSource(42)
	x2 := rand.New(g2)
	fmt.Print(x2.Intn(100), ",")
	fmt.Print(x2.Intn(100))
	fmt.Println()
	g3 := rand.NewSource(42)
	x3 := rand.New(g3)
	fmt.Print(x3.Intn(100), ",")
	fmt.Print(x3.Intn(100))
	fmt.Println("")

	fmt.Println("#Parsing Numbers--------------------")
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)
	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)
	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)
	k, _ := strconv.Atoi("135")
	fmt.Println(k)
	_, e := strconv.Atoi("wat")
	fmt.Println(e)
}
