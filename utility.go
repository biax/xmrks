package main

import (
	"fmt"
	"math"
	"os"
)

func pr(f string, v ...interface{}) {
	fmt.Printf(f+"\n", v...)
}

func check(e error) {
	if e != nil {
		pr("Error: %s", e.Error())
		os.Exit(2)
	}
}

func exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func humanNumber(n int, si bool) string {
	num := float64(n)
	var unit float64 = 1024
	if !si {
		unit = 1000
	}
	if num < unit {
		return fmt.Sprintf("%dB", int(num))
	}
	exp := int(math.Log(num) / math.Log(unit))
	pre := "kMGTPE"
	if si {
		pre = "KMGTPE"
		pre = pre[exp-1:exp] + "i"
	} else {
		pre = pre[exp-1 : exp]
	}
	r := n / int(math.Pow(unit, float64(exp)))
	return fmt.Sprintf("%d %sB", r, pre)
}
