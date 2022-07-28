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
	MaxWidth         int
	MaxDepth         int
	ChildProbability float64
}

func NameNode(t *tree.Tree) {
	name := ""
	p, ok := t.Parent()
	if ok {
		val := p.Val()
		pname, _ := val.(tree.NodeString)
		name += string(pname)
	}
	p.SetVal(tree.NodeString(name))

}

func Mutate(t *tree.Tree, depth int, mo MutateOpts) {
	log.Printf("Mutating tree at depth %d\n", depth)

	if depth > mo.MaxDepth {
		log.Printf("trying to exit recursion: %t\n", depth > mo.MaxDepth)
		return
	}

	for i := 0; i < mo.MaxWidth; i++ {
		if rand.Float64() < mo.ChildProbability {
			continue
		}
		log.Printf("Creating child node (%d%s)\n", depth, i)
		subtree := tree.NewTree(tree.NodeString(fmt.Sprintf("%s%d%s", depth, chars[i])))
		NameNode(subtree)
		t.AddChildNode(subtree)
	}

	log.Printf("Added %d children at depth %d\n", len(t.Children()), depth)

	for _, c := range t.Children() {
		Mutate(c, depth+1, mo)
	}
}

func main() {

	var ()

	rand.Seed(time.Now().UnixNano())
	t := tree.NewTree(tree.NodeString("1"))
	mo := MutateOpts{
		MaxDepth:         6,
		MaxWidth:         2,
		ChildProbability: .4,
	}
	log.Println("mutations starting")
	Mutate(t, 0, mo)
	log.Println("mutations complete")
	fmt.Println(t)
}
