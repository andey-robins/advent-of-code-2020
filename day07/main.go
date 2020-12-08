package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

// Bag is a struct that tracks the adjective and color of a bag
type Bag struct {
	adj   string
	color string
	held  []string
}

// BagTreeNode is a node in a tree of bags
type BagTreeNode struct {
	bag      Bag
	children []BagTreeNode
}

// BagTreeNodeSearchError is an error raised by bagInTree if the bag can't be found
type BagTreeNodeSearchError struct {
	bag  Bag
	prob string
}

func (e *BagTreeNodeSearchError) Error() string {
	return fmt.Sprintf("%s - %s", e.bag, e.prob)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// a traversal that will return nil if the bag is in the tree and an error if it is not
func (b *BagTreeNode) bagInTree(bag Bag) error {
	if b.bag.adj == bag.adj && b.bag.color == bag.color {
		return nil
	}
	for _, subBag := range b.children {
		return subBag.bagInTree(bag)
	}
	return &BagTreeNodeSearchError{bag, "Unable to find bag"}
}

// returns the number of leaves attached to a node
func (b *BagTreeNode) countTreeLeaves() int {

	leaves := 0

	if len(b.children) == 0 {
		return 1
	}

	for _, child := range b.children {
		leaves += child.countTreeLeaves()
	}

	return leaves
}

// removes a bag from a list of bags then returns the newly modified list
func removeBagFromList(list []Bag, index int) []Bag {
	list[index] = list[len(list)-1]
	return list[:len(list)-1]
}

// given a list of bags to remove (such as the one provided by findBagsThatCanHold), return a list with all of those indices removed
// respects their position even after beginning to delete bags by iterating over indices in reverse (i.e. from high to low)
func removeListOfBagsFromList(list []Bag, indices []int) []Bag {
	output := list
	for i := len(indices) - 1; i >= 0; i++ {
		output = removeBagFromList(list, indices[i])
	}
	return output
}

// returns a list of bags that can hold the argument heldBag and the indexes of them in the list
func findBagsThatCanHold(list []Bag, heldBag Bag) ([]Bag, []int) {
	bagsThatHold := make([]Bag, 0)
	indexes := make([]int, 0)
	for i, bag := range list {
		// this is only checking if the bag is already there (which of course it wont be)
		if bag.adj == heldBag.adj && bag.color == heldBag.color {
			bagsThatHold = append(bagsThatHold, Bag{adj: bag.adj, color: bag.color})
			indexes = append(indexes, i)
		}

		foundAdj := ""
		for j := 0; j < len(bag.held); j++ {
			if j%4 == 1 && heldBag.adj == bag.held[j] {
				foundAdj = bag.held[j]
			} else if j%4 == 2 && heldBag.color == bag.held[j] && foundAdj == heldBag.adj {
				bagsThatHold = append(bagsThatHold, Bag{adj: foundAdj, color: bag.held[j]})
			}
		}
	}

	fmt.Println(bagsThatHold)
	return bagsThatHold, indexes
}

func main() {
	// select part one or two
	partOne := true
	fileName := "./rules.txt"
	if len(os.Args) == 2 && os.Args[1] == "2" {
		partOne = false
	} else if len(os.Args) == 2 && os.Args[1] == "s" {
		fileName = "./rulessmall.txt"
	}

	f, err := os.Open(fileName)
	check(err)
	defer func() {
		err := f.Close()
		check(err)
	}()

	wordRegex := regexp.MustCompile("\\w*")

	fmt.Println("Beginning to load data from file")

	// read data in from the input file
	s := bufio.NewScanner(f)
	allBags := make([]Bag, 0)
	shinyGoldIndex := -1
	indexCounter := 0
	for s.Scan() {
		line := s.Text()
		words := wordRegex.FindAllString(line, -1)
		lineBag := Bag{
			adj:   words[0],
			color: words[1],
			held:  words[4:],
		}

		if lineBag.adj == "shiny" && lineBag.color == "gold" {
			shinyGoldIndex = indexCounter
		}

		allBags = append(allBags, lineBag)
		indexCounter++
	}

	fmt.Println("Reading complete.")

	if partOne {
		root := BagTreeNode{
			bag: allBags[shinyGoldIndex],
		}

		allBags = removeBagFromList(allBags, shinyGoldIndex)
		queue := make([]BagTreeNode, 0)
		queueIndex := 0
		queue = append(queue, root)

		for len(allBags) != 0 {
			activeNode := queue[queueIndex]
			queueIndex++ // simulate a pop

			fmt.Printf("Beginning searching for children for bag %s %s\n", activeNode.bag.adj, activeNode.bag.color)

			children, indices := findBagsThatCanHold(allBags, activeNode.bag)

			fmt.Println(children)

			// create tree nodes for children bags
			activeNodesChildren := make([]BagTreeNode, 0)
			for _, childBag := range children {
				childBagNode := BagTreeNode{bag: childBag}
				queue = append(queue, childBagNode)
				activeNodesChildren = append(activeNodesChildren, childBagNode)
			}
			activeNode.children = activeNodesChildren
			allBags = removeListOfBagsFromList(allBags, indices)
		}

		// go over every bag to see if it can hold a shiny gold bag
		// if it can, add it to the next level of the tree
		// remove the bag from the list of bags
		// go down a level and for each bag, repeat the process

		fmt.Printf("Part 1 solution: %v", root.countTreeLeaves())
	} else {
		fmt.Println("Part 2")
	}
}
