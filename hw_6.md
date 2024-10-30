# Индексы


1. В контейнере с постгресом создаю файл workload.sql с содержимым

       WITH all_place AS (
       SELECT count(s.id) as all_place, s.fkbus as fkbus
       FROM book.seat s
       group by s.fkbus
       ),
       order_place AS (
       SELECT count(t.id) as order_place, t.fkride
       FROM book.tickets t
       group by t.fkride
       )
       SELECT r.id, r.startdate as depart_date, bs.city || ', ' || bs.name as busstation,  
       t.order_place, st.all_place
       FROM book.ride r
       JOIN book.schedule as s
       on r.fkschedule = s.id
       JOIN book.busroute br
       on s.fkroute = br.id
       JOIN book.busstation bs
       on br.fkbusstationfrom = bs.id
       JOIN order_place t
       on t.fkride = r.id
       JOIN all_place st
       on r.fkbus = st.fkbus
       GROUP BY r.id, r.startdate, bs.city || ', ' || bs.name, t.order_place,st.all_place
       ORDER BY r.startdate
       limit 10;

2. Используя утилиту pgbench, измеряю производительность

       /usr/lib/postgresql/17/bin/pgbench -c 8 -j 4 -T 10 -f ~/workload.sql -n -U postgres -p 5432 thai
       pgbench (17.0 (Debian 17.0-1.pgdg120+1))
    
       transaction type: /var/lib/postgresql/workload.sql
       scaling factor: 1
       query mode: simple
       number of clients: 8
       number of threads: 4
       maximum number of tries: 1
       duration: 10 s
       number of transactions actually processed: 40
       number of failed transactions: 0 (0.000%)
       latency average = 2740.087 ms
       initial connection time = 5.387 ms
       tps = 2.919616 (without initial connection time)

    Получили 2.9 транзакции в секунду

3. Навешиваю индексы на все внешние ключи

       create index on book.tickets (fkride);
       create index on book.tickets (fkseat);

       create index on book.ride (fkbus);
       create index on book.ride (fkschedule);
    
       create index on book.schedule (fkroute);
    
       create index on book.busroute (fkbusstationfrom);
       create index on book.busroute (fkbusstationto);
    
       create index on book.seat (fkbus);

при аналогичном тесте pgbench получаю:

    postgres@d1d7425547eb:/data/postgres$ /usr/lib/postgresql/17/bin/pgbench -c 8 -j 4 -T 10 -f ~/workload.sql -n -U postgres -p 5432 thai
    pgbench (17.0 (Debian 17.0-1.pgdg120+1))
    transaction type: /var/lib/postgresql/workload.sql
    scaling factor: 1
    query mode: simple
    number of clients: 8
    number of threads: 4
    maximum number of tries: 1
    duration: 10 s
    number of transactions actually processed: 33
    number of failed transactions: 0 (0.000%)
    latency average = 2995.766 ms
    initial connection time = 6.712 ms
    tps = 2.670435 (without initial connection time)

Получаем падение производительности ~20%))

При повторных попытках измерить производительность при помощи pgbench получаю разброс от 2 до 3 tps с и без/индексов.
Поэтому провожу тесты самостоятельно


Пробую замерить скорость выполнения руками с и без индексов

Без индексов получаю время выполнения 10 запросов: 14.295s
С индексами получаю время выполнения 10 запросов: 9.345s

Код программы на Golang, которая использовалась для тестирования можно найти [здесь](./assets/go/hw_6/main.go)
