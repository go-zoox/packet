package socksz

// Reference:
//   SOCKS5: https://www.quarkay.com/code/383/socks5-protocol-rfc-chinese-traslation
//   SHADOWSOCKS5: https://www.ichenxiaoyu.com/ss/
//   SOCKS6: https://datatracker.ietf.org/doc/html/draft-olteanu-intarea-socks-6
//   VMESS: https://github.com/v2ray/manual/blob/master/eng_en/protocols/vmess.md
//   mKCP: https://github.com/v2ray/manual/blob/master/eng_en/protocols/mkcp.md
//   MUXCOOL: https://github.com/v2ray/manual/blob/master/eng_en/protocols/muxcool.md

// USER
//  CLIENT_ID
//  CLIENT_SECRET
//  PAIR_KEY

// PACKET Protocol:
//  VER | CMD | CRYPTO | COMPRESS | DATA
//   1  |  1  |  1     |   1      | -
const (
	// LengthVer ...
	LengthVer = 1
	// LengthCmd ...
	LengthCmd = 1
	// LengthCrypto ...
	LengthCrypto = 1
	// LengthCompression ...
	LengthCompression = 1
)

// DATA Protocol:
//
// AUTHENTICATE DATA:
// request:  USER_CLIENT_ID | TIMESTAMP | NONCE | SIGNATURE
//             10           |    13     |   6   |  64 HMAC_SHA256
// response: STATUS | MESSAGE
//            1     |  -
const (
	// LengthUserClientID is the byte length of USER_CLIENT_ID
	LengthUserClientID = 10
	// LengthTimestamp is the byte length of TIMESTAMP
	LengthTimestamp = 13
	// LengthNonce is the byte length of NONCE
	LengthNonce = 6
	// LengthSignature is the byte length of SIGNATURE
	LengthSignature = 64
	// LengthStatus is the byte length of STATUS
	LengthStatus = 1
)

// Handshake DATA:
// request:  CONNECTION_ID | TARGET_USER_CLIENT_ID | TARGET_USER_PAIR_KEY |  NETWORK   | ATYP                 | DST.ADDR 							 | DST.PORT
//					       13      |       10              |					10          | 1(tcp/udp) | 1(IPv4/IPv6/Domain)  |   4 or 16 or domain    |    2
// response: CONNECTION_ID | STATUS | MESSAGE
//                 13      |  1     |  -
const (
	// LengthConnectionID ...
	LengthConnectionID = 13
	// LengthTargetUserClientID ...
	LengthTargetUserClientID = 10
	// LengthTargetUserPairSignature ...
	LengthTargetUserPairSignature = 64
	// LengthNetwork ...
	LengthNetwork = 1
	// LengthATyp ...
	LengthATyp = 1

	// LengthDSTAddr = 4

	// LengthDSTPort ...
	LengthDSTPort = 2

	// // LengthStatus ...
	// LengthStatus = 1
)

// CONNECTION TRANSMISIION DATA:
// request:  CONNECTION_ID | DATA
//					       13      |  -

// CONNECTION CLOSE:
// request:  CONNECTION_ID
//                 13
