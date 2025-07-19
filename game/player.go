package game

import (
	"my-game/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	image         *ebiten.Image
	position      Vector
	game          *Game
	laserCooldown *Timer
}

func NewPlayer(game *Game) *Player {
	image := assets.GopherPlayer
	bounds := image.Bounds()
	halfW := float64(bounds.Dx()) / 2

	position := Vector{
		X: (screenWidth / 2) - halfW,
		Y: 500,
	}

	return &Player{
		image:         image,
		game:          game,
		position:      position,
		laserCooldown: NewTimer(16),
	}
}

func (p *Player) Update() {
	speed := 6.0
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		p.position.X -= speed
	} else if ebiten.IsKeyPressed(ebiten.KeyD) {
		p.position.X += speed
	}

	p.laserCooldown.Update()
	if ebiten.IsKeyPressed(ebiten.KeySpace) && p.laserCooldown.IsReady() {
		p.laserCooldown.Reset()
		bound := p.image.Bounds()
		halfW := float64(bound.Dx()) / 2 //metade da largura da imagem do laser
		halfH := float64(bound.Dy()) / 2 //metade da altura da imagem do laser

		spawnPos := Vector{
			p.position.X + halfW,
			p.position.Y - halfH/2,
		}

		laser := NewLaser(spawnPos)
		p.game.AddLasers(laser)

	}

}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	//Posição X e Y que a imagem sera desenhada na tela
	op.GeoM.Translate(p.position.X, p.position.Y)
	screen.DrawImage(p.image, op)

}

func (p *Player) Collider() Rect {
	bounds := p.image.Bounds()
	return NewRect(p.position.X,
		p.position.Y,
		float64(bounds.Dx()),
		float64(bounds.Dy()))
}
