package main

import (
	"flag"
	"fyne.io/fyne/container"
	"image"
	"image/draw"
	"image/gif"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/widget"
)

func main() {
	flag.Parse()

	if flag.NArg() == 0 {
		// Read from stdin
		if err := display(os.Stdin); err != nil {
			log.Fatal(err)
		}
	} else {
		for _, filename := range flag.Args() {
			if strings.HasPrefix(filename, "http://") || strings.HasPrefix(filename, "https://") {
				res, err := http.Get(filename)
				if err != nil {
					log.Fatal(err)
				}
				defer res.Body.Close()

				if err := display(res.Body); err != nil {
					log.Fatal(err)
				}
			} else {
				// Skip errors and directories
				if fi, err := os.Stat(filename); err != nil || fi.IsDir() {
					continue
				}

				f, err := os.Open(filename)
				if err != nil {
					log.Fatal(err)
				}
				defer f.Close()

				if err := display(f); err != nil {
					log.Fatal(err)
				}
			}
		}
	}
}

func display(r io.Reader) error {
	anigif, err := gif.DecodeAll(r)
	if err != nil {
		return err
	}
	loop := gifLoop{
		src:      anigif,
		dst:      canvas.NewImageFromImage(nil),
		stopping: false,
	}

	a := app.New()
	w := a.NewWindow("Hello")

	hello := &widget.Label{Text: "Hello Fyne!"}
	//anigif.Config.Width
	w.SetContent(container.NewGridWithColumns(1,

		hello,
		&widget.Button{Text: "Hi!", OnTapped: func() {
			hello.SetText("Welcome :)")
		}},
		loop.dst,
	))

	w.ShowAndRun()
	return nil
}

type gifLoop struct {
	src      *gif.GIF
	dst      *canvas.Image
	stopping bool
}

func (g *gifLoop) stop() {
	g.stopping = true
}

func (g *gifLoop) run() {
	size := g.src.Image[0].Bounds().Size()
	overpaintImage := image.NewRGBA(image.Rect(0, 0, size.X, size.Y))
	draw.Draw(overpaintImage, overpaintImage.Bounds(), g.src.Image[0], image.ZP, draw.Src)
	g.dst.Image = overpaintImage

	for {
		for c, srcImg := range g.src.Image {
			if g.stopping {
				return
			}

			draw.Draw(overpaintImage, overpaintImage.Bounds(), srcImg, image.ZP, draw.Over)
			canvas.Refresh(g.dst)

			time.Sleep(time.Millisecond * time.Duration(g.src.Delay[c]) * 10)
		}
	}
}
