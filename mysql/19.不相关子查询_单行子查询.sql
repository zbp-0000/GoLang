-- 引入子查询：
-- 查询所有比“CLARK”工资高的员工的信息  
-- 步骤1：“CLARK”工资
select sal from emp where ename = 'CLARK'  -- 2450
-- 步骤2：查询所有工资比2450高的员工的信息  
select * from emp where sal > 2450;
-- 两次命令解决问题 --》效率低 ，第二个命令依托于第一个命令，第一个命令的结果给第二个命令使用，但是
-- 因为第一个命令的结果可能不确定要改，所以第二个命令也会导致修改

-- 将步骤1和步骤2合并 --》子查询：
select * from emp where sal > (select sal from emp where ename = 'CLARK');
-- 一个命令解决问题 --》效率高


-- 单行子查询：
-- 查询工资高于平均工资的雇员名字和工资。
select ename,sal
from emp
where sal > (select avg(sal) from emp);

-- 查询和CLARK同一部门且比他工资低的雇员名字和工资。
select ename,sal
from emp
where deptno = (select deptno from emp where ename = 'CLARK') 
			and 
      sal < (select sal from emp where ename = 'CLARK')

-- 查询职务和SCOTT相同，比SCOTT雇佣时间早的雇员信息  
select * 
from emp
where job = (select job from emp where ename = 'SCOTT') 
      and 
			hiredate < (select hiredate from emp where ename = 'SCOTT')

