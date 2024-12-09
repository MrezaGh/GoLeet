package primitives

import (
	"io"
	"net"
	"testing"
)

// func TestCond(t *testing.T) {
// cond
// queuesim()
// btnSim()
// }
func init() {
	//fmt.Println(runtime.NumCPU())
	wg := startNetworkDaemon()
	wg.Wait()
}
func BenchmarkPoolServer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		conn, err := net.Dial("tcp", "localhost:8070")
		if err != nil {
			b.Fatal("cant connect to host", err)
		}
		if _, err = io.ReadAll(conn); err != nil {
			b.Fatal("cannot read, err:", err)
		}
		conn.Close()
	}
}
