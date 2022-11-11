package connect

// reference: https://www.quarkay.com/code/383/socks5-protocol-rfc-chinese-traslation

// Connect Request Protocol:
//   VER | CMD | RSV  | ATYP | DST.ADDR | DST.PORT
//    1  |  1  | 0x00 |  1   | Variable | 2
//
//   VER - 本次请求的协议版本号，取固定值 0x05（表示socks 5）
//   CMD - 连接方式，0x01:CONNECT, 0x02:BIND, 0x03:UDP ASSOCAITE
//   RSV - 保留字段，没用
//   ATYP - 地址类型，0x01:IPv4, 0x03:域名, 0x04:IPv6
//			IPV4　X‘01’
//			域名　X‘03’
//			IPV6　X‘04’
//   DST.ADDR - 目标地址
//   DST.PORT - 目标端口，2字节，网络字节序（network octec order）
//
//
// 	Connect Response Protocl:
//  	VER | REP | RSV  | ATYP | BND.ADDR | BND.PORT
//    1   |  1  | 0x00 |  1   | Variable | 2
//
//    VER  - 本次请求的协议版本号，取固定值 0x05（表示socks 5）
//    REP  - 状态码，0x00:成功，0x01:失败
//			X‘00’　成功
// 			X‘01’　常规 SOCKS 服务故障
// 			X‘02’　规则不允许的连接
// 			X‘03’　网络不可达
// 			X‘04’　主机无法访问
// 			X‘05’　拒绝连接
// 			X‘06’　连接超时
// 			X‘07’　不支持的命令
// 			X‘08’　不支持的地址类型
// 			X‘09’　到　X’FF’　未定义
//    RSV  - 保留字段，没用
//    ATYP - 地址类型，0x01:IPv4, 0x03:域名, 0x04:IPv6
//			IPV4　X‘01’
//			域名　X‘03’
//			IPV6　X‘04’
//    BND.ADDR - 服务端绑定地址
//    BND.PORT - 服务端绑定端口 （网络字节序）
//
