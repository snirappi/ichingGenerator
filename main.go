package main

import (
	"iching"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixMicro())
}
func main() {
	bst := iching.InitializeIchingBST()
	hexagram, changedHexagram := iching.GenerateHexagram()
	iching.PrintHexagram(hexagram)
	iching.FindHexagram(bst, hexagram)
	iching.PrintChangedLines(hexagram, changedHexagram)
	iching.PrintHexagram(changedHexagram)
	iching.FindHexagram(bst, changedHexagram)

}
