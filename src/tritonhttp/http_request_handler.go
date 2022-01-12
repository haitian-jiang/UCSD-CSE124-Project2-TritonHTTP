package tritonhttp

import (
	"io"
	"log"
	"net"
	"regexp"
	"strings"
	"time"
)

const READ_TIMEOUT = 5 * time.Second
const DELIM = "\r\n"

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
	for {
		// Set a timeout for read operation
		if err := conn.SetReadDeadline(time.Now().Add(READ_TIMEOUT)); err != nil {
			log.Println("SetReadDeadline failed:", err)
			return
		}
		// If reusing read buffer, truncate it before next read
		connRecv := []byte{}
		buf := make([]byte, 1024)
		// Read from the connection socket into a buffer
		for {
			size, readErr := conn.Read(buf)
			if readErr != nil {
				if readErr == io.EOF {
					break // of no use... really break by judging size < len(buf)
				}
				if netErr, ok := readErr.(net.Error); ok && netErr.Timeout() {
					if len(connRecv) > 0 {
						// bad request: time out with incomplete request
						hs.handleBadRequest(conn) // respond sth before closing connection
					}
					return // close connection for time out and bad request
				}
			}
			connRecv = append(connRecv, buf[:size]...)
			if size < len(buf) {
				log.Println("Received one request.")
				break
			}
		}
		// Validate the request lines that were read
		requestHeader, ok := hs.parseRequest(string(connRecv))
		if !ok {
			hs.handleBadRequest(conn)
			return
		} else {
			// Handle any complete requests
			closed := hs.handleResponse(&requestHeader, conn)
			if closed == "close" {
				return
			}
		}
		// Update any ongoing requests
	}
}

func (hs *HttpServer) parseRequest(request string) (HttpRequestHeader, bool) {
	requestHeader := *new(HttpRequestHeader)
	// verify CRLF
	lines := strings.Split(request, DELIM)
	for i, line := range lines {
		if i < len(lines)-2 && line == "" {
			return requestHeader, false
		}
		if i >= len(lines)-2 && line != "" {
			return requestHeader, false
		}
	}
	// verify first line
	firstLine := strings.Split(lines[0], " ")
	if len(firstLine) != 3 || firstLine[0] != "GET" ||
		firstLine[1][0] != '/' || firstLine[2] != "HTTP/1.1" {
		return requestHeader, false
	}

	requestHeader.URL = firstLine[1]
	pairPattern := regexp.MustCompile(`(?P<key>[\w-]+):\s*(?P<value>.*)`)
	for _, line := range lines[1 : len(lines)-2] {
		match := pairPattern.FindStringSubmatch(line)
		if len(match) == 0 {
			return requestHeader, false
		}
		requestHeader.KeyValue = make(map[string]string)
		switch strings.ToLower(match[1]) {
		case "host":
			requestHeader.Host = match[2]
		case "connection":
			if strings.ToLower(match[2]) != "close" {
				return requestHeader, false
			}
			requestHeader.Connection = "close"
		default:
			requestHeader.KeyValue[match[1]] = match[2]
		}
	}
	if requestHeader.Host == "" {
		return requestHeader, false
	}
	return requestHeader, true
}
