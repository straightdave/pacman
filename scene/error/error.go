package error

import (
	g "github.com/straightdave/pacman/game"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type SceneError struct{}

func init() {
	g.RegisterScene(
		"error",
		func() (g.Scene, error) {
			return &SceneError{}, nil
		})
}

func (s *SceneError) IsActive() bool { return false }
func (s *SceneError) Activate()      {}
func (s *SceneError) Deactivate()    {}

func (s *SceneError) Update(_ *g.Context) error { return nil }

func (s *SceneError) Draw(ctx *g.Context, screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Error: "+ctx.Message)
}
