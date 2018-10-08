package usi

var (
	CmdUsi     = []byte("usi")
	CmdIsReady = []byte("isready")
	CmdNew     = []byte("usinewgame")

	CmdQuit = []byte("quit")

	ResOk    = []byte("usiok")
	ResReady = []byte("readyok")

	StartCmds = [][]byte{
		CmdUsi,
		CmdIsReady,
		CmdNew,
	}
)
