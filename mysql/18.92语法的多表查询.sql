-- 查询员工的编号，员工姓名，薪水，员工部门编号，部门名称：
select e.empno,e.ename,e.sal,e.deptno,d.dname
from emp e,dept d
-- 相当于99语法中的cross join ,出现笛卡尔积，没有意义
select e.empno,e.ename,e.sal,e.deptno,d.dname
from emp e,dept d
where e.deptno = d.deptno;
-- 相当于99语法中的natural join 


-- 查询员工的编号，员工姓名，薪水，员工部门编号，部门名称，查询出工资大于2000的员工
select e.empno,e.ename,e.sal,e.deptno,d.dname
from emp e,dept d
where e.deptno = d.deptno and e.sal > 2000;

-- 查询员工的名字，岗位，上级编号，上级名称（自连接）：
select e1.ename,e1.job,e1.mgr ,e2.ename 
from emp e1,emp e2
where e1.mgr = e2.empno;

-- 查询员工的编号、姓名、薪水、部门编号、部门名称、薪水等级
select e.empno,e.ename,e.sal,e.deptno,d.dname,s.grade 
from emp e,dept d,salgrade s
where e.deptno = d.deptno and e.sal >= s.losal and e.sal <= s.hisal;

-- 总结：
-- 1.92语法麻烦 
-- 2.92语法中 表的连接条件 和  筛选条件  是放在一起的没有分开
-- 3.99语法中提供了更多的查询连接类型：cross,natural,inner,outer 
