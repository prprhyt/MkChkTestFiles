package main

import (
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	hashTableFilePath := "src/hashtable.txt"
	dstDirPath := "dst/"
	if(3==len(os.Args)){
		hashTableFilePath = os.Args[1]
		dstDirPath = os.Args[2]
		if('/'!=dstDirPath[len(dstDirPath)-1]){
			dstDirPath = dstDirPath + "/"
		}
	}

	b, err := ioutil.ReadFile(hashTableFilePath)
	if err != nil {
		fmt.Print(err)
	}
	collisionFile := []string{}
	missFile := []string{}
	isExistMap := make(map[string]bool)
	str := string(b)
	hashtableString := strings.Split(str,"\n")
	fileCount :=0
	isCollectFileCount := 0
	for i:=1;i<=1000;i++{
		b, err := ioutil.ReadFile(dstDirPath+strconv.Itoa(i)+".bin")
		if(nil!=err){
			fmt.Println("failed:"+ strconv.Itoa(i)+".bin")
			continue
		}
		fileCount++
		a := sha512.Sum512(b)
		s := base64.StdEncoding.EncodeToString(a[:])
		if(contains(hashtableString,s)){
			_, exist := isExistMap[s]
			if(exist){
				collisionFile = append(collisionFile, "collision: "+dstDirPath+strconv.Itoa(i)+".bin")
				continue
			}
			fmt.Println("correct: "+dstDirPath+strconv.Itoa(i)+".bin")
			isExistMap[s] = true
			isCollectFileCount++
		}else{
			missFile = append(missFile, "miss:"+dstDirPath+strconv.Itoa(i)+".bin")
		}
	}
	fmt.Println("")
	for _,e:= range collisionFile{
		fmt.Println(e)
	}
	fmt.Println("")
	for _,e:= range missFile{
		fmt.Println(e)
	}
	fmt.Println("")
	fmt.Println("correct/receive="+strconv.Itoa(isCollectFileCount)+"/"+strconv.Itoa(fileCount))
}

func contains(s []string, e string) bool {
	for _, v := range s {
		if e == v {
			return true
		}
	}
	return false
}