package bucket

import (
	"container/list"
	"github.com/ndf1d/easydht/id"
	"github.com/ndf1d/easydht/peer"
	"sync"
)

// Bucket contains list of k-buckets.
type Bucket struct {
	list *list.List
	lock sync.RWMutex
}

// New returns new Bucket.
func New() *Bucket {
	b := new(Bucket)
	b.list = list.New()
	return b
}

// Has checks if bucket has or not id.
func (b *Bucket) Has(id id.ID) bool {
	b.lock.RLock()
	defer b.lock.RUnlock()
	for e := b.list.Front(); e != nil; e = e.Next() {
		if id.Equals(e.Value.(*peer.Peer).ID) {
			return true
		}
	}
	return false
}

func (b *Bucket) listElement(id id.ID) (*list.Element, bool) {
	b.lock.Lock()
	defer b.lock.Unlock()
	var p *peer.Peer
	for e := b.list.Front(); e != nil; e = e.Next() {
		p = e.Value.(*peer.Peer)
		if id.Equals(p.ID) {
			return e, true
		}
	}
	return nil, false
}

// Remove removes id from a bucket.
func (b *Bucket) Remove(id id.ID) {
	b.lock.Lock()
	defer b.lock.Unlock()
	if elem, ok := b.listElement(id); ok {
		b.list.Remove(elem)
	}
}

// MoveToFront moves peer to front of a k-bucket.
func (b *Bucket) MoveToFront(id id.ID) {
	b.lock.Lock()
	defer b.lock.Unlock()
	if elem, ok := b.listElement(id); ok {
		b.list.MoveToFront(elem)
	}
}

// PushFront inserts peer at the front of k-bucket.
func (b *Bucket) PushFront(p *peer.Peer) {
	b.lock.Lock()
	b.list.PushFront(p)
	b.lock.Unlock()
}

// PopBack removes last peer in the k-bucket.
func (b *Bucket) PopBack() *peer.Peer {
	b.lock.Lock()
	defer b.lock.Unlock()
	last := b.list.Back()
	b.list.Remove(last)
	p := last.Value.(peer.Peer)
	return &p
}

// Len returns k-bucket list
func (b *Bucket) Len() int {
	b.lock.RLock()
	defer b.lock.RUnlock()
	return b.list.Len()
}

// Front returns first peer in the k-bucket.
func (b *Bucket) Front() *list.Element {
	b.lock.RLock()
	defer b.lock.RUnlock()
	return b.list.Front()
}

// Peers returns k-bucket peers.
func (b *Bucket) Peers() []*peer.Peer {
	var peers []*peer.Peer
	var p *peer.Peer
	for e := b.list.Front(); e != nil; e = e.Next() {
		p = e.Value.(*peer.Peer)
		peers = append(peers, p)
	}
	return peers
}
