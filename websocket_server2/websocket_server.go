package main

import (
	"bufio"
	"container/list"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/wtmmac/go.net/websocket"
)

var (
	connid int
	conns  *list.List
)

type conn struct {
	ws     *websocket.Conn
	liveid int
	uid    int
}

func MainHandler(w http.ResponseWriter, r *http.Request) {
	html := `<!DOCTYPE html>
	<html>
	<head>
		<title>chat demo</title>
	</head>
	<body>
	<h1>直播弹幕接收测试页</h1>
		<div id="msg"></div>
		<script src="http://code.jquery.com/jquery-1.10.1.min.js"></script>
	<script type="text/javascript" >
	thisHost = location.host;
	href = location.href;
	href_split = href.split("?");
	liveid = href_split[1];
	if (!liveid)
	liveid = 658;
	
	(function() {
		var $msg = $('#msg');
		var $text = $('#text');
	
		var WebSocket = window.WebSocket || window.MozWebSocket;
		if (WebSocket) {
			try {
				var socket = new WebSocket('ws://' + thisHost + '/dmlive/new-msg/socket/' + liveid);
			} catch (e) {}
		}
	
		if (socket) {
			socket.onmessage = function(event) {
				$msg.append('<p>' + event.data + '</p>');
			}
	
			$('form').submit(function() {
				socket.send($text.val());
				$text.val('').select();
				return false;
			});
		}
	})();
	</script>
	</body>
	</html>`

	io.WriteString(w, html)
}

// 在线人数
func StatusHandler(w http.ResponseWriter, r *http.Request) {
	number := conns.Len()
	io.WriteString(w, fmt.Sprintf("%d", number))
	fmt.Println(number)
}

// 输出跨域XML
func CrossDomain(w http.ResponseWriter, r *http.Request) {
	html := `<?xml version="1.0" encoding="UTF-8"?>
	<!DOCTYPE cross-domain-policy SYSTEM "http://www.adobe.com/xml/dtds/cross-domain-policy.dtd">
	<!-- Policy file for xmlsocket://socks.adobe.com -->
	<cross-domain-policy>
		<site-control permitted-cross-domain-policies="all" />
		<allow-access-from domain="*" to-ports="*" />
		<allow-http-request-headers-from domain="*" headers="*"/>
	</cross-domain-policy`

	io.WriteString(w, html)
}

// websocket房间
func ChatroomServer(ws *websocket.Conn) {
	defer ws.Close()
	connid++
	fmt.Printf("connection id: %d\n", connid)

	vars := mux.Vars(ws.Request())
	liveid, err := strconv.Atoi(vars["liveid"])
	if err != nil {
		return
	}
	fmt.Printf("connected to liveid: %d\n", liveid)
	conn_instance := &conn{ws: ws, liveid: liveid, uid: connid}
	item := conns.PushBack(conn_instance)
	// name := fmt.Sprintf("user%d", id)
	// SendMessage(nil, fmt.Sprintf("welcome %s join\n", name))
	r := bufio.NewReader(ws)
	for {
		data, err := r.ReadBytes('\n')
		if err != nil {
			// fmt.Printf("disconnected id: %d\n", id)
			// SendMessage(nil, fmt.Sprintf("%s offline\n", name))
			conns.Remove(item)
			return
		}
		fmt.Printf("%d: %s", connid, data)
		// SendMessage(item, fmt.Sprintf("%s\t> %s", name, data), liveid)
	}
}

// 消息发送
func SendMessage(self *list.Element, data string, liveid int) {
	for item := conns.Front(); item != nil; item = item.Next() {
		conn_instance, ok := item.Value.(*conn)

		if !ok {
			panic("item not *websocket.Conn")
		}

		//		if item == self {
		//			continue
		//		}

		// liveid为0时发送给所有客户端，liveid大于0时只发送给等于liveid的客户端
		if liveid == 0 || liveid > 0 && liveid == conn_instance.liveid {
			io.WriteString(conn_instance.ws, data)
		}
	}
}

func Client(w http.ResponseWriter, r *http.Request) {
	html := "测试客户端"
	io.WriteString(w, html)
	SendMessage(nil, html, 0)
}

// HTTP接收
func ReceiveHandler(w http.ResponseWriter, r *http.Request) {
	queryForm, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		return
	}

	if len(queryForm["liveid"]) <= 0 {
		return
	}

	liveid, err := strconv.Atoi(queryForm["liveid"][0])
	if err != nil {
		return
	}

	if len(queryForm["dm"]) <= 0 {
		return
	}

	dm := queryForm["dm"][0]

	fmt.Printf("%d&%s\n", liveid, dm)

	SendMessage(nil, dm, liveid)
}

// 主函数
func main() {
	fmt.Printf(`Welcome chatroom server! `)
	connid = 0
	conns = list.New()

	r := mux.NewRouter()
	r.HandleFunc("/dmlive/test", MainHandler)
	r.HandleFunc("/dmlive/status", StatusHandler)
	r.HandleFunc("/dmlive/receive", ReceiveHandler)
	r.Handle("/dmlive/new-msg/socket/{liveid:[0-9]+}", websocket.Handler(ChatroomServer))
	r.HandleFunc("/client", Client)
	r.HandleFunc("/crossdomain.xml", CrossDomain)
	http.Handle("/", r)

	err := http.ListenAndServe(":7001", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
