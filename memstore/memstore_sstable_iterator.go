package memstore

import (
	"errors"
	"github.com/thomasjungblut/go-sstables/v2/skiplist"
	"github.com/thomasjungblut/go-sstables/v2/sstables"
)

type SkipListSStableIterator struct {
	iterator skiplist.IteratorI[[]byte, ValueStruct]
}

func (s SkipListSStableIterator) Next() ([]byte, []byte, error) {
	key, val, err := s.iterator.Next()
	if err != nil {
		if errors.Is(err, skiplist.Done) {
			return nil, nil, sstables.Done
		} else {
			return nil, nil, err
		}
	}
	return key, *val.value, nil
}
