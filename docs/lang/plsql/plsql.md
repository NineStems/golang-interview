# SQL Basic
## SQL JOIN  
### Типы соединений  
**Пример соединения**  
```
select 
* 
from a
join b on a.f1 = b.f1

```
- a – outer-таблица, по которой мы производим чтение
- b – inner-таблица, в которой мы ищем соответствующие записи
- on a.f1 = b.f1 - условие соединения таблицы  
**Пример соединения**  
- Inner join: соединение таблиц происходит по тем строкам, для которых полностью выполняется условие соединения
-  Outer joins:
   -  left outer join – тип соединения, при котором выводятся все строки из outer-таблицы и соответствующие условию соединения строки из inner-таблицы
   -  right outer join – тип соединения, при котором выводятся все строки из inner-таблицы и соответствующие условию соединения строки из outer-таблицы
   -  full outer join - тип соединения, который представляет собой объединение left и right outer joins

- cross join – тип соединения, в котором результатом выборки будет являться декартовое произведение строк
аналог cross join:
```
select *
from tableA, tableB;

select *
from tableA
cross join tableB
```

### Методы соединений
**Nested loop**  
Для каждой строки outer-таблицы по условию соединения происходит последовательный поиск строк из inner-таблицы. Применяется в том случае, когда в условии соединения участвуют проиндексированные столбцы. В таком случае, основной таблицей будет считаться та, у которой столбец, по которому идет соединение, не проиндексирован. В случае, если столбцы в обеих таблицах имеют индексы, основной будет считаться та, у которой размер индекса больше.
- Плюсы: быстрая выборка в том случае, когда соединение идет по индексированным столбцам
- Минусы: многократное чтение неосновной таблицы  

**Hash join**  
Применяется в том случае, когда соединяется большой набор данных. В данном случае на основании наименьшей таблицы во внутренней памяти строится хэш-таблица. Хэш-таблица представляет собой ассоциативный массив, в котором индексом служат значения ключа соединения (набор значений столбцов второй таблицы, по которым идет соединение), а значением элемента массива является результат хэш-функции. После построения хэш-таблицы производится сканирование основной таблицы, в рамках которого ищется соответствующее значение в хэш-таблице.
- Плюсы: быстрая выборка по сравнению с NL в том случае:
    - когда размер таблиц большой (или одна большая и одна маленькая)
    - когда соединение идет не по индексу
- Минусы:
    - для выполнения данного метода требуется много внутренней памяти. Есть риск того, что она закончится, в таком случае запрос не выполнится.

**Sort merge join**  
Вариация NL, при которой таблицы сначала сортируются по столбцам, входящим в условие соединения.
   - Плюсы: если таблицы заранее отсортированы, то соединение будет выполняться быстрее, чем NL за счет того, что данные из второй таблицы будут перечитываться не с начала, а от последнего выбранного значения.
   - Минусы: если таблицы не отсортированы, будет тратиться дополнительное время на сортировку

**Cartesian join (merge join cartesian)**  
Используется в том случае, когда нет условий соединения хотя бы с одной из таблиц в запросе. Типичные примеры таблиц, получаемых в результате данного вида соединения: таблица умножения, ежедневник с почасовой разбивкой.

### Порядок соединения
Порядок основывается на сравнении следующих параметров: 
- Размер таблиц
- Индексы таблицах
- Возможность использования индексов в текущем запросе
- Количество строк, которые нужно просканировать для каждой таблицы в различных вариантах соединения.
   - В большинстве случаев порядок соединения таблиц определяется количеством строк, которые предполагается выбрать (на основе cardinality) — та таблица, в которой будет меньше строк, будет основной.
   - Исключением является тот случай, когда в условии соединения участвуют проиндексированные столбцы. В таком случае основной таблицей будет считаться та, у которой вернется больше строк. Если оба столбца проиндексированы, то в данном случае основной будет считаться та таблица, размер которой меньше.

Оптимизатор не может менять порядок соединения в том случае, если используются outer-join’ы (left,right,full). В данном случае порядок соединения будет именно такой, какой указан в запросе.

**Heap-Organized Tables**  
- Full Table Scan применяется в том случае, когда:
   - у таблицы нет индексов
   - по условию запроса к индексированному столбцу применяется функция
   - если в запросе применяется функция count(<индексированный столбец>), и в данном столбце содержатся пустые значения
   - в случае составного индекса на таблице, в условии запроса не используются первые столбцы, входящие в индекс (может применяться также index skip scan)
   - запрос является неселективным (когда оптимизатор понимает, что для выполнения запроса необходимо прочитать бОльшую часть таблицы)
   - устаревшая статистика
   - маленький размер таблицы (тот случай, когда всю таблицу прочитать дешевле, чем ее индекс и потом по нему вытягивать значения других столбцов)
   - таблица имеет высокую степень параллелизма
   - применяется хинт /*+ full(<синоним таблицы>)*/  
