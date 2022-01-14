package tritonhttp

import (
	"log"
	"net"
)

// NewHttpdServer Initialize the tritonhttp server by populating HttpServer structure
func NewHttpdServer(port string, docRoot map[string]string, mimePath string) (*HttpServer, error) {
	// Initialize mimeMap for server to refer
	mimeMap, err := ParseMIME(mimePath)
	if err != nil {
		return nil, err
	}
	// Return pointer to HttpServer
	httpServer := HttpServer{
		ServerPort: port,
		DocRoot:    docRoot,
		MIMEPath:   mimePath,
		MIMEMap:    mimeMap,
	}
	return &httpServer, nil
}

// Start the tritonhttp server
func (hs *HttpServer) Start() (err error) {
	// Start listening to the server port
	l, err := net.Listen("tcp", hs.ServerPort)
	if err != nil {
		return err
	}
	log.Println("Listening to connections on port", hs.ServerPort)
	defer func() {
		if err = l.Close(); err != nil {
			log.Println(err)
		}
	}()
	// Accept connection from client
	for {
		conn, err := l.Accept()
		if err != nil {
			return err
		}
		// Spawn a go routine to handle request
		go hs.handleConnection(conn)
	}
}
