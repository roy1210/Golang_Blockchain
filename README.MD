# Blockchain

## Description

Making blockchain with POW system.

**_DEMO:_**

<img src="https://github.com/roy1210/Golang_Blockchain/raw/master/GoBlockchain.jpg" width="80%">

## Usage

1. Open main.go
2. Add some words in args of `chain.AddBlock` function in `func main()`
3. Run: `$ run main.go`
4. Word added in #2 will link with previous hash

```
func main() {
	chain := blockchain.InitBlockChain()
	chain.AddBlock("First Block after Genesis block")
	chain.AddBlock("Second Block after Genesis block")
	chain.AddBlock("Third Block after Genesis block")

	for _, block := range chain.Blocks {
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Date in Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)

		pow := blockchain.NewProof(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
```

then run: `$ run main.go`

## Requirement

- Golang

## Licence

[MIT](./LICENSE.txt)

## Author

[Shoe Kure](https://github.com/roy1210)
