package CH9

import (
	"fmt"
	"github.com/beevik/etree"
	"strings"
	"testing"
)

func TestEtree(t *testing.T) {
	// 初始化根节点
	doc := etree.NewDocument()
	if err := doc.ReadFromFile("mode.xml"); err != nil {
		panic(err)
	}
	root := doc.SelectElement("rss")
	servers := root.SelectElement("channel")
	for _, server := range servers.SelectElements("item") {
		fmt.Println(server.SelectElement("title").Text())
	}

}
