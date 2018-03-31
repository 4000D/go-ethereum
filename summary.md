# go-ethereum 주요 구조체
- `cmd/geth/main.go` : `geth` 명령어 main package
- `node/node.go` : `type Node struct {...}`
  - ethereum 풀 노드 구현체
  - `Node.server`: p2p 통신을 위한 서버 (`p2p/server.go`)
  - `Node.services`: full node 에서 관리하는 서비스들
    - `eth/backend.go` : `type Ethereum struct {...}`
      - ethereum full node 서비스.
    - `les/backend.go`: `type LightEthereum strcut {...}`
      - light client node 서비스.
    - `whisper/whisperv5/whisper.go`: `type Whisper struct {...}`
      - shh 서비스
- `p2p/server.go`: `type Server struct {...}`
  - p2p 서버 구현체. `node.Node.server` 로 사용됨
  - `p2p/peer.go`: `type Peer struct {...}`
    - p2p 통신에 관한 peer 구조체.
    - *prptocol peer* 는 별도로 구현되어있음.
    - *protocol peer* 는 `p2p/rlpx.go` 의 `func (t *rlpx) doProtoHandshake(our *protoHandshake)` 를 통해 handshake
    - `eth/peer.go`: `type peer struct {...}`
      - eth 서비스 피어
    - `les/peer.go`: `type peer struct {...}`
      - les 서비스 피어
    - `whisper/whisperv5/peer.go`: `type Peer struct {...}`
      - ssh 서비스 피어

## entrypoint of command `geth`
### cmd/geth/main.go
```golang
func geth(ctx *cli.Context) error {
	node := makeFullNode(ctx)
	startNode(ctx, node)
	node.Wait()
	return nil
}
```

# Full node implementation
### node/node.go : Full node implementation
```golang
// Node is a container on which services can be registered.
type Node struct {
	accman   *accounts.Manager

	server       *p2p.Server // Currently running P2P networking layer

	serviceFuncs []ServiceConstructor     // Service constructors (in dependency order)
	services     map[reflect.Type]Service // Currently running services

	rpcAPIs       []rpc.API   // List of APIs currently provided by the node
	inprocHandler *rpc.Server // In-process RPC request handler to process the API requests

	ipcEndpoint string       // IPC endpoint to listen at (empty = IPC disabled)
	ipcListener net.Listener // IPC RPC listener socket to serve API requests
	ipcHandler  *rpc.Server  // IPC RPC request handler to process the API requests
}
```

# node.Node.services
### eth/backend.go : ethereum service as node.Node.services
```golang
// Ethereum implements the Ethereum full node service.
type Ethereum struct {
	config      *Config
	chainConfig *params.ChainConfig

	// Handlers
	txPool          *core.TxPool
	blockchain      *core.BlockChain
	protocolManager *ProtocolManager
	lesServer       LesServer

	// DB interfaces
	chainDb ethdb.Database // Block chain database

	engine         consensus.Engine
	accountManager *accounts.Manager

	bloomRequests chan chan *bloombits.Retrieval // Channel receiving bloom data retrieval requests
	bloomIndexer  *core.ChainIndexer             // Bloom indexer operating during block imports

	ApiBackend *EthApiBackend

	miner     *miner.Miner
	gasPrice  *big.Int
	etherbase common.Address

	networkId     uint64
	netRPCService *ethapi.PublicNetAPI
}
```

### les/backend.go : lesServer as eth.Ethereum.lesServer or ethereum service as node.Node.services
```golang
type LightEthereum struct {
	config *eth.Config

	chainConfig *params.ChainConfig

	// Handlers
	peers           *peerSet
	txPool          *light.TxPool
	blockchain      *light.LightChain
	protocolManager *ProtocolManager
	serverPool      *serverPool
	reqDist         *requestDistributor
	retriever       *retrieveManager

	// DB interfaces
	chainDb ethdb.Database // Block chain database

	bloomRequests                              chan chan *bloombits.Retrieval // Channel receiving bloom data retrieval requests
	bloomIndexer, chtIndexer, bloomTrieIndexer *core.ChainIndexer

	ApiBackend *LesApiBackend

	engine         consensus.Engine
	accountManager *accounts.Manager

	networkId     uint64
	netRPCService *ethapi.PublicNetAPI
}
```

