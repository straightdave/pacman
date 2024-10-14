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

	pacman *Pacman
}

func NewGame(screenWidth, screenHeight int) *Game {
	g := &Game{
		sWidth:  screenWidth,
		sHeight: screenHeight,

		inited: false,
		layers: levelMap1,
		rect:   ebiten.NewImage(32, 32),

		pacman: NewPacman(screenWidth/2, screenHeight/2),
	}
	g.rect.Fill(color.RGBA{0, 0, 255, 1})
	return g
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		g.pacman.moveLeft()
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		g.pacman.moveUp()
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		g.pacman.moveRight()
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		g.pacman.moveDown()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.drawLevel(screen)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("pos=%v", g.pacman.Pos()))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.sWidth, g.sHeight
}

func (g *Game) drawLevel(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	for i, r := range g.layers {
		for j, v := range r {
			op.GeoM.Reset()
			op.GeoM.Translate(float64(i*32), float64(j*32))

			if v == 1 {
				screen.DrawImage(g.rect, op)
			}

			if v == 2 && !g.inited {
				g.pacman.Put(i*32, j*32)
				g.inited = true
			}
		}
	}
	g.pacman.Draw(screen)
}
