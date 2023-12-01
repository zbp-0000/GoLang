-- 函数举例：
select empno,ename,lower(ename),upper(ename),sal from emp;
-- 函数的功能：封装了特定的一些功能，我们直接拿过来使用，可以实现对应的功能
-- 函数作用：为了提高select的能力
-- 注意：函数没有改变数据自身的值，而是在真实数据的上面进行加工处理，展示新的结果而已。

select max(sal),min(sal),count(sal),sum(sal),avg(sal) from emp;
-- 函数的分类：
-- lower(ename),upper(ename) ：改变每一条结果，每一条数据对应一条结果  -- 单行函数
-- max(sal),min(sal),count(sal),sum(sal),avg(sal):多条数据，最终展示一个结果  -- 多行函数


-- 单行函数包含：
-- 1.字符串函数
select ename,length(ename),substring(ename,2,3) from emp;
-- substring字符串截取，2:从字符下标为2开始，3：截取长度3    （下标从1开始）
-- 2.数值函数
select abs(-5),ceil(5.3),floor(5.9),round(3.14) from dual; -- dual实际就是一个伪表 
select abs(-5) 绝对值,ceil(5.3) 向上取整,floor(5.9) 向下取整,round(3.14) 四舍五入;  -- 如果没有where条件的话，from dual可以省略不写
select ceil(sal) from emp;
select 10/3,10%3,mod(10,3) ;
-- 3.日期与时间函数 
select * from emp;
select curdate(),curtime() ; -- curdate()年月日 curtime()时分秒
select now(),sysdate(),sleep(3),now(),sysdate() from dual; -- now(),sysdate() 年月日时分秒

insert into emp values (9999,'lili','SALASMAN',7698,now(),1000,null,30);
-- now()可以表示年月日时分秒，但是插入数据的时候还是要参照表的结构的
desc emp;
-- 4.流程函数
-- if相关
select empno,ename,sal,if(sal>=2500,'高薪','底薪') as '薪资等级' from emp; -- if-else 双分支结构
select empno,ename,sal,comm,sal+ifnull(comm,0) from emp; -- 如果comm是null，那么取值为0 -- 单分支
select nullif(1,1),nullif(1,2) from dual; --  如果value1等于value2，则返回null，否则返回value1  
-- case相关：
-- case等值判断
select empno,ename,job,case job 
 when 'CLERK' then '店员'
 when 'SALESMAN'  then '销售'
 when 'MANAGER' then '经理'
 else '其他'
end '岗位',
sal from emp;

-- case区间判断:
select empno,ename,sal,
case 
 when sal<=1000 then 'A'
 when sal<=2000 then 'B'
 when sal<=3000 then 'C'
 else 'D'
end '工资等级',deptno from emp;



from emp;


-- 5.JSON函数  

-- 6.其他函数
select database(),user(),version() from dual;



-- 多行函数：
select max(sal),min(sal),count(sal),sum(sal),sum(sal)/count(sal),avg(sal) from emp;
select * from emp;
-- 多行函数自动忽略null值
select max(comm),min(comm),count(comm),sum(comm),sum(comm)/count(comm),avg(comm) from emp;
-- max(),min(),count()针对所有类型   sum(),avg() 只针对数值型类型有效
select max(ename),min(ename),count(ename),sum(ename),avg(ename) from emp;


-- count --计数   
-- 统计表的记录数：方式1：
select * from emp;
select count(ename) from emp;
select count(*) from emp;

-- 统计表的记录数：方式2
select 1 from dual;
select 1 from emp;
select count(1) from emp;






