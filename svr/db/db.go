package db

import "sync"

var Set = boltSet
var Get = boltGet
var Delete = boltDelete
var writeLock sync.Mutex
