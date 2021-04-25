package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	player1Speed = 0.25 //Скорость передвижения машины
	player1Size  = 100  //Размер игрока №1
)

type player1 struct { //Структура игрока №1
	tex  *sdl.Texture
	x, y float64 //Координаты для перемещения игрока (машины)
}

func newPlayer1(renderer *sdl.Renderer) (p player1, err error) {
	img, err := sdl.LoadBMP("sprites/Car1.bmp") //Загрузка картинки машины игрока №1 (в bmp формате)
	if err != nil {
		return player1{}, fmt.Errorf("Загрузка спрайта игрока №1: %v", err)

	}
	defer img.Free()
	p.tex, err = renderer.CreateTextureFromSurface(img)
	if err != nil {
		return player1{}, fmt.Errorf("Создание текстуры игрока №1: %v", err)
	}

	p.x = player1Size
	p.y = screenHeight/2.0 - 15

	return p, nil
}

func (p *player1) draw(renderer *sdl.Renderer) {
	//Конвертируем координаты игрока в верхний левый угол
	x := p.x - player1Size/2.0
	y := p.y - player1Size/2.0

	renderer.Copy(p.tex,
		&sdl.Rect{X: 0, Y: 0, W: 250, H: 125},
		&sdl.Rect{X: int32(x), Y: int32(y), W: 250, H: 125}) //Координаты перемещения и размеры машины

}

func (p *player1) update() {
	rand.Seed(time.Now().UTC().UnixNano())
	keys := sdl.GetKeyboardState() //Привязка клавиш для управления движением игрока
	if keys[sdl.SCANCODE_LEFT] == 1 {
		p.x -= player1Speed //Движение игрока влево
	} else if keys[sdl.SCANCODE_1] == 1 {
		gaz1 := rand.Intn(2)
		if gaz1 == 0 {
			p.x += player1Speed - 0.04 //Движение игрока вправо
		}
	} else if keys[sdl.SCANCODE_2] == 1 {
		gaz2 := rand.Intn(2)
		if gaz2 == 0 {
			p.x += player1Speed - 0.1 //Движение игрока вправо
		}
	} else if keys[sdl.SCANCODE_3] == 1 {
		gaz3 := rand.Intn(2)
		if gaz3 == 0 {
			p.x += player1Speed - 0.09 //Движение игрока вправо
		}
	} else if keys[sdl.SCANCODE_4] == 1 {
		gaz4 := rand.Intn(2)
		if gaz4 == 0 {
			p.x += player1Speed - 0.15 //Движение игрока вправо
		}
	} else if keys[sdl.SCANCODE_5] == 1 {
		gaz5 := rand.Intn(2)
		if gaz5 == 0 {
			p.x += player1Speed - 0.06 //Движение игрока вправо
		}
	}

}