### whisper/whisperv5/whisper.go : whisper service as node.Node.services
```golang
// Whisper represents a dark communication interface through the Ethereum
// network, using its very own P2P communication layer.
type Whisper struct {
	protocol p2p.Protocol // Protocol description and parameters
	filters  *Filters     // Message filters installed with Subscribe function

	privateKeys map[string]*ecdsa.PrivateKey // Private key storage
	symKeys     map[string][]byte            // Symmetric key storage
	keyMu       sync.RWMutex                 // Mutex associated with key storages

	poolMu      sync.RWMutex              // Mutex to sync the message and expiration pools
	envelopes   map[common.Hash]*Envelope // Pool of envelopes currently tracked by this node
	expirations map[uint32]*set.SetNonTS  // Message expiration pool

	peerMu sync.RWMutex       // Mutex to sync the active peer set
	peers  map[*Peer]struct{} // Set of currently active peers

	messageQueue chan *Envelope // Message queue for normal whisper messages
	p2pMsgQueue  chan *Envelope // Message queue for peer-to-peer messages (not to be forwarded any further)
	quit         chan struct{}  // Channel used for graceful exit

	settings syncmap.Map // holds configuration settings that can be dynamically changed

	statsMu sync.Mutex // guard stats
	stats   Statistics // Statistics of whisper node

	mailServer MailServer // MailServer interface
}
```

### ethstats/ethstats.go : ethstat service as node.Node.services
```golang
// Service implements an Ethereum netstats reporting daemon that pushes local
// chain statistics up to a monitoring server.
type Service struct {
	server *p2p.Server        // Peer-to-peer server to retrieve networking infos
	eth    *eth.Ethereum      // Full Ethereum service if monitoring a full node
	les    *les.LightEthereum // Light Ethereum service if monitoring a light node
	engine consensus.Engine   // Consensus engine to retrieve variadic block fields

	node string // Name of the node to display on the monitoring page
	pass string // Password to authorize access to the monitoring page
	host string // Remote address of the monitoring service

	pongCh chan struct{} // Pong notifications are fed into this channel
	histCh chan []uint64 // History request block numbers are fed into this channel
}
```

# node.Node.server related
### p2p/server.go : p2p server implementation for full node
```golang
// Config holds Server options.
type Config struct {
	// This field must be set to a valid secp256k1 private key.
	PrivateKey *ecdsa.PrivateKey `toml:"-"`

	// MaxPeers is the maximum number of peers that can be
	// connected. It must be greater than zero.
	MaxPeers int

	// MaxPendingPeers is the maximum number of peers that can be pending in the
	// handshake phase, counted separately for inbound and outbound connections.
	// Zero defaults to preset values.
	MaxPendingPeers int `toml:",omitempty"`

	// DialRatio controls the ratio of inbound to dialed connections.
	// Example: a DialRatio of 2 allows 1/2 of connections to be dialed.
	// Setting DialRatio to zero defaults it to 3.
	DialRatio int `toml:",omitempty"`

	// NoDiscovery can be used to disable the peer discovery mechanism.
	// Disabling is useful for protocol debugging (manual topology).
	NoDiscovery bool

	// DiscoveryV5 specifies whether the the new topic-discovery based V5 discovery
	// protocol should be started or not.
	DiscoveryV5 bool `toml:",omitempty"`

	// Name sets the node name of this server.
	// Use common.MakeName to create a name that follows existing conventions.
	Name string `toml:"-"`

	// BootstrapNodes are used to establish connectivity
	// with the rest of the network.
	BootstrapNodes []*discover.Node

	// BootstrapNodesV5 are used to establish connectivity
	// with the rest of the network using the V5 discovery
	// protocol.
	BootstrapNodesV5 []*discv5.Node `toml:",omitempty"`

	// Static nodes are used as pre-configured connections which are always
	// maintained and re-connected on disconnects.
	StaticNodes []*discover.Node

	// Trusted nodes are used as pre-configured connections which are always
	// allowed to connect, even above the peer limit.
	TrustedNodes []*discover.Node

	// Connectivity can be restricted to certain IP networks.
	// If this option is set to a non-nil value, only hosts which match one of the
	// IP networks contained in the list are considered.
	NetRestrict *netutil.Netlist `toml:",omitempty"`

	// NodeDatabase is the path to the database containing the previously seen
	// live nodes in the network.
	NodeDatabase string `toml:",omitempty"`

	// Protocols should contain the protocols supported
	// by the server. Matching protocols are launched for
	// each peer.
	Protocols []Protocol `toml:"-"`

	// If ListenAddr is set to a non-nil address, the server
	// will listen for incoming connections.
	//
	// If the port is zero, the operating system will pick a port. The
	// ListenAddr field will be updated with the actual address when
	// the server is started.
	ListenAddr string

	// If set to a non-nil value, the given NAT port mapper
	// is used to make the listening port available to the
	// Internet.
	NAT nat.Interface `toml:",omitempty"`

	// If Dialer is set to a non-nil value, the given Dialer
	// is used to dial outbound peer connections.
	Dialer NodeDialer `toml:"-"`

	// If NoDial is true, the server will not dial any peers.
	NoDial bool `toml:",omitempty"`

	// If EnableMsgEvents is set then the server will emit PeerEvents
	// whenever a message is sent to or received from a peer
	EnableMsgEvents bool

	// Logger is a custom logger to use with the p2p.Server.
	Logger log.Logger `toml:",omitempty"`
}

// Server manages all peer connections.
type Server struct {
	// Config fields may not be modified while the server is running.
	Config

	// Hooks for testing. These are useful because we can inhibit
	// the whole protocol stack.
	newTransport func(net.Conn) transport
	newPeerHook  func(*Peer)

	lock    sync.Mutex // protects running
	running bool

	ntab         discoverTable
	listener     net.Listener
	ourHandshake *protoHandshake
	lastLookup   time.Time
	DiscV5       *discv5.Network

	// These are for Peers, PeerCount (and nothing else).
	peerOp     chan peerOpFunc
	peerOpDone chan struct{}

	quit          chan struct{}
	addstatic     chan *discover.Node
	removestatic  chan *discover.Node
	posthandshake chan *conn
	addpeer       chan *conn
	delpeer       chan peerDrop
	loopWG        sync.WaitGroup // loop, listenLoop
	peerFeed      event.Feed
	log           log.Logger
}
```

