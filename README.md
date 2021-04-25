# Race2D
Эта игра создана как курсовой проект Калашникова Василия Сергеевича - студента высшего учебного заведения МГТУ им. Н.Э. Баумана.

Руководство по игре:

Чтобы поиграть в сетевую игру "Race2D", следует сделать следующее:
1. Установить на свой компьютер язык программирования "Go"
2. Скачать сборку файлов "Race2D", закреплённых в репозитории, и закинуть её в папку /go/src/github.com/
3. Скачать библиотеку github.com/veandco/go-sdl2/sdl на свой компьютер
4. Подключиться к локальной сети и запустить клиент
5. Перейти в браузере по адресу http://0.0.0.0:3000/Race2D и наслаждаться игрой! 

*ДИСКЛЕЙМЕР*

Игра "Race2D" создана чисто в развлекательных целях и не является чьим-то плагиатом. Все события в данной игре вымышленные, любое сходство с реальными событиями случайно. Автор не несёт ответственности за причиннёный вред своей игрой. Всё сказанное является исключительно личным мнением автора.

Описание игры:
"Race2D" - это сетевая 2D игра, которая описывает историю одной знаменитой в мире гонки. Вы и ваш соперник играете за машинки синего и красного цвета. Вы, собственно, как и ваш соперник очень жадите победы и хотите выйграть любой ценой! По канону игры, в вечер перед гонкой вы проникли к своему сопернику в гараж, где стояла его гоночная машина и испортили в ней коробку передач, но вы и догадаться не могли, что ваш противник сделает тоже самое. И вот долгожданная встреча - красный против синего! Впереди потрясающая гонка, но кто придёт к финишу первым?

Правила игры:
1. Перед тем как начать гонку, стоит дождаться подключения второго игрока (можно играть и с одноого устройства)
2. Убедиться, что оба игрока готовы к игре
3. Кнопки: 1,2,3,4,5 - управляют синей машиной, а кнопки: g,h,j,k,l - управляют красной машиной (но никто не знает: какая именно из этих кнопок отвечает за какую передачу, вам стоит это выяснить по ходу гонки) 
4. Оба игрока должны стартовать вместе по согласованности!
5. Побеждает тот, у кого надпись "Победа" появиться на экране первой (полностью)!
6. В случае ничьи переиграть гонку

ЗАПРЕЩАЕТСЯ
1. Стартовать раньше согласованной команды
2. Если вы играете по сети, использовать кнопки одновременно и красной,и синей машины
3. Продолжать гонку, если чья-то надпись "Победа" уже полностью появилась на экране
