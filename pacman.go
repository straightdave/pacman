package main

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	UP = iota
	DOWN
	LEFT
	RIGHT
)

type Pacman struct {
	i      *ebiten.Image
	op     *ebiten.DrawImageOptions
	w, h   int
	x, y   int
	vx, vy int
	dir    int
}

func NewPacman(x, y int) *Pacman {
	cImage := readImage("example.png")
	w, h := cImage.Bounds().Dx(), cImage.Bounds().Dy()

	return &Pacman{
		i:  ebiten.NewImageFromImage(cImage),
		op: &ebiten.DrawImageOptions{},
		w:  w,
		h:  h,

		// speed
		vx: 1,
		vy: 1,

		// initial position and direction
		x:   x,
		y:   y,
		dir: RIGHT,
	}
}

func (p *Pacman) moveLeft() {
	p.x -= p.vx
	p.dir = LEFT
}

func (p *Pacman) moveUp() {
	p.y -= p.vy
	p.dir = UP
}

func (p *Pacman) moveRight() {
	p.x += p.vy
	p.dir = RIGHT
}

func (p *Pacman) moveDown() {
	p.y += p.vy
	p.dir = DOWN
}

func (p *Pacman) Pos() []int {
	return []int{p.x, p.y}
}

func (p *Pacman) Put(x, y int) {
	p.x, p.y = x, y
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