### p2p/peer.go
> geth p2p peer 구조체. p2p.discovery 를 통해서 노드 탐색

```golang
// Peer represents a connected remote node.
type Peer struct {
	rw      *conn // 해당 피어와 데이터 전송(read / write) 위한 connection
	running map[string]*protoRW
	log     log.Logger
	created mclock.AbsTime

	wg       sync.WaitGroup
	protoErr chan error
	closed   chan struct{}
	disc     chan DiscReason

	// events receives message send / receive events if set
	events *event.Feed
}
```

### eth/peer.go
> eth 프로토콜 peer

```golang
type peer struct {
	id string

	*p2p.Peer
	rw p2p.MsgReadWriter

	version  int         // Protocol version negotiated
	forkDrop *time.Timer // Timed connection dropper if forks aren't validated in time

	head common.Hash
	td   *big.Int
	lock sync.RWMutex

	knownTxs    *set.Set // Set of transaction hashes known to be known by this peer
	knownBlocks *set.Set // Set of block hashes known to be known by this peer
}
```

### les/peer.go
> les 프로토콜 peer

```golang
type peer struct {
	*p2p.Peer
	pubKey *ecdsa.PublicKey

	rw p2p.MsgReadWriter

	version int    // Protocol version negotiated
	network uint64 // Network ID being on

	announceType, requestAnnounceType uint64

	id string

	headInfo *announceData
	lock     sync.RWMutex

	announceChn chan announceData
	sendQueue   *execQueue

	poolEntry      *poolEntry
	hasBlock       func(common.Hash, uint64) bool
	responseErrors int

	fcClient       *flowcontrol.ClientNode // nil if the peer is server only
	fcServer       *flowcontrol.ServerNode // nil if the peer is client only
	fcServerParams *flowcontrol.ServerParams
	fcCosts        requestCostTable
}
```

### whisper/whisperv5/peer.go
> shh 프로토콜 peer

```golang
// peer represents a whisper protocol peer connection.
type Peer struct {
	host    *Whisper
	peer    *p2p.Peer
	ws      p2p.MsgReadWriter
	trusted bool

	known *set.Set // Messages already known by the peer to avoid wasting bandwidth

	quit chan struct{}
}
```
