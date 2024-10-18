package main

import (
	"fmt"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	UP = iota
	DOWN
	LEFT
	RIGHT
)

const V = 2

type Pacman struct {
	i    *ebiten.Image
	op   *ebiten.DrawImageOptions
	w, h int

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

func NewPacman(logicalX, logicalY int) *Pacman {
	cImage := readImage("example.png")
	w, h := cImage.Bounds().Dx(), cImage.Bounds().Dy()
	x, y := logicalX*32, logicalY*32

	return &Pacman{
		i:  ebiten.NewImageFromImage(cImage),
		op: &ebiten.DrawImageOptions{},
		w:  w,
		h:  h,

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

func (p *Pacman) Debug() string {
	return fmt.Sprintf(
		"pos=%v, is_moving=%t, stop=%v \n lpos=%v lstop=%v",
		p.Pos(),
		p.moving,
		[]int{p.stopX, p.stopY},
		[]int{p.lx, p.ly},
		[]int{p.stopLX, p.stopLY},
	)
}

func (p *Pacman) Update(wallTest func(int, int) bool) {
	if p.moving {
		p.move()
		return
	}

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

func (p *Pacman) startMoving(dir int) {
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

func (p *Pacman) move() {
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

func (p *Pacman) Pos() []int {
	return []int{p.x, p.y}
}

func (p *Pacman) LogicalPos() (int, int) {
	return p.lx, p.ly
}

func (p *Pacman) Draw(screen *ebiten.Image) {
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
	screen.DrawImage(p.i, p.op)
}

func (p *Pacman) rotateInPlace(geoM *ebiten.GeoM, degree int) *ebiten.GeoM {
	geoM.Translate(-float64(p.w)/2, -float64(p.h)/2)
	geoM.Rotate(2 * math.Pi * float64(degree) / 360)
	geoM.Translate(float64(p.w)/2, float64(p.h)/2)
	return geoM
}
