package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"scorpion/pkg/bmp"
	"scorpion/pkg/jpeg"
	"scorpion/pkg/png"
	"slices"
	"sync"
	"syscall"
	"time"
)

var validExts = []string{
	".jpg", ".jpeg", ".png", ".bmp",
}

func Usage() {
	fmt.Fprintln(os.Stderr,
		"usage: "+
			""+
			"./scorpion FILE [...FILE]",
	)
}

func commonData(file string, info os.FileInfo) string {
	stat := info.Sys().(*syscall.Stat_t)
	return fmt.Sprintf(
		"File Name:                   %s\n"+
			"Directory:                   %s\n"+
			"File Size:                   %d bytes\n"+
			"File Modification Date/Time: %s\n"+
			"File Access Date/Time        %s\n"+
			"File Inode Change Date/Time  %s\n"+
			"File Permissions             %s\n"+
			"File Type Extension:         %s\n",
		filepath.Base(file),
		filepath.Dir(file),
		info.Size(),
		info.ModTime().String(),
		time.Unix(stat.Atim.Sec, stat.Atim.Nsec).String(),
		time.Unix(stat.Ctim.Sec, stat.Ctim.Nsec).String(),
		info.Mode().Perm(),
		filepath.Ext(file),
	)
}

func main() {
	if len(os.Args[1:]) < 1 {
		fmt.Fprintln(os.Stderr, "Not enough arguments")
		Usage()
		os.Exit(1)
	}
	wg := &sync.WaitGroup{}
	c, cancel := context.WithCancel(context.Background())
	errs := make(chan error, 50)
	res := make(chan string, 50)
	for _, file := range os.Args[1:] {
		wg.Go(func() {
			ext := filepath.Ext(file)
			if !slices.Contains(validExts, ext) {
				errs <- fmt.Errorf("%s: won't do", file)
				return
			}
			info, err := os.Stat(file)
			if err != nil {
				errs <- err
				return
			}
			f, err := os.OpenFile(file, os.O_RDONLY, 0644)
			defer f.Close()
			if err != nil {
				errs <- err
				return
			}
			var imgInfo string
			switch ext {
			case ".png":
				pngInfo, err := png.Png(f)
				if err != nil {
					errs <- fmt.Errorf("%s: %s", file, err)
					return
				}
				imgInfo = pngInfo
			case ".jpeg", ".jpg":
				jpegInfo, err := jpeg.Jpeg(f)
				if err != nil {
					errs <- fmt.Errorf("%s: %s", file, err)
				}
				imgInfo = jpegInfo
			case ".bmp":
				bmpInfo, err := bmp.Bmp(f)
				if err != nil {
					errs <- fmt.Errorf("%s: %s", file, err)
					return
				}
				imgInfo = bmpInfo
			default:
				errs <- fmt.Errorf("%s: won't do", file)
				return
			}
			output := commonData(file, info)
			res <- output + imgInfo
		})
	}
	go func() {
		wg.Wait()
		close(errs)
		close(res)
		cancel()
	}()
	for {
		select {
		case e := <-errs:
			if e != nil {
				fmt.Println(e.Error())
			}
		case r, ok := <-res:
			if !ok {
				return
			}
			fmt.Println(r)
		case <-c.Done():
			return
		}
	}
}
