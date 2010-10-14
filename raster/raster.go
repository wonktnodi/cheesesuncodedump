package raster

import (
    "image"
    "math"

    "sdl"
)

// non-bressenham line
// TODO: figure out just how much better bressenham is
func line(dst *sdl.Surface, p0, p1 Point, c image.Color) {
    d := p1.Sub(p0)
    l := d.ToFloat().Length()
    inc := int32(1)

    if d.X*d.X > d.Y*d.Y {
        if d.X < 0 {
            inc = -1
        }
        m := d.ToFloat().Mult(1.0/l).Slope()
        for x := int32(0); x != d.X; x += inc {
            y := int32(m * float64(x))
            dst.Set(int(x + p0.X), int(y + p0.Y), c)
        }
    } else {
        if d.Y < 0 {
            inc = -1
        }
        m := d.ToFloat().Inverse().Mult(1.0/l).Slope()
        for y := int32(0); y != d.Y; y += inc {
            x := int32(m * float64(y))
            dst.Set(int(x + p0.X), int(y + p0.Y), c)
        }
    }
}

func filledCircle(dst *sdl.Surface, org Point, r int) {
    colour := image.RGBAColor{0, 255, 0, 255}
    // add a bias for smoother poles
    br := float64(r) + 0.5
    for x := -r; x <= r; x++ {
        dy := int(math.Sqrt(br*br - float64(x*x)))
        for y := -dy; y <= dy; y++ {
            dst.Set(int(org.X) + x, int(org.Y) + y, colour)
        }
    }
}

func circle(dst *sdl.Surface, org Point, r int) {
    colour := image.RGBAColor{0, 255, 255, 0}

    // set the poles
    dst.Set(int(org.X), int(org.Y) + r, colour)
    dst.Set(int(org.X), int(org.Y) - r, colour)
    dst.Set(int(org.X) + r, int(org.Y), colour)
    dst.Set(int(org.X) - r, int(org.Y), colour)

    // add a bias for smoother poles
    br := float64(r) + 0.5

    // draw and rotate an octant
    var x int32 = 1
    for {
        y := int32(math.Sqrt(br*br - float64(x*x)))
        if x < y {
            dst.Set(int(org.X + x), int(org.Y + y), colour)
            dst.Set(int(org.X - x), int(org.Y + y), colour)
            dst.Set(int(org.X + x), int(org.Y - y), colour)
            dst.Set(int(org.X - x), int(org.Y - y), colour)
            dst.Set(int(org.X + y), int(org.Y + x), colour)
            dst.Set(int(org.X - y), int(org.Y + x), colour)
            dst.Set(int(org.X + y), int(org.Y - x), colour)
            dst.Set(int(org.X - y), int(org.Y - x), colour)
        } else {
            if x == y {
                // draw the NE, SE, SW, and NW pixels
                dst.Set(int(org.X + x), int(org.Y + y), colour)
                dst.Set(int(org.X - x), int(org.Y + y), colour)
                dst.Set(int(org.X + x), int(org.Y - y), colour)
                dst.Set(int(org.X - x), int(org.Y - y), colour)
            }
            break
        }

        x++
    }
}

func filledEllipse(dst *sdl.Surface, org Point, a, b int) {
    colour := image.RGBAColor{0, 255, 0, 255}
    // add a bias to a and b for smoother poles
    ba := float64(a) + 0.5
    bb := float64(b) + 0.5
    for x := -a; x <= a; x++ {
        dy := int(math.Sqrt((bb*bb) * (1.0 - float64(x*x)/(ba*ba))))
        for y := -dy; y <= dy; y++ {
            dst.Set(int(org.X) + x, int(org.Y) + y, colour)
        }
    }
}

func ellipse(dst *sdl.Surface, org Point, a, b int) {
    colour := image.RGBAColor{0, 255, 0, 255}

    // set the poles
    dst.Set(int(org.X), int(org.Y) + b, colour)
    dst.Set(int(org.X), int(org.Y) - b, colour)
    dst.Set(int(org.X) + a, int(org.Y), colour)
    dst.Set(int(org.X) - a, int(org.Y), colour)

    // add a bias to a and b for smoother poles
    ba := float64(a) + 0.5
    bb := float64(b) + 0.5

    // draw and rotate a quadrant
    for x := 1; x < a; x++ {
        y1 := (-bb*float64(x)) / (ba * math.Sqrt(ba*ba - float64(x*x)))
        y := int(math.Sqrt((bb*bb) * (1.0 - float64(x*x)/(ba*ba))))
        if y1 > -1.0 {
            dst.Set(int(org.X) + x, int(org.Y) + y, colour)
            dst.Set(int(org.X) - x, int(org.Y) + y, colour)
            dst.Set(int(org.X) + x, int(org.Y) - y, colour)
            dst.Set(int(org.X) - x, int(org.Y) - y, colour)
        } else {
            for dy := 1; dy <= y; dy++ {
                dx := int(math.Sqrt((ba*ba) * (1.0 - float64(dy*dy)/(bb*bb))))
                dst.Set(int(org.X) + dx, int(org.Y) + dy, colour)
                dst.Set(int(org.X) - dx, int(org.Y) + dy, colour)
                dst.Set(int(org.X) + dx, int(org.Y) - dy, colour)
                dst.Set(int(org.X) - dx, int(org.Y) - dy, colour)
            }
            break
        }
    }
}

