-- 【1】查询最高工资的员工  （不相关子查询）
select * from emp where sal = (select max(sal) from emp)


-- 【2】查询本部门最高工资的员工   （相关子查询）
-- 方法1：通过不相关子查询实现：
select * from emp where deptno = 10 and sal = (select max(sal) from emp where deptno = 10)
union
select * from emp where deptno = 20 and sal = (select max(sal) from emp where deptno = 20)
union
select * from emp where deptno = 30 and sal = (select max(sal) from emp where deptno = 30)
-- 缺点：语句比较多，具体到底有多少个部分未知

-- 方法2： 相关子查询
select * from emp e where sal = (select max(sal) from emp where deptno = e.deptno) order by deptno

-- 【3】查询工资高于其所在岗位的平均工资的那些员工  （相关子查询）
-- 不相关子查询：
select * from emp where job = 'CLERK' and sal >= (select avg(sal) from emp where job = 'CLERK')
union ......
-- 相关子查询：
select * from emp e where sal >= (select avg(sal) from emp e2 where e2.job = e.job)









