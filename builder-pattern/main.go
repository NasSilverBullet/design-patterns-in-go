package main

import (
	"fmt"
	"os"

	"github.com/NasSilverBullet/design-patterns-in-go/builder-pattern/director"
	"github.com/NasSilverBullet/design-patterns-in-go/builder-pattern/htmlbuilder"
	"github.com/NasSilverBullet/design-patterns-in-go/builder-pattern/textbuilder"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	os.Exit(0)
}

func run() error {
	if len(os.Args) < 2 {
		usage()
		return nil
	}

	switch os.Args[1] {
	case "plain":
		makePlain()
		return nil

	case "HTML":
		if err := makeHTML(); err != nil {
			return err
		}
		return nil
	}

	usage()
	return nil
}

func makePlain() {
	tb := textbuilder.New()
	d := director.New(tb)
	d.Construct()
	fmt.Println(tb.GetResult())
}

func makeHTML() error {
	f, err := os.OpenFile("index.html", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer f.Close()
	hb := htmlbuilder.New(f)
	d := director.New(hb)
	d.Construct()
	fmt.Println(f.Name())
	return nil
}

func usage() {
	fmt.Println("Usage: go run main.go plain プレーンテキストで文章を作成")
	fmt.Println("Usage: go run main.go HTML  HTMLファイルで文章を作成")
}
