package joinAsAgent

// BASE:
//  VER | CMD | CRYPTO | COMPRESS | DATA
//   1  |  1  |  1     |   1      | -

// Request DATA:
//  Type | ID_LENGTH | ID | TIMESTAMP | NONCE | SIGNATURE
// 	  1  |       1   | -  |    13     |   6   |  64 (HMAC_SHA256)
//
//  Type						- 类型: 用户(user, 0x00) | ROOM(room, 0x01)
//  ID_LENGTH 			- 如果类型是用户，则 ID 为 用户 User ClientID 的长度；如果类型是 ROOM，则 ID 为 RoomID 的长度
// 	ID 							- 如果类型是用户，则 ID 为 用户 User ClientID；如果类型是 ROOM，则 ID 为 RoomID
//  TIMESTAMP 			- 时间戳，毫秒
//  NONCE 					- 随机数
//	SIGNATURE				- 数据签名，算法：HMAC_SHA256(USER_CLIENT_ID+TIMESTAMP+NONCE)

// Response DATA:
//   STATUS | MESSAGE
//      1   |  -
//
//   STATUS 	- 认证状态
// 			0x00: 成功
//	 	  0x01: 用户 Client ID 或者 Room ID 无效
//			0x02: 数据签名无效
//
//   MESSAGE  - 错误信息
//
