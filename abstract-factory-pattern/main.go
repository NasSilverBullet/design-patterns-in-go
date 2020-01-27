package main

import (
	"fmt"
	"os"

	"github.com/NasSilverBullet/design-patterns-in-go/abstract-factory-pattern/factory"
	"github.com/NasSilverBullet/design-patterns-in-go/abstract-factory-pattern/listfactory"
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
	case "list":
		if err := makeList(); err != nil {
			return err
		}
		return nil

	case "table":
		if err := makeTable(); err != nil {
			return err
		}
	}

	usage()
	return nil
}

func makeList() error {
	f := factory.GetFactory(listfactory.New())

	asahi := f.Factory.CreateLink("朝日新聞", "https://www.asahi.com/")
	yomiuri := f.Factory.CreateLink("読売新聞", "https://www.yomiuri.co.jp/")

	usYahoo := f.Factory.CreateLink("Yahho!", "https://www.yahoo.com/")
	jpYahoo := f.Factory.CreateLink("Yahho!Japan", "https://www.yahoo.co.jp/")
	excite := f.Factory.CreateLink("Excite", "https://www.excite.com/")
	google := f.Factory.CreateLink("Google", "https://google.com/")

	trayNews := f.Factory.CreateTray("新聞")
	trayNews.Add(asahi)
	trayNews.Add(yomiuri)

	trayYahoo := f.Factory.CreateTray("Yahoo!")
	trayYahoo.Add(usYahoo)
	trayYahoo.Add(jpYahoo)

	traySearch := f.Factory.CreateTray("サーチエンジン")
	traySearch.Add(trayYahoo)
	traySearch.Add(excite)
	traySearch.Add(google)

	p := f.Factory.CreatePage("LinkPage", "結城 浩")
	p.Add(trayNews)
	p.Add(traySearch)

	file, err := os.OpenFile("index.html", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	p.Output(file)
	fmt.Println("index.html created")

	return nil
}

func makeTable() error {
	return nil
}

func usage() {
	fmt.Println("Usage: go run main.go list")
	fmt.Println("Usage: go run main.go table")
}
