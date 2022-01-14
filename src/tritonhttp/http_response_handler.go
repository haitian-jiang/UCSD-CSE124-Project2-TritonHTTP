package tritonhttp

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	HTML404Path = "./src/404.html"
	HTML400Path = "./src/400.html"
)

func (hs *HttpServer) handleBadRequest(conn net.Conn) {
	// todo: add body
	now := time.Now().UTC().Format("Mon, 02 Jan 2006 15:04:05 GMT")
	respHeader := HttpResponseHeader{
		StatusCode:  400,
		Description: "Bad Request",
		Date:        now,
		Connection:  "close",
		FilePath:    HTML400Path,
	}
	hs.sendResponse(respHeader, conn)
}

func (hs *HttpServer) handleFileNotFoundRequest(requestHeader *HttpRequestHeader, conn net.Conn) {
	// todo: add body
	now := time.Now().UTC().Format("Mon, 02 Jan 2006 15:04:05 GMT")
	responseHeader := HttpResponseHeader{
		StatusCode:  404,
		Description: "Not Found",
		Date:        now,
		FilePath:    HTML404Path,
	}
	hs.sendResponse(responseHeader, conn)
}

func (hs *HttpServer) handleResponse(requestHeader *HttpRequestHeader, conn net.Conn) (result string) {
	//panic("todo - handleResponse")
	now := time.Now().UTC().Format("Mon, 02 Jan 2006 15:04:05 GMT")
	responseHeader := *new(HttpResponseHeader)
	responseHeader.Date = now

	// response for closing connection
	if requestHeader.Connection == "close" {
		responseHeader.StatusCode = 200
		responseHeader.Description = "OK"
		responseHeader.Connection = "close"
		hs.sendResponse(responseHeader, conn)
		return "close"
	}

	// check hostname
	host := strings.Split(requestHeader.Host, ":")[0]
	docRoot, ok := hs.DocRoot[strings.ToLower(host)]
	if !ok {
		hs.handleBadRequest(conn)
		return "400"
	}

	// check requested file
	url := filepath.Clean(requestHeader.URL)
	if strings.HasSuffix(url, "/") {
		url += "index.html"
	}
	file := docRoot + url
	_, err := os.Stat(file)
	if err != nil {
		if os.IsNotExist(err) {
			// 404 file not found
			hs.handleFileNotFoundRequest(requestHeader, conn)
			return "404"
		} else {
			log.Println(err)
		}
	}
	// todo nomal
	responseHeader.StatusCode = 200
	responseHeader.Description = "OK"
	responseHeader.FilePath = file
	hs.sendResponse(responseHeader, conn)
	return "200"
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
