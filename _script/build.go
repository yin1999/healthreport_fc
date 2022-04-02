package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"os/exec"
)

func main() {
	var targets []string
	if len(os.Args) == 1 {
		targets = []string{"aliyun", "tencent"}
	} else {
		targets = os.Args[1:]
	}
	os.Setenv("GOOS", "linux")
	os.Setenv("GOARCH", "amd64")
	for _, target := range targets {
		var out string
		switch target {
		case "aliyun":
			out = "main"
		case "tencent":
			out = "bootstrap"
		default:
			panic("not supported platform")
		}
		cmd := exec.Command("go",
			"build",
			"-trimpath",
			"-ldflags", "-s -w -buildid=",
			"-tags", target,
			"-o", out,
		)
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Printf("%s\n", string(output))
			panic(err)
		}
		err = zipFile(target+"-serverless.zip", out)
		if err != nil {
			fmt.Printf("%s\n", string(output))
			panic(err)
		}
		os.Remove(out)
	}
}

func zipFile(output, file string) (err error) {
	var out *os.File
	out, err = os.OpenFile(output, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return
	}
	defer func() {
		out.Close()
		if err != nil {
			os.Remove(output)
		}
	}()
	w := zip.NewWriter(out)
	var writer io.Writer
	header := &zip.FileHeader{
		Name:   file,
		Method: zip.Deflate,
	}
	header.SetMode(0755) // set executable permission
	writer, err = w.CreateHeader(header)
	if err != nil {
		return
	}
	var in *os.File
	in, err = os.Open(file)
	if err != nil {
		return
	}
	defer in.Close()
	if _, err = io.Copy(writer, in); err != nil {
		return
	}
	err = w.Close()
	return
}
