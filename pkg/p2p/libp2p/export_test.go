package libp2p

import (
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/primevprotocol/mev-commit/pkg/p2p"
)

func (s *Service) Addrs() ([]byte, error) {
	info := s.host.Peerstore().PeerInfo(s.host.ID())
	return info.MarshalJSON()
}

func (s *Service) Peer() p2p.Peer {
	return p2p.Peer{
		EthAddress: s.ethAddress,
		Type:       s.peerType,
	}
}

func (s *Service) HostID() peer.ID {
	return s.host.ID()
}