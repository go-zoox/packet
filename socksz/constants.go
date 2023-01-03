package socksz

const (
	// VER ...
	VER = 1
)

const (
	// CommandAuthenticate client <-> server
	CommandAuthenticate = 0x01
	// CommandHandshakeRequest client <-server-> client
	CommandHandshakeRequest = 0x02
	// CommandHandshakeResponse client <-server-> client
	CommandHandshakeResponse = 0x03
	// CommandForward client <-> client
	CommandForward = 0x04
	// CommandClose client <-> client
	CommandClose = 0x05
	// CommandBind client <-> server
	CommandBind = 0x06
	// CommandJoinAsAgent client <-> server
	CommandJoinAsAgent = 0x07
)

const (
	// StatusOK means status ok
	StatusOK = 0x00
	// StatusInvalidUserClientID means invalid user client id
	StatusInvalidUserClientID = 0x01
	// StatusInvalidSignature means invalid signature
	StatusInvalidSignature = 0x02
	// StatusUserNotOnline means user is not online
	StatusUserNotOnline = 0x03
	// StatusFailedToPair means failed to pair with signature
	StatusFailedToPair = 0x04
	// StatusFailedToHandshake means failed to handshake
	StatusFailedToHandshake = 0x05
)