- Table Access by Rowid применяется в том случае, когда:
   - результат запроса можно получить при использовании индекса. В случае, если все требуемые столбцы входят в индекс, будет использоваться Index Fast Full Scan
- Sample Table Scan  
   - Index-organized tables:  
   - Index unique Scan применяется в том случае, если условие запроса можно выполнить после сканирования уникального индекса  
   - Index range scan применяется в том случае, когда:  
     - один или более столбцов индекса (ведущих) используются в условии запроса
     - по условию запроса индекс может вернуть 0, 1 и более значений
   - index full scan применяется в том случае, когда:
     - в условии запроса есть сортировка на проиндексированные столбцы, не содержащие пустых значений
   - index fast full scan применяется в том случае, когда в запросе участвуют только столбцы, входящие в индекс, и при этом сортировка не важна.
   - index skip scan применяется в том случае, когда:
     - ведущий столбец индекса не участвует в запросе
     - ведущий столбец индекса имеет всего несколько уникальных значений, в то время когда ведомые столбцы индекса содержат множество различных значений.
   - index join scan применяется в том случае, когда:
     - hash-join нескольких индексов может вернуть все столбцы, указанные в запросе, без чтения таблиц
     - стоимость чтения строк таблицы будет больше, чем чтение индексов 

### Group/Analytic functions
**count**  
Возвращает количество ненулевых значений столбца (внутри можно использовать distinct, в таком случае функция вернет количество уникальных значений). Всегда возвращает значение.

**min/max/avg/sum**  
Вычисление минимального/максимального/среднего/суммы значений в столбце  

**Аналитические функции** — это не скалярные функции (за исключением аналитических статистических, скалярных по результату), которые в отличие от стандартных агрегатных могут употребляться только во фразах SELECT и ORDER BY, так как применяются к уже отобранному результату.  
Свое название получили по той причине, что позволяют средствами SQL (в Oracle) строить запросы, анализирующие данные в БД. Являются вариацией "оконных функций", вошедших в SQL:2003.
Функции этой категории иногда называют "функциями OLAP" ввиду того, что они хорошо подходят для систем типа OLAP (On-Line Analytical Processing), аналитических систем и "аналитических баз данных".

В Oracle они могут быть следующих видов:
- функции ранжирования;
- статистические функции для плавающего интервала;
- функции подсчета долей;
- статистические функции LAG/LEAD с запаздывающим/опережающим аргументом;
- статистические функции (линейная регрессия и т. д.).  
Техническая цель введения аналитических функций — дать лаконичную формулировку и увеличить скорость выполнения "аналитических запросов" к БД, то есть запросов, имеющих смыслом выявление внутренних соотношений и зависимостей в данных. Более точно, пользование аналитическими функциями может дать следующие выгоды перед обычными SQL-операторами:
- Лаконичную и простую формулировку. Многие аналитические запросы к БД традиционными средствами сложно формулируются, а потому с трудом осмысливаются и плохо отлаживаются.
- Снижение нагрузки на сеть. То, что раньше могло формулироваться только серией запросов, сворачивается в один запрос. По сети только отправляется запрос и получается окончательный результат.
- Перенос вычислений на сервер. С использованием аналитических функций нет нужды организовывать расчеты на клиенте; они полностью проводятся на сервере, ресурсы которого могут быть более подходящи для быстрой обработки больших объемов данных.
- Лучшую эффективность обработки запросов. Аналитические функции имеют алгоритмы вычисления, неразрывно связанные со специальными планами обработки запросов, оптимизированными для большей скорости получения результата.

**Особенности обработки**  
Разбиение данных на группы для вычислений
Аналитические функции агрегируют данные порциями (partitions; группами), количество и размер которых можно регулировать специальной синтаксической конструкцией. Ниже она указана на примере агрегирующей функции SUM:   
```SUM(выражение 1) OVER([PARTITION BY выражение 2 [, выражение 3 [, …]]])```  
Если PARTITION BY не указано, то в качестве единственной группы для вычислений будет взят полный набор строк:

**Упорядочение в границах отдельной группы**  
С помощью синтаксической конструкции ORDER BY строки в группах вычислений можно упорядочивать. Синтаксис иллюстрируется на примере агрегирующей функции SUM:
```
SUM(выражение 1) OVER([PARTITION …]
ORDER BY выражение 2 [,…] [{ASC|DESC}] [{NULLS FIRST|NULLS LAST}])
```