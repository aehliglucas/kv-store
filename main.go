package main

import (
	"fmt"
)

func main() {
	kvstore := KVStore{
		memtable: make(map[string]string),
		write_counter: 0,
		threshold: 3,
		sst_dir: "/tmp/kv-store",
	}

	kvstore.Put("c", "still works!")
	kvstore.Put("a", "works!")
	kvstore.Put("b", "0815")
	kvstore.Put("d", "i have been flushed!")
	fmt.Printf("Obtained from memtable: %s\n", kvstore.Get("d"))
	fmt.Printf("Obtained from disk: %s", kvstore.Get("a"))
}