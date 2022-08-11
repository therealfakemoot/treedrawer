package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
	// "strings"

	"github.com/therealfakemoot/treedrawer/tree"
)

var chars = map[int]string{
	1:  "a",
	2:  "b",
	3:  "c",
	4:  "d",
	5:  "e",
	6:  "f",
	7:  "g",
	8:  "h",
	9:  "i",
	10: "j",
	11: "k",
	12: "l",
	13: "m",
	14: "n",
	15: "o",
	16: "p",
	17: "q",
	18: "s",
	19: "r",
	20: "t",
	21: "u",
	22: "v",
	23: "w",
	24: "x",
	25: "y",
	26: "z",
}

type MutateOpts struct {
	MaxWidth int
	MaxDepth int
}

func Mutate(t *tree.Tree, depth int, mo MutateOpts) {
	log.Printf("Mutating tree at depth %d\n", depth)

	if depth > mo.MaxDepth {
		log.Printf("Mutate recursion depth exceeeded: %t\n", depth > mo.MaxDepth)
		return
	}

	numChildren := rand.Intn(mo.MaxWidth + 1)
	for i := 0; i < numChildren; i++ {
		parentName := t.Val()
		subtree := tree.NewTree(tree.NodeString(fmt.Sprintf("%s%d%s", parentName, depth+1, chars[i+1])))
		t.AddChildNode(subtree)
	}

	for _, c := range t.Children() {
		Mutate(c, depth+1, mo)
	}
}

func main() {

	var ()

	rand.Seed(time.Now().UnixNano())
	t := tree.NewTree(tree.NodeString(""))
	mo := MutateOpts{
		MaxDepth: 999,
		MaxWidth: 300,
	}
	log.Println("mutations starting")
	Mutate(t, 0, mo)
	log.Println("mutations complete")
	// fmt.Println(t)
}
