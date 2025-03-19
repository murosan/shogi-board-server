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
	OnReceive(func([]byte) bool) // if true, continue receiving
	UnsetOnReceive()
	Writer() io.Writer
}

type connector struct {
	sync.RWMutex
	cmd       Cmd
	logger    logger.Logger
	onReceive func([]byte) bool
	isClosed  bool
}

// NewConnector returns new Connector.
func NewConnector(cmd Cmd, logger logger.Logger) Connector {
	return &connector{
		cmd:       cmd,
		logger:    logger,
		onReceive: func(b []byte) bool { return true },
	}
}

func (conn *connector) Connect() error {
	conn.Lock()
	err := conn.cmd.Start()
	conn.Unlock()

	if err != nil {
		return err
	}

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
	conn.Lock()
	conn.onReceive = block
	conn.Unlock()
}

func (conn *connector) UnsetOnReceive() {
	conn.OnReceive(func(b []byte) bool { return true })
}

func (conn *connector) Writer() io.Writer { return conn.cmd }

func (conn *connector) receive() {
	sc := conn.cmd.Scanner()
	if sc == nil {
		panic("scanner is nil. connect to cmd first")
	}

	for sc.Scan() {
		b := sc.Bytes()

		conn.RLock()
		f := conn.onReceive
		conn.RUnlock()

		if !f(b) {
			conn.UnsetOnReceive()
		}
	}

	if err := sc.Err(); err != nil {
		conn.logger.Warn("connection pipe broken", zap.Error(err))
	}

	if err := conn.Close(3 * time.Second); err != nil {
		conn.logger.Warn("closing error")
	}
}
