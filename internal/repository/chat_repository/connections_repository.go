package chatrepository

import (
	"fmt"

	"github.com/gorilla/websocket"
)

type ConnectionsRepository struct {
	connections     map[string]*websocket.Conn
	userConnections map[string][]*websocket.Conn
}

func NewConnectionsRepository() IConnectionsRepository {
	return &ConnectionsRepository{
		connections:     make(map[string]*websocket.Conn),
		userConnections: make(map[string][]*websocket.Conn),
	}
}

func (c *ConnectionsRepository) GetConnection(key string) *websocket.Conn {
	return c.connections[key]
}

func (c *ConnectionsRepository) GetConnectionsByUserKey(userKey string) []*websocket.Conn {
	return c.userConnections[userKey]
}

func (c *ConnectionsRepository) AddConnection(key string, conn *websocket.Conn) {
	fmt.Println("Adding connection with key: ", key, " and connection.remoteAddr: ", conn.RemoteAddr().String())
	c.connections[key] = conn
}

func (c *ConnectionsRepository) AddUserConnection(userKey string, conn *websocket.Conn) {
	if _, ok := c.userConnections[userKey]; !ok {
		c.userConnections[userKey] = []*websocket.Conn{}
	}
	c.userConnections[userKey] = append(c.userConnections[userKey], conn)
}

func (c *ConnectionsRepository) DeleteConnection(key string) {
	delete(c.connections, key)
}

func (c *ConnectionsRepository) DeleteUserConnection(userKey string, conn *websocket.Conn) {
	if _, ok := c.userConnections[userKey]; ok {
		for _, c := range c.userConnections[userKey] {
			if c == conn {
				fmt.Println("Deleting connection with key: ", userKey, " and connection.remoteAddr: ", conn.RemoteAddr().String())
			}
		}
	}
}

func (c *ConnectionsRepository) DeleteUserConnections(userKey string, conn *websocket.Conn) {
	delete(c.userConnections, userKey)
}

func (c ConnectionsRepository) GetKeyFromConnection(conn *websocket.Conn) string {
	for key, connection := range c.connections {
		if connection == conn {
			return key
		}
	}
	return ""
}
