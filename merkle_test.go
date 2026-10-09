package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestBuildMerkleTree(t *testing.T) {
	leaves := []string{"one", "two", "three", "four"}
	expectedTree := [][]string{}
	expectedTree = append(expectedTree, []string{"44633b02bb638a76c8616dcaf18285b024e690fd33b73572d4fa18938d67137b"})
	expectedTree = append(expectedTree, []string{"298315b0a8b450097e70dd056648c5ab87b6aaa2bf31f75465906c243cb36145",
		"f284fdf4379eaec2d455f7e589618c8b0055b26a7e106b0f19afedaa1605befd"})
	expectedTree = append(expectedTree, []string{"7692c3ad3540bb803c020b3aee66cd8887123234ea0c6e7143c0add73ff431ed",
		"3fc4ccfe745870e2c0d99f71f30ff0656c8dedd41cc1d7d3d376b0dbe685e2f3",
		"8b5b9db0c13db24256c829aa364aa90c6d2eba318b9232a4ab9313b954d3555f",
		"04efaf080f5a3e74e1c29d1ca6a48569382cbbcd324e8d59d2b83ef21c039f00"})
	expectedTree = append(expectedTree, []string{"one", "two", "three", "four"})
	tree := buildMerkleTree(leaves)
	if fmt.Sprint(tree) != fmt.Sprint(expectedTree) {
		t.Errorf("bad merkle tree")
	}
}

func TestBuildMerkleProof(t *testing.T) {
	leaves := []string{"one", "two", "three", "four"}
	tree := buildMerkleTree(leaves)
	proof, sides := merkleProof(tree, 1)
	expectedProof := []string{"two", "7692c3ad3540bb803c020b3aee66cd8887123234ea0c6e7143c0add73ff431ed",
		"f284fdf4379eaec2d455f7e589618c8b0055b26a7e106b0f19afedaa1605befd"}
	expectedSides := []bool{false, true}
	if strings.Join(proof, ",") != strings.Join(expectedProof, ",") {
		t.Errorf("wrong steps give for merkle tree proof")
	}
	for i := range sides {
		if sides[i] != expectedSides[i] {
			t.Errorf("wrong side calculated")
		}
	}
}

func TestMerkleTreeVerify(t *testing.T) {
	leaves := []string{"first", "second", "third", "fourth", "fifth", "sixth",
		"seventh", "eighth"}
	tree := buildMerkleTree(leaves)
	for i := 0; i < len(leaves); i++ {
		if !verifyMerkleTree(tree, i) {
			t.Errorf("merkle tree verification failed")
		}
	}
}
