package main

import (
	"fmt"
	"hash/fnv"
)

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

func main() {

	keys := []string{
		"key1",
		"key2",
		"key3",
		"key4",
		"key5",
		"key6",
		"key7",
		"key8",
	}
	N := 10

	for i := range keys {
		servcerIndex := hash(keys[i]) % uint32(N)
		fmt.Printf("Index\t[%d]\t%s\t%d\t%d\n", i, keys[i], hash(keys[i]), servcerIndex)
	}
}
