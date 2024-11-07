create table docs (
                      id serial primary key,
                      data jsonb
);
drop table docs;

CREATE TABLE docs AS
SELECT i AS id, (SELECT jsonb_object_agg(j, j) FROM generate_series(1, 500) j) data
FROM generate_series(1, 100000) i;

do
$$
    begin
        for i in  1..1000000 loop
                insert into docs (data) values ('{"doc_id": 1, "employee_id": 123, "comment": "dfiiiiiiiiiiiiiiiiiiiiiiiiiiiiiugbbbaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa312312312312312kkkkaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaakkkkkkbbbbbbbbbbbbbbbbbbbdababa", "ext":"21333333333333333333333333333333333333333nnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnn"}');
            end loop;
    end;
$$
language plpgsql;

drop table docs;

SELECT relname, reltoastrelid FROM pg_class WHERE relname = 'docs';
SELECT relname FROM pg_class WHERE oid = 40965;
select * from pg_toast.pg_toast_40961;

create index on docs(data);

drop table docs;

SELECT oid::regclass AS heap_rel,
       pg_size_pretty(pg_relation_size(oid)) AS heap_rel_size,
       reltoastrelid::regclass AS toast_rel,
       pg_size_pretty(pg_relation_size(reltoastrelid)) AS toast_rel_size
FROM pg_class WHERE relname = 'docs';

update docs set data = (SELECT jsonb_object_agg(j, j) FROM generate_series(1, 1000) j);



select * from docs;


