package server

type Response struct {
	Status          bool `json:"status"`
	ConsensusHeight uint64  `json:"consensus height"`
	BridgeHeight    uint64  `json:"bridge height"`
}
