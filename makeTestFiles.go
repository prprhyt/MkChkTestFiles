package main

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
	"os"
	"strconv"
)

func main() {
	createFile()
}

func createFile(){
	hash_file, _ := os.Create("src/hashtable.txt")
	for i:=1;i<=1000;i++{
		b := make([]byte,100000+i-1)
		rand.Read(b)
		file,_ :=os.Create("src/"+strconv.Itoa(i)+".bin")
		file.Write(b)
		file.Close()
		a := sha512.Sum512(b)
		hash_file.WriteString(base64.StdEncoding.EncodeToString(a[:])+"\n")
	}
	hash_file.Close()
}