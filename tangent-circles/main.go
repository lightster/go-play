package main

import (
    "github.com/ajstarks/svgo"
    "log"
    "math"
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
    svgW = 800 // 1680 // 7680
    svgH = 600 // 1050 // 4320
    strokeW = 0 // 2
    minR = 3
    maxR = 30
)

var (
    colors = [3]string{"rgb(204, 204, 255)", "rgb(204, 255, 204)", "rgb(204, 234, 255)"}
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
        fill := colors[rand.Intn(len(colors))]
        circleStyle := "fill:" + fill + ";"
        if strokeW > 0 {
            circleStyle += "stroke:black;stroke-width:" + strconv.Itoa(strokeW)
        }
        s.Circle(circle.X, circle.Y, circle.R, circleStyle)
    }
    s.End()
}

func generateCircles() []circle {
    var circles []circle

    first := generateFirstCircle()
    circles = append(circles, first)

    second := generateSecondCircle(first)
    circles = append(circles, second)

    return circles
}

func generateFirstCircle() circle {
    r := minR + rand.Intn(maxR - minR)
    x := 1 + r + rand.Intn((svgW - 2 * strokeW) - (2 * r))
    y := 1 + r + rand.Intn((svgH - 2 * strokeW) - (2 * r))

    return circle{x, y, r}
}

func generateSecondCircle(first circle) circle {
    second := circle{}

    second.R = minR + rand.Intn(maxR - minR)

    angle := rand.Float64() * math.Pi / 2
    rSum := float64(first.R + second.R)

    signX := 1.0
    if first.X > svgW / 2 {
        signX = -1
    }
    signY := 1.0
    if first.Y > svgH / 2 {
        signY = -1
    }

    second.X = int(math.Floor(float64(first.X) + signX * rSum * math.Sin(angle)))
    second.Y = int(math.Floor(float64(first.Y) + signY * rSum * math.Cos(angle)))

    return second
}
