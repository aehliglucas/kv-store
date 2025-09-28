package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"os"
	"path/filepath"
	// "reflect"
	// "sort"
	"time"
)

type KVStore struct {
	memtable map[string]string
	write_counter int
	threshold int
	sst_dir string
}

func (store *KVStore) Get(key string) string {
	value, ok := store.memtable[key]
	var err error
	if !ok {
		value, err = store.querySSTs(key)
		if err != nil {
			log.Fatal(err)
		}
	}
	return value
}

func (store *KVStore) Put(key string, value string) {
	if (store.write_counter >= store.threshold) {
		store.flush()
	}
	store.memtable[key] = value
	store.write_counter++
}

func (store *KVStore) flush() {
	// TODO: Replace with Skiplist
	// keys := reflect.ValueOf(store.memtable).MapKeys()
	// resultKeyList := []string{}

	// for _, key := range keys {
	// 	resultKeyList = append(resultKeyList, key.Interface().(string))
	// }
	// sort.Strings(resultKeyList)

	filePath := filepath.Join(store.sst_dir, time.Now().Format(time.RFC3339))
	fileHandler, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR, 0640)
	if err != nil {
		log.Fatal(err)
	}
	defer fileHandler.Close()

	encoder := gob.NewEncoder(fileHandler)
	err = encoder.Encode(store.memtable)
	if err != nil {
		log.Fatal(err)
	}

	store.memtable = make(map[string]string)
}

func (store *KVStore) querySSTs(key string) (string, error) {
	files, _ := os.ReadDir(store.sst_dir)

	for _, file := range files {
		result, err := store.deserializeSST(file.Name())
		if err != nil {
			return "", err
		}
		if value, ok := result[key]; ok {
			return value, nil
		}
	}

	return "", fmt.Errorf("Key %s not found.", key)
}

func (store *KVStore) deserializeSST(fileName string) (map[string]string, error) {
	filePath := filepath.Join(store.sst_dir, fileName)
	fileHandler, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR, 0640)
	if err != nil {
		return nil, err
	}
	defer fileHandler.Close()
	
	var result map[string]string
	decoder := gob.NewDecoder(fileHandler)
	err = decoder.Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}