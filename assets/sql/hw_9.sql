with latest_salaries as (
    select id, amount
    from employee,
         lateral
         ( select amount
           from salary
           where fk_employee = id
           order by to_date desc
           limit 1
             )
    group by id, amount
), previous_salaries as (
    select id, amount
    from employee,
         lateral
         ( select lag(amount) over (partition by fk_employee) as amount
           from salary
           where fk_employee = id
           order by to_date desc
           limit 1
             )
    group by id, amount
)
select latest_salaries.id, coalesce(latest_salaries.amount - previous_salaries.amount, 0) as salary_diff
from previous_salaries join latest_salaries on latest_salaries.id = previous_salaries.id;
