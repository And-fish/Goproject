// package main

// func main() {
// 	/*
// 		在浏览器中键入URL，按下回车之后会经历什么?
// 			1.浏览器向DNS服务器请求解析该URL中的域名所对应的IP地址
// 			2.解析出IP地址后，根据IP地址和默认端口80，和服务器建立TCP连接
// 			3.浏览器发出读取文件的HTTP请求，该请求报文作为TCP三次握手的第三个报文的数据发送给服务器
// 			4.服务器对浏览器请求做出响应，并把对应的html文本发送给浏览器
// 			5.释放TCP连接
// 			6.浏览器将html文本并显示内容

// 		基于 请求-响应 的模式:
// 			HTTP协议规定，请求从客户端发出，最后服务器端响应该请求并返回。
// 			即，服务端在没有接受到请求之前不会发送响应

// 		无状态保存:
// 			HTTP是一种不保存状态,即无状态协议。HTTP协议自身不对请求和响应之间的通信状态进行保存。
// 			即，在HTTP这个级别，协议遂于发送的请求和响应都不做持久化处理

// 		无连接:
// 			每次连接只处理以恶搞请求。服务器处理完客户的请求，并受到客户的应答后，即断开连接

// 		HTTP的请求方法:
// 			1.GET（查）：向指定的资源发出“显示”请求。GET方法只用在读取数据。
// 			2.HEAD：和GET一样都是向指定资源发出请求。之不够服务器不创会资源的文本部分。可以在不必传输全部内容的情况下，就可以获取其中“关于该资源的信息”
// 			3.POST（增）：向指定资源提交数据，请求服务器进行处理
// 			4.PUT（改）：向指定资源位置上传其最新内容
// 			5.DELETE（删）：删除指定资源
// 			6.TRACE：回显服务器收到的请求，主要用于测试或诊断
// 			7.CONNECT：用于SSL加密服务器的链接
// 			8.OPTIONS：使服务器传回该资源所支持的所有HTTP请求方法

// 		GET和POST的区别：
// 			1.GET提交的数据会放在URL之后，也就是请求行里面，以?分割URL和传输数据，参数之间以&相连，如EditBook?name=xxh&id=123456
// 			  POST是把提交到数据放在HTTP包的请求体中
// 			2.GET提交的大小有限制，而POST方法提交的数据没有限制
// 			3.GET与POST请求在服务端获取请求数据方式不同

// 		HTTP状态码:
// 			1xx消息：请求已经被服务器接收
// 			2xx成功：请求已成功被服务器接收、理解、并接受
// 			3xx重定向：需要后续操作才能完成这一请求
// 			4xx请求错误：请求含有此法错误或者无法被执行
// 			5xx服务器错误：服务器在处理某个正确请求时发送的错误

// 		const (
// 			StatusContinue           = 100 // RFC 7231, 6.2.1
// 			StatusSwitchingProtocols = 101 // RFC 7231, 6.2.2
// 			StatusProcessing         = 102 // RFC 2518, 10.1
// 			StatusEarlyHints         = 103 // RFC 8297

// 			StatusOK                   = 200 // RFC 7231, 6.3.1
// 			StatusCreated              = 201 // RFC 7231, 6.3.2
// 			StatusAccepted             = 202 // RFC 7231, 6.3.3
// 			StatusNonAuthoritativeInfo = 203 // RFC 7231, 6.3.4
// 			StatusNoContent            = 204 // RFC 7231, 6.3.5
// 			StatusResetContent         = 205 // RFC 7231, 6.3.6
// 			StatusPartialContent       = 206 // RFC 7233, 4.1
// 			StatusMultiStatus          = 207 // RFC 4918, 11.1
// 			StatusAlreadyReported      = 208 // RFC 5842, 7.1
// 			StatusIMUsed               = 226 // RFC 3229, 10.4.1

