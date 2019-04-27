package blockchain

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nonce    int
}

type BlockChain struct {
	Blocks []*Block
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	// add latest blocks array and new block.
	chain.Blocks = append(chain.Blocks, new)
}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{
		Hash:     []byte{},
		Data:     []byte(data),
		PrevHash: prevHash,
		Nonce:    0,
	}
	// block.DeriverHash()
	pow := NewProof(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}

/* No need due to use the ProofOfWork algorism
L 190423: Below method can make any hash which no target & Difficulty things.

func (b *Block) DeriverHash() {
// [][]byte because join both Data and PrevHash is []byte type. 2 types so 2-dimension `[][]` byte`
info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
hash := sha256.Sum256(info)
b.Hash = hash[:]
}
*/
