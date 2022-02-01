package iching

import (
	"fmt"
	"math/rand"
	"time"
)

/*
true = unbroken line
false = broken line
0 = 6 broken -> unbroken
1 = 7 unbroken
2 = 8 broken
3 = 9 unbroken -> broken
*/

type BSTnode struct {
	value int      //0 when not at end
	left  *BSTnode //false
	right *BSTnode //true
}

type BST struct {
	root *BSTnode
}

func InitializeIchingBST() BST {
	rootNode := &BSTnode{value: 0, left: nil, right: nil}
	bst := &BST{root: rootNode}
	initializeIchingBSTRec(bst.root, 0)
	insertHexagram(bst, [6]bool{true, true, true, true, true, true}, 1)
	insertHexagram(bst, [6]bool{false, false, false, false, false, false}, 2)
	insertHexagram(bst, [6]bool{true, false, false, false, true, false}, 3)
	insertHexagram(bst, [6]bool{false, true, false, false, false, true}, 4)
	insertHexagram(bst, [6]bool{true, true, true, false, true, false}, 5)
	insertHexagram(bst, [6]bool{false, true, false, true, true, true}, 6)
	insertHexagram(bst, [6]bool{false, true, false, false, false, false}, 7)
	insertHexagram(bst, [6]bool{false, false, false, false, true, false}, 8)
	insertHexagram(bst, [6]bool{true, true, true, false, true, true}, 9)
	insertHexagram(bst, [6]bool{true, true, false, true, true, true}, 10)
	insertHexagram(bst, [6]bool{true, true, true, false, false, false}, 11)
	insertHexagram(bst, [6]bool{false, false, false, true, true, true}, 12)
	insertHexagram(bst, [6]bool{true, false, true, true, true, true}, 13)
	insertHexagram(bst, [6]bool{true, true, true, true, false, true}, 14)
	insertHexagram(bst, [6]bool{false, false, true, false, false, false}, 15)
	insertHexagram(bst, [6]bool{false, false, false, true, false, false}, 16)
	insertHexagram(bst, [6]bool{true, false, false, true, true, false}, 17)
	insertHexagram(bst, [6]bool{false, true, true, false, false, true}, 18)
	insertHexagram(bst, [6]bool{true, true, false, false, false, false}, 19)
	insertHexagram(bst, [6]bool{false, false, false, false, true, true}, 20)
	insertHexagram(bst, [6]bool{true, false, false, true, false, true}, 21)
	insertHexagram(bst, [6]bool{true, false, true, false, false, true}, 22)
	insertHexagram(bst, [6]bool{false, false, false, false, false, true}, 23)
	insertHexagram(bst, [6]bool{true, false, false, false, false, false}, 24)
	insertHexagram(bst, [6]bool{true, false, false, true, true, true}, 25)
	insertHexagram(bst, [6]bool{true, true, true, false, false, true}, 26)
	insertHexagram(bst, [6]bool{true, false, false, false, false, true}, 27)
	insertHexagram(bst, [6]bool{false, true, true, true, true, false}, 28)
	insertHexagram(bst, [6]bool{false, true, false, false, true, false}, 29)
	insertHexagram(bst, [6]bool{true, false, true, true, false, true}, 30)
	insertHexagram(bst, [6]bool{false, false, true, true, true, false}, 31)
	insertHexagram(bst, [6]bool{false, true, true, true, false, false}, 32)
	insertHexagram(bst, [6]bool{false, false, true, true, true, true}, 33)
	insertHexagram(bst, [6]bool{true, true, true, true, false, false}, 34)
	insertHexagram(bst, [6]bool{false, false, false, true, false, true}, 35)
	insertHexagram(bst, [6]bool{true, false, true, false, false, false}, 36)
	insertHexagram(bst, [6]bool{true, false, true, false, true, true}, 37)
	insertHexagram(bst, [6]bool{true, true, false, true, false, true}, 38)
	insertHexagram(bst, [6]bool{false, false, true, false, true, false}, 39)
	insertHexagram(bst, [6]bool{false, true, false, true, false, false}, 40)
	insertHexagram(bst, [6]bool{true, true, false, false, false, true}, 41)
	insertHexagram(bst, [6]bool{true, false, false, false, true, true}, 42)
	insertHexagram(bst, [6]bool{true, true, true, true, true, false}, 43)
	insertHexagram(bst, [6]bool{false, true, true, true, true, true}, 44)
	insertHexagram(bst, [6]bool{false, false, false, true, true, false}, 45)
	insertHexagram(bst, [6]bool{false, true, true, false, false, false}, 46)
	insertHexagram(bst, [6]bool{false, true, false, true, true, false}, 47)
	insertHexagram(bst, [6]bool{false, true, true, false, true, false}, 48)
	insertHexagram(bst, [6]bool{true, false, true, true, true, false}, 49)
	insertHexagram(bst, [6]bool{false, true, true, true, false, true}, 50)
	insertHexagram(bst, [6]bool{true, false, false, true, false, false}, 51)
	insertHexagram(bst, [6]bool{false, false, true, false, false, true}, 52)
	insertHexagram(bst, [6]bool{false, false, true, false, true, true}, 53)
	insertHexagram(bst, [6]bool{true, true, false, true, false, false}, 54)
	insertHexagram(bst, [6]bool{true, false, true, true, false, false}, 55)
	insertHexagram(bst, [6]bool{false, false, true, true, false, true}, 56)
	insertHexagram(bst, [6]bool{false, true, true, false, true, true}, 57)
	insertHexagram(bst, [6]bool{true, true, false, true, true, false}, 58)
	insertHexagram(bst, [6]bool{false, true, false, false, true, true}, 59)
	insertHexagram(bst, [6]bool{true, true, false, false, true, false}, 60)
	insertHexagram(bst, [6]bool{true, true, false, false, true, true}, 61)
	insertHexagram(bst, [6]bool{false, false, true, true, false, false}, 62)
	insertHexagram(bst, [6]bool{true, false, true, false, true, false}, 63)
	insertHexagram(bst, [6]bool{false, true, false, true, false, true}, 64)

	return *bst
}

