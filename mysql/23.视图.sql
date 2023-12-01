-- 创建/替换单表视图：
create or replace view myview01
as
select empno,ename,job,deptno 
from emp
where deptno = 20
with check option;

-- 查看视图：
select * from myview01;

-- 在视图中插入数据：
insert into myview01 (empno,ename,job,deptno) values (9999,'lili','CLERK',20);
insert into myview01 (empno,ename,job,deptno) values (8888,'nana','CLERK',30);
insert into myview01 (empno,ename,job,deptno) values (7777,'feifei','CLERK',30); 
-- > 1369 - CHECK OPTION failed 'mytestdb.myview01'


-- 创建/替换多表视图：

create or replace view myview02
as 
select e.empno,e.ename,e.sal,d.deptno,d.dname
from emp e
join dept d
on e.deptno = d.deptno
where sal > 2000 ;

select * from myview02;


-- 创建统计视图：
create or replace view myview03
as
select e.deptno,d.dname,avg(sal),min(sal),count(1)
from emp e
join dept d
using(deptno)
group by e.deptno ;


select * from myview03;


-- 创建基于视图的视图：


create or replace view myview04
as
select * from myview03 where deptno = 20;

select * from myview04;












