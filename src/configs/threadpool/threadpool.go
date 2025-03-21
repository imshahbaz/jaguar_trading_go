package threadpool

import (
	"log"

	"github.com/panjf2000/ants/v2"
)

var Pool *ants.Pool

func InitPool() {
	var err error
	Pool, err = ants.NewPool(10, ants.WithOptions(ants.Options{Nonblocking: false}))
	if err != nil {
		log.Panicf("Error creating thread pool %v", err)
	}
}
