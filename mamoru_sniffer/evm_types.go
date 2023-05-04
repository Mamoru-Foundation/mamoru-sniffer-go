package mamoru_sniffer

type Block struct {
	BlockIndex      uint64
	Hash            string
	ParentHash      string
	StateRoot       string
	Nonce           uint64
	Status          string
	Timestamp       uint64
	BlockReward     []byte
	FeeRecipient    string
	TotalDifficulty uint64
	Size            float64
	GasUsed         uint64
	GasLimit        uint64
}

type Transaction struct {
	TxIndex    uint32
	TxHash     string
	Type       uint8
	Nonce      uint64
	Status     uint64
	BlockIndex uint64
	From       string
	To         string
	Value      uint64
	Fee        uint64
	GasPrice   uint64
	GasLimit   uint64
	GasUsed    uint64
	Size       float64
	Input      []byte
}

type Event struct {
	Index       uint32
	TxIndex     uint32
	Address     string
	BlockNumber uint64
	TxHash      string
	BlockHash   string
	Topic0      []byte
	Topic1      []byte
	Topic2      []byte
	Topic3      []byte
	Topic4      []byte
	Data        []byte
}

type CallTrace struct {
	Seq        uint32
	Depth      uint32
	TxIndex    uint32
	BlockIndex uint64
	Type       string
	From       string
	To         string
	Value      uint64
	GasLimit   uint64
	GasUsed    uint64
	Input      []byte
}
