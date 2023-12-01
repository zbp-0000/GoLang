-- 单表查询练习：
-- 列出工资最小值小于2000的职位
select job,min(sal)
from emp
group by job
having min(sal) < 2000 ;


-- 列出平均工资大于1200元的部门和工作搭配组合
select deptno,job,avg(sal)
from emp
group by deptno,job
having avg(sal) > 1200
order by deptno;


-- 统计[人数小于4的]部门的平均工资。 
select deptno,count(1),avg(sal)
from emp
group by deptno
having count(1) < 4



-- 统计各部门的最高工资，排除最高工资小于3000的部门。
select deptno,max(sal)
from emp 
group by deptno
having max(sal) < 3000;


