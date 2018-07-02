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
    minR = 10
    maxR = 30
)

var (
    // colors = [3]string{"rgba(204, 204, 255, .5)", "rgba(204, 255, 204, .5)", "rgba(204, 234, 255, .5)"}
    colors = []string{"rgba(255, 0, 0, .5)", "rgba(255, 170, 0, .5)", "rgba(204, 204, 0, .5)", "rgba(0, 255, 0, .5)", "rgba(0, 0, 255, .5)", "rgba(204, 0, 255, .5)"}
    s *svg.SVG
)

type circle struct {
    X int
    Y int
    R int
}

func renderCircles(w http.ResponseWriter, req *http.Request) {
    w.Header().Set("Content-Type", "image/svg+xml")
    s = svg.New(w)
    s.Start(svgW, svgH)

    circles := generateCircles()

    for i := 0; i < len(circles); i++ {
        circle := circles[i]
        fill := colors[i % len(colors)] // rand.Intn(len(colors))]
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

     // var tryCircles [][2]circle
     // tryCircles = append(tryCircles, [2]circle{first, second})
     // for len(tryCircles) > 0 && len(circles) < 1000 {
     //     var try [2]circle
     //     try, tryCircles = tryCircles[0], tryCircles[1:]

     //     var newCircle circle
     //     newCircle = generateCirclesTangentToTwoCircles(try[0], try[1], 1)
     //     if circleFits(circles, newCircle) {
     //         circles = append(circles, newCircle)
     //         tryCircles = append(tryCircles, [2]circle{try[0], newCircle})
     //         tryCircles = append(tryCircles, [2]circle{try[1], newCircle})
     //         newCircle = generateSecondCircle(newCircle)
     //         if circleFits(circles, newCircle) {
     //             circles = append(circles, newCircle)
     //         }
     //     }
     //     newCircle = generateCirclesTangentToTwoCircles(try[0], try[1], -1)
     //     if circleFits(circles, newCircle) {
     //         circles = append(circles, newCircle)
     //         tryCircles = append(tryCircles, [2]circle{try[0], newCircle})
     //         tryCircles = append(tryCircles, [2]circle{try[1], newCircle})
     //         newCircle = generateSecondCircle(newCircle)
     //         if circleFits(circles, newCircle) {
     //             circles = append(circles, newCircle)
     //         }
     //     }
     //     log.Println(strconv.Itoa(len(circles)) + " " + strconv.Itoa(len(tryCircles)))
     // }

    third := generateCirclesTangentToTwoCircles(first, second, 1)
    circles = append(circles, third)

    fourth := generateCirclesTangentToTwoCircles(first, second, -1)
    circles = append(circles, fourth)

    // var circle circle
    // circle = generateCirclesTangentToTwoCircles(first, third, 1)
    // if (circleFits(circles, circle)) {
    //     circles = append(circles, circle)
    // }
    // circle = generateCirclesTangentToTwoCircles(first, third, -1)
    // if (circleFits(circles, circle)) {
    //     circles = append(circles, circle)
    // }
    // circle = generateCirclesTangentToTwoCircles(second, third, 1)
    // if (circleFits(circles, circle)) {
    //     circles = append(circles, circle)
    // }
    // circle = generateCirclesTangentToTwoCircles(second, third, -1)
    // if (circleFits(circles, circle)) {
    //     circles = append(circles, circle)
    // }
    // circle = generateCirclesTangentToTwoCircles(first, fourth, 1)
    // if (circleFits(circles, circle)) {
    //     circles = append(circles, circle)
    // }
    // circle = generateCirclesTangentToTwoCircles(first, fourth, -1)
    // if (circleFits(circles, circle)) {
    //     circles = append(circles, circle)
    // }
    // circle = generateCirclesTangentToTwoCircles(second, fourth, 1)
    // if (circleFits(circles, circle)) {
    //     circles = append(circles, circle)
    // }
    // circle = generateCirclesTangentToTwoCircles(second, fourth, -1)
    // if (circleFits(circles, circle)) {
    //     circles = append(circles, circle)
    // }
    // circles = append(circles, generateCirclesTangentToTwoCircles(first, third, -1))
    // circles = append(circles, generateCirclesTangentToTwoCircles(second, third, 1))

    return circles
}

func generateFirstCircle() circle {
    r := 60
    r = minR + rand.Intn(maxR - minR)
    x := 100
    x = 1 + r + rand.Intn((svgW - 2 * strokeW) - (2 * r))
    y := 200
    y = 1 + r + rand.Intn((svgH - 2 * strokeW) - (2 * r))

    return circle{x, y, r}
}

func generateSecondCircle(first circle) circle {
    second := circle{}

    second.R = 40
    second.R = minR + rand.Intn(maxR - minR)

    angle := .75 * math.Pi
    angle = rand.Float64() * math.Pi / 2
    rSum := float64(first.R + second.R)

    signX := 1.0
    if first.X > svgW / 2 {
        signX = -1
    }
    signY := 1.0
    if first.Y > svgH / 2 {
        signY = -1
    }

    second.X = int(math.Round(float64(first.X) + signX * rSum * math.Sin(angle)))
    second.Y = int(math.Round(float64(first.Y) + signY * rSum * math.Cos(angle)))

    return second
}

// http://jwilson.coe.uga.edu/EMAT6680Su06/Swanagan/Assignment7/BSAssignment7.html
func generateCirclesTangentToTwoCircles(first circle, second circle, angleOffset float64) circle {
    smaller, larger := first, second
    if smaller.R > larger.R {
       smaller, larger = larger, smaller
    }

    third := circle{}
    overlap := circle{}

    angle := math.Atan2(float64(smaller.Y - larger.Y), float64(smaller.X - larger.X))
    overlapAngle := angle + math.Pi * (rand.Float64() * .15 + .15) * angleOffset

    overlap.R = smaller.R
    overlapX := float64(larger.X) + float64(larger.R) * math.Cos(overlapAngle)
    overlapY := float64(larger.Y) + float64(larger.R) * math.Sin(overlapAngle)
    overlap.X = int(math.Round(overlapX))
    overlap.Y = int(math.Round(overlapY))

    // y = mx + b
    largerM := (overlapY - float64(larger.Y)) / (overlapX - float64(larger.X))
    largerB := float64(overlapY) - largerM * float64(overlapX)

    innerX := math.Cos(overlapAngle) * float64(larger.R - overlap.R) + float64(larger.X)
    innerY := math.Sin(overlapAngle) * float64(larger.R - overlap.R) + float64(larger.Y)

    smallerM := (innerY - float64(smaller.Y)) / (innerX - float64(smaller.X))
    // smallerB := float64(innerY) - smallerM * float64(innerX)

    distInnerSmall := math.Sqrt(math.Pow(float64(smaller.X) - innerX, 2) + math.Pow(float64(smaller.Y) - innerY, 2))

    smallAngle := math.Atan2(float64(smaller.Y) - innerY, float64(smaller.X) - innerX)
    tangentX := (float64(innerX) + float64(distInnerSmall / 2) * math.Cos(smallAngle))
    tangentY := (float64(innerY) + float64(distInnerSmall / 2) * math.Sin(smallAngle))
    perpendicularM := -1 / smallerM
    perpendicularB := tangentY - perpendicularM * tangentX

    // m1 * x + b1 = m2 * x + b2
    // m1 * x - m2 * x = b2 - b1
    // x = (b2 - b1) / (m1 - m2)
    thirdX := (perpendicularB - largerB) / (largerM - perpendicularM)
    thirdY := perpendicularM * float64(thirdX) + perpendicularB
    third.X = int(math.Round(thirdX))
    third.Y = int(math.Round(thirdY))
    third.R = int(math.Round(math.Sqrt(math.Pow(float64(larger.X) - thirdX, 2) + math.Pow(float64(larger.Y) - thirdY, 2)))) - larger.R

    // s.Circle(larger.X, larger.Y, 3, "stroke: black; fill: rgba(0, 0, 0, .2)")
    // s.Circle(overlap.X, overlap.Y, overlap.R, "stroke: black; fill: rgba(0, 0, 0, .2)")
    // s.Circle(overlap.X, overlap.Y, 3, "stroke: black; fill: rgba(0, 0, 0, .2)")
    // s.Line(0, int(largerM * 0 + largerB), svgW, int(largerM * svgW + largerB), "stroke:red; stroke-dasharray: 100,0;")
    // s.Circle(int(math.Cos(overlapAngle) * float64(larger.R - overlap.R)) + larger.X, int(math.Sin(overlapAngle) * float64(larger.R - overlap.R)) + larger.Y, 3, "stroke:black; fill: rgba(0, 0, 0, .2)")
    // s.Line(0, int(smallerM * 0 + smallerB), svgW, int(smallerM * svgW + smallerB), "stroke:black; stroke-dasharray: 100,0;")
    // s.Line(0, int(perpendicularM * 0 + perpendicularB), svgW, int(perpendicularM * svgW + perpendicularB), "stroke:black; stroke-dasharray: 100,0;")
    // s.Circle(third.X, third.Y, 3, "stroke: black; fill: rgba(0, 0, 0, .2)")

    return third
}

// https://en.wikipedia.org/wiki/Descartes%27_theorem
// https://math.stackexchange.com/questions/44406/how-do-i-get-the-square-root-of-a-complex-number/44414#44414?newreg=04ecb91435294d93a890c1e7228d912e
// https://en.wikipedia.org/wiki/De_Moivre%27s_formula
// https://www.khanacademy.org/math/precalculus/imaginary-and-complex-numbers/multiplying-and-dividing-complex-numbers-in-polar-form/a/complex-number-polar-form-review
// https://www.varsitytutors.com/hotmath/hotmath_help/topics/polar-form-of-a-complex-number


func circleFits(circles []circle, proposed circle) bool {
    for _, circle := range circles {
        if circlesIntersect(proposed, circle) {
            return false
        }
    }

    return true
}

func circlesIntersect(proposed circle, existing circle) bool {
    return math.Sqrt(math.Pow(float64(proposed.X - existing.X), 2) + math.Pow(float64(proposed.Y - existing.Y), 2)) + .2 < float64(proposed.R + existing.R)
}