// 			StatusMultipleChoices   = 300 // RFC 7231, 6.4.1
// 			StatusMovedPermanently  = 301 // RFC 7231, 6.4.2
// 			StatusFound             = 302 // RFC 7231, 6.4.3
// 			StatusSeeOther          = 303 // RFC 7231, 6.4.4
// 			StatusNotModified       = 304 // RFC 7232, 4.1
// 			StatusUseProxy          = 305 // RFC 7231, 6.4.5
// 			_                       = 306 // RFC 7231, 6.4.6 (Unused)
// 			StatusTemporaryRedirect = 307 // RFC 7231, 6.4.7
// 			StatusPermanentRedirect = 308 // RFC 7538, 3

// 			StatusBadRequest                   = 400 // RFC 7231, 6.5.1
// 			StatusUnauthorized                 = 401 // RFC 7235, 3.1
// 			StatusPaymentRequired              = 402 // RFC 7231, 6.5.2
// 			StatusForbidden                    = 403 // RFC 7231, 6.5.3
// 			StatusNotFound                     = 404 // RFC 7231, 6.5.4
// 			StatusMethodNotAllowed             = 405 // RFC 7231, 6.5.5
// 			StatusNotAcceptable                = 406 // RFC 7231, 6.5.6
// 			StatusProxyAuthRequired            = 407 // RFC 7235, 3.2
// 			StatusRequestTimeout               = 408 // RFC 7231, 6.5.7
// 			StatusConflict                     = 409 // RFC 7231, 6.5.8
// 			StatusGone                         = 410 // RFC 7231, 6.5.9
// 			StatusLengthRequired               = 411 // RFC 7231, 6.5.10
// 			StatusPreconditionFailed           = 412 // RFC 7232, 4.2
// 			StatusRequestEntityTooLarge        = 413 // RFC 7231, 6.5.11
// 			StatusRequestURITooLong            = 414 // RFC 7231, 6.5.12
// 			StatusUnsupportedMediaType         = 415 // RFC 7231, 6.5.13
// 			StatusRequestedRangeNotSatisfiable = 416 // RFC 7233, 4.4
// 			StatusExpectationFailed            = 417 // RFC 7231, 6.5.14
// 			StatusTeapot                       = 418 // RFC 7168, 2.3.3
// 			StatusMisdirectedRequest           = 421 // RFC 7540, 9.1.2
// 			StatusUnprocessableEntity          = 422 // RFC 4918, 11.2
// 			StatusLocked                       = 423 // RFC 4918, 11.3
// 			StatusFailedDependency             = 424 // RFC 4918, 11.4
// 			StatusTooEarly                     = 425 // RFC 8470, 5.2.
// 			StatusUpgradeRequired              = 426 // RFC 7231, 6.5.15
// 			StatusPreconditionRequired         = 428 // RFC 6585, 3
// 			StatusTooManyRequests              = 429 // RFC 6585, 4
// 			StatusRequestHeaderFieldsTooLarge  = 431 // RFC 6585, 5
// 			StatusUnavailableForLegalReasons   = 451 // RFC 7725, 3

// 			StatusInternalServerError           = 500 // RFC 7231, 6.6.1
// 			StatusNotImplemented                = 501 // RFC 7231, 6.6.2
// 			StatusBadGateway                    = 502 // RFC 7231, 6.6.3
// 			StatusServiceUnavailable            = 503 // RFC 7231, 6.6.4
// 			StatusGatewayTimeout                = 504 // RFC 7231, 6.6.5
// 			StatusHTTPVersionNotSupported       = 505 // RFC 7231, 6.6.6
// 			StatusVariantAlsoNegotiates         = 506 // RFC 2295, 8.1
// 			StatusInsufficientStorage           = 507 // RFC 4918, 11.5
// 			StatusLoopDetected                  = 508 // RFC 5842, 7.2
// 			StatusNotExtended                   = 510 // RFC 2774, 7
// 			StatusNetworkAuthenticationRequired = 511 // RFC 6585, 6
// 		)

// 		URL：
// 			例：https://github.com:80/docker/compose/releases?page=3
// 			https，是协议
// 			github.com，是服务器
// 			80，是网络端口号，默认不显示
// 			/docker/compose/releases，是路径
// 			?page=3，是查询

// 		Restful、SOAP、XML-RPC三种web编程交互方案。Restful简单明了，更流行

// 	*/
// }
