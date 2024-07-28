package freecache

import (
	"github.com/coocood/freecache"
	"github.com/senizdegen/sdu-housing/api-gateway/pkg/cache"
)

type iterator struct {
	iter *freecache.Iterator
}

func (i *iterator) Next() *cache.Entry {
	entry := i.iter.Next()
	if entry == nil {
		return nil
	}

	return &cache.Entry{
		Key:   entry.Key,
		Value: entry.Value,
	}
}
