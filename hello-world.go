package main

import "fmt"
//import "math/big"
import "crypto/elliptic"
import "crypto/sha256"
import "crypto/rand"
import "crypto/ecdsa"

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
}
