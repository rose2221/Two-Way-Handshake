package relay

import (
	"context"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	contracts "github.com/you/handshake-relayer/contracts"
	"../internal/eth"
	"github.com/you/handshake-relayer/internal/persistence"
)

type Loop struct {
	ctx        context.Context
	cancel     context.CancelFunc
	db         *persistence.LevelDB
	a, b       *contracts.ContractA      // bindings *on their native chains*
	authA, authB *bind.TransactOpts
}

func NewLoop(
	ctx context.Context,
	db *persistence.LevelDB,
	a *contracts.ContractA, b *contracts.ContractB,
	authA, authB *bind.TransactOpts,
) *Loop {
	c, cancel := context.WithCancel(ctx)
	return &Loop{ctx:c, cancel:cancel, db:db, a:a, b:b, authA:authA, authB:authB}
}

func (l *Loop) Start() {
	go l.watchSyn()
	go l.watchAck()
	go l.watchSynAck()
	<-l.ctx.Done()
}

/*────────── event → tx pipelines ─────────*/

func (l *Loop) watchSyn() {
	sink := make(chan *contracts.ContractASyn)
	sub, err := l.a.WatchSyn(&bind.WatchOpts{Context:l.ctx}, sink, nil)
	if err != nil { log.Fatalf("watch syn: %v", err) }
	log.Println("Listening for Syn…")

	for {
		select {
		case ev := <-sink:
			log.Printf("[Syn] session=%d ↦ ACK on B", ev.SessionId)
			tx, err := l.b.Acknowledge(l.authB, ev.SessionId)
			if err != nil {
				log.Printf("ack tx err: %v", err); continue
			}
			log.Printf("…sent tx %s", tx.Hash())
		case err := <-sub.Err():
			log.Printf("syn sub err: %v", err)
			return
		case <-l.ctx.Done():
			return
		}
	}
}

func (l *Loop) watchAck() {
	sink := make(chan *contracts.ContractBAck)
	sub, err := l.b.WatchAck(&bind.WatchOpts{Context:l.ctx}, sink, nil)
	if err != nil { log.Fatalf("watch ack: %v", err) }
	log.Println("Listening for Ack…")

	for {
		select {
		case ev := <-sink:
			log.Printf("[Ack] session=%d ↦ SYN-ACK on A", ev.SessionId)
			tx, err := l.a.Acknowledge(l.authA, ev.SessionId)
			if err != nil {
				log.Printf("syn-ack err: %v", err); continue
			}
			log.Printf("…sent tx %s", tx.Hash())
		case err := <-sub.Err():
			log.Printf("ack sub err: %v", err)
			return
		case <-l.ctx.Done():
			return
		}
	}
}

func (l *Loop) watchSynAck() {
	sink := make(chan *contracts.ContractASynAck)
	sub, err := l.a.WatchSynAck(&bind.WatchOpts{Context:l.ctx}, sink, nil)
	if err != nil { log.Fatalf("watch synack: %v", err) }
	log.Println("Listening for SynAck…")

	for {
		select {
		case ev := <-sink:
			log.Printf("[SynAck] session=%d ↦ Confirm on B", ev.SessionId)
			tx, err := l.b.Confirm(l.authB, ev.SessionId)
			if err != nil {
				log.Printf("confirm err: %v", err); continue
			}
			log.Printf("…sent tx %s", tx.Hash())
		case err := <-sub.Err():
			log.Printf("synack sub err: %v", err)
			return
		case <-l.ctx.Done():
			return
		}
	}
}
