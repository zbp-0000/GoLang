-- 查询员工的编号、姓名、上级编号,上级的姓名
select * from emp;

select e1.empno 员工编号,e1.ename 员工姓名,e1.mgr 领导编号,e2.ename 员工领导姓名
from emp e1
inner join emp e2
on e1.mgr = e2.empno;

-- 左外连接：
select e1.empno 员工编号,e1.ename 员工姓名,e1.mgr 领导编号,e2.ename 员工领导姓名
from emp e1
left outer join emp e2
on e1.mgr = e2.empno;