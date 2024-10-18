package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	sWidth, sHeight int

	inited bool
	layers [][]int
	rect   *ebiten.Image
	dot    *ebiten.Image
	score  int

	pacman *Pacman
}

const (
	screenWidth  = 320
	screenHeight = 320
)

func NewGame() *Game {
	g := &Game{
		sWidth:  screenWidth,
		sHeight: screenHeight,

		inited: false,
		layers: levelMap1,
		rect:   ebiten.NewImage(32, 32),
		dot:    ebiten.NewImage(32, 32),
	}
	g.rect.Fill(color.RGBA{0, 0, 255, 1})
	g.dot.Fill(color.RGBA{255, 255, 255, 1})
	return g
}

func (g *Game) Update() error {
	if !g.inited {
		return nil
	}

	g.pacman.Update(g.wallTest)

	lx, ly := g.pacman.LogicalPos()
	if g.layers[lx][ly] == 3 {
		g.layers[lx][ly] = 0
		g.score++
	}

	return nil
}

func (g *Game) wallTest(lx, ly int) bool {
	return g.layers[lx][ly] == 1
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.draw(screen)
	debugMsg := fmt.Sprintf("score=%d %v", g.score, g.pacman.Debug())
	ebitenutil.DebugPrint(screen, debugMsg)
}

func (g *Game) draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	for i, r := range g.layers {
		for j, v := range r {
			op.GeoM.Reset()
			op.GeoM.Translate(float64(i*32), float64(j*32))

			if v == 1 {
				screen.DrawImage(g.rect, op)
			}

			if v == 2 && !g.inited {
				g.pacman = NewPacman(i, j)
				g.inited = true
			}

			if v == 3 {
				screen.DrawImage(g.dot, op)
			}
		}
	}
	g.pacman.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.sWidth, g.sHeight
}
