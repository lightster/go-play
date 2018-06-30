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
    svgW = 1680 // 7680
    svgH = 1050 // 4320
    strokeW = 0 // 2
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

    for len(circles) < 6500 { // }(svgW + svgH) {
        radius := 3 + rand.Intn(100 - 3)
        x := 1 + radius + rand.Intn((svgW - 2 * strokeW) - (2 * radius))
        y := 1 + radius + rand.Intn((svgH - 2 * strokeW) - (2 * radius))
        circle := circle{x, y, radius}
        if circleFits(circle, circles) {
            circles = append(circles, circle)
        }
    }

    return circles
}

func circleFits(proposed circle, circles []circle) bool {
    for _, circle := range circles {
        if circlesIntersect(proposed, circle) {
            return false
        }
    }

    return true
}

func circlesIntersect(proposed circle, existing circle) bool {
    return math.Sqrt(math.Pow(float64(proposed.X - existing.X), 2) + math.Pow(float64(proposed.Y - existing.Y), 2)) < float64(proposed.R + existing.R + strokeW)
}
