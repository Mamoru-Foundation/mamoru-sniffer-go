package cosmos

type Block struct {
	Seq                                          uint64
	Height                                       int64
	Hash                                         []byte
	VersionBlock                                 uint64
	VersionApp                                   uint64
	ChainId                                      string
	Time                                         int64
	LastBlockIdHash                              []byte
	LastBlockIdPartSetHeaderTotal                uint32
	LastBlockIdPartSetHeaderHash                 []byte
	LastCommitHash                               []byte
	DataHash                                     []byte
	ValidatorsHash                               []byte
	NextValidatorsHash                           []byte
	ConsensusHash                                []byte
	AppHash                                      []byte
	LastResultsHash                              []byte
	EvidenceHash                                 []byte
	ProposerAddress                              []byte
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
	PubKey []byte
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
	Seq             uint64
	BlockSeq        uint64
	Validator       []byte
	SignedLastBlock bool
}

type Transaction struct {
	Seq       uint64
	Tx        []byte
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
