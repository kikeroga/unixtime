package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	as := os.Args
	if len(as) < 2 {
		fmt.Println("Argument does not exist.")
		return
	}
	t := as[1]

	us, err := strconv.ParseInt(t, 10, 64)
	if err != nil {
		fmt.Println(fmt.Sprintf("Invalid Unix time: %s", t))
		return
	}

	u := time.Unix(us, 0)
	fmt.Println(u)
}
