-- 查询员工的编号、姓名、薪水、部门编号、部门名称、薪水等级
select * from emp;
select * from dept;
select * from salgrade;


select e.ename,e.sal,e.empno,e.deptno,d.dname,s.* 
from emp e
right outer join dept d
on e.deptno = d.deptno
inner join salgrade s 
on e.sal between s.losal and s.hisal

