package tritonhttp

type HttpServer struct {
	ServerPort string
	DocRoot    map[string]string
	MIMEPath   string
	MIMEMap    map[string]string
}

type HttpResponseHeader struct {
	// Add any fields required for the response here
	StatusCode    int
	Description   string
	Date          string
	LastModified  string
	ContentType   string
	ContentLength int64
	Connection    string
	FilePath      string
}

type HttpRequestHeader struct {
	// Add any fields required for the request here
	URL        string
	Host       string
	Connection string
	KeyValue   map[string]string
}
