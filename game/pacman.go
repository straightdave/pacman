package game

import (
	"image"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	UP = iota
	DOWN
	LEFT
	RIGHT
)

const (
	V       = 2
	CWidth  = 32
	CHeight = 32
)

type Pacman struct {
	Animate

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
	cImage := readImage("assets/pacman.png")
	img := ebiten.NewImageFromImage(cImage)
	x, y := logicalX*CWidth, logicalY*CHeight

	return &Pacman{
		Animate: Animate{
			imgRef: img,
			framesInSprites: []image.Rectangle{
				image.Rect(0, 0, CWidth, CHeight),
				image.Rect(0, CHeight, CWidth, 2*CHeight),
			},
			ticksPerFrame: 5,
		},

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

func (p *Pacman) Update(wallTest func(int, int) bool) {
	if p.moving {
		p.AnimeTick()
		p.move()
		return
	}

	p.AnimeTickReset()

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
	op := &ebiten.DrawImageOptions{}

	switch p.dir {
	case UP:
		p.rotateInPlace(&op.GeoM, -90)
	case LEFT:
		p.rotateInPlace(&op.GeoM, 180)
	case DOWN:
		p.rotateInPlace(&op.GeoM, 90)
	default:
		p.rotateInPlace(&op.GeoM, 0)
	}
	op.GeoM.Translate(float64(p.x), float64(p.y))

	screen.DrawImage(p.ImageToDraw(), op)
}

func (p *Pacman) rotateInPlace(geoM *ebiten.GeoM, degree int) *ebiten.GeoM {
	geoM.Translate(-float64(CWidth)/2, -float64(CHeight)/2)
	geoM.Rotate(2 * math.Pi * float64(degree) / 360)
	geoM.Translate(float64(CWidth)/2, float64(CHeight)/2)
	return geoM
}
