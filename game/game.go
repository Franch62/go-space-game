package game

import (
	"fmt"
	"image/color"
	"my-game/assets"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

type Game struct {
	player         *Player
	lasers         []*Laser
	meteors        []*Meteor
	meteorCooldown *Timer
	stars          []*Star
	starCooldown   *Timer
	score          int
}

func NewGame() *Game {
	g := &Game{
		meteorCooldown: NewTimer(24),
		starCooldown:   NewTimer(24),
	}
	player := NewPlayer(g)
	g.player = player
	return g
}

// responsavel por atualizar a logica do jogo
// é chamada 60 vezes por segundo, pois o jogo roda a 60fps
func (g *Game) Update() error {
	g.player.Update()

	for _, l := range g.lasers {
		l.Update()
	}

	g.meteorCooldown.Update()
	if g.meteorCooldown.IsReady() {
		g.meteorCooldown.Reset()

		m := NewMeteor()
		g.meteors = append(g.meteors, m)
	}

	g.starCooldown.Update()
	if g.starCooldown.IsReady() {
		g.starCooldown.Reset()

		s := NewStar()
		g.stars = append(g.stars, s)
	}

	for _, m := range g.meteors {
		m.Update()
	}

	for _, s := range g.stars {
		s.Update()
	}

	for _, m := range g.meteors {
		if m.Collider().Intersects(g.player.Collider()) {
			g.Reset()
		}
	}

	for i, m := range g.meteors {
		for j, l := range g.lasers {
			if m.Collider().Intersects(l.Collider()) {
				g.meteors = append(g.meteors[:i], g.meteors[i+1:]...)
				g.lasers = append(g.lasers[:j], g.lasers[j+1:]...)
				g.score += 1
			}
		}
	}

	return nil
}

// esta também desenha objetos a 60fps
func (g *Game) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)

	for _, l := range g.lasers {
		l.Draw(screen)
	}

	for _, m := range g.meteors {
		m.Draw(screen)
	}

	for _, s := range g.stars {
		s.Draw(screen)
	}

	text.Draw(screen, fmt.Sprintf("Pontos: %d", g.score), assets.FontUi, 16, 100, color.White)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) AddLasers(laser *Laser) {
	g.lasers = append(g.lasers, laser)
}

func (g *Game) AddMeteors(meteor *Meteor) {
	g.meteors = append(g.meteors, meteor)
}

func (g *Game) Reset() {
	g.player = NewPlayer(g)
	g.meteors = nil
	g.lasers = nil
	g.meteorCooldown.Reset()
	g.score = 0
	g.starCooldown.Reset()
}
