package usi

var (
	CmdUsi     = []byte("usi")
	CmdIsReady = []byte("isready")
	CmdNew     = []byte("usinewgame")

	ResOk    = []byte("usiok")
	ResReady = []byte("readyok")

	startCmds = [][]byte{
		CmdUsi,
		CmdIsReady,
		CmdNew,
	}
)
