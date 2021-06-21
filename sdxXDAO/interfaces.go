// Copyright 2019 The Tomochain Authors
// This file is part of the Core Tomochain infrastructure
// https://tomochain.com
// Package sdxxDAO provides an interface to work with sdxx database, including leveldb for masternode and mongodb for SDK node
package sdxxDAO

import (
	"github.com/tomochain/tomochain/common"
	"github.com/tomochain/tomochain/ethdb"
)

const defaultCacheLimit = 1024

type sdxXDAO interface {
	// for both leveldb and mongodb
	IsEmptyKey(key []byte) bool
	Close() error

	// mongodb methods
	HasObject(hash common.Hash, val interface{}) (bool, error)
	GetObject(hash common.Hash, val interface{}) (interface{}, error)
	PutObject(hash common.Hash, val interface{}) error
	DeleteObject(hash common.Hash, val interface{}) error // won't return error if key not found
	GetListItemByTxHash(txhash common.Hash, val interface{}) interface{}
	GetListItemByHashes(hashes []string, val interface{}) interface{}
	DeleteItemByTxHash(txhash common.Hash, val interface{})

	// basic sdxx
	InitBulk()
	CommitBulk() error

	// sdxx lending
	InitLendingBulk()
	CommitLendingBulk() error

	// leveldb methods
	Put(key []byte, value []byte) error
	Get(key []byte) ([]byte, error)
	Has(key []byte) (bool, error)
	Delete(key []byte) error
	NewBatch() ethdb.Batch
	HasAncient(kind string, number uint64) (bool, error)
	Ancient(kind string, number uint64) ([]byte, error)
	Ancients() (uint64, error)
	AncientSize(kind string) (uint64, error)
	AppendAncient(number uint64, hash, header, body, receipt, td []byte) error
	TruncateAncients(n uint64) error
	Sync() error
	NewIterator(prefix []byte, start []byte) ethdb.Iterator

	Stat(property string) (string, error)
	Compact(start []byte, limit []byte) error
}

// use alloc to prevent reference manipulation
func EmptyKey() []byte {
	key := make([]byte, common.HashLength)
	return key
}
