package shared

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"strings"
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

func (id Id) String() string {
	return hex.EncodeToString(id[:])
}
func ParseId(value string) (id Id, err error) {
	if len(value) == 0 {
		err = fmt.Errorf("Invalid id: value is empty")
		return
	}

	var b []byte
	orgValue := value

	if len(value) != 32 {
		value = strings.Map(func(r rune) rune {
			if r == '-' || r == '{' || r == '}' {
				return -1
			}
			return r
		}, value)
	}

	if b, err = hex.DecodeString(value); err != nil {
		err = fmt.Errorf("invalid id %v: %v", orgValue, err.Error())
		return
	}

	if len(b) != 16 {
		err = fmt.Errorf("invalid id %v: did not convert to a 16 byte array", orgValue)
		return
	}

	for index, value := range b {
		id[index] = value
	}

	return
}

// JSON marshalling
func (id Id) MarshalJSON() ([]byte, error) {

	jsonString := `"` + hex.EncodeToString(id[:]) + `"`
	return []byte(jsonString), nil
}

func (id *Id) UnmarshalJSON(data []byte) error {
	jsonString := string(data)
	valueString := strings.Trim(jsonString, "\"")

	value, err := ParseId(valueString)
	if err != nil {
		return err
	}

	*id = value
	return nil
}
