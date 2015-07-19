package main

import (
	"bufio"
	"container/list"
	"fmt"
	"github.com/wtmmac/go.net/websocket"
	"io"
	"net/http"
)

var connid int
var conns *list.List

func ChatroomServer(ws *websocket.Conn) {
	defer ws.Close()
	connid++
	id := connid
	fmt.Printf("connection id: %d\n", id)
	item := conns.PushBack(ws)
	//conns.Remove(item)
	name := fmt.Sprintf("user%d", id)
	SendMessage(nil, fmt.Sprintf("welcome %s join\n", name))
	r := bufio.NewReader(ws)
	for {
		data, err := r.ReadBytes('\n')
		if err != nil {
			fmt.Printf("disconnected id: %d\n", id)
			SendMessage(item, fmt.Sprintf("%s offline\n", name))
			break
		}
		fmt.Printf("%s: %s", name, data)
		SendMessage(item, fmt.Sprintf("%s\t> %s", name, data))
	}
}

func SendMessage(self *list.Element, data string) {
	print("\nenter sendmessage\n")

	for item := conns.Front(); item != nil; item = item.Next() {

		ws, ok := item.Value.(*websocket.Conn)

		if !ok {
			panic("item not *websocket.Conn")
		}

		if item == self {
			continue
		}

		io.WriteString(ws, data)
		print("write......\n")
	}
	print("exit sendmessage\n")
}

func Client(w http.ResponseWriter, r *http.Request) {
	html := "这里是一段html代码，呵呵！这个html代码中含有javascript脚本，脚本中含有创建websocket的代码，由于微博会将这段代码渲染成网页，所以暂时消除这段代码！"
	io.WriteString(w, html)
	SendMessage(nil, "ttttttttt")
	print("client......\n")
}

func main() {
	fmt.Printf(`Welcome chatroom server! `)
	connid = 0
	conns = list.New()
	http.Handle("/chatroom", websocket.Handler(ChatroomServer))
	http.HandleFunc("/client", Client)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
