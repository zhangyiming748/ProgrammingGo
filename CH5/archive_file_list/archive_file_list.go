package main

import (
	"archive/tar"
	"archive/zip"
	"compress/bzip2"
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

var FunctionForSuffix = map[string]func(string) ([]string, error){
	".gz": GzipFileList, ".tar": TarFileList, ".tar.gz": TarFileList,
	".tar.bz2": TarFileList, ".tgz": TarFileList, ".zip": ZipFileList}

func main() {
	if len(os.Args) == 1 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Printf("usage: %s archive1 [archive2 [... archiveN]]\n",
			filepath.Base(os.Args[0]))
		os.Exit(1)

	}
	args := commandLineFiles(os.Args[1:])
	for _, filename := range args {
		fmt.Print(filename)
		lines, err := ArchiveFileList(filename)
		if err != nil {
			fmt.Println(" ERROR:", err)
		} else {
			fmt.Println()
			for _, line := range lines {
				fmt.Println(" ", line)
			}
		}
	}
}

func commandLineFiles(files []string) []string {
	if runtime.GOOS == "windows" {
		args := make([]string, 0, len(files))
		for _, name := range files {
			if matches, err := filepath.Glob(name); err != nil {
				args = append(args, name) // Invalid pattern
			} else if matches != nil { // At least one match
				args = append(args, matches...)
			}
		}
		return args
	}
	return files
}

func ArchiveFileList(file string) ([]string, error) {
	if function, ok := FunctionForSuffix[Suffix(file)]; ok {
		return function(file)
	}
	return nil, errors.New("unrecognized archive")
}

func Suffix(file string) string {
	file = strings.ToLower(filepath.Base(file))
	if i := strings.LastIndex(file, "."); i > -1 {
		if file[i:] == ".bz2" || file[i:] == ".gz" || file[i:] == ".xz" {
			if j := strings.LastIndex(file[:i], ".");
				j > -1 && strings.HasPrefix(file[j:], ".tar") {
				return file[j:]
			}
		}
		return file[i:]
	}
	return file
}

func ZipFileList(filename string) ([]string, error) {
	zipReader, err := zip.OpenReader(filename)
	if err != nil {
		return nil, err
	}
	defer zipReader.Close()
	var files []string
	for _, file := range zipReader.File {
		files = append(files, file.Name)
	}
	return files, nil
}

func GzipFileList(filename string) ([]string, error) {
	reader, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	gzipReader, err := gzip.NewReader(reader)
	if err != nil {
		return nil, err
	}
	return []string{gzipReader.Header.Name}, nil
}

func TarFileList(filename string) ([]string, error) {
	reader, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	var tarReader *tar.Reader
	if strings.HasSuffix(filename, ".gz") ||
		strings.HasSuffix(filename, ".tgz") {
		gzipReader, err := gzip.NewReader(reader)
		if err != nil {
			return nil, err
		}
		tarReader = tar.NewReader(gzipReader)
	} else if strings.HasSuffix(filename, ".bz2") {
		bz2Reader := bzip2.NewReader(reader)
		tarReader = tar.NewReader(bz2Reader)
		openbz2(tarReader)
	} else {
		tarReader = tar.NewReader(reader)
	}
	var files []string
	for {
		header, err := tarReader.Next()
		if err != nil {
			if err == io.EOF {
				break
			}
			return files, err
		}
		if header == nil {
			break
		}
		files = append(files, header.Name)
	}
	return files, nil
}
func openbz2(tr *tar.Reader) {

	//var srcFile = "passwd.tar"
	//
	//// 将 tar 包打开
	//fr, err := os.Open(srcFile)
	//ErrPrintln(err)
	//defer fr.Close()
	//
	//// 通过 fr 创建一个 tar.*Reader 结构，然后将 tr 遍历，并将数据保存到磁盘中
	//tr := tar.NewReader(fr)

	for hdr, err := tr.Next(); err != io.EOF; hdr, err = tr.Next() {
		// 处理 err ！= nil 的情况
		ErrPrintln(err)
		// 获取文件信息
		fi := hdr.FileInfo()

		// 创建一个空文件，用来写入解包后的数据
		fw, err := os.Create(fi.Name())
		ErrPrintln(err)

		// 将 tr 写入到 fw
		n, err := io.Copy(fw, tr)
		ErrPrintln(err)
		log.Printf("解包： %s  ，共处理了 %d 个字符的数据。", fi.Name(), n)

		// 设置文件权限，这样可以保证和原始文件权限相同，如果不设置，会根据当前系统的 umask 来设置。
		os.Chmod(fi.Name(), fi.Mode().Perm())

		// 注意，因为是在循环中，所以就没有使用 defer 关闭文件
		// 如果想使用 defer 的话，可以将文件写入的步骤单独封装在一个函数中即可
		fw.Close()
	}
}

func ErrPrintln(err error) {
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
}
