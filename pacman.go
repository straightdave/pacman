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
	dir           int
	startX, stopX int
	startY, stopY int
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

		dir:    LEFT,
		startX: x,
		stopX:  x,
		startY: y,
		stopY:  y,
	}
}

func (p *Pacman) Debug() string {
	return fmt.Sprintf(
		"pos=%v, is_moving=%t, stop=%v \n lpos=%v",
		p.Pos(),
		p.isMoving(),
		[]int{p.stopX, p.stopY},
		[]int{p.lx, p.ly},
	)
}

func (p *Pacman) Update(wallTest func(int, int) bool) {
	if p.isMoving() {
		p.move()
		return
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		p.dir = LEFT
		if !wallTest(p.lx-1, p.ly) {
			p.startMoving(LEFT)
			p.lx -= 1
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		p.dir = UP
		if !wallTest(p.lx, p.ly-1) {
			p.startMoving(UP)
			p.ly -= 1
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		p.dir = RIGHT
		if !wallTest(p.lx+1, p.ly) {
			p.startMoving(RIGHT)
			p.lx += 1
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		p.dir = DOWN
		if !wallTest(p.lx, p.ly+1) {
			p.startMoving(DOWN)
			p.ly += 1
		}
	}
}

func (p *Pacman) startMoving(dir int) {
	switch dir {
	case UP:
		p.stopX, p.stopY = p.x, p.y-32
	case RIGHT:
		p.stopX, p.stopY = p.x+32, p.y
	case DOWN:
		p.stopX, p.stopY = p.x, p.y+32
	default:
		p.stopX, p.stopY = p.x-32, p.y
	}
}

func (p *Pacman) move() {
	switch p.dir {
	case UP:
		if p.y < p.stopY {
			p.y = p.stopY
		} else {
			p.y -= V
		}

	case RIGHT:
		if p.x > p.stopX {
			p.x = p.stopX
		} else {
			p.x += V
		}

	case DOWN:
		if p.y > p.stopY {
			p.y = p.stopY
		} else {
			p.y += V
		}

	default:
		if p.x < p.stopX {
			p.x = p.stopX
		} else {
			p.x -= V
		}
	}
}

func (p *Pacman) Pos() []int {
	return []int{p.x, p.y}
}

func (p *Pacman) isMoving() bool {
	return p.x != p.stopX || p.y != p.stopY
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
