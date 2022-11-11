package authenticate

// reference: https://www.quarkay.com/code/383/socks5-protocol-rfc-chinese-traslation

// Auth Request Protocol:
// 	VER | NMETHODS | METHODS
// 	1   | 1        | 1
//
// 	VER - 本次请求的协议版本号，取固定值 0x05（表示socks 5）
//  NMETHODS - 客户端支持的认证方式数量，可取值 1~255
//  METHODS - 可用的认证方式列表
//
//		0x00 NO AUTHENTICATION REQUIRED 无需认证
// 		0x01 GSSAPI
// 		0x02 USERNAME/PASSWORD 无需认证
// 		0x03 to 0x7F IANA ASSIGNED
// 		0x80 to 0xFE REVERSED FOR PRIVATE METHODS
// 		0xFF NO ACCEPTABLE METHODS
//
//
// Auth Response Protocol:
//   VER | METHOD
//   1   |   1
//
//   VER 		- 协议版本
//   METHOD - 服务端期望的认证方式
//
