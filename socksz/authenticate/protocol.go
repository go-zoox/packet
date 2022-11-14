package authenticate

// BASE:
//  VER | CMD | CRYPTO | COMPRESS | DATA
//   1  |  1  |  1     |   1      | -

// Auth Request DATA:
// 	USER_CLIENT_ID | TIMESTAMP | NONCE | SIGNATURE
// 	      10       |    13     |   6   |  64 (HMAC_SHA256)
//
// 	USER_CLIENT_ID 	- 用户的 Client ID
//  TIMESTAMP 			- 时间戳，毫秒
//  NONCE 					- 随机数
//	SIGNATURE				- 数据签名，算法：HMAC_SHA256(USER_CLIENT_ID+TIMESTAMP+NONCE)

// Auth Response DATA:
//   STATUS | MESSAGE
//      1   |  -
//
//   STATUS 	- 认证状态
// 			0x00: 认证成功
//	 	  0x01: 用户 Client ID 无效
//			0x02: 数据签名无效
//
//   MESSAGE  - 错误信息
//
