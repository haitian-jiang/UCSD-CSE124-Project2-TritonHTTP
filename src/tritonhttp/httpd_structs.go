package tritonhttp

type HttpServer	struct {
	ServerPort	string
	DocRoot		map[string]string
	MIMEPath	string
	MIMEMap		map[string]string
}

type HttpResponseHeader struct {
	// Add any fields required for the response here
	StatusCode     int
	Date           string
	LastModified   string
	ContentType    string
	ContentLength  int
	Connection     string
}

type HttpRequestHeader struct {
	// Add any fields required for the request here

}
