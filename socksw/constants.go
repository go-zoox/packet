package socksz

const (
	VERSION = 1
)

const (
	// client <-> server
	COMMAND_AUTHENTICATE = 0x01
	// client <-server-> client
	COMMAND_HANDSHAKE_REQUEST  = 0x02
	COMMAND_HANDSHAKE_RESPONSE = 0x03
	// client <-> client
	COMMAND_FORWARD = 0x04
	// client <-> client
	COMMAND_CLOSE = 0x05
	// client <-> server
	COMMAND_BIND = 0x06
)

const (
	STATUS_OK                     = 0x00
	STATUS_INVALID_USER_CLIENT_ID = 0x01
	STATUS_INVALID_SIGNATURE      = 0x02
	STATUS_USER_NOT_ONLINE        = 0x03
	STATUS_FAILED_TO_PAIR         = 0x04
	STATUS_FAILED_TO_HANDSHAKE    = 0x05
)
