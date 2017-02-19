package main

import "fmt"
//import "math/big"
import "crypto/elliptic"
import "crypto/sha256"
import "crypto/rand"
import "crypto/ecdsa"
import "io/ioutil"
import "os"

func main() {
	h := sha256.New()
	h.Write([]byte("hello world\n"))
	fmt.Println("HASHING")
	fmt.Println(h.Sum(nil))
	myCurve := elliptic.P256()
	priv,_ := ecdsa.GenerateKey(myCurve,rand.Reader)
	fmt.Println("PRIV")
	fmt.Println(priv)
	pub := priv.PublicKey
	fmt.Println("PUB")
	fmt.Println(pub)
	r,s,_ := ecdsa.Sign(rand.Reader,priv,[]byte("herro"))
	fmt.Println("R")
	fmt.Println(r)
	fmt.Println("S")
	fmt.Println(s)
	fmt.Println("VERIFIED?")
	fmt.Println(ecdsa.Verify(&pub,[]byte("herro"),r,s))

	ioutil.WriteFile("/home/pi/gopath/ioutiltest.txt",[]byte("hello"),1911)
	txt,_ := ioutil.ReadFile("/home/pi/gopath/ioutiltest.txt")
	fmt.Println("FILE TEXT USING IOUTIL")
	fmt.Println(string(txt))

	f,_ := os.Create("/home/pi/gopath/ostest.txt")
	f.Write([]byte("hello"))
	f.WriteString(" it's me")
	data := make([]byte,13)
	f,_ = os.Open("/home/pi/gopath/ostest.txt")
	f.Read(data)
	fmt.Println("FILE TEXT USING OS")
	fmt.Println(string(data))
	f.Close()
}
