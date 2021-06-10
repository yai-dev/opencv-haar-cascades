package main

import (
    "fmt"
    "image/color"
    "log"

    "gocv.io/x/gocv"
)

func main() {
    cam, err := gocv.VideoCaptureDevice(0)
    if err != nil {
        log.Fatalf("error captured when open cam [0]: %s", err)
    }

    defer func(cam *gocv.VideoCapture) {
        err := cam.Close()
        if err != nil {
            panic(fmt.Sprintf("panic when closing the opened cam [0]: %s", err))
        }
    }(cam)

    img := gocv.NewMat()
    defer func(img *gocv.Mat) {
        err := img.Close()
        if err != nil {
            panic(fmt.Sprintf("panic when close matrix: %s", err))
        }
    }(&img)

    window := gocv.NewWindow("OpenCV Haar Cascades Demo")

    defer func(window *gocv.Window) {
        err := window.Close()
        if err != nil {
            panic(fmt.Sprintf("panic when closing opened window: %s", err))
        }
    }(window)

    harrCascade := "haarcascade_frontalface_default.xml"

    classifier := gocv.NewCascadeClassifier()
    classifier.Load(harrCascade)
    defer func(classifier *gocv.CascadeClassifier) {
        err := classifier.Close()
        if err != nil {
            panic(fmt.Sprintf("panic when closing opened Cascade Classifier: %s", err))
        }
    }(&classifier)

    c := color.RGBA{0, 255, 0, 0}

    for {
        if ok := cam.Read(&img); !ok || img.Empty() {
            log.Println("[Cam] Unable to read from cam [0]")
            continue
        }

        rects := classifier.DetectMultiScale(img)
        for _, rect := range rects {
            log.Printf("[Classifier] Face detected: %s\n", rect)
            gocv.Rectangle(&img, rect, c, 3)
        }

        window.IMShow(img)
        window.WaitKey(10)
    }
}
