package main

import (
	"flag"
	"log"
	"os"
	"runtime/debug"
	_ "net/http/pprof"
	"net/http"
)

import (
	"github.com/sternix/wl"
	"github.com/sternix/wl/ui"
)

func init() {
	flag.Parse()
	log.SetFlags(0)
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			debug.PrintStack()
		}
	}()

	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()


	exitChan := make(chan bool, 10)

	if flag.NArg() == 0 {
		log.Fatalf("usage: %s imagefile", os.Args[0])
	}

	img, err := ImageFromFile(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}

	display, err := ui.Connect("")
	if err != nil {
		log.Fatal(err)
	}

	b := img.Bounds()
	w := int32(b.Dx())
	h := int32(b.Dy())
	window, err := display.NewWindow(w, h)
	if err != nil {
		log.Fatal(err)
	}

	keh := func(ev interface{}) {
		if kev, ok := ev.(wl.KeyboardKeyEvent); ok {
			if kev.Key == 16 {
				exitChan <- true
			}
		}
	}

	kh := wl.HandlerFunc(keh)
	display.Keyboard().AddKeyHandler(kh)

	window.Draw(img)

loop:
	for {
		select {
		case <-exitChan:
			break loop
		case display.Dispatch() <- true:
		}
	}

	log.Print("Loop finished")
	window.Dispose()
	display.Disconnect()
}
