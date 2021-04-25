package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenWidth  = 1200 //Ширина окна, в котором воспроизводится игра
	screenHeight = 800  //Длина окна, в котором воспроизводится игра
)

func main() {
	handler := http.NewServeMux()                                                //Выделяет и возвращает новый ServeMux
	handler.HandleFunc("/Race2D", func(w http.ResponseWriter, r *http.Request) { //Регистрирует функцию-обработчик для данного шаблона
		rand.Seed(time.Now().UTC().UnixNano())

		if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
			fmt.Println("Инициализация SDL:", err)
			return
		}

		window, err := sdl.CreateWindow(
			"Race2D", //Название игры в верхней части окна
			sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
			screenWidth, screenHeight,
			sdl.WINDOW_OPENGL)
		if err != nil {
			fmt.Println("Инициализация window:", err)
			return
		}
		defer window.Destroy()

		renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
		if err != nil {
			fmt.Println("Инициализация renderer:", err)
			return
		}
		defer renderer.Destroy()

		plr1, err := newPlayer1(renderer)
		if err != nil {
			fmt.Println("Создание Игрока №1:", err)
			return
		}

		plr2, err := newPlayer2(renderer)
		if err != nil {
			fmt.Println("Создание Игрока №2:", err)
			return
		}

		fn, err := newFON(renderer)
		if err != nil {
			fmt.Println("Создание Трассы (фона):", err)
			return
		}

		wn2, err := newWinner2(renderer)
		if err != nil {
			fmt.Println("Создание надписи Победа для красной машинки:", err)
			return
		}

		wn1, err := newWinner1(renderer)
		if err != nil {
			fmt.Println("Создание надписи Победа для синей машинки:", err)
			return
		}

		for { //Бесконечный цикл, для постоянной работы программы (событие выхода из SDL)
			for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
				switch event.(type) {
				case *sdl.QuitEvent:
					return
				}
			}

			renderer.SetDrawColor(255, 255, 255, 255) //Первоначальный цвет окна игры (белый)
			renderer.Clear()

			fn.draw(renderer) //Отрисовка Трассы (фона)

			plr1.draw(renderer) //Отрисовка игрока №1
			plr1.update()

			plr2.draw(renderer) //Отрисовка игрока №2
			plr2.update()

			wn1.draw(renderer) //Отрисовка надписи "Победа" игрока №1
			wn1.update()

			wn2.draw(renderer) //Отрисовка надписи "Победа" игрока №2
			wn2.update()

			renderer.Present()
		}
	})
	s := &http.Server{
		ReadHeaderTimeout: 30 * time.Second,
		ReadTimeout:       30 * time.Second,
		WriteTimeout:      30 * time.Second,
		Handler:           handler,
		Addr:              ":3000",
	}

	log.Fatal(s.ListenAndServe())

}
