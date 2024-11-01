package main

import (
	"fmt"
	"image"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Ghost struct {
	i  *ebiten.Image
	op *ebiten.DrawImageOptions

	// animation
	animeTick  int
	animeState int

	// logical pos
	lx, ly int

	// current pos
	x, y int

	// moving state
	moving         bool
	dir            int
	stopX, stopY   int
	stopLX, stopLY int
}

func NewGhost(logicalX, logicalY int) *Ghost {
	cImage := readImage("ghost.png")
	x, y := logicalX*CWidth, logicalY*CHeight

	return &Ghost{
		i:  ebiten.NewImageFromImage(cImage),
		op: &ebiten.DrawImageOptions{},

		animeTick:  0,
		animeState: 0,

		lx: logicalX,
		ly: logicalY,

		x: x,
		y: y,

		moving: false,
		dir:    RIGHT,
		stopX:  x,
		stopY:  y,
		stopLX: logicalX,
		stopLY: logicalY,
	}
}

func (p *Ghost) Debug() string {
	return fmt.Sprintf(
		"pos=%v, is_moving=%t, stop=%v \n lpos=%v lstop=%v",
		p.Pos(),
		p.moving,
		[]int{p.stopX, p.stopY},
		[]int{p.lx, p.ly},
		[]int{p.stopLX, p.stopLY},
	)
}

func (p *Ghost) Update(wallTest func(int, int) bool) {
	if p.moving {
		p.animeTick++
		p.move()
		return
	}

	p.animeTick = 0

	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		p.dir = LEFT
		if !wallTest(p.lx-1, p.ly) {
			p.startMoving(LEFT)
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		p.dir = UP
		if !wallTest(p.lx, p.ly-1) {
			p.startMoving(UP)
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		p.dir = RIGHT
		if !wallTest(p.lx+1, p.ly) {
			p.startMoving(RIGHT)
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		p.dir = DOWN
		if !wallTest(p.lx, p.ly+1) {
			p.startMoving(DOWN)
		}
	}
}

func (p *Ghost) startMoving(dir int) {
	if p.moving {
		return
	}

	p.moving = true

	switch dir {
	case UP:
		p.stopY = p.y - 32
		p.stopLY = p.ly - 1
	case RIGHT:
		p.stopX = p.x + 32
		p.stopLX = p.lx + 1
	case DOWN:
		p.stopY = p.y + 32
		p.stopLY = p.ly + 1
	default:
		p.stopX = p.x - 32
		p.stopLX = p.lx - 1
	}
}

func (p *Ghost) move() {
	if !p.moving {
		return
	}

	switch p.dir {
	case UP:
		p.y -= V
		if p.y <= p.stopY {
			p.y = p.stopY
			p.ly = p.stopLY
			p.moving = false
		}

	case RIGHT:
		p.x += V
		if p.x >= p.stopX {
			p.x = p.stopX
			p.lx = p.stopLX
			p.moving = false
		}

	case DOWN:
		p.y += V
		if p.y >= p.stopY {
			p.y = p.stopY
			p.ly = p.stopLY
			p.moving = false
		}

	default:
		p.x -= V
		if p.x <= p.stopX {
			p.x = p.stopX
			p.lx = p.stopLX
			p.moving = false
		}
	}
}

func (p *Ghost) Pos() []int {
	return []int{p.x, p.y}
}

func (p *Ghost) LogicalPos() (int, int) {
	return p.lx, p.ly
}

func (p *Ghost) Draw(screen *ebiten.Image) {
	p.op.GeoM.Reset()

	switch p.dir {
	case UP:
		p.rotateInPlace(&p.op.GeoM, -90)
	case LEFT:
		p.rotateInPlace(&p.op.GeoM, 180)
	case DOWN:
		p.rotateInPlace(&p.op.GeoM, 90)
	default:
		p.rotateInPlace(&p.op.GeoM, 0)
	}
	p.op.GeoM.Translate(float64(p.x), float64(p.y))

	i := (p.animeTick / 5) % 2
	sx, sy := 0, i*CHeight
	screen.DrawImage(
		p.i.SubImage(image.Rect(sx, sy, sx+CWidth, sy+CHeight)).(*ebiten.Image),
		p.op,
	)
}

func (p *Ghost) rotateInPlace(geoM *ebiten.GeoM, degree int) *ebiten.GeoM {
	geoM.Translate(-float64(CWidth)/2, -float64(CHeight)/2)
	geoM.Rotate(2 * math.Pi * float64(degree) / 360)
	geoM.Translate(float64(CWidth)/2, float64(CHeight)/2)
	return geoM
}
