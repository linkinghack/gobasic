package main

import (
	"sync"
	"io"
	"log"
	"sync/atomic"
	"./pool"
)

var idCounter int32

func createConnection() (io.Closer, error) {
	id := atomic.AddInt32(&idCounter, 1)
	log.Println("Create: New Connection", id)

	return &dbConnection{id}, nil
}

var wg sync.WaitGroup
wd.Add(maxGoroutines)

p, err := pool.New(createConnection, pooledResources)
if err != nil {
	log.Println(err)
}

for query := 0;