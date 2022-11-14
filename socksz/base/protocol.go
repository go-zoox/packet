package base

// BASE:
//  VER | CMD | CRYPTO | COMPRESS | DATA
//   1  |  1  |  1     |   1      | -
//
//  VER 		- 版本
//  CMD 		- 命令
//  CRYPTO 	- 加密算法
//    0x00: 不加密
//    0x01: AES-128-CBC
//    0x02: AES-192-CBC
//    0x03: AES-256-CBC
//    0x04: AES-128-CFB
//    0x05: AES-192-CFB
//  COMPRESS - 压缩算法
//    0x00: 不压缩
//    0x01: GZIP
//    0x02: ZLIB
//    0x03: ZIP
//    0x04: BZIP2
//  DATA     - 数据
//    认证(authenticate)：客户端与服务端认证
//    握手(handshake)：客户端与客户端握手
//    转发(forward)：客户端与客户端通信
