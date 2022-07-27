package main

import (
	"fmt"
	"log"
	"math/rand"
	// "strings"

	"github.com/therealfakemoot/treedrawer/tree"
)

type MutateOpts struct {
	MinWidth, MaxWidth int
	MinDepth, MaxDepth int
}

func Mutate(t *tree.Tree, depth int, mo MutateOpts) {
	log.Printf("Mutating tree at depth %d\n", depth)
	if depth > mo.MaxDepth {
		log.Printf("trying to exit recursion: %t\n", depth > mo.MaxDepth)
		return
	}
	for i := 0; i < mo.MaxWidth; i++ {
		if rand.Float64() > .5 {
			return
		}
		log.Printf("Creating node (%d,%d)\n", depth, i)
		subtree := tree.NewTree(tree.NodeString(fmt.Sprintf("%d,%d", depth, i)))
		Mutate(subtree, depth+1, mo)
		t.AddChildNode(subtree)
	}

}

func main() {

	var ()
	t := tree.NewTree(tree.NodeString("0"))
	mo := MutateOpts{
		MinDepth: 4,
		MinWidth: 2,
		MaxDepth: 9,
		MaxWidth: 6,
	}
	Mutate(t, 0, mo)
	fmt.Println(t)
}
