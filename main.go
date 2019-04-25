package main

import (
	"fmt"

	"github.com/iiewad/gofiice-docker/goffice"
)

func main() {
	fmt.Println("Start")
	args := []string{
		"/Users/iewad/go/src/github.com/iiewad/gofiice-docker/test-files/SamplePPTFile_500kb.ppt=>pdf",
		"/Users/iewad/go/src/github.com/iiewad/gofiice-docker/test-files/SampleDOCFile_1000kb.doc=>pdf",
		"/Users/iewad/go/src/github.com/iiewad/gofiice-docker/test-files/SamplePPTFile_1000kb.ppt=>pdf",
		"/Users/iewad/go/src/github.com/iiewad/gofiice-docker/test-files/SampleXLSFile_904kb.xls=>txt"}
	goffice.SofficeConvert(args...)
}
