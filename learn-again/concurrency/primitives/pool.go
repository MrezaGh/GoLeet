package primitives

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

func connectToService() interface{} {
	time.Sleep(1 * time.Second)
	return struct{}{}
}

func warmServiceConnCache() *sync.Pool {
	p := &sync.Pool{
		New: connectToService,
	}
	for i := 0; i < 10; i++ {
		p.Put(p.New())
	}
	return p
}

func startNetworkDaemon() *sync.WaitGroup {
	var w8forServer sync.WaitGroup
	w8forServer.Add(1)
	go func() {
		connPool := warmServiceConnCache()

		server, err := net.Listen("tcp", "localhost:8070")
		if err != nil {
			log.Fatal("nope! err:", err)
		}
		defer server.Close()
		w8forServer.Done()

		for {
			conn, err := server.Accept()
			if err != nil {
				log.Println("con ops, err:", err)
				continue
			}

			//connectToService()
			svcConn := connPool.Get()
			fmt.Fprintln(conn, "")
			connPool.Put(svcConn)
			conn.Close()
		}
	}()
	return &w8forServer
}
