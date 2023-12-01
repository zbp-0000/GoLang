-- 查询员工的编号，姓名，部门编号：
select * from emp;
select empno,ename,deptno from emp;

-- 查询员工的编号，姓名，部门编号,部门名称：
select * from emp; -- 14条记录
select * from dept; -- 4条记录 


-- 多表查询 ：
-- 交叉连接：cross join
select * 
from emp
cross join dept; -- 14*4 = 56条 笛卡尔乘积 ： 没有实际意义，有理论意义

select * 
from emp
join dept; -- cross 可以省略不写，mysql中可以，oracle中不可以


-- 自然连接：natural join 
-- 优点：自动匹配所有的同名列 ,同名列只展示一次 ，简单
select * 
from emp
natural join dept;



select empno,ename,sal,dname,loc 
from emp
natural join dept;

-- 缺点： 查询字段的时候，没有指定字段所属的数据库表，效率低
-- 解决： 指定表名：
select emp.empno,emp.ename,emp.sal,dept.dname,dept.loc,dept.deptno
from emp
natural join dept;


-- 缺点：表名太长
-- 解决：表起别名
select e.empno,e.ename,e.sal,d.dname,d.loc,d.deptno
from emp e
natural join dept d;


-- 自然连接 natural join 缺点：自动匹配表中所有的同名列，但是有时候我们希望只匹配部分同名列：
-- 解决： 内连接 - using子句：
select * 
from emp e
inner join dept d -- inner可以不写
using(deptno) -- 这里不能写natural join了 ,这里是内连接

-- using缺点：关联的字段，必须是同名的 
-- 解决： 内连接 - on子句：
select * 
from emp e
inner join dept d
on (e.deptno = d.deptno);

-- 多表连接查询的类型： 1.交叉连接  cross join  2. 自然连接  natural join  
-- 3. 内连接 - using子句   4.内连接 - on子句
-- 综合看：内连接 - on子句

select * 
from emp e
inner join dept d
on (e.deptno = d.deptno)
where sal > 3500;

-- 条件：
-- 1.筛选条件  where  having
-- 2.连接条件 on,using,natural 
-- SQL99语法 ：筛选条件和连接条件是分开的






