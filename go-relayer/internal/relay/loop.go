package relay

import (
	"context"
	"log"
"github.com/ethereum/go-ethereum/accounts/abi/bind"
	contracts "github.com/rose2221/Two-Way-Handshake/go-relayer/internal/contracts"
	// "github.com/rose2221/Two-Way-Handshake/go-relayer/internal/eth"
	// "github.com/rose2221/Two-Way-Handshake/go-relayer/internal/persistence"
)

type Loop struct {
	ctx        context.Context
	cancel     context.CancelFunc
	// db         *persistence.LevelDB
	a          *contracts.ContractA    
	b          *contracts.ContractB      
	auth_a, auth_b *bind.TransactOpts
}

func NewLoop(
	ctx context.Context,
	// db *persistence.LevelDB,
	a *contracts.ContractA, b *contracts.ContractB,
	auth_a, auth_b *bind.TransactOpts,
) *Loop {
	con, cancel := context.WithCancel(ctx)
	return &Loop{ctx:con, cancel:cancel, /*db:db, */ a:a, b:b, auth_a:auth_a, auth_b:auth_b}
}

func (l *Loop) Start() {
	go l.see_Syn()
	go l.see_Ack()
	go l.see_SynAck()
	<-l.ctx.Done()
}


func (l *Loop) see_Syn() {
	sink1 := make(chan *contracts.ContractASyn)
	
	sub1, err := l.a.WatchSyn(&bind.WatchOpts{Context:l.ctx}, sink1, nil, nil)
	if err != nil { log.Fatalf("got error while watching syn- %v", err) }
	log.Println("listening has been initiated for Syn events....")

	for {
		select {
		case ev := <-sink1:
			log.Printf("Syn session=%d -> will lead to ACK on B", ev.SessionId)
			txn, err := l.b.Acknowledge(l.auth_b, ev.SessionId)
			if err != nil {
				log.Printf("ack txn err: %v", err); continue
			}
			log.Printf("sent txn %s", txn.Hash())
		case err := <-sub1.Err():
			log.Printf("syn sub err: %v", err)
			return
		case <-l.ctx.Done():
			return
		}
	}
}

func (l *Loop) see_SynAck() {
	sink := make(chan *contracts.ContractASynAck)
	sub, err := l.a.WatchSynAck(&bind.WatchOpts{Context:l.ctx}, sink, nil)
	if err != nil { log.Fatalf("got error while watching synack: %v", err) }
	log.Println("Listening initiated for SynAck event...")

	for {
		select {
		case ev := <-sink:
			log.Printf("SynAck event session=%d -> Confirm on B", ev.SessionId)
			txn, err := l.b.Confirm(l.auth_b, ev.SessionId)
			if err != nil {
				log.Printf("confirm err: %v", err); continue
			}
			log.Printf("sent txn %s", txn.Hash())
		case err := <-sub.Err():
			log.Printf("got a synack sub err: %v", err)
			return
		case <-l.ctx.Done():
			return
		}
	}
}
func (l *Loop) see_Ack() {
	sink2 := make(chan *contracts.ContractBAck)
	sub2, err := l.b.WatchAck(&bind.WatchOpts{Context:l.ctx}, sink2, nil)
	if err != nil { log.Fatalf("watching ack: %v", err) }
	log.Println("listening for Ack…")

	for {
		select {
	case ev := <-sink2:
			log.Printf("Ack session=%d -> SYN-ACK on A", ev.SessionId)
		txn, err := l.a.Acknowledge(l.auth_a, ev.SessionId)
			 if err != nil {
				log.Printf("syn-ack err: %v", err); continue
			}
			log.Printf("sent tx %s", txn.Hash())
		case err := <-sub2.Err():
			log.Printf("ack sub err: %v", err)
			return
		case <-l.ctx.Done():
			return
		}
	}
}


