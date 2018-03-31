package plasma

import (
	"bytes"
	"crypto/ecdsa"
	"fmt"
	"net"
	"sync"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/p2p/discover"
	"github.com/ethereum/go-ethereum/p2p/nat"
)

const NumNodes = 16 // must not exceed the number of keys (32)

type TestNode struct {
	pls    *Plasma
	server *p2p.Server
}

var result TestData
var nodes [NumNodes]*TestNode

// This test does the following:
// 1. creates a chain of whisper nodes,
// 2. installs the filters with shared (predefined) parameters,
// 3. each node sends a number of random (undecryptable) messages,
// 4. first node sends one expected (decryptable) message,
// 5. checks if each node have received and decrypted exactly one message.
func TestSimulation(t *testing.T) {
	initialize(t)

	for i := 0; i < NumNodes; i++ {
		sendMsg(t, false, i)
	}

	sendMsg(t, true, 0)
	checkPropagation(t)
	stopServers()
}

func initialize(t *testing.T) {
	var err error
	ip := net.IPv4(127, 0, 0, 1)
	port0 := 30303

	for i := 0; i < NumNodes; i++ {
		var node TestNode
		node.shh = New(&DefaultConfig)
		node.shh.SetMinimumPoW(0.00000001)
		node.shh.Start(nil)
		topics := make([]TopicType, 0)
		topics = append(topics, sharedTopic)
		f := Filter{KeySym: sharedKey}
		f.Topics = [][]byte{topics[0][:]}
		node.filerId, err = node.shh.Subscribe(&f)
		if err != nil {
			t.Fatalf("failed to install the filter: %s.", err)
		}
		node.id, err = crypto.HexToECDSA(keys[i])
		if err != nil {
			t.Fatalf("failed convert the key: %s.", keys[i])
		}
		port := port0 + i
		addr := fmt.Sprintf(":%d", port) // e.g. ":30303"
		name := common.MakeName("whisper-go", "2.0")
		var peers []*discover.Node
		if i > 0 {
			peerNodeId := nodes[i-1].id
			peerPort := uint16(port - 1)
			peerNode := discover.PubkeyID(&peerNodeId.PublicKey)
			peer := discover.NewNode(peerNode, ip, peerPort, peerPort)
			peers = append(peers, peer)
		}

		node.server = &p2p.Server{
			Config: p2p.Config{
				PrivateKey:     node.id,
				MaxPeers:       NumNodes/2 + 1,
				Name:           name,
				Protocols:      node.shh.Protocols(),
				ListenAddr:     addr,
				NAT:            nat.Any(),
				BootstrapNodes: peers,
				StaticNodes:    peers,
				TrustedNodes:   peers,
			},
		}

		err = node.server.Start()
		if err != nil {
			t.Fatalf("failed to start server %d.", i)
		}

		nodes[i] = &node
	}
}

func stopServers() {
	for i := 0; i < NumNodes; i++ {
		n := nodes[i]
		if n != nil {
			n.shh.Unsubscribe(n.filerId)
			n.shh.Stop()
			n.server.Stop()
		}
	}
}

func checkPropagation(t *testing.T) {
	if t.Failed() {
		return
	}

	const cycle = 100
	const iterations = 100

	for j := 0; j < iterations; j++ {
		time.Sleep(cycle * time.Millisecond)

		for i := 0; i < NumNodes; i++ {
			f := nodes[i].shh.GetFilter(nodes[i].filerId)
			if f == nil {
				t.Fatalf("failed to get filterId %s from node %d.", nodes[i].filerId, i)
			}

			mail := f.Retrieve()
			if !validateMail(t, i, mail) {
				return
			}

			if isTestComplete() {
				return
			}
		}
	}

	t.Fatalf("Test was not complete: timeout %d seconds.", iterations*cycle/1000)
}

func validateMail(t *testing.T, index int, mail []*ReceivedMessage) bool {
	var cnt int
	for _, m := range mail {
		if bytes.Equal(m.Payload, expectedMessage) {
			cnt++
		}
	}

	if cnt == 0 {
		// no messages received yet: nothing is wrong
		return true
	}
	if cnt > 1 {
		t.Fatalf("node %d received %d.", index, cnt)
		return false
	}

	if cnt > 0 {
		result.mutex.Lock()
		defer result.mutex.Unlock()
		result.counter[index] += cnt
		if result.counter[index] > 1 {
			t.Fatalf("node %d accumulated %d.", index, result.counter[index])
		}
	}
	return true
}

func isTestComplete() bool {
	result.mutex.RLock()
	defer result.mutex.RUnlock()

	for i := 0; i < NumNodes; i++ {
		if result.counter[i] < 1 {
			return false
		}
	}

	for i := 0; i < NumNodes; i++ {
		envelopes := nodes[i].shh.Envelopes()
		if len(envelopes) < 2 {
			return false
		}
	}

	return true
}

func sendMsg(t *testing.T, expected bool, id int) {
	if t.Failed() {
		return
	}

	opt := MessageParams{KeySym: sharedKey, Topic: sharedTopic, Payload: expectedMessage, PoW: 0.00000001, WorkTime: 1}
	if !expected {
		opt.KeySym[0]++
		opt.Topic[0]++
		opt.Payload = opt.Payload[1:]
	}

	msg, err := NewSentMessage(&opt)
	if err != nil {
		t.Fatalf("failed to create new message with seed %d: %s.", seed, err)
	}
	envelope, err := msg.Wrap(&opt)
	if err != nil {
		t.Fatalf("failed to seal message: %s", err)
	}

	err = nodes[id].shh.Send(envelope)
	if err != nil {
		t.Fatalf("failed to send message: %s", err)
	}
}

func TestPeerBasic(t *testing.T) {
	InitSingleTest()

	params, err := generateMessageParams()
	if err != nil {
		t.Fatalf("failed generateMessageParams with seed %d.", seed)
	}

	params.PoW = 0.001
	msg, err := NewSentMessage(params)
	if err != nil {
		t.Fatalf("failed to create new message with seed %d: %s.", seed, err)
	}
	env, err := msg.Wrap(params)
	if err != nil {
		t.Fatalf("failed Wrap with seed %d.", seed)
	}

	p := newPeer(nil, nil, nil)
	p.mark(env)
	if !p.marked(env) {
		t.Fatalf("failed mark with seed %d.", seed)
	}
}
