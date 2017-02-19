package main

import "fmt"
import "math/big"
import "crypto/elliptic"
import "crypto/sha256"
import "crypto/rand"
import "crypto/ecdsa"
import "io/ioutil"
import "os"

//TXNS,BLOCKS,AND CHAIN
type Otp struct {
	priv ecdsa.PrivateKey
	amt int
}
type TxnSig struct {
	pub ecdsa.PublicKey
	r big.Int
	s big.Int
}
type Txn struct { //hash before and after transaction signatures can be calculated
	inputs []Txn
	outputs []Otp
	transactionSigs []TxnSig
}
type Block struct { //mining reward, hit, difficulty, and hash can be calculated
	transactions []Txn //are both this and merkle tree hash necessary?
	merkleTreeHash []byte
	miner ecdsa.PublicKey
	timeSinceGenesis big.Int //lots of big ints below, are they all necessary?
	blockNum big.Int
	target big.Int
	lastBlockSignedHashR big.Int
	lastBlockSignedHashS big.Int
	hash []byte
}
type Chain struct {
	blocks []Block
	totalDifficulty big.Int
	totalSupply big.Int
	genesisTime big.Int //genesis unix time
}
//BASIC FUNCTIONS
func (txn Txn)toByte(beforeSigs bool) []byte {
	return []byte("do this")
} 
func (block Block)toByte() []byte {
	return []byte("do this")
}

func (txn Txn)hash(beforeSigs bool) []byte {
	sha := sha256.New()
	sha.Write(txn.toByte(beforeSigs))
	return sha.Sum(nil)
}

func addBlock(file os.File,block Block) (int,error) {
	return file.Write(block.toByte())
}
//MAIN LOOP
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
