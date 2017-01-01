package elkm1grpc

import (
	"io"
	"log"
	"golang.org/x/net/context"
	"fmt"
	"time"
	"errors"
	"sync"
)

func NewServer(stream io.ReadWriter) (s *Server) {
	s = new(Server)
	s.stream = stream
	s.requestChan = make(chan string)
	s.listenChan = make(chan string)
	s.armingStatusChan = make(map[chan *ArmingStatusReport]bool)
	s.armingStatusChanLock = new(sync.Mutex)
	go s.listen()
	go s.handle()
	return
}

type elkRequest struct {
	Command  string
	Response chan string
}

type Server struct {
	ElkConn              *io.ReadWriter
	stream               io.ReadWriter
	requestChan          chan string
	listenChan           chan string
	armingStatusChan     map[chan *ArmingStatusReport]bool
	armingStatusChanLock *sync.Mutex
}

func (s *Server) readArmingStatus() chan *ArmingStatusReport {
	c := make(chan *ArmingStatusReport)
	s.armingStatusChanLock.Lock()
	s.armingStatusChan[c] = true
	s.armingStatusChanLock.Unlock()
	return c
}

func (s *Server) asciiLookup(Type, Address int) chan string {
	s.request(fmt.Sprintf("sd%03d%03d00", Type, Address))
	// TODO finish this
	return nil
}

func (s *Server) ArmStay(context context.Context, args *ArmArgs) (*ArmingStatusArea, error) {
	s.request(fmt.Sprintf("a%s%d%06d00",
		string(ArmingStatusArea_ARM_STEP_TO_NEXT_STAY_MODE + elkEnumConst),
		args.Area,
		args.Pin,
	))

	t := time.After(2 * time.Second)
	select {
	case as := <-s.readArmingStatus():
		return as.Areas[args.Area], nil
	case <-t:
	}
	return nil, errors.New("timeout")
}

func (s *Server) ArmAway(context context.Context, args *ArmArgs) (*ArmingStatusArea, error) {
	s.request(fmt.Sprintf("a%s%d%06d00",
		string(ArmingStatusArea_ARM_STEP_TO_NEXT_AWAY_MODE + elkEnumConst),
		args.Area,
		args.Pin,
	))

	t := time.After(2 * time.Second)
	select {
	case as := <-s.readArmingStatus():
		return as.Areas[args.Area], nil
	case <-t:
	}
	return nil, errors.New("timeout")
}

func (s *Server) Disarm(context context.Context, args *ArmArgs) (*ArmingStatusArea, error) {
	s.request(fmt.Sprintf("a%s%d%06d00",
		string(ArmingStatusArea_DISARMED + elkEnumConst),
		args.Area,
		args.Pin,
	))
	t := time.After(2 * time.Second)
	select {
	case as := <-s.readArmingStatus():
		return as.Areas[args.Area], nil
	case <-t:
	}
	return nil, errors.New("timeout")
}

func (s *Server) ArmingStatusChange(args *ArmingChangeArgs, stream ElkGRPC_ArmingStatusChangeServer) error {
	for {
		err := stream.Send(<-s.readArmingStatus())
		if err != nil {
			log.Println("unable to send arming status change", err)
		}
	}
}

func (s *Server) ZoneChange(args *ZoneChangeArgs, stream ElkGRPC_ZoneChangeServer) error {
	return nil
}
