package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type pixels map[int]map[int]string

func main() {
	buf, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(countConnected(string(buf)))
}

func countConnected(body string) int {
	p := make(pixels)
	p.bind(body)
	count := 0
	for y, row := range p {
		for x := range row {
			if p.affectAt(x, y) > 0 {
				count++
			}
		}
	}
	return count
}

func (p *pixels) bind(body string) {
	x := 0
	y := 0
	for _, v := range body {
		if fmt.Sprintf("%c", v) == "\n" {
			y++
			x = 0
			continue
		}
		if _, ok := (*p)[y]; !ok {
			(*p)[y] = make(map[int]string)
		}
		(*p)[y][x] = string(v)
		x++
	}
}

func (p *pixels) affectAt(x, y int) (affected int) {
	v, ok := (*p)[y][x]
	if !ok {
		return 0
	}
	if v == "." {
		return 0
	}
	{
		// if "0", change to "." (affected)
		(*p)[y][x] = "."
		affected++
	}
	for ty := y - 1; ty <= y+1; ty++ {
		for tx := x - 1; tx <= x+1; tx++ {
			affected += p.affectAt(tx, ty)
		}
	}
	return affected
}
