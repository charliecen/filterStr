package src

import (
	"bufio"
	"errors"
	"github.com/syyongx/go-wordsfilter"
	"io"
	"os"
	"strings"
)

type filterStr struct {}

var FilterStr = new(filterStr)
var sensitivityConfigInit []string = FilterStr.sensitivityConfigInit("./config/sensitivity.txt")

func (s *filterStr) FilStr(name string) error {
	wf := wordsfilter.New()
	root := wf.Generate(sensitivityConfigInit)
	check := wf.Contains(name, root)
	if check {
		return errors.New("存在敏感词")
	}
	return nil
}

func (s *filterStr) sensitivityConfigInit(paths string) []string {
	texts := make([]string, 0)
	f, _ := os.Open(paths)
	defer f.Close()
	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n')
		line = strings.Trim(line, "\n")
		sliceStr := strings.Split(line, "|")
		texts = append(texts, sliceStr[0])
		if err != nil || err == io.EOF {
			if err != nil {
				return texts
			}
			break
		}
	}
	return texts
}