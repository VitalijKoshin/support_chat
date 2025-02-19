package chatrepository

import "github.com/gorilla/websocket"

type IConnectionsRepository interface {
	GetConnection(key string) *websocket.Conn
	GetConnectionsByUserKey(userKey string) []*websocket.Conn
	AddConnection(key string, conn *websocket.Conn)
	AddUserConnection(userKey string, conn *websocket.Conn)
	DeleteConnection(key string)
	DeleteUserConnection(userKey string, conn *websocket.Conn)
	DeleteUserConnections(userKey string, conn *websocket.Conn)
}
