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

	now := time.Now().UTC().Format("Mon, 02 Jan 2006 15:04:05 GMT")
	responseHeader := *new(HttpResponseHeader)
	responseHeader.Date = now

	// check hostname
	host := strings.Split(requestHeader.Host, ":")[0]
	docRoot, ok := hs.DocRoot[strings.ToLower(host)]
	if !ok {
		hs.handleBadRequest(conn)
		return "400"
	}

	// check requested file
	url := filepath.Clean(requestHeader.URL)
	file := docRoot + url
	stat, err := os.Stat(file)
	if err != nil {
		if os.IsNotExist(err) {
			// 404 file not found
			hs.handleFileNotFoundRequest(requestHeader, conn)
			return "404"
		} else {
			log.Println(err)
		}
	} else if stat.IsDir() {
		file += "/index.html"
	}

	// acceptable requests
	responseHeader.StatusCode = 200
	responseHeader.Description = "OK"
	responseHeader.FilePath = file
	responseHeader.Connection = requestHeader.Connection
	hs.sendResponse(responseHeader, conn)
	if responseHeader.Connection == "close" {
		return "close"
	}
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

	// construct headers
	headers := fmt.Sprintf(
		"HTTP/1.1 %d %s\r\nDate: %s\r\n",
		responseHeader.StatusCode,
		responseHeader.Description,
		responseHeader.Date,
	)

	// extension
	ext := filepath.Ext(responseHeader.FilePath)
	contentType, ok := hs.MIMEMap[ext]
	if ok {
		responseHeader.ContentType = contentType
		headers += fmt.Sprintf("Content-Type: %s\r\n", responseHeader.ContentType)
	}

	// file meta info
	stat, _ := os.Stat(responseHeader.FilePath)
	responseHeader.LastModified = stat.ModTime().Format("Mon, 02 Jan 2006 15:04:05 GMT")
	responseHeader.ContentLength = stat.Size()
	headers += fmt.Sprintf("Last-Modified: %s\r\n", responseHeader.LastModified)
	headers += fmt.Sprintf("Content-Length: %d\r\n", responseHeader.ContentLength)

	// header end
	if responseHeader.Connection == "close" {
		headers += "Connection: close\r\n\r\n"
	} else {
		headers += "\r\n"
	}

	if _, err := writer.WriteString(headers); err != nil {
		log.Println(err)
	}

	// Send file if required
	content, err := os.ReadFile(responseHeader.FilePath) // go 1.16 or newer, otherwise use ioutil.ReadFile
	if err != nil {
		log.Println(err)
	}
	if _, err = writer.Write(content); err != nil {
		log.Println(err)
	}
	// Hint - Use the bufio package to write response
}
