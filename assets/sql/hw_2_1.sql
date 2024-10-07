create table pets (
                      id serial primary key,
                      pet_species text,
                      pet_name text
);

select * from pets;

insert into pets (id, pet_species, pet_name) values (default, 'cat', 'Fluffy');
insert into pets (id, pet_species, pet_name) values (default, 'dog', 'Bobby');
insert into pets (id, pet_species, pet_name) values (default, 'parrot', 'Jack');

SHOW default_transaction_isolation;

begin;
insert into pets (id, pet_species, pet_name) values (default, 'mouse', 'Stewart');
select * from pets;
commit;

begin;
insert into pets (id, pet_species, pet_name) values (default, 'snake', 'Liquid');
commit;