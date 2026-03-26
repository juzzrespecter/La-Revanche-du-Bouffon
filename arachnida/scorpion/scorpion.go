package main

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"scorpion/pkg/bmp"
	"scorpion/pkg/exif"
	"scorpion/pkg/png"
	"slices"
	"sync"
	"syscall"
	"time"
)

var validExts = []string{
	".jpg", ".jpeg", ".png", ".bmp",
}

var magicNumbers = map[string][]byte{
	".jpg":  {0xFF, 0xD8},
	".jpeg": {0xFF, 0xD8},
	".png":  {0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A},
	".bmp":  {0x42, 0x4D},
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
			f, err := os.OpenFile(file, os.O_RDONLY, 644)
			if err != nil {
				errs <- err
				return
			}

			switch ext {
			case ".png":
				magic := make([]byte, 8)
				f.Read(magic)
				if !bytes.Equal(magic, magicNumbers[ext]) {
					errs <- fmt.Errorf("%s: not a png file", file)
					return
				}
				pngInfo, err := png.Png(f)
				if err != nil {
					errs <- err
					return
				}
				output := commonData(file, info)
				res <- output + pngInfo

			case ".jpeg", ".jpg":
				magic := make([]byte, 2)
				f.Read(magic)
				if !bytes.Equal(magic, magicNumbers[ext]) {
					errs <- fmt.Errorf("%s: not a jpeg file", file)
					return
				}
				jpegInfo, err := exif.Exif(f)
				if err != nil {
					errs <- err
					return
				}
				output := commonData(file, info)
				res <- output + jpegInfo
			case ".bmp":
				magic := make([]byte, 2)
				f.Read(magic)
				if !bytes.Equal(magic, magicNumbers[ext]) {
					errs <- fmt.Errorf("%s: not a bmp file", file)
					return
				}
				bmpInfo, err := bmp.Bmp(f)
				if err != nil {
					errs <- err
					return
				}
				output := commonData(file, info)
				res <- output + bmpInfo
			default:
				fmt.Println("Diantres")
			}

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
		case r := <-res:
			fmt.Println(r)
		case <-c.Done():
			return
		}
	}
}
