package win

import (
	g "github.com/straightdave/pacman/game"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type SceneWin struct{}

func init() {
	g.RegisterScene("win", func() (g.Scene, error) {
		return &SceneWin{}, nil
	})
}

func (s *SceneWin) IsActive() bool { return false }
func (s *SceneWin) Activate()      {}
func (s *SceneWin) Deactivate()    {}

func (s *SceneWin) Update(_ *g.Context) error { return nil }

func (s *SceneWin) Draw(_ *g.Context, screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "WIN!")
}
