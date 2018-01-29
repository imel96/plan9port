package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	var nflg = flag.Bool("n", false, "Numeric Unix timestamp")
	var uflg = flag.Bool("u", false, "Unix date in GMT")
	var now time.Time

	flag.Parse()
	s := flag.Arg(0)

	if len(s) > 0 {
		sec, _ := strconv.ParseInt(s, 10, 64)
		now = time.Unix(sec, 0)

	} else {
		now = time.Now()
	}
	if *nflg {
		fmt.Println(now.Unix())

	} else if *uflg {
		loc, _ := time.LoadLocation("GMT")
		fmt.Println(now.In(loc).Format(time.UnixDate))

	} else {
		fmt.Println(now.Format(time.UnixDate))
	}
	os.Exit(0)
}
