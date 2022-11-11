package socksz

const (
	// VER ...
	VER = 1
)

const (
	// CommandAUTHENTICATE client <-> server
	CommandAUTHENTICATE = 0x01
	// CommandHANDSHAKERequest client <-server-> client
	CommandHANDSHAKERequest = 0x02
	// CommandHANDSHAKEResponse client <-server-> client
	CommandHANDSHAKEResponse = 0x03
	// CommandFORWARD client <-> client
	CommandFORWARD = 0x04
	// CommandCLOSE client <-> client
	CommandCLOSE = 0x05
	// CommandBIND client <-> server
	CommandBIND = 0x06
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
