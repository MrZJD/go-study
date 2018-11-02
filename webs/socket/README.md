## Socket编程

Socket

> Socket起源于Unix，一切皆文件，open -> write/read -> close

流式SOCK_STREAM: tcp
数据报式SOCK_DGRAM: udp

TCP/UDP协议: 确定一个进程 需要一个三元组: 协议 IP地址 端口

WebSocket

> 基于TCP http -> switch protocol -> ws(tcp)