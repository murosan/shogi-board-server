package infrastructure

import (
	"io"
	"sync"
	"time"

	"go.uber.org/zap"

	"github.com/murosan/shogi-board-server/app/logger"
)

// Connector is a communicator with shogi engine service and os cmd.
type Connector interface {
	Connect() error
	Close(timeout time.Duration) error
	OnReceive(func([]byte) bool)
	Writer() io.Writer
}

type connector struct {
	sync.Mutex
	cmd      Cmd
	logger   logger.Logger
	receiver chan []byte
	isClosed bool
}

// NewConnector returns new Connector.
func NewConnector(cmd Cmd, logger logger.Logger) Connector {
	return &connector{
		cmd:      cmd,
		logger:   logger,
		receiver: make(chan []byte),
	}
}

func (conn *connector) Connect() error {
	conn.Lock()
	if err := conn.cmd.Start(); err != nil {
		return err
	}
	conn.Unlock()

	// receive on background until the pipe broken
	go conn.receive()
	return nil
}

func (conn *connector) Close(timeout time.Duration) error {
	conn.Lock()
	defer conn.Unlock()

	if conn.isClosed {
		return nil
	}
	conn.isClosed = true

	return conn.cmd.Wait(timeout)
}

func (conn *connector) OnReceive(block func([]byte) bool) {
	for {
		b, ok := <-conn.receiver
		if !ok {
			break // channel is closed
		}

		if !block(b) {
			break
		}
	}
}

func (conn *connector) Writer() io.Writer { return conn.cmd }

func (conn *connector) receive() {
	sc := conn.cmd.Scanner()
	if sc == nil {
		panic("scanner is nil. connect to cmd first")
	}

	for sc.Scan() {
		conn.receiver <- sc.Bytes()
	}

	if err := sc.Err(); err != nil {
		conn.logger.Warn("connection pipe broken", zap.Error(err))
	}

	close(conn.receiver)
	if err := conn.Close(3 * time.Second); err != nil {
		conn.logger.Warn("closing error")
	}
}
