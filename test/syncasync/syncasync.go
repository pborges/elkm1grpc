package main

import (
	"net"
	"bufio"
	"log"
	"os"
	"io"
	"regexp"
	"time"
	"strings"
	"fmt"
)

func main() {
	c, err := net.Dial("tcp", "stratus:3000")
	if err != nil {
		panic(err)
	}

	m := NewModem(c)

	r := &request{
		Request: "0Bsd010010065",
		Filter:  regexp.MustCompile("SD"),
		Response:make(chan string),
	}
	m.writeChan <- r
	fmt.Println("RESPONSE: ", <-r.Response)

	bufio.NewReader(os.Stdin).ReadLine()
}

func NewModem(stream io.ReadWriter) *Modem {
	m := &Modem{
		stream:     stream,
		readChan:   make(chan string, 1),
		writeChan:  make(chan *request),
	}
	go m.read()
	go m.write()
	return m
}

type request struct {
	Request  string
	Filter   *regexp.Regexp
	Response chan string
}

type Modem struct {
	stream    io.ReadWriter
	readChan  chan string
	writeChan chan *request
}

func (m *Modem) read() {
	r := bufio.NewReader(m.stream)
	for {
		line, err := r.ReadString(0x0a)
		if err != nil {
			log.Println("readChan:", err)
			continue
		}
		m.readChan <- strings.Trim(line, "\r\n")
	}
}

func (m *Modem) write() {
	for {
		select {
		case req := <-m.writeChan:
			_, err := m.stream.Write([]byte(req.Request + "\r\n"))
			if err != nil {
				log.Println("writeChan:", err)
				continue
			}
			t := time.After(1 * time.Second)
				select {
				case msg := <-m.readChan:
					if req.Filter != nil && req.Response != nil && req.Filter.MatchString(msg) {
						req.Response <- msg
						m.readChan <- msg
					}
				case <-t:
					log.Println("writeChan: timeout")
				}
		case msg := <-m.readChan:
			log.Println(msg)
		}
	}
}
