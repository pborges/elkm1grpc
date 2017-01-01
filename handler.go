package elkm1grpc

import (
	"bufio"
	"log"
	"os"
	"fmt"
)

var Log *log.Logger

func init() {
	Log = log.New(os.Stdout, "[elkm1grpc] ", log.LstdFlags | log.Lshortfile)
}

func (e *Server) listen() {
	scanner := bufio.NewScanner(e.stream)
	for scanner.Scan() {
		s := scanner.Text()
		e.listenChan <- s
	}
	if err := scanner.Err(); err != nil {
		Log.Println("Scanner error", err)
	}
}

func (e *Server)handle() {
	for {
		select {
		case cmd := <-e.requestChan:
			le := "SEND: " + cmd
			_, err := e.stream.Write([]byte(cmd))
			if err != nil {
				le += " error: " + err.Error()
			}
			Log.Print(le)
		case event := <-e.listenChan:
			switch event[2:4] {
			case "AS":// ARMING STATUS REPORT UPDATE
				e.armingStatusChanLock.Lock()
				for c := range e.armingStatusChan {
					select {
					case c <- decodeArmingStatusReport(event):
						close(c)
						delete(e.armingStatusChan, c)
					default:
					}
				}
				e.armingStatusChanLock.Unlock()
			}
			Log.Println("EVNT[", event[2:4], "]", event)
		}
	}
}

func (e *Server)request(cmd string) {
	c := fmt.Sprintf("%02X%s", len(cmd) + 2, cmd)
	e.requestChan <- fmt.Sprintf("%s%02X\r\n", c, crc([]byte(c)))
}