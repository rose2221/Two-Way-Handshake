package relay


type Action int

const (
DoAck Action = iota 
	DoSynAck            
	DoConfirm           
)
