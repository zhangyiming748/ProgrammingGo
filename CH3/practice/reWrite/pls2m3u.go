package reWrite

import (
	"fmt"
	"log"
	"path/filepath"
	"strconv"
	"strings"
)

type song struct {
	Title    string
	Filename string
	Seconds  int
}

func pls2m3u(fp string) {
	song := readpls(fp)
	writeM3Us(song)
}
func readpls(fp string) (songs []song) {
	var s song
	for _, line := range strings.Split(fp, "\n") {
		if line = strings.TrimSpace(line); line == "" {
			continue
		}
		switch n, val := parseLine(line); n {
		case "File":
			s.Filename = strings.Map(
				mapPlatformDirSeparator, val)
		case "Title":
			s.Title = val
		case "Length":
			var err error
			if s.Seconds, err = strconv.Atoi(val); err != nil {
				log.Printf("failed to read the duration for '%s': %v\n",
					s.Title, err)
				s.Seconds = -1
			}
		}
		if s.Filename != "" && s.Title != "" && s.Seconds != 0 {
			songs = append(songs, s)
			s = song{} //初始化进入下次循环
		}
	}
	return songs
}

func mapPlatformDirSeparator(char rune) rune {
	if char == '/' || char == '\\' {
		return filepath.Separator
	}
	return char
}
func parseLine(line string) (name, value string) {
	const separator = "="
	if i := strings.Index(line, separator); i > -1 {
		if j := strings.IndexAny(line, "0123456789"); j > -1 && j < i {
			name = line[:j]
			value = line[i+len(separator):]
		}
	}
	return name, value
}
func writeM3Us(songs []song) {
	fmt.Println("#EXTM3U")
	for _, song := range songs {
		fmt.Printf("#EXTINF:%d,%s\n", song.Seconds, song.Title)
		fmt.Println(song.Filename)
	}
}
