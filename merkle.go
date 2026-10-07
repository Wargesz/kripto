package main

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"math"
	"slices"
	"strings"
)

type tree [][]string

func (t *tree) String() string {
	builder := strings.Builder{}
	for i, v := range *t {
		builder.WriteString(fmt.Sprintln(i, v))
	}
	return builder.String()
}

func buildMerkleTree(values []string) tree {
	treeHeight := calculateTreeHeight(values)
	tree := make([][]string, treeHeight)
	tree[treeHeight-1] = hashValues(values)
	for i := treeHeight - 2; i >= 0; i-- {
		tree[i] = buildParents(tree[i+1])
	}
	tree = append(tree, values)
	return tree
}

func hashValues(values []string) []string {
	var hashes []string
	for _, v := range values {
		bytes := sha256.Sum256([]byte(v))
		hashes = append(hashes, fmt.Sprintf("%x", bytes))
	}
	return hashes
}

func buildParents(values []string) []string {
	var parents []string
	for i := range len(values) / 2 {
		hash := sha256.Sum256([]byte(values[2*i] + values[2*i+1]))
		parents = append(parents, fmt.Sprintf("%x", hash))
	}
	return parents
}

func calculateTreeHeight(values []string) int {
	return int(math.Log2(float64(len(values)))) + 1
}

func merkleProof(tree tree, index int) ([]string, []bool) {
	height := len(tree)
	var values []string
	var sides []bool
	values = append(values, tree[height-1][index])
	parent := tree[height-2][index]
	for parent != tree[0][0] {
		value, side := getSibling(tree, parent)
		values = append(values, value)
		sides = append(sides, side == 1)
		parent = getParent(tree, parent)
	}
	return values, sides
}

func joinTree(tree tree) []string {
	var slice []string
	for i := range tree {
		slice = slices.Concat(slice, tree[i])
	}
	return slice
}

func getParent(tree tree, value string) string {
	joinedTree := joinTree(tree)
	parentIndex := (joinedIndexOf(joinedTree, value) - 1) / 2
	return joinedTree[parentIndex]
}

func getSibling(tree tree, value string) (string, int) {
	side := -1
	if arrIndexOf(tree, value)%2 == 0 {
		side = 1
	}
	return tree[depthOf(tree, value)][arrIndexOf(tree, value)+side], side
}

func arrIndexOf(tree tree, target string) int {
	for i, item := range tree[depthOf(tree, target)] {
		if item == target {
			return i
		}
	}
	must(errors.New("arrIndexOf"))
	return -1
}

func joinedIndexOf(items []string, target string) int {
	for i, item := range items {
		if item == target {
			return i
		}
	}
	must(errors.New("joinedIndexOf"))
	return -1
}

func verifyMerkleTree(tree tree, index int) bool {
	proof, sides := merkleProof(tree, index)
	bytes := sha256.Sum256([]byte(proof[0]))
	hash := fmt.Sprintf("%x", bytes)
	var concat string
	for i, value := range proof[1:] {
		if sides[i] {
			concat = hash + value
		} else {
			concat = value + hash
		}
		bytes = sha256.Sum256([]byte(concat))
		hash = fmt.Sprintf("%x", bytes)
	}
	return hash == tree[0][0]
}

func depthOf(tree tree, target string) int {
	for i := range tree {
		if slices.Contains(tree[i], target) {
			return i
		}
	}
	must(errors.New("depthOf: " + target))
	return -1
}
