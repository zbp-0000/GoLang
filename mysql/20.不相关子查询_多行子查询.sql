-- 多行子查询：
-- 【1】查询【部门20中职务同部门10的雇员一样的】雇员信息。
-- 查询雇员信息
select * from emp;
-- 查询部门20中的雇员信息
select * from emp where deptno = 20;-- CLERK,MANAGER,ANALYST
-- 部门10的雇员的职务：
select job from emp where deptno = 10; -- MANAGER,PRESIDENT,CLERK
-- 查询部门20中职务同部门10的雇员一样的雇员信息。
select * from emp 
where deptno = 20 
and job in (select job from emp where deptno = 10)
-- > Subquery returns more than 1 row
select * from emp 
where deptno = 20 
and job = any(select job from emp where deptno = 10)


-- 【2】查询工资比所有的“SALESMAN”都高的雇员的编号、名字和工资。
-- 查询雇员的编号、名字和工资
select empno,ename,sal from emp
-- “SALESMAN”的工资：
select sal from emp where job = 'SALESMAN'
-- 查询工资比所有的“SALESMAN”都高的雇员的编号、名字和工资。
-- 多行子查询：
select empno,ename,sal 
from emp 
where sal > all(select sal from emp where job = 'SALESMAN');
-- 单行子查询：
select empno,ename,sal 
from emp 
where sal > (select max(sal) from emp where job = 'SALESMAN');


-- 【3】查询工资低于任意一个“CLERK”的工资的雇员信息。  
-- 查询雇员信息
select * from emp;
-- 查询工资低于任意一个“CLERK”的工资的雇员信息
select * 
from emp
where sal < any(select sal from emp where job = 'CLERK')
and job != 'CLERK'

-- 单行子查询：
select * 
from emp
where sal < (select max(sal) from emp where job = 'CLERK')
and job != 'CLERK'

