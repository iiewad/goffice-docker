package goffice

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

const (
	// ShareAndOutDir For Docker share Dir
	ShareAndOutDir = "/home/docker/ocr_storage"
	// ContainerNamePrefix For Docker container
	ContainerNamePrefix = "libreoffice"
)

// SofficeConvert Func
func SofficeConvert(args ...string) {
	start := time.Now()
	ch := make(chan string)

	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	for i, arg := range args {
		argArr := strings.Split(arg, "=>")
		filePath := argArr[0]
		targetFormat := argArr[1]
		convertFilePath := strings.Replace(filePath, currentDir, ShareAndOutDir, -1)
		go convert(convertFilePath, targetFormat, i+1, ch)
	}
	for range args {
		fmt.Println(<-ch)
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func convert(filePath, targetFormat string, index int, ch chan<- string) {
	start := time.Now()
	sofficeContainer := ContainerNamePrefix + strconv.Itoa(index)
	cmd := exec.Command("docker", "exec", sofficeContainer, "soffice", "--headless", "--convert-to", targetFormat, "--outdir", ShareAndOutDir, filePath)
	err := cmd.Run()
	if err != nil {
		ch <- fmt.Sprintf("cmd.Run failed with %s\n", err)
	}
	ch <- fmt.Sprintf("%.2fs elapsed\n", time.Since(start).Seconds())
}
