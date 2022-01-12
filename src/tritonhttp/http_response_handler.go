package tritonhttp

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func (hs *HttpServer) handleBadRequest(conn net.Conn) {
	now := time.Now().UTC().Format("Mon, 02 Jan 2006 15:04:05 GMT")
	respHeader := HttpResponseHeader{
		StatusCode:  400,
		Description: "Bad Request",
		Date:        now,
		Connection:  "close",
	}
	hs.sendResponse(respHeader, conn)
}

func (hs *HttpServer) handleFileNotFoundRequest(requestHeader *HttpRequestHeader, conn net.Conn) {
	panic("todo - handleFileNotFoundRequest")
}

func (hs *HttpServer) handleResponse(requestHeader *HttpRequestHeader, conn net.Conn) (result string) {
	//panic("todo - handleResponse")
	now := time.Now().UTC().Format("Mon, 02 Jan 2006 15:04:05 GMT")
	responseHeader := *new(HttpResponseHeader)
	responseHeader.Date = now
	if requestHeader.Connection == "close" {
		responseHeader.StatusCode = 200
		responseHeader.Description = "OK"
		responseHeader.Connection = "close"
		hs.sendResponse(responseHeader, conn)
		return "close"
	}
	conn.Write([]byte("normal"))
	return "1"
}

func (hs *HttpServer) sendResponse(responseHeader HttpResponseHeader, conn net.Conn) {
	//panic("todo - sendResponse")
	writer := bufio.NewWriter(conn)
	defer func() {
		if err := writer.Flush(); err != nil {
			log.Println(err)
		}
	}()

	// Send headers
	headers := fmt.Sprintf(
		"HTTP/1.1 %d %s\r\nDate: %s\r\n",
		responseHeader.StatusCode,
		responseHeader.Description,
		responseHeader.Date,
	)

	if responseHeader.Connection == "close" {
		headers += "Connection: close\r\n\r\n"
		if _, err := writer.WriteString(headers); err != nil {
			log.Println(err)
		}
		return
	}
	if _, err := writer.WriteString(headers); err != nil {
		log.Println(err)
	}

	// Send file if required

	// Hint - Use the bufio package to write response

}
