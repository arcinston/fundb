package main

import (
	"bufio"
	"os"
	"regexp"
)

type KVStore struct {
	filepath string
}

func NewKVStore(filepath string) *KVStore {
	store := &KVStore{
		filepath: filepath + "kvstore.db",
	}
	// check if the file exists
	_, err := os.Create(filepath + "kvstore.db")

	if err != nil {
		if os.IsExist(err) {
			print("file already exists")
		} else {
			print("error creating file")
		}
	}

	return store
}

func (kv *KVStore) loadfromfile() {

}

func (kv *KVStore) saveToFile() error {

}

func (kv *KVStore) set(key string, value string) bool {

	file, err := os.OpenFile(kv.filepath, os.O_APPEND, 0644)

	if err != nil {
		print("error opening file")
		return false
	}

	// we wrill write key=value\n to the file

	file.Write([]byte(key + "=" + value + "\n"))

	file.Close()
	return true
}

func (kv *KVStore) get(key string) (string, bool) {

	file, err := os.OpenFile(kv.filepath, os.O_RDONLY, 0644)

	if err != nil {
		print("error opening file")
		return "", false
	}

	buffer := make([]byte, 1024)

	for {
		n, err := file.Read(buffer)
		if err != nil {
			break
		}
		if n > 0 {
			print(string(buffer))
			// now we have this buffer , we need to read every line and check if the key is present

		}

	}

	return value, exists
}
