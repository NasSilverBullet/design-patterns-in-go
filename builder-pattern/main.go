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
		if err := makePlain(); err != nil {
			return err
		}
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

func makePlain() error {
	tb := textbuilder.New()
	d := director.New(tb)
	if err := d.Construct(); err != nil {
		return err
	}
	fmt.Println(tb.GetResult())
	return nil
}

func makeHTML() error {
	hb := htmlbuilder.New()
	d := director.New(hb)
	if err := d.Construct(); err != nil {
		return err
	}
	fmt.Println(hb.GetResult())
	return nil
}

func usage() {
	fmt.Println("Usage: go run main.go plain プレーンテキストで文章を作成")
	fmt.Println("Usage: go run main.go HTML  HTMLファイルで文章を作成")
}
