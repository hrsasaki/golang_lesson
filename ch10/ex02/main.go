package main

import (
	"archive/tar"
	"archive/zip"
	"flag"
	"fmt"
	"io"
	"os"
)

var format = flag.String("f", "zip", "file format to encode. supported: zip, tar (default: zip)")

func main() {
	flag.Parse()
	fmt.Println(os.Args[1])
	switch *format {
	case "zip":
		if err := unzip(os.Args[1]); err != nil {
			fmt.Fprintf(os.Stderr, "unpack archived file: %v\n", err)
			os.Exit(1)
		}
	case "tar":
		if err := unpackTar(os.Args[1]); err != nil {
			fmt.Fprintf(os.Stderr, "unpack archived file: %v\n", err)
			os.Exit(1)
		}
	}
	if err := unzip(os.Args[1]); err != nil {
		fmt.Fprintf(os.Stderr, "unpack archived file: %v\n", err)
		os.Exit(1)
	}
}

func unzip(src string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		fmt.Printf("Contents of %s:\n", f.Name)
		rc, err := f.Open()
		if err != nil {
			return err
		}
		_, err = io.CopyN(os.Stdout, rc, 68)
		if err != nil {
			return err
		}
		rc.Close()
		fmt.Println()
	}
	return nil
}

func unpackTar(src string) error {
	f, err := os.Open(src)
	if err != nil {
		return err
	}
	r := tar.NewReader(f)

	for {
		header, err := r.Next()
		if err == io.EOF {
			// end of tar archive
			break
		}
		if err != nil {
			return err
		}
		fmt.Printf("Contents of %s:\n", header.Name)
		if _, err := io.Copy(os.Stdout, r); err != nil {
			return err
		}
		fmt.Println()
	}
	return nil
}
