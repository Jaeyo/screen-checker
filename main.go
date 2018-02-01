package main

import (
	"github.com/kbinani/screenshot"
	"fmt"
	"os"
	"image/png"
	"gopkg.in/alecthomas/kingpin.v2"
	"time"
	"image"
)

var (
	period = kingpin.Flag("period", "period").Default("5m").Short('p').Duration()
)

func main() {
	kingpin.Parse()
	println("period: ", period.String())
	mkdirs()

	for {
		err := createScreenShot()
		if err != nil {
			panic(err)
		}

		time.Sleep(*period)
	}
}

func mkdirs() {
	_ = os.Mkdir("./pics", os.ModePerm)
}

func createScreenShot() (error) {
	n := screenshot.NumActiveDisplays()
	for i := 0; i < n; i++ {
		img, err := captureScreenShot(i)
		if err != nil {
			return err
		}

		err = saveScreenShot(i, img)
		if err != nil {
			return err
		}
	}
	return nil
}

func captureScreenShot(screenNum int) (img *image.RGBA, err error){
	bounds := screenshot.GetDisplayBounds(screenNum)
	img, err = screenshot.CaptureRect(bounds)
	return
}

func saveScreenShot(screenNum int, img *image.RGBA) (err error) {
	now := time.Now().Format("2006-01-02_15:04:05")
	fileName := fmt.Sprintf("./pics/%s_%d.png", now, screenNum)
	println("filename: ", fileName)

	file, err := os.Create(fileName)
	if err != nil {
		return
	}

	defer file.Close()
	err = png.Encode(file, img)

	return
}
