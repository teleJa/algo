package hash

import (
	"log"
	"testing"
)

func TestChainHashMap(t *testing.T) {

	hashMap := NewChainHashMap(10)
	hashMap.Put("a", 1)
	hashMap.Put("b", 2)
	hashMap.Put("c", 3)
	hashMap.Put("e", 5)

	hashMap.print()

	hashMap.Remove("a")
	log.Println("=======================>")
	hashMap.print()

	log.Println("=======================>")
	exist, i := hashMap.Get("e")
	log.Println("exist:", exist, "val:", i)

}
