package bucket

import (
	"github.com/ndf1d/easydht/id"
	"github.com/ndf1d/easydht/peer"
	"testing"
)

var testID = id.ID{0, 0, 1}

func TestBucket_Has(t *testing.T) {
	b := New()
	if b.Has(testID) {
		t.FailNow()
	}
	c := &peer.Peer{
		ID: testID,
	}
	b.PushFront(c)
	if b.list.Len() != 1 {
		t.Fail()
	}
	if !b.Has(testID) {
		t.Fail()
	}
}
