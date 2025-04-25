package relay

// Finite-state automaton driven by events.
type Action int

const (
	DoAck Action = iota // call ContractB.acknowledge
	DoSynAck            // call ContractA.acknowledge
	DoConfirm           // call ContractB.confirm
)