func initializeIchingBSTRec(node *BSTnode, depth int) {

	if depth == 6 {
		return
	}
	node.left = &BSTnode{value: 0, left: nil, right: nil}
	node.right = &BSTnode{value: 0, left: nil, right: nil}
	initializeIchingBSTRec(node.left, depth+1)
	initializeIchingBSTRec(node.right, depth+1)
}

func insertHexagram(bst *BST, hexagram [6]bool, number int) {
	currentNode := bst.root
	for i := 0; i < len(hexagram); i++ {
		if hexagram[i] {
			currentNode = currentNode.right
		} else {
			currentNode = currentNode.left
		}
	}
	currentNode.value = number
}

func FindHexagram(bst BST, hexagram [6]bool) {
	currentNode := bst.root
	for i := 0; i < len(hexagram); i++ {
		if hexagram[i] {
			currentNode = currentNode.right
		} else {
			currentNode = currentNode.left
		}
	}
	fmt.Println(currentNode.value)
}

func FlipCoins() int {
	rand.NewSource(time.Now().UnixNano())
	flip := rand.Intn(4)
	return flip
}

func PrintHexagram(hexagram [6]bool) {
	for i := len(hexagram) - 1; i >= 0; i-- {
		if hexagram[i] {
			fmt.Println("--------")
		} else {
			fmt.Println("---  ---")
		}
	}
	fmt.Println()
}

func PrintChangedLines(hexagram [6]bool, changedHexagram [6]bool) {
	fmt.Print("Changing Lines: ")
	for i := 0; i < len(hexagram); i++ {
		if hexagram[i] != changedHexagram[i] {
			fmt.Printf("%v ", i+1)
		}
	}
	fmt.Println()
}

func GenerateHexagram() ([6]bool, [6]bool) {
	var hexagram [6]bool
	var changedHexagram [6]bool
	for i := 0; i < 6; i++ {
		flipValue := FlipCoins()
		if flipValue == 0 {
			hexagram[i] = true
			changedHexagram[i] = false
		}
		if flipValue == 1 {
			hexagram[i] = false
			changedHexagram[i] = false
		}
		if flipValue == 2 {
			hexagram[i] = true
			changedHexagram[i] = true
		}
		if flipValue == 3 {
			hexagram[i] = false
			changedHexagram[i] = true
		}
	}
	return hexagram, changedHexagram
}
