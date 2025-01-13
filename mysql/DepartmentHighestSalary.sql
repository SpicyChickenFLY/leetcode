# Write your MySQL query statement below
SELECT d.Name as Department, e.Name as Employee, m.Salary 
FROM Employee e, Department d, (
	SELECT MAX(Salary) as Salary, DepartmentId 
	FROM Employee 
	GROUP BY DepartmentId) m
WHERE e.DepartmentId = d.Id 
	AND d.Id = m.DepartmentId
	AND e.Salary = m.Salary
				    
