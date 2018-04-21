package peer

import (
	"github.com/ndf1d/easydht/id"
	"strings"
)

// Peer contains peer information.
type Peer struct {
	ID   id.ID
	Host string
	Port string
}

// Address returns string address representation.
func (ni *Peer) Address() string {
	return strings.Join([]string{ni.Host, ":", ni.Port}, "")
}

// LocalPeer returns peer with public ip.
func LocalPeer(localPeerID id.ID, port string) (*Peer, error) {
	externalIP, err := GetMyExternalIP()
	if err != nil {
		return nil, err
	}
	return &Peer{
		ID:   localPeerID,
		Host: externalIP,
		Port: port,
	}, nil
}
