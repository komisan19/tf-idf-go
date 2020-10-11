package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"github.com/ikawaha/kagome/tokenizer"
)

var sc = bufio.NewScanner(os.Stdin)

func nextLine() string {
    sc.Scan()
    return sc.Text()
}

func main(){
	dic := tokenizer.SysDicSimple()
	t := tokenizer.NewWithDic(dic)
	tokens := t.Tokenize(nextLine())
	for _, token := range tokens {
		features := strings.Join(token.Features(), ",")
		fmt.Printf("%s\t%v\n", token.Surface, features)
	}
}
