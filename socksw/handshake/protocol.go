package handshake

// BASE:
//  VER | CMD | CRYPTO | COMPRESS | DATA
//   1  |  1  |  1     |   1      | -

// Auth Request DATA:
// 	CONNECTION_ID | TARGET_USER_CLIENT_ID | TARGET_USER_PAIR_SIGNATURE | NETWORK | ATYP | DST.ADDR  | DST.PORT
// 	      21      |       10              |					    64             |    1    |   1  |  Variable |    2
//
// 	CONNECTION_ID 							- 连接 ID
//  TARGET_USER_CLIENT_ID 			- 目标用户 Client ID
//  TARGET_USER_PAIR_SIGNATURE  - 目标用户配对签名，签名算法: HMAC_SHA256(CONNECTION_ID + TARGET_USER_CLIENT_ID)
//  NETWORK 										- 连接网络类型
//    0x01: TCP
//		0x02: UDP
//  ATYP - 地址类型，0x01:IPv4, 0x03:域名, 0x04:IPv6
//		0x01: IPv4
//		0x03: 域名
//	  0x04: IPv6
//   DST.ADDR - 目标地址
//	   当 ATYP 是 IPv4 时，长度时 4 byte
//		 当 ATYP 是 IPv6 时，长度时 16 byte
//     当 ATYP 是 域名 时，第一个 byte 时域名的长度
//   DST.PORT - 目标端口，2字节，网络字节序（network octec order）

// Auth Response DATA:
//   STATUS | MESSAGE
//      1   |  -
//
//   STATUS 	- 认证状态
// 			0x00: 认证成功
//	 	  0x03: 目标用户不在线，无法沟通
//			0x04: 无效的配对签名，与目标用户配对失败
//		  0x05: 握手失败
//   MESSAGE  - 错误信息
//
