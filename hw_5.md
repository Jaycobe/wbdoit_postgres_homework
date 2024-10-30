1.  Поднимаю сингл инстанс в докер контейнере:


   #### Результаты тестирования

   workload.sql и workload2.sql из лекции

   чтение :
      
      pgbench (17.0 (Debian 17.0-1.pgdg120+1))
      transaction type: /var/lib/postgresql/workload.sql
      scaling factor: 1
      query mode: simple
      number of clients: 8
      number of threads: 4
      maximum number of tries: 1
      duration: 10 s
      number of transactions actually processed: 1601873
      number of failed transactions: 0 (0.000%)
      latency average = 0.050 ms
      initial connection time = 10.313 ms
      tps = 160351.339420 (without initial connection time)


   результаты тестирования на запись:

      pgbench (17.0 (Debian 17.0-1.pgdg120+1))
      transaction type: /var/lib/postgresql/workload2.sql
      scaling factor: 1
      query mode: simple
      number of clients: 8
      number of threads: 4
      maximum number of tries: 1
      duration: 10 s
      number of transactions actually processed: 146146
      number of failed transactions: 0 (0.000%)
      latency average = 0.547 ms
      initial connection time = 11.055 ms
      tps = 14630.228010 (without initial connection time)



2. Используем базу `thai`

    Измеряем время выполнения запроса `select * from book.tickets` (время 10, 50, 100, 500, 1000, 5000 запросов подряд с одного клиента)

    На сингл инстансе:
      чтение

        10 queries time: 46 ms
        50 queries time: 325 ms
        100 queries time: 269 ms
        500 queries time: 976 ms
        1000 queries time: 2534 ms
        5000 queries time: 7467 ms

      запись

        10 queries time: 67 ms
        50 queries time: 315 ms
        100 queries time: 335 ms
        500 queries time: 905 ms
        1000 queries time: 809 ms
        5000 queries time: 4164 ms

3. Поднимаем в докере второй кластер:
   создаём сеть в докере

   `docker network create postgres`

   в контейнере создаём пользователя-репликатора

   `createuser -U postgresadmin -P -c 5 --replication replicationUser`

