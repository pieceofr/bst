package bst

import (
	"fmt"
	"testing"
	//"log"
)

//Test number an "array" of integer
var testKeys = [...]int{11, 6, 8, 19, 4, 10, 5, 17, 43, 49, 31}
var preOrderRes = [...]int{11, 6 ,4 ,5 ,8 ,10, 19, 17,43, 31, 49}

func TestInsert(t *testing.T) {
	btree := BST{}

	for _, key := range testKeys {
		btree.Insert(key)
	}
	list := btree.PreOrderPrint()
	for idx, key := range list {
		if preOrderRes[idx] != key {
			fmt.Printf("preOrderRes[%d]=%d != %d \n", idx, preOrderRes[idx], key)
			t.Failed()
		}
	}
}
