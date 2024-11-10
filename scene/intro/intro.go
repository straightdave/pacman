package intro

import (
	g "github.com/straightdave/pacman/game"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type SceneIntro struct {
	isActive bool
}

func init() {
	g.RegisterScene("intro", func() (g.Scene, error) {
		return &SceneIntro{isActive: true}, nil
	})
}

func (s *SceneIntro) IsActive() bool { return s.isActive }
func (s *SceneIntro) Activate()      { s.isActive = true }
func (s *SceneIntro) Deactivate()    { s.isActive = false }

func (s *SceneIntro) Update(ctx *g.Context) error {
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		ctx.NextScene = "play"
		s.Deactivate()
	}
	return nil
}

func (s *SceneIntro) Draw(_ *g.Context, screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Press SPACE To Start")
}
