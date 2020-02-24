package handle

import (
	"fmt"
	"testing"
)

func TestHandle(t *testing.T) {
	urls := make([]string, 0)
	urls = getUrls("https://so.gushiwen.org/authors/authorvsw_b90660e3e492A1.aspx", 10)
	for i:=0; i<10; i++{
		fmt.Println(urls[i])
	}

}
