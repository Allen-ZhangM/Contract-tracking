package service

import "github.com/kongyixueyuan.com/bill/blockchain"

type FabricSetupService struct {
	Fabric *blockchain.FabricSetup
}

type Bill struct {
	BillInfoID   string `json:"BillInfoID"`   // 票据号码
	BillInfoAmt  string `json:"BillInfoAmt"`  // 票据金额
	BillInfoType string `json:"BillInfoType"` // 票据类型

	BillInfoIsseDate string `json:"BillInfoisseDate"` // 出票日期
	BillInfoDueDate  string `json:"BillInfoDueDate"`  // 到期日期

	DrwrAcct string `json:"DrwrAcct"` // 出票人名称
	DrwrCmID string `json:"DrwrCmID"` // 出票人证件号码

	AccptrAcct string `json:"AccptrAcct"` // 承兑人名称
	AccptrCmID string `json:"AccptrCmID"` // 承兑人证件号码

	PyeeAcct string `json:"PyeeAcct"` // 收款人名称
	PyeeCmID string `json:"PyeeCmID"` // 收款人证件号码

	HoldrAcct string `json:"HoldrAcct"` // 当前持票人名称
	HoldrCmID string `json:"HoldrCmID"` // 当前持票人证件号码

	WaitEndorseAcct string `json:"WaitEndorseAcct"` // 待背书人名称
	WaitEndorseCmID string `json:"WaitEndorseCmID"` // 待背书人证件号码

	RejectEndorseAcct string `json:"RejectEndorseAcct"` // 拒绝背书人名称
	RejectEndorseCmID string `json:"RejectEndorseCmID"` // 拒绝背书人证件号码

	State    string        `json:"State"` // 票据状态
	Historys []HistoryItem // 当前票据的历史流转记录
}

type HistoryItem struct {
	TxId string
	Bill Bill
}
