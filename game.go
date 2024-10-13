package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	sWidth, sHeight int

	rect   *ebiten.Image
	rectOp *ebiten.DrawImageOptions

	pacman *Pacman
}

func NewGame(screenWidth, screenHeight int) *Game {
	g := &Game{
		sWidth:  screenWidth,
		sHeight: screenHeight,
		rect:    ebiten.NewImage(64, 48),
		rectOp:  &ebiten.DrawImageOptions{},
		pacman:  NewPacman(screenWidth/2, screenHeight/2),
	}
	g.rect.Fill(color.RGBA{0, 0, 255, 100})
	g.rectOp.GeoM.Translate(100.0, 100.0)
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
	g.pacman.Draw(screen)
	screen.DrawImage(g.rect, g.rectOp)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("pos=%v", g.pacman.Pos()))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.sWidth, g.sHeight
}
