package shared

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"sync/atomic"
	"time"

	"code.google.com/p/go-uuid/uuid"
)

var (
	EmptyId Id
	seq     uint32
	node    = nodeId()
)

type Id [16]byte

func nodeId() uint32 {
	n := uuid.NodeID()
	return binary.BigEndian.Uint32(n)
}

func NewId() Id {
	var uuid [16]byte
	var now = time.Now().UTC()

	nano := now.UnixNano()
	incr := atomic.AddUint32(&seq, 1)

	binary.BigEndian.PutUint64(uuid[0:], uint64(nano))
	binary.BigEndian.PutUint32(uuid[8:], incr)
	binary.BigEndian.PutUint32(uuid[12:], node)

	return uuid
}

func (id Id) Bytes() []byte {
	return id[:]
}

func (id Id) Time() time.Time {
	bytes := id[:]
	nsec := binary.BigEndian.Uint64(bytes)
	return time.Unix(0, int64(nsec)).UTC()
}

func (s Id) Equals(other Id) (eq bool) {
	eq = bytes.Equal(s.Bytes(), other.Bytes())
	return
}

func (this Id) String() string {
	return hex.EncodeToString(this[:])
}
