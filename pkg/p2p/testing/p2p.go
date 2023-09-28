package p2ptest

import (
	"context"
	"errors"

	"github.com/primevprotocol/mev-commit/pkg/p2p"
)

func NewDuplexStream() (out *testStream, in *testStream) {
	out = newStream()
	in = newStream()
	pipe(out, in)
	return
}

type testStream struct {
	in  chan []byte
	out chan []byte
}

func newStream() *testStream {
	return &testStream{
		in:  make(chan []byte, 8),
		out: make(chan []byte, 8),
	}
}

func (s *testStream) ReadMsg() ([]byte, error) {
	return <-s.in, nil
}

func (s *testStream) WriteMsg(msg []byte) error {
	s.out <- msg
	return nil
}

func (s *testStream) Close() error {
	close(s.out)
	return nil
}

func (s *testStream) Reset() error {
	close(s.out)
	return nil
}

func pipe(a, b *testStream) {
	go func() {
		for {
			msg, ok := <-a.out
			if !ok {
				return
			}
			b.in <- msg
		}
	}()

	go func() {
		for {
			msg, ok := <-b.out
			if !ok {
				return
			}
			a.in <- msg
		}
	}()
}

type Option func(*P2PTest)

type P2PTest struct {
	handlers        map[string]p2p.ProtocolSpec
	connectFunc     func([]byte) (p2p.Peer, error)
	addressbookFunc func(p2p.Peer) ([]byte, error)
}

func WithConnectFunc(fn func([]byte) (p2p.Peer, error)) Option {
	return func(p *P2PTest) {
		p.connectFunc = fn
	}
}

func WithAddressbookFunc(fn func(p p2p.Peer) ([]byte, error)) Option {
	return func(p *P2PTest) {
		p.addressbookFunc = fn
	}
}

func New(opts ...Option) *P2PTest {
	p := &P2PTest{
		handlers: make(map[string]p2p.ProtocolSpec),
	}

	for _, opt := range opts {
		opt(p)
	}

	return p
}

func (p *P2PTest) SetPeerHandler(peer p2p.Peer, proto p2p.ProtocolSpec) {
	p.handlers[peer.EthAddress.Hex()] = proto
}

func (p *P2PTest) Connect(_ context.Context, addr []byte) (p2p.Peer, error) {
	if p.connectFunc != nil {
		return p.connectFunc(addr)
	}

	return p2p.Peer{}, errors.New("connect not implemented")
}

func (p *P2PTest) GetPeerInfo(peer p2p.Peer) ([]byte, error) {
	if p.addressbookFunc != nil {
		return p.addressbookFunc(peer)
	}

	return nil, errors.New("addressbook not implemented")
}

func (p *P2PTest) NewStream(
	_ context.Context,
	peer p2p.Peer,
	proto, version, stream string,
) (p2p.Stream, error) {
	sHandlers, found := p.handlers[peer.EthAddress.Hex()]
	if !found {
		return nil, errors.New("peer not found")
	}

	var handler p2p.Handler
	for _, h := range sHandlers.StreamSpecs {
		if h.Name == stream {
			handler = h.Handler
			break
		}
	}

	if handler == nil {
		return nil, errors.New("stream not found")
	}

	out, in := NewDuplexStream()

	go func() {
		defer in.Close()

		err := handler(context.Background(), peer, in)
		if err != nil {
			panic(err)
		}
	}()

	return out, nil
}