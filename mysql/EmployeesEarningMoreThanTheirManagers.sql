

select a.Name as Employee from Employee a left join Employee b on a.ManagerId = b.Id 
    where a.Salary > b.Salary;