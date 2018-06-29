package main

import (
	"github.com/ajstarks/svgo"
	"log"
	"math/rand"
	"net/http"
)

func main() {
    http.Handle("/circles", http.HandlerFunc(renderCircles))
    log.Println("Open http://localhost:2003/circles in a browser")
    err := http.ListenAndServe(":2003", nil)
    if err != nil {
        log.Fatal("ListenAndServe:", err)
    }
}

type circle struct {
    X int
    Y int
    R int
}


func renderCircles(w http.ResponseWriter, req *http.Request) {
    circles := generateCircles()

    w.Header().Set("Content-Type", "image/svg+xml")
    s := svg.New(w)
    s.Start(500, 500)
    for i := 0; i < len(circles); i++ {
        circle := circles[i]
        s.Circle(circle.X, circle.Y, circle.R, "fill:none;stroke:black")
    }
    s.End()
}

func generateCircles() []circle {
    var circles []circle;

    for i := 0; i < 100; i++ {
        radius := 3 + rand.Intn(100 - 3)
        circles = append(circles, circle{1 + radius + rand.Intn(498 - 2 * radius), 1 + radius + rand.Intn(498 - 2 * radius), radius})
    }

    return circles
}
