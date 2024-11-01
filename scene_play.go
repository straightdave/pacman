package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type ScenePlay struct {
	isActive bool

	layers [][]int
	rect   *ebiten.Image
	dot    *ebiten.Image
	score  int
	pacman *Pacman
}

func NewScenePlay() *ScenePlay {
	// init logical pos of pacman
	lx, ly := 5, 5
	levelMap1[lx][ly] = 2

	s := &ScenePlay{
		isActive: true,
		layers:   levelMap1,
		rect:     ebiten.NewImage(32, 32),
		dot:      ebiten.NewImage(32, 32),
		score:    0,
		pacman:   NewPacman(lx, ly),
	}
	s.rect.Fill(color.RGBA{0, 0, 255, 1})
	s.dot.Fill(color.RGBA{255, 255, 255, 1})
	return s
}

func (s *ScenePlay) IsActive() bool { return s.isActive }
func (s *ScenePlay) Activate()      { s.isActive = true }
func (s *ScenePlay) Deactivate()    { s.isActive = false }

func (s *ScenePlay) Update(ctx *Context) error {
	if s.score >= 5 {
		ctx.NextScene = "win"
		s.Deactivate()
		return nil
	}

	if s.pacman == nil {
		return nil
	}

	s.pacman.Update(s.wallTest)

	lx, ly := s.pacman.LogicalPos()
	if s.layers[lx][ly] == 3 {
		s.layers[lx][ly] = 0
		s.score++
	}

	return nil
}

func (s *ScenePlay) wallTest(lx, ly int) bool {
	return s.layers[lx][ly] == 1
}

func (s *ScenePlay) Draw(_ *Context, screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	for i, r := range s.layers {
		for j, v := range r {
			op.GeoM.Reset()
			op.GeoM.Translate(float64(i*32), float64(j*32))

			if v == 1 {
				screen.DrawImage(s.rect, op)
			}

			if v == 3 {
				screen.DrawImage(s.dot, op)
			}
		}
	}

	if s.pacman != nil {
		s.pacman.Draw(screen)
	}

	ebitenutil.DebugPrint(screen, fmt.Sprintf("score: %d", s.score))
}
