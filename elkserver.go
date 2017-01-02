package elkm1grpc

import (
	"io"
	"log"
	"golang.org/x/net/context"
	"fmt"
	"time"
	"errors"
	"sync"
	"bufio"
	"strings"
)

const WaitDuration = 5 * time.Second

func NewServer(stream io.ReadWriter) (s *Server) {
	s = &Server{
		stream:     stream,
		readChan:   make(chan string),
		writeChan:  make(chan string),
	}
	s.interestChannels = make(map[string]map[chan string]bool)
	s.interestLock = new(sync.Mutex)

	s.asciiTable = make(map[int]map[int]string)
	s.asciiTableLock = new(sync.Mutex)

	go s.read()
	go s.write()
	return s
}

type Server struct {
	stream    io.ReadWriter
	readChan  chan string
	writeChan chan string

	interestChannels map[string]map[chan string]bool
	interestLock     *sync.Mutex

	asciiTable     map[int]map[int]string
	asciiTableLock *sync.Mutex
}

func (s *Server) read() {
	r := bufio.NewReader(s.stream)
	for {
		line, err := r.ReadString(0x0a)
		if err != nil {
			log.Panic("readChan:", err)
		}
		s.readChan <- strings.Trim(line, "\r\n")
	}
}

func (s *Server) write() {
	for {
		select {
		case req := <-s.writeChan:
			r := fmt.Sprintf("%02X%s", len(req)+2, req)
			r = fmt.Sprintf("%s%02X", r, crc([]byte(r)))
			log.Println("WRITE:", r)
			_, err := s.stream.Write([]byte(r + "\r\n"))
			if err != nil {
				log.Panic("writeChan:", err)
			}
		case msg := <-s.readChan:
			log.Println("READ:", msg)
			s.interestLock.Lock()
			if chn, ok := s.interestChannels[msg[2:4]]; ok {
				for c := range chn {
					select {
					case c <- msg:
					default:
					}
				}
			}
			s.interestLock.Unlock()
		}
	}
}

func (s *Server) registerInterest(msg string) chan string {
	c := make(chan string)
	s.interestLock.Lock()
	defer s.interestLock.Unlock()
	if _, ok := s.interestChannels[msg]; !ok {
		s.interestChannels[msg] = make(map[chan string]bool)
	}
	s.interestChannels[msg][c] = true
	return c
}

func (s *Server) unregisterInterest(msg string, c chan string) {

	if _, ok := s.interestChannels[msg]; ok {
		if _, ok := s.interestChannels[msg][c]; ok {
			close(c)
			delete(s.interestChannels[msg], c)
			return
		}
	}
	log.Panic("tried to unregister a channel that was not registered...")
}

func (s *Server) asciiLookup(Type, Address int) string {
	if _, ok := s.asciiTable[Type]; !ok {
		s.asciiTable[Type] = make(map[int]string)
	}
	if _, ok := s.asciiTable[Type][Address]; !ok {
		c := s.registerInterest("SD")
		defer s.unregisterInterest("SD", c)
		s.writeChan <- fmt.Sprintf("sd%02d%03d00", Type, Address)
		t := time.After(1 * time.Second)
		select {
		case des := <-c:
			s.asciiTable[Type][Address] = strings.Trim(des[9:25], " ")
		case <-t:
			return "UNKNOWN"
		}
	}
	return s.asciiTable[Type][Address]
}

func (s *Server) ArmStay(context context.Context, args *ArmArgs) (*ArmingStatusArea, error) {
	s.writeChan <- fmt.Sprintf("a%s%d%06d00",
		string(ArmingStatusArea_ARM_STEP_TO_NEXT_STAY_MODE + elkEnumConst),
		args.Area,
		args.Pin,
	)

	c := s.registerInterest("AS")
	defer s.unregisterInterest("AS", c)

	t := time.After(WaitDuration)
	select {
	case as := <-c:
		return decodeArmingStatusReport(as).Areas[args.Area], nil
	case <-t:
	}

	return nil, errors.New("timeout")
}

func (s *Server) ArmAway(context context.Context, args *ArmArgs) (*ArmingStatusArea, error) {
	s.writeChan <- fmt.Sprintf("a%s%d%06d00",
		string(ArmingStatusArea_ARM_STEP_TO_NEXT_AWAY_MODE + elkEnumConst),
		args.Area,
		args.Pin,
	)

	c := s.registerInterest("AS")
	defer s.unregisterInterest("AS", c)

	t := time.After(WaitDuration)
	select {
	case as := <-c:
		return decodeArmingStatusReport(as).Areas[args.Area], nil
	case <-t:
	}

	return nil, errors.New("timeout")
}

func (s *Server) Disarm(context context.Context, args *ArmArgs) (as *ArmingStatusArea, err error) {
	s.writeChan <- fmt.Sprintf("a%s%d%06d00",
		string(ArmingStatusArea_DISARMED + elkEnumConst),
		args.Area,
		args.Pin,
	)

	c := s.registerInterest("AS")
	defer s.unregisterInterest("AS", c)

	t := time.After(WaitDuration)
	select {
	case as := <-c:
		return decodeArmingStatusReport(as).Areas[args.Area], nil
	case <-t:
	}

	return nil, errors.New("timeout")
}

func (s *Server) ZoneStatus(context context.Context, args *ZoneStatusArgs) (zsr *ZoneStatusReport, err error) {
	s.writeChan <- "zs00"

	c := s.registerInterest("ZS")
	defer s.unregisterInterest("ZS", c)

	t := time.After(WaitDuration)
	select {
	case r := <-c:
		zsr = new(ZoneStatusReport)
		zsr.Zones = make([]*Zone, 0)
		for i := 0; i < 208; i++ {
			z := new(Zone)
			z.Zone = int32(i)
			status, sub := decodeZoneStatus(int(r[4 + i]))
			z.Status = Zone_ZoneStatus(status)
			z.SubStatus = Zone_ZoneSubStatus(sub)
			if z.SubStatus != Zone_UNCONFIGURED {
				z.Description = s.asciiLookup(0, i)
				zsr.Zones = append(zsr.Zones, z)
			} else {
				break
			}
		}
		return
	case <-t:
		err = errors.New("timeout")
	}

	return
}

func (s *Server) ArmingStatus(context context.Context, args *ArmingStatusArgs) (asr *ArmingStatusReport, err error) {
	s.writeChan <- "as00"

	c := s.registerInterest("AS")
	defer s.unregisterInterest("AS", c)

	t := time.After(5 * time.Second)
	select {
	case r := <-c:
		asr = decodeArmingStatusReport(r)
		return
	case <-t:
		err = errors.New("timeout")
	}

	return
}

func (s *Server) ArmingStatusChange(args *ArmingChangeArgs, stream ElkGRPC_ArmingStatusChangeServer) error {
	c := s.registerInterest("AS")
	defer s.unregisterInterest("AS", c)
	for {
		err := stream.Send(decodeArmingStatusReport(<-c))
		if err != nil {
			log.Println("unable to send arming status change", err)
		}
	}
	return nil
}

func (s *Server) ZoneChange(args *ZoneChangeArgs, stream ElkGRPC_ZoneChangeServer) error {
	zc := s.registerInterest("ZC")
	defer s.unregisterInterest("ZC", zc)
	for {
		z := decodeZoneUpdate(<-zc)
		z.Description = s.asciiLookup(0, int(z.Zone))
		err := stream.Send(z)
		if err != nil {
			log.Println("unable to send zone update change", err)
		}
	}
	return nil
}
