-- inner join - on子句： 显示的是所有匹配的信息
select * 
from emp e
inner join dept d
on e.deptno = d.deptno;

select * from emp;
select * from dept;
-- 问题：
-- 1.40号部分没有员工，没有显示在查询结果中
-- 2.员工scott没有部门，没有显示在查询结果中

-- 外连接：除了显示匹配的数据之外，还可以显示不匹配的数据



-- 左外连接： left outer join   -- 左面的那个表的信息，即使不匹配也可以查看出效果
select * 
from emp e
left outer join dept d
on e.deptno = d.deptno;



-- 右外连接： right outer join   -- 右面的那个表的信息，即使不匹配也可以查看出效果
select * 
from emp e
right outer join dept d
on e.deptno = d.deptno;

-- 全外连接  full outer join -- 这个语法在mysql中不支持，在oracle中支持 -- 展示左，右表全部不匹配的数据 
-- scott ，40号部门都可以看到
select * 
from emp e
full outer join dept d
on e.deptno = d.deptno;


-- 解决mysql中不支持全外连接的问题：

select * 
from emp e
left outer join dept d
on e.deptno = d.deptno
union -- 并集 去重 效率低
select * 
from emp e
right outer join dept d
on e.deptno = d.deptno;




select * 
from emp e
left outer join dept d
on e.deptno = d.deptno
union all-- 并集 不去重 效率高
select * 
from emp e
right outer join dept d
on e.deptno = d.deptno;

-- mysql中对集合操作支持比较弱，只支持并集操作，交集，差集不支持（oracle中支持）
-- outer可以省略不写
