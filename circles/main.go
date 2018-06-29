package main

import (
    "github.com/ajstarks/svgo"
    "log"
    "math/rand"
    "net/http"
    "strconv"
)

func main() {
    http.Handle("/circles", http.HandlerFunc(renderCircles))
    log.Println("Open http://localhost:2003/circles in a browser")
    err := http.ListenAndServe(":2003", nil)
    if err != nil {
        log.Fatal("ListenAndServe:", err)
    }
}

const (
    svgW = 7680
    svgH = 4320
    strokeW = 2
)

type circle struct {
    X int
    Y int
    R int
}


func renderCircles(w http.ResponseWriter, req *http.Request) {
    circles := generateCircles()

    w.Header().Set("Content-Type", "image/svg+xml")
    s := svg.New(w)
    s.Start(svgW, svgH)
    for i := 0; i < len(circles); i++ {
        circle := circles[i]
        circleStyle := "fill:none;stroke:black;stroke-width:" + strconv.Itoa(strokeW)
        s.Circle(circle.X, circle.Y, circle.R, circleStyle)
    }
    s.End()
}

func generateCircles() []circle {
    var circles []circle;

    for i := 0; i < 1000; i++ {
        radius := 3 + rand.Intn(100 - 3)
        x := 1 + radius + rand.Intn((svgW - 2 * strokeW) - (2 * radius))
        y := 1 + radius + rand.Intn((svgH - 2 * strokeW) - (2 * radius))
        circle := circle{x, y, radius}
        circles = append(circles, circle)
    }

    return circles
}
