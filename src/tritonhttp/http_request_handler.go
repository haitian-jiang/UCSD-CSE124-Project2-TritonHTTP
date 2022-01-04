package tritonhttp

import (
	"io"
	"log"
	"net"
	"time"
)

/* 
For a connection, keep handling requests until 
	1. a timeout occurs or
	2. client closes connection or
	3. client sends a bad request
*/
func (hs *HttpServer) handleConnection(conn net.Conn) {
	log.Println("Accepted new connection.")
	defer conn.Close()
	defer log.Println("Closed connection.")

	// Start a loop for reading requests continuously
	for  {
		// Set a timeout for read operation
		if err := conn.SetReadDeadline(time.Now().Add(20 * time.Second)); err != nil {
			log.Println("SetReadDeadline failed:", err)
			return
		}
		connRecv := []byte{}
		buf := make([]byte, 1024)
		// Read from the connection socket into a buffer
		for  {
			size, readErr := conn.Read(buf)
			if readErr != nil {
				if readErr == io.EOF {
					break
				}
				if netErr, ok := readErr.(net.Error); ok && netErr.Timeout() {
					if len(connRecv) > 0 {
						// bad request: time out with incomplete request
						hs.handleBadRequest(conn)
					} else {
						// time out, just close connection
						return
					}
				}
			}
			connRecv = append(connRecv, buf[:size]...)
			if size < len(buf) {
				break
			}
		}
		print(string(connRecv), "\n")
		_, err := conn.Write(connRecv)
		if err != nil {
			break
		}
		if string(connRecv) == "close" {
			break
		}
		connRecv = []byte{}
	}
		
		// Validate the request lines that were read

		// Handle any complete requests
		
		// Update any ongoing requests
		
		// If reusing read buffer, truncate it before next read
	
}
