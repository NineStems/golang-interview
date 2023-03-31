## [Справочные материалы](https://github.com/NineStems/junior-backend-developer-topics-and-questions/blob/main/lib/golang.md)

## [Вопросы по языку Golang](./question.md)
- [Что из себя представляет тип данных string в языке Golang? Можно ли изменить определенный символ в строке? Что происходит при склеивании строк?](question.md#1-------string---golang------------)
- [Как эффективно склеивать множество строк?](./question.md#2-----)
- [Как устроен тип карты? Что такое эвакуация?](./question.md#3-------)
- [Какие нюансы при работе с картами в многопоточном режиме?](./question.md#4---------)
- [Расскажите про многопоточность в golang](./question.md#5-----golang)
- [Что такое планировщик и как он работает](./question.md#6-------)
- _
- [Как задать направление канала?](./question.md#8----)
- [Напишите собственную функцию Sleep, используя time.After](./question.md#9----sleep--timeafter)
- [Что такое буферизированный канал? Как создать такой канал с ёмкостью в 20 сообщений?](./question.md#10------------20-)
- [Напишите программу, которая меняет местами два числа](./question.md#11--------x--1-y--2-swapx-y---x2--y1)
- [Практика на чтение кода #1](./question.md#12------x---)
- [Практика на чтение кода #2](./question.md#13-----true--false--false--true--false--false)
- [Практика на чтение кода #3](./question.md#14----)
- [Практика на чтение кода #4](./question.md#15----)
- [Как работает Garbage Collection в Go?](./question.md#16---garbage-collection--go)
- [Что такое interface, как они работают в Go?](./question.md#17---interface-----go)
- [Что такое slice, как устроены и чем отличаются от массивов?](./question.md#18---slice-------)
- [Что такое len и capacity в slice Go?](./question.md#19---len--capacity--slice-go)
- [Возможно ли предугадать, что GC отработает за константное время N?](./question.md#20-----gc-----n)
- [Что будет, если создать канал и отправить туда запись, но у него нет читателей?](./question.md#21--------------)
- [Практика на чтение кода #5](./question.md#22--------)
- [Практика на чтение кода #6](./question.md#23---------)
- [Практика на чтение кода #7](./question.md#24--------3----6)
- [Практика на чтение кода #8](./question.md#25------)
- [Практика на чтение кода #9](./question.md#26-----------)
- [Практика на чтение кода #10](./question.md#27-------------)

## [Вопросы от Данила Подольского ](./question-podolsky.md)
- [Go — императивный или декларативный? А в чем разница?](./question-podolsky.md#1-go--------)
- [Что такое type switch?](./question-podolsky.md#2---type-switch)
- [Как сообщить компилятору, что наш тип реализует интерфейс?](./question-podolsky.md#3--------)
- [Как работает append?](./question-podolsky.md#4---append)
- [Какое у slice zero value? Какие операции над ним возможны?](./question-podolsky.md#5---slice-zero-value-----)
- [Как устроен тип map?](./question-podolsky.md#6----map) 
- [Каков порядок перебора map?](./question-podolsky.md#7----map)
- [Что будет, если читать из закрытого канала?](./question-podolsky.md#8-------)
- [Что будет, если писать в закрытый канал?](./question-podolsky.md#9-------)
- [Как вы отсортируете массив структур по алфавиту по полю Name?](./question-podolsky.md#10----------name)
- [Что такое сериализация? Зачем она нужна?](./question-podolsky.md#11------)
- [Сколько времени в минутах займет у вас написание процедуры обращения односвязного списка?](./question-podolsky.md#12------------)
- [Где следует поместить описание интерфейса: в пакете с реализацией или в пакете, где этот интерфейс используется? Почему?](./question-podolsky.md#13-----------------)
- [Предположим, ваша функция должна возвращать детализированные Recoverable и Fatal ошибки. Как это реализовано в пакете net? Как это надо делать в современном Go?](./question-podolsky.md#14-------recoverable--fatal-------net-------go)
- [Главный недостаток стандартного логгера?](./question-podolsky.md#15----)
- [Есть ли для Go хороший orm? Ответ обоснуйте.](./question-podolsky.md#16----go--orm--)
- [Какой у вас любимый линтер?](./question-podolsky.md#17-----)
- [Можно ли использовать один и тот же буфер []byte в нескольких горутинах?](./question-podolsky.md#18---------byte---)
- [Какие типы мьютексов предоставляет stdlib?](./question-podolsky.md#19-----stdlib)
- [Что такое lock-free структуры данных, и есть ли в Go такие?](./question-podolsky.md#20---lock-free-------go-)
- [Способы поиска проблем производительности на проде?](./question-podolsky.md#21------)
- [Стандартный набор метрик prometheus в Go -программе?](./question-podolsky.md#22----prometheus--go-)
- [Как встроить стандартный профайлер в свое приложение?](./question-podolsky.md#23-------)
- [Overhead от стандартного профайлера?](./question-podolsky.md#24-overhead---)
- [Почему встраивание — не наследование?](./question-podolsky.md#25-----)
- [Какие средства обобщенного программирования есть в Go?](./question-podolsky.md#26-------go)
- [Какие технологические преимущества языка Go вы можете назвать?](./question-podolsky.md#27-----go---)
- [Какие технологические недостатки языка Go вы можете назвать?](./question-podolsky.md#28-----go---) 