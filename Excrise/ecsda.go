package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"os"
	"io/ioutil"
)

func main() {
	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	file, _ := os.Create("key.dat")
	defer file.Close()
	err := ioutil.WriteFile("key.dat", privateKey.D.Bytes(), 0644)
	if err!= nil{
		fmt.Println(err)
	}
}
