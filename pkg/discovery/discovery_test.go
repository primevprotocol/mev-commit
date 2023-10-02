package discovery_test

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/primevprotocol/mev-commit/pkg/discovery"
	"github.com/primevprotocol/mev-commit/pkg/p2p"
	p2ptest "github.com/primevprotocol/mev-commit/pkg/p2p/testing"
)

type testTopo struct {
	mu    sync.Mutex
	peers []p2p.Peer
}

func (t *testTopo) AddPeers(peers ...p2p.Peer) {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.peers = append(t.peers, peers...)
}

func (t *testTopo) IsConnected(addr common.Address) bool {
	t.mu.Lock()
	defer t.mu.Unlock()

	for _, p := range t.peers {
		if p.EthAddress == addr {
			return true
		}
	}
	return false
}

func newTestLogger(w io.Writer) *slog.Logger {
	testLogger := slog.NewTextHandler(w, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})
	return slog.New(testLogger)
}

func TestDiscovery(t *testing.T) {
	t.Parallel()

	t.Run("ok", func(t *testing.T) {
		client := p2p.Peer{
			EthAddress: common.HexToAddress("0x1"),
			Type:       p2p.PeerTypeBuilder,
		}
		server := p2p.Peer{
			EthAddress: common.HexToAddress("0x2"),
			Type:       p2p.PeerTypeBuilder,
		}

		svc := p2ptest.New(
			&client,
			p2ptest.WithConnectFunc(func(addr []byte) (p2p.Peer, error) {
				if string(addr) != "test" {
					return p2p.Peer{}, errors.New("invalid address")
				}
				return client, nil
			}),
		)

		topo := &testTopo{}
		d := discovery.New(topo, svc, newTestLogger(os.Stdout))
		t.Cleanup(func() {
			err := d.Close()
			if err != nil {
				t.Fatal(err)
			}
		})

		svc.SetPeerHandler(server, d.Protocol())

		err := d.BroadcastPeers(context.Background(), server, []p2p.PeerInfo{
			{
				EthAddress: common.HexToAddress("0x1"),
				Underlay:   []byte("test"),
			},
		})
		if err != nil {
			t.Fatal(err)
		}

		start := time.Now()
		for {
			if time.Since(start) > 5*time.Second {
				t.Fatal("timed out")
			}
			if len(topo.peers) == 1 {
				break
			}
			time.Sleep(100 * time.Millisecond)
		}
	})
}
