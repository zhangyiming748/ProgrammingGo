package reWrite

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestMaster(t *testing.T) {
	if len(os.Args) == 1 ||
		(!strings.HasSuffix(os.Args[1], ".pls") &&
			!strings.HasSuffix(os.Args[1], ".m3u")) {
		fmt.Printf("usage: %s <file.[pls|m3u]>\n",
			filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	if rawBytes, err := ioutil.ReadFile(os.Args[1]); err != nil {
		log.Fatal(err)
	} else {
		data := string(rawBytes)
		if strings.HasSuffix(os.Args[1], "pls") {
			songs := readpls(data)
			writeM3Us(songs)
		} else if strings.HasSuffix(os.Args[1], "m3u") {
			songs := readM3uPlaylist(data)
			writePlsPlaylist(songs)
		}
	}
}
