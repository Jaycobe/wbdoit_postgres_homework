create table pets (
                      id serial primary key,
                      pet_species text,
                      pet_name text
);

insert into pets(pet_species, pet_name)
select 'turtle', 'Leonardo'
from generate_series(0, 1000000);


SELECT pg_size_pretty( pg_total_relation_size('pets') );

select n_live_tup, n_dead_tup, relname, last_vacuum, last_autovacuum
from pg_stat_all_tables
where relname = 'pets';

do language plpgsql
$$
    declare
    begin
        for i in 1..5 loop
                raise notice 'iteration: %', i;
                update pets
                set pet_name = 'Leonardoseff';
            end loop;
    end
$$;

