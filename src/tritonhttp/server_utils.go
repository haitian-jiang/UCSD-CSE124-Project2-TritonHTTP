package tritonhttp

import (
	"bufio"
	"log"
	"os"
	"strings"
)

/** 
	Load and parse the mime.types file 
**/
func ParseMIME(MIMEPath string) (MIMEMap map[string]string, err error) {
	log.Println("Reading MIME file " + MIMEPath)
	MIMEMap = make(map[string]string)

	f, err :=os.Open(MIMEPath)
	if err != nil {
		return MIMEMap, err
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		line := strings.TrimSpace(s.Text())
		fields := strings.Split(line, " ")
		extension := strings.TrimSpace(fields[0])
		mimeType := strings.TrimSpace(fields[1])
		MIMEMap[extension] = mimeType
	}

	if err = s.Err(); err != nil {
		return MIMEMap, err
	}
	return MIMEMap, nil
}

