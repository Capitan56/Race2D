package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	player2Speed = 0.25 //Скорость передвижения машины
	player2Size  = 100  //Размер игрока №2
)

type player2 struct { //Структура игрока №2
	tex  *sdl.Texture
	x, y float64 //Координаты для перемещения игрока (машины)
}

func newPlayer2(renderer *sdl.Renderer) (p player2, err error) {
	img, err := sdl.LoadBMP("sprites/Car2.bmp") //Загрузка картинки машины игрока №2 (в bmp формате)
	if err != nil {
		return player2{}, fmt.Errorf("Загрузка спрайта игрока №2: %v", err)

	}
	defer img.Free()
	p.tex, err = renderer.CreateTextureFromSurface(img)
	if err != nil {
		return player2{}, fmt.Errorf("Создание текстуры игрока №2: %v", err)
	}

	p.x = player2Size + 4
	p.y = screenHeight/2.0 + 150

	return p, nil
}

func (p *player2) draw(renderer *sdl.Renderer) {
	//Конвертируем координаты игрока в верхний левый угол
	x := p.x - player2Size/2.0
	y := p.y - player2Size/2.0
	renderer.Copy(p.tex,
		&sdl.Rect{X: 0, Y: 0, W: 250, H: 225},
		&sdl.Rect{X: int32(x), Y: int32(y), W: 250, H: 225}) //Координаты перемещения и размеры машины
}

func (p *player2) update() {
	rand.Seed(time.Now().UTC().UnixNano())
	keys := sdl.GetKeyboardState() //Привязка клавиш для управления движением игрока
	if keys[sdl.SCANCODE_LEFT] == 1 {
		p.x -= player1Speed //Движение игрока влево
	} else if keys[sdl.SCANCODE_G] == 1 {
		gaz11 := rand.Intn(2)
		if gaz11 == 1 {
			p.x += player1Speed - 0.06 //Движение игрока вправо
		}
	} else if keys[sdl.SCANCODE_H] == 1 {
		gaz22 := rand.Intn(2)
		if gaz22 == 1 {
			p.x += player1Speed - 0.1 //Движение игрока вправо
		}
	} else if keys[sdl.SCANCODE_J] == 1 {
		gaz33 := rand.Intn(2)
		if gaz33 == 1 {
			p.x += player1Speed - 0.15 //Движение игрока вправо
		}
	} else if keys[sdl.SCANCODE_K] == 1 {
		gaz44 := rand.Intn(2)
		if gaz44 == 1 {
			p.x += player1Speed - 0.04 //Движение игрока вправо
		}
	} else if keys[sdl.SCANCODE_L] == 1 {
		gaz55 := rand.Intn(2)
		if gaz55 == 1 {
			p.x += player1Speed - 0.09 //Движение игрока вправо
		}
	}

}
