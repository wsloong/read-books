package server

import (
	"log"
	"net/http"

	"github.com/wsloong/chatroom/logic"

	"nhooyr.io/websocket/wsjson"

	"nhooyr.io/websocket"
)

func WebSocketHandleFunc(w http.ResponseWriter, r *http.Request) {
	// Accept 从客户端接收 WebSocket 握手，并将链接升级为 WebScoket
	// 如果 Origin 域名与主机不同，Accept 将拒绝握手，除非设置了 InsecureSkipVerify 选项(通过第三个参数 AcceptOptions 来设置)
	// 换句话说，默认情况下，不允许跨域请求，如果发生错误，Accept 将始终写入适当的响应
	conn, err := websocket.Accept(w, r, nil)
	if err != nil {
		log.Println("websocket accept error: ", err)
		return
	}

	// 1 新用户进来，构造该用户的实例
	nickname := r.FormValue("nickname")
	if l := len(nickname); l < 2 || l > 20 {
		log.Println("nickname illegal: ", nickname)
		wsjson.Write(r.Context(), conn, logic.NewMessage("非法昵称,昵称长度: 4-20"))
		conn.Close(websocket.StatusUnsupportedData, "nickname illegal!")
		return
	}

	if !logic.Broadcaster.CanEnterRoom(nickname) {
		log.Println("昵称已经存在: ", nickname)
		wsjson.Write(r.Context(), conn, logic.NewMessage("该昵称已经存在!"))
		conn.Close(websocket.StatusUnsupportedData, "nickname exists!")
		return
	}

}
