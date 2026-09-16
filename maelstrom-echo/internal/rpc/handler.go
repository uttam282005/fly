package rpc

import "sync"
import "log"

type globalCounter struct {
	count int64
	mu sync.Mutex
}

var gc globalCounter

func(l *LocalService) GenerateUID(args *Args, reply *Reply) error {
	reply.Id = generateUID()
	return nil
}

func generateUID() int64 {
	gc.mu.Lock()
	defer gc.mu.Unlock()
	
	gc.count++
	log.Println(gc.count)
	return gc.count
}