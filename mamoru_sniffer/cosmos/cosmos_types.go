package cosmos

type Block struct {
	Seq                                          uint64
	Height                                       int64
	Hash                                         string
	VersionBlock                                 uint64
	VersionApp                                   uint64
	ChainId                                      string
	Time                                         int64
	LastBlockIdHash                              string
	LastBlockIdPartSetHeaderTotal                uint32
	LastBlockIdPartSetHeaderHash                 string
	LastCommitHash                               string
	DataHash                                     string
	ValidatorsHash                               string
	NextValidatorsHash                           string
	ConsensusHash                                string
	AppHash                                      string
	LastResultsHash                              string
	EvidenceHash                                 string
	ProposerAddress                              string
	LastCommitInfoRound                          int32
	ConsensusParamUpdatesBlockMaxBytes           int64
	ConsensusParamUpdatesBlockMaxGas             int64
	ConsensusParamUpdatesEvidenceMaxAgeNumBlocks int64
	ConsensusParamUpdatesEvidenceMaxAgeDuration  int64
	ConsensusParamUpdatesEvidenceMaxBytes        int64
	ConsensusParamUpdatesValidatorPubKeyTypes    string
	ConsensusParamUpdatesVersionApp              uint64
}

type ValidatorUpdate struct {
	Seq    uint64
	PubKey []byte //Ed25519 PubKey
	Power  int64
}

type Misbehavior struct {
	Seq              uint64
	BlockSeq         uint64
	Typ              string
	ValidatorPower   int64
	ValidatorAddress string
	Height           int64
	Time             int64
	TotalVotingPower int64
}

type VoteInfo struct {
	Seq              uint64
	BlockSeq         uint64
	ValidatorAddress string
	ValidatorPower   int64
	SignedLastBlock  bool
}

type Transaction struct {
	Seq       uint64
	Tx        []byte
	TxHash    string
	TxIndex   uint32
	Code      uint32
	Data      []byte
	Log       string
	Info      string
	GasWanted int64
	GasUsed   int64
	Codespace string
}

type Event struct {
	Seq       uint64
	EventType string
}

type EventAttribute struct {
	Seq      uint64
	EventSeq uint64
	Key      string
	Value    string
	Index    bool
}

type EvmCallTrace struct {
	TxHash       string
	TxIndex      uint32
	BlockIndex   int64
	Depth        uint32
	Type         string
	From         string
	To           string
	Value        uint64
	GasLimit     uint64
	GasUsed      uint64
	Input        []byte
	Output       string
	Error        string
	RevertReason string
}
