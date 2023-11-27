# Redis 介绍

官网：https://redis.io/

Redis 可作为数据库、缓存、流引擎和消息代理的开源内存数据存储。被用在不计其数的应用中。Redis 连续 5 年被评为最受欢迎的数据库，是开发人员、架构师和开源贡献者参与社区的中心。

Redis 是一个开源（BSD 许可）的内存数据结构存储，用作数据库、缓存、消息代理和流引擎。 Redis 提供数据结构，例如字符串、散列、列表、集合、具有范围查询的排序集合、位图、Hyper日志、地理空间索引和流。 Redis 具有内置复制、Lua 脚本、LRU 驱逐、事务和不同级别的磁盘持久性，并通过 Redis Sentinel 和 Redis Cluster 自动分区提供高可用性。

您可以对这些类型运行原子操作，例如附加到字符串；增加哈希值；将元素推送到列表中；计算集交、并、差；或获取排序集中排名最高的成员。
为了达到最佳性能，Redis 使用内存数据集。根据您的用例，Redis 可以通过定期将数据集转储到磁盘或将每个命令附加到基于磁盘的日志来持久化您的数据。如果您只需要一个功能丰富的网络内存缓存，您也可以禁用持久性。

Redis 支持异步复制，具有快速非阻塞同步和自动重新连接以及网络拆分上的部分重新同步。

Redis 还包括：

- 事务
- 发布/订阅
- Lua 脚本
- 有限生命周期的 Key
- LRU 策略的 Key 清理
- 自动故障转移

Redis 支持多种语言客户端，支持几乎全部的主流语言。参见：https://redis.io/docs/clients/

Redis 适合的典型业务逻辑场景：

- 热数据，临时缓存
- 计数、统计相关，布隆过滤
- 排行榜类
- 分布式锁，秒杀
- 队列，流数据处理

# Go 操作 Redis 的客户端

## 常用包

有很多开源好用的 Go 操作 Redis 的包，列举Gihub上Star超过500的包如下：

| 包                                      | Star  |          |
| --------------------------------------- | ----- | -------- |
| https://github.com/go-redis/redis       | 15.8k | 我们选择 |
| https://github.com/gomodule/redigo      | 9.3k  |          |
| https://github.com/mediocregopher/radix | 587   |          |
| https://github.com/rueian/rueidis       | 720   |          |
| https://github.com/hoisie/redis         | 592   |          |

选择 https://github.com/go-redis/redis 包，完成 Redis 操作，go-redis 的主要特点是：

- 默认支持连接池
- 类型安全
- 内置 Cluster、Sentinel、Ring 等多种类型客户端
- 支持 OpenTelemetry 指标数据统计
- 用基数大
- 还可以作为 kvrocks 的客户端使用

https://redis.uptrace.dev/

https://pkg.go.dev/github.com/go-redis/redis/v9

## 安装 go-redis

go-redis/redis 包有不同版本，其中：

如果使用 Redis6，则：

```shell
go get github.com/go-redis/redis/v8
```

如果使用 Redis 7（Redis7于2022 Apr 27 发布），则：

```shell
go get github.com/go-redis/redis/v9
```

如果不能确定Redis版本，使用命令 `redis-server --version` 或客户端登录后，使用名 `info server` 可以查询：

```shell
# docker exec -it redis-dev redis-server --version
Redis server v=7.0.5 sha=00000000:0 malloc=jemalloc-5.2.1 bits=64 build=a507a0b937977997
```

```shell
127.0.0.1:6379>info server
# Server
redis_version:7.0.5

```

## 连接单点模式服务

单点模式服务指的是使用单一的redis服务器实现Redis服务，最简单的一种模式。

### 启动单点模式服务

启动，Docker 方式：

```shell
$ docker run --rm --name redis-single -p 6379:6379 -d \
	redis redis-server --requirepass yourPassword

```

### 连接（使用连接池）

经典的方式是通过地址建立连接：

```go
import "github.com/go-redis/redis/v9"

rdb := redis.NewClient(&redis.Options{
		Addr:        "192.168.157.135:6379",
		Username:    "default",       // default user
		Password:    "some-pass",     // no password set
		DB:          0,               // use default DB
		DialTimeout: 1 * time.Second, // 1 second
	})

```

另一种方式是基于连接字符串建立连接：

```go
import "github.com/go-redis/redis/v9"

opt, err := redis.ParseURL("redis://default:some-pass@192.168.157.135:6379/0?dial_timeout=1")
	if err != nil {
		panic(err)
	}

rdb := redis.NewClient(opt)
```

常用选项：

```go
Network string // 网络类型 Default is tcp
Addr string // 地址 host:port
Username string // ACL 用户名
Password string // ACL 密码
DB int // 所选数据库
DialTimeout time.Duration // 拨号连接超时时间，default 5 seconds
ReadTimeout time.Duration // 读操作超时时间
WriteTimeout time.Duration // 写操作超时时间
PoolTimeout time.Duration // 连接池等待超时时间
PoolSize int // 连接池大小
```

注意，NewClient() 操作，不会立即对 Redis 服务器建立连接，仅用来配置连接选项。只有当对 Redis 服务器发出第一次操作时，才会建立连接。例如，`rdb.Ping()` 可用来测试服务器的连接情况：

```go
status := rdb.Ping(context.Background())
fmt.Println(status.Result())
```

go-redis 采用连接池方案管理连接，做到连接复用。

## 连接集群模式服务

## 连接哨兵模式服务

## 连接池

## 单连接

## 使用 TLS

## 通过 SSH

# 执行命令

执行 Redis 命令，需要选择特定命令，同时传递 context 对象。

## Context

## redis.Nil

redis.Nil 是 go-redis 定义的特定错误，用于表示获取的 Key 不存在。通常获取类的命令都会使用 redis.Nil，例如 GET、BLPOP、ZSCORE 等。

```go
const Nil = RedisError("redis: nil") // nolint:errname
```

之所以需要定义 redis.Nil，是因为 go 在获取结果时，不容易区分是空字符串“”还是 Key 不存在，例如：

```go
val, _ := client.Get(ctx, "someKey").Result()
fmt.Printf("%#v", val) // ""
```

在 someKey 不存在的情况下，value 的值是 ""，也就是如果不去判断错误，不能确定是 someKey 不存在或者是 someKey 的值本就是“”。

因此，当需要确定 key 是否存在时，通常的逻辑如下：

```go
val, err := client.Get(ctx, "someKey").Result()
switch {
    case err == redis.Nil:
    fmt.Println("key does not exist")
    case err != nil:
    fmt.Println("Get failed", err)
    case val == "":
    fmt.Println("value is empty")
}
```

# String字符串操作

## String 介绍

Redis字符串存储字节序列，包括文本、序列化对象和二进制数组。因此，字符串是最基本的Redis数据类型。它们通常用于缓存，但它们支持额外的功能，允许您实现计数器并执行按位操作。

默认情况下，string 的最大尺寸为：512MB。

大多数的 string 操作时间复杂度为：O(1)。

但 SUBSTR, GETRANGE， SETRANGE 命令的时间复杂度为O(n)。

## 操作方法说明

### 设置

```go
// 设置
func (c Client) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *StatusCmd
// 设置，使用参数配置。Mode（`NX` or `XX` or empty），TTL，ExpireAt，Get，KeepTTL
func (c Client) SetArgs(ctx context.Context, key string, value interface{}, a SetArgs) *StatusCmd

// 设置，同时设置有效期
func (c Client) SetEX(ctx context.Context, key string, value interface{}, expiration time.Duration) *StatusCmd
// 设置，前提是Key不存在
func (c Client) SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) *BoolCmd
// 设置，前提是Key存在
func (c Client) SetXX(ctx context.Context, key string, value interface{}, expiration time.Duration) *BoolCmd
// 注意：SetEx,SetNx,SetXX 的功能可以由 Set、SetArgs完成，因此建议以使用 Set、SetArgs 为主

// 设置多个
func (c Client) MSet(ctx context.Context, values ...interface{}) *StatusCmd
// 设置多个，前提是全部的key都要不存在
func (c Client) MSetNX(ctx context.Context, values ...interface{}) *BoolCmd

```

其中：

- NX 表示Not Exists，不存在时才执行
- XX 表示 Exists，存在时才执行

### 获取

```go
// 获取
func (c Client) Get(ctx context.Context, key string) *StringCmd
// 获取并删除
func (c Client) GetDel(ctx context.Context, key string) *StringCmd
// 获取并设置有效期
func (c Client) GetEx(ctx context.Context, key string, expiration time.Duration) *StringCmd

// 获取多个
func (c Client) MGet(ctx context.Context, keys ...string) *SliceCmd

// 获取字符串长度
func (c Client) StrLen(ctx context.Context, key string) *IntCmd
```

### 追加和递增递减

```go
// 追加字符串
func (c Client) Append(ctx context.Context, key, value string) *IntCmd

// 递减，字符串作为 64位整数存储
func (c Client) Decr(ctx context.Context, key string) *IntCmd
func (c Client) DecrBy(ctx context.Context, key string, decrement int64) *IntCmd

// 递增
func (c Client) Incr(ctx context.Context, key string) *IntCmd
func (c Client) IncrBy(ctx context.Context, key string, value int64) *IntCmd
func (c Client) IncrByFloat(ctx context.Context, key string, value float64) *FloatCmd
```

其中：

- value 被认为是64bit的整数

### 子串操作

```go
// 设置特定字符串的特定索引内容
func (c Client) SetRange(ctx context.Context, key string, offset int64, value string) *IntCmd
// 获取特定范围的字符串（substr）
func (c Client) GetRange(ctx context.Context, key string, start, end int64) *StringCmd
```

其中：

- start，end，是索引，都是闭区间
- 负值表示从后计数第几个，-1表示倒数第一个字节
- 不会越界

### Do

```
LCS key1 key2 [LEN] [IDX] [MINMATCHLEN len] [WITHMATCHLEN]
```

# Bitmap位图操作

## Bitmap 介绍

![image.png](https://fynotefile.oss-cn-zhangjiakou.aliyuncs.com/fynote/fyfile/13080/1667211781083/63b128a7d7a945f9ac5d06f6b52a752b.png)

Redis位图是字符串数据类型的扩展，可以将字符串视为位向量。您还可以对一个或多个字符串执行按位操作。位图用例的一些示例包括：

- 集合的成员对应于整数0-N的情况下的有效集合表示。
- 对象权限，其中每个位代表一个特定的权限，类似于文件系统存储权限的方式。

限制：

- 最大支持 2^32 个bit，也就是字符串的限制 512 MB

性能：

- SETBIT 和 GETBIT 时间复杂度是 O(1)
- BITOP 时间复杂度是 O(n)

## 操作方法说明

### 设置位

```go
// 设置特定位
func (c Client) SetBit(ctx context.Context, key string, offset int64, value int) *IntCmd

```

### 获取位

```go
// 获取特定位
func (c Client) GetBit(ctx context.Context, key string, offset int64) *IntCmd
// 统计位数量
func (c Client) BitCount(ctx context.Context, key string, bitCount *BitCount) *IntCmd
// 获取特定位的位置
func (c Client) BitPos(ctx context.Context, key string, bit int64, pos ...int64) *IntCmd

```

### 字符串操作位图

```go
// 设置
func (c Client) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *StatusCmd
// 获取
func (c Client) Get(ctx context.Context, key string) *StringCmd
```

字符串：

- "\xHH" 支持16进制编码字符转义
- 支持常规字符串，ASCII字符

### 位运算

```go
// 位与
func (c Client) BitOpAnd(ctx context.Context, destKey string, keys ...string) *IntCmd
// 位非
func (c Client) BitOpNot(ctx context.Context, destKey string, key string) *IntCmd
// 位或
func (c Client) BitOpOr(ctx context.Context, destKey string, keys ...string) *IntCmd
// 位异或
func (c Client) BitOpXor(ctx context.Context, destKey string, keys ...string) *IntCmd
```

### 多位复合操作

```go
// 一次性设置多个位
func (c Client) BitField(ctx context.Context, key string, args ...interface{}) *IntSliceCmd
```

支持设置、递增、获取特定的位。同时支持整数的有无符号和宽度编码。

# 通用操作

通用操作指的是绝大多数数据类型都支持的操作，例如拷贝、设置有效期、删除、判定是否存在等。

## 常用

```go
// 拷贝
func (c Client) Copy(ctx context.Context, sourceKey, destKey string, db int, replace bool) *IntCmd
// 删除
func (c Client) Del(ctx context.Context, keys ...string) *IntCmd
// 异步删除
func (c Client) Unlink(ctx context.Context, keys ...string) *IntCmd
// 判定是否存在
func (c Client) Exists(ctx context.Context, keys ...string) *IntCmd
// 重命名
func (c Client) Rename(ctx context.Context, key, newkey string) *StatusCmd
// 获取数据类型
func (c Client) Type(ctx context.Context, key string) *StatusCmd
// 返回值的序列化内容
func (c Client) Dump(ctx context.Context, key string) *StringCmd
// 基于串行化的数据重新存储
func (c Client) Restore(ctx context.Context, key string, ttl time.Duration, value string) *StatusCmd
// 获取全部key
func (c Client) Keys(ctx context.Context, pattern string) *StringSliceCmd
// 分片遍历key
func (c Client) Scan(ctx context.Context, cursor uint64, match string, count int64) *ScanCmd
// 随机获取key
func (c Client) RandomKey(ctx context.Context) *StringCmd
```

## 有效期

```go
// 设置有效期，时间段
func (c Client) Expire(ctx context.Context, key string, expiration time.Duration) *BoolCmd
// 设置有效期，要求新有效期大于原来有效期
func (c Client) ExpireGT(ctx context.Context, key string, expiration time.Duration) *BoolCmd
// 设置有效期，要求新有效期小于原来有效期
func (c Client) ExpireLT(ctx context.Context, key string, expiration time.Duration) *BoolCmd
// 设置有效期，要求key不存在
func (c Client) ExpireNX(ctx context.Context, key string, expiration time.Duration) *BoolCmd
// 设置有效期，要求key存在
func (c Client) ExpireXX(ctx context.Context, key string, expiration time.Duration) *BoolCmd
// 设置有效期，时间
func (c Client) ExpireAt(ctx context.Context, key string, tm time.Time) *BoolCmd
// 获取有效期
func (c Client) TTL(ctx context.Context, key string) *DurationCmd
// 移除有效期
func (c Client) Persist(ctx context.Context, key string) *BoolCmd

// P前缀，表示毫秒时间级别 milliseconds
func (c Client) PExpire(ctx context.Context, key string, expiration time.Duration) *BoolCmd
func (c Client) PExpireAt(ctx context.Context, key string, tm time.Time) *BoolCmd
func (c Client) PTTL(ctx context.Context, key string) *DurationCmd

// 更新key的访问时间
func (c Client) Touch(ctx context.Context, keys ...string) *IntCmd
```

## 排序

```go
// 排序
func (c Client) Sort(ctx context.Context, key string, sort *Sort) *StringSliceCmd
func (c Client) SortStore(ctx context.Context, key, store string, sort *Sort) *IntCmd
func (c Client) SortInterfaces(ctx context.Context, key string, sort *Sort) *SliceCmd
```

## 内部信息

```go
// 获取对象内部编码
func (c Client) ObjectEncoding(ctx context.Context, key string) *StringCmd
// 获取对象的空闲时间
func (c Client) ObjectIdleTime(ctx context.Context, key string) *DurationCmd
// 获取对象的引用计数
func (c Client) ObjectRefCount(ctx context.Context, key string) *IntCmd
```

## 转移

```go
// 转移到不同库
func (c Client) Move(ctx context.Context, key string, db int) *BoolCmd
// 转移到不同服务器
func (c Client) Migrate(ctx context.Context, host, port, key string, db int, timeout time.Duration) *StatusCmd
```

## Do

```
OBJECT FREQ key
```

# SCAN, SSCAN, HSCAN, ZSCAN 命令

Redis 的 SCAN 命令及其相关命令 SSCAN, HSCAN ZSCAN 命令都是用于增量遍历集合中的元素。

- `SCAN` 命令用于迭代当前数据库中的数据库键
- `SSCAN` 命令用于迭代集合键中的元素
- `HSCAN` 命令用于迭代哈希键中的键值对
- `ZSCAN` 命令用于迭代有序集合中的元素（包括元素成员和元素分值）

# List列表操作

## List 介绍

List 是字符串值的链表。List 经常用于：

- 实现堆栈stacks 和队列queues。
- 通过索引检索元素

List 允许的最大长度（元素个数）：2^32 -1。

List 支持：

- 通过索引操作元素
- 在头部和尾部插入和取出元素
- 在特定索引位置插入元素
- 支持特定命令的阻塞操作，命令通常以B（Block）开头

性能：

- 通常的时间复杂度为O(n)，例如 LINDEX
- 头部和尾部元素操作为O(1)，例如 LPUSH

因此，若使用大 List，尽量使用头部和尾部操作为主的设计。

## 操作方法说明

| left | Element | right |
| ---- | ------- | ----- |

### 添加元素

```go
// 在头部插入元素
func (c Client) LPush(ctx context.Context, key string, values ...interface{}) *IntCmd
// 在头部插入元素，前提是 key 对应的 List 已经存在
func (c Client) LPushX(ctx context.Context, key string, values ...interface{}) *IntCmd
// 在尾部插入元素
func (c Client) RPush(ctx context.Context, key string, values ...interface{}) *IntCmd
// 在尾部插入元素，前提是 key 对应的 List 已经存在
func (c Client) RPushX(ctx context.Context, key string, values ...interface{}) *IntCmd
```

### 取出元素

```go
// 弹出头部的第一个元素
func (c Client) LPop(ctx context.Context, key string) *StringCmd
// 弹出头部的特定数量的元素
func (c Client) LPopCount(ctx context.Context, key string, count int) *StringSliceCmd
// 弹出尾部的第一个元素
func (c Client) RPop(ctx context.Context, key string) *StringCmd
// 弹出尾部的特定数量的元素
func (c Client) RPopCount(ctx context.Context, key string, count int) *StringSliceCmd

```

### 删除元素

```go
// 删除特定数量的特定元素
func (c Client) LRem(ctx context.Context, key string, count int64, value interface{}) *IntCmd
```

### 获取长度

```go
// 获取长度
func (c Client) LLen(ctx context.Context, key string) *IntCmd
```

### 基于索引操作

```go
// 利用索引获取元素
func (c Client) LIndex(ctx context.Context, key string, index int64) *StringCmd
// 设置特定索引对应的元素
func (c Client) LSet(ctx context.Context, key string, index int64, value interface{}) *StatusCmd
// 在某个元素前或后插入元素
func (c Client) LInsert(ctx context.Context, key, op string, pivot, value interface{}) *IntCmd
// 在某个元素后插入
func (c Client) LInsertAfter(ctx context.Context, key string, pivot, value interface{}) *IntCmd
// 在某个元素前插入
func (c Client) LInsertBefore(ctx context.Context, key string, pivot, value interface{}) *IntCmd
// 获取匹配元素的一个索引
func (c Client) LPos(ctx context.Context, key string, value string, a LPosArgs) *IntCmd
// 获取匹配元素的多个索引
func (c Client) LPosCount(ctx context.Context, key string, value string, count int64, a LPosArgs) *IntSliceCmd

// 基于索引裁剪 List
func (c Client) LTrim(ctx context.Context, key string, start, stop int64) *StatusCmd
```

### 操作部分元素

```go
// 获取指定范围的元素
func (c Client) LRange(ctx context.Context, key string, start, stop int64) *StringSliceCmd
```

### 多队列间操作

```go
// 移动元素
func (c Client) LMove(ctx context.Context, source, destination, srcpos, destpos string) *StringCmd
// 从 source 弹出，然后插入到 destination 中
func (c Client) RPopLPush(ctx context.Context, source, destination string) *StringCmd
```

### 阻塞操作

```go
// 支持阻塞的 LMove，当 source 为空时，会阻塞连接，直到其他客户端插入元素或超时
func (c Client) BLMove(ctx context.Context, source, destination, srcpos, destpos string, ...) *StringCmd
// 支持阻塞的 LPop
func (c Client) BLPop(ctx context.Context, timeout time.Duration, keys ...string) *StringSliceCmd
// 支持阻塞的 RPop
func (c Client) BRPop(ctx context.Context, timeout time.Duration, keys ...string) *StringSliceCmd
// 支持阻塞的 RPopLPush
func (c Client) BRPopLPush(ctx context.Context, source, destination string, timeout time.Duration) *StringCmd
```

## 操作方法说明

### 添加元素

```go
// 在头部插入元素
func (c Client) LPush(ctx context.Context, key string, values ...interface{}) *IntCmd
// 在头部插入元素，前提是 key 对应的 List 已经存在
func (c Client) LPushX(ctx context.Context, key string, values ...interface{}) *IntCmd
// 在尾部插入元素
func (c Client) RPush(ctx context.Context, key string, values ...interface{}) *IntCmd
// 在尾部插入元素，前提是 key 对应的 List 已经存在
func (c Client) RPushX(ctx context.Context, key string, values ...interface{}) *IntCmd
```

### 获取长度

```go
// 获取长度
func (c Client) LLen(ctx context.Context, key string) *IntCmd
```

### 取出元素

```go
// 弹出头部的第一个元素
func (c Client) LPop(ctx context.Context, key string) *StringCmd
// 弹出头部的特定数量的元素
func (c Client) LPopCount(ctx context.Context, key string, count int) *StringSliceCmd
// 弹出尾部的第一个元素
func (c Client) RPop(ctx context.Context, key string) *StringCmd
// 弹出尾部的特定数量的元素
func (c Client) RPopCount(ctx context.Context, key string, count int) *StringSliceCmd
// 删除特定数量的特定元素
func (c Client) LRem(ctx context.Context, key string, count int64, value interface{}) *IntCmd
```

### 基于索引操作

```go
// 利用索引获取元素
func (c Client) LIndex(ctx context.Context, key string, index int64) *StringCmd
// 设置特定索引对应的元素
func (c Client) LSet(ctx context.Context, key string, index int64, value interface{}) *StatusCmd
// 在某个元素前或后插入元素
func (c Client) LInsert(ctx context.Context, key, op string, pivot, value interface{}) *IntCmd
// 在某个元素后插入
func (c Client) LInsertAfter(ctx context.Context, key string, pivot, value interface{}) *IntCmd
// 在某个元素前插入
func (c Client) LInsertBefore(ctx context.Context, key string, pivot, value interface{}) *IntCmd
// 获取匹配元素的一个索引
func (c Client) LPos(ctx context.Context, key string, value string, a LPosArgs) *IntCmd
// 获取匹配元素的多个索引
func (c Client) LPosCount(ctx context.Context, key string, value string, count int64, a LPosArgs) *IntSliceCmd

// 基于索引裁剪 List
func (c Client) LTrim(ctx context.Context, key string, start, stop int64) *StatusCmd
```

### 操作部分元素

```go
// 获取指定范围的元素
func (c Client) LRange(ctx context.Context, key string, start, stop int64) *StringSliceCmd
```

### 多队列间操作

```go
// 移动元素
func (c Client) LMove(ctx context.Context, source, destination, srcpos, destpos string) *StringCmd
```

### 阻塞操作

```go
// 支持阻塞的 LMove，当 source 为空时，会阻塞连接，直到其他客户端插入元素或超时
func (c Client) BLMove(ctx context.Context, source, destination, srcpos, destpos string, ...) *StringCmd
// 支持阻塞的 LPop
func (c Client) BLPop(ctx context.Context, timeout time.Duration, keys ...string) *StringSliceCmd
// 支持阻塞的 RPop
func (c Client) BRPop(ctx context.Context, timeout time.Duration, keys ...string) *StringSliceCmd
// 支持阻塞的 RPopLPush
func (c Client) BRPopLPush(ctx context.Context, source, destination string, timeout time.Duration) *StringCmd
```

# Set集合操作

## Set 介绍

![image.png](https://fynotefile.oss-cn-zhangjiakou.aliyuncs.com/fynote/fyfile/13080/1667211781083/29cc2aaf22914a9191bfbb55c5911b6c.png)

Redis 集是唯一字符串（成员）的无序集合。您可以使用 Redis 集高效地：

- 跟踪唯一项目（例如，跟踪访问给定博客文章的所有唯一 IP 地址）。
- 表示关系（例如，具有给定角色的所有用户的集合）。
- 执行常见的集合运算，例如交集、并集和差集。

Set 的最大成员数量为：2^32-1。

性能：

大多数的操作时间复杂度为 O(1)。

但获取全部成员 SMEMBERS 类的操作时间复杂度为O(n)，因此使用时要格外注意，如果可以，SSCAN 也是一个选择。

## 操作方法说明

### 添加成员

```go
// 添加成员
func (c Client) SAdd(ctx context.Context, key string, members ...interface{}) *IntCmd

```

### 获取成员

```go
// 随机获取成员
func (c Client) SRandMember(ctx context.Context, key string) *StringCmd
// 随机获取N个成员
func (c Client) SRandMemberN(ctx context.Context, key string, count int64) *StringSliceCmd
// 随机获取并删除成员
func (c Client) SPop(ctx context.Context, key string) *StringCmd
// 随机获取并删除N个成员
func (c Client) SPopN(ctx context.Context, key string, count int64) *StringSliceCmd

// 获取全部成员
func (c Client) SMembers(ctx context.Context, key string) *StringSliceCmd
// 获取全部成员，返回 Map 结构
func (c Client) SMembersMap(ctx context.Context, key string) *StringStructMapCmd

// 分片获取成员
func (c Client) SScan(ctx context.Context, key string, cursor uint64, match string, count int64) *ScanCmd
```

### 删除成员

```go
// 删除一个或多个成员
func (c Client) SRem(ctx context.Context, key string, members ...interface{}) *IntCmd
```

### 统计

```go
// 统计成员数量
func (c Client) SCard(ctx context.Context, key string) *IntCmd
// 判定成员是否存在
func (c Client) SIsMember(ctx context.Context, key string, member interface{}) *BoolCmd
// 判定给出的每个成员是否存在
func (c Client) SMIsMember(ctx context.Context, key string, members ...interface{}) *BoolSliceCmd
```

### 移动成员

```go
// 将成员从一个集合移动到另一个集合
func (c Client) SMove(ctx context.Context, source, destination string, member interface{}) *BoolCmd
```

### 集合运算

```go
// 差集
func (c Client) SDiff(ctx context.Context, keys ...string) *StringSliceCmd
// 计算差集并存储结果
func (c Client) SDiffStore(ctx context.Context, destination string, keys ...string) *IntCmd
// 交集
func (c Client) SInter(ctx context.Context, keys ...string) *StringSliceCmd
// 计算交集并存储结果
func (c Client) SInterStore(ctx context.Context, destination string, keys ...string) *IntCmd
// 并集
func (c Client) SUnion(ctx context.Context, keys ...string) *StringSliceCmd
// 计算并集并存储结果
func (c Client) SUnionStore(ctx context.Context, destination string, keys ...string) *IntCmd
```

![image.png](https://fynotefile.oss-cn-zhangjiakou.aliyuncs.com/fynote/fyfile/13080/1667211781083/367362997de748798d0e6621ec9bcd9f.png)

# Sorted Set排序集合操作

## Sorted Set 介绍

![image.png](https://fynotefile.oss-cn-zhangjiakou.aliyuncs.com/fynote/fyfile/13080/1667211781083/f298b33f8a0e49fa994016b3cafe97dd.png)

Redis 排序集是由相关分数排序的唯一字符串（成员）的集合。当多个字符串具有相同的分数时，这些字符串按字典顺序排列。排序集的一些用例包括：

- 排行榜。例如，您可以使用排序集轻松维护大型在线游戏中最高分的有序列表。
- 速率限制器。特别是，您可以使用排序集来构建滑动窗口速率限制器，以防止过多的 API 请求。

性能：

大多数的命令时间复杂度为：O(long(n))。若执行 ZRANGE 命令，获取 m 个元素时，时间复杂度为：O(log(n) + m)。

## 操作方法说明

### 设置成员

```go
// 添加成员
func (c Client) ZAdd(ctx context.Context, key string, members ...*Z) *IntCmd
// 添加成员，指定选项，支持：NX，XX，LT, GT, Ch, Members 选项，LT，GT，NX 为互斥选项
func (c Client) ZAddArgs(ctx context.Context, key string, args ZAddArgs) *IntCmd
// 递增成员的分值
func (c Client) ZIncrBy(ctx context.Context, key string, increment float64, member string) *FloatCmd
// 添加成员，指定选项，支持：NX，XX，LT, GT, Ch, Members 选项，做递增操作
func (c Client) ZAddArgsIncr(ctx context.Context, key string, args ZAddArgs) *FloatCmd

// 添加成员，返回被修改（添加的和更新的）的成员数量
func (c Client) ZAddCh(ctx context.Context, key string, members ...*Z) *IntCmd
// 添加成员，仅添加新成员，不更新成员
func (c Client) ZAddNX(ctx context.Context, key string, members ...*Z) *IntCmd
// 添加成员，仅添加新成员，不更新成员，返回添加的成员数量
func (c Client) ZAddNXCh(ctx context.Context, key string, members ...*Z) *IntCmd
// 添加成员，仅更新成员，不添加新成员
func (c Client) ZAddXX(ctx context.Context, key string, members ...*Z) *IntCmd
// 添加成员，仅更新成员，不添加新成员，返回更新的成员数量
func (c Client) ZAddXXCh(ctx context.Context, key string, members ...*Z) *IntCmd

// `ZADD key INCR score member`， 将弃用
func (c Client) ZIncr(ctx context.Context, key string, member *Z) *FloatCmd
// `ZADD key NX INCR score member`， 将弃用
func (c Client) ZIncrNX(ctx context.Context, key string, member *Z) *FloatCmd
// `ZADD key XX INCR score member`， 将弃用
func (c Client) ZIncrXX(ctx context.Context, key string, member *Z) *FloatCmd
```

### 获取成员

```go
// 获取给定的成员分值
func (c Client) ZScore(ctx context.Context, key, member string) *FloatCmd
// 获取给定的成员（多个）分值
func (c Client) ZMScore(ctx context.Context, key string, members ...string) *FloatSliceCmd

// 随机获取 count 个成员
func (c Client) ZRandMember(ctx context.Context, key string, count int) *StringSliceCmd
// 随机获取 count 个成员，返回带有分值
func (c Client) ZRandMemberWithScores(ctx context.Context, key string, count int) *ZSliceCmd
```

### Range 操作

```go
// 获取特定范围的成员
func (c Client) ZRange(ctx context.Context, key string, start, stop int64) *StringSliceCmd
// 获取特定范围的成员，利用参数获取，参数：Key，Start，Stop，ByScore，ByLex，Rev，Offset，Count
func (c Client) ZRangeArgs(ctx context.Context, z ZRangeArgs) *StringSliceCmd

// 获取特定范围的成员，返回成员及分数
func (c Client) ZRangeWithScores(ctx context.Context, key string, start, stop int64) *ZSliceCmd
// 获取特定范围的成员，利用参数获取，返回成员及分数
func (c Client) ZRangeArgsWithScores(ctx context.Context, z ZRangeArgs) *ZSliceCmd

// 获取特定范围的成员并存储
func (c Client) ZRangeStore(ctx context.Context, dst string, z ZRangeArgs) *IntCmd

// 获取特定范围的成员，利用参数获取,ByLex = true
func (c Client) ZRangeByLex(ctx context.Context, key string, opt *ZRangeBy) *StringSliceCmd
// 获取特定范围的成员，利用参数获取,ByScore = true
func (c Client) ZRangeByScore(ctx context.Context, key string, opt *ZRangeBy) *StringSliceCmd
// 获取特定范围的成员，利用参数获取,ByScore = true，返回成员及分数
func (c Client) ZRangeByScoreWithScores(ctx context.Context, key string, opt *ZRangeBy) *ZSliceCmd
// 以下Rev的都是对应方法的倒序版
func (c Client) ZRevRange(ctx context.Context, key string, start, stop int64) *StringSliceCmd
func (c Client) ZRevRangeByLex(ctx context.Context, key string, opt *ZRangeBy) *StringSliceCmd
func (c Client) ZRevRangeByScore(ctx context.Context, key string, opt *ZRangeBy) *StringSliceCmd
func (c Client) ZRevRangeByScoreWithScores(ctx context.Context, key string, opt *ZRangeBy) *ZSliceCmd
func (c Client) ZRevRangeWithScores(ctx context.Context, key string, start, stop int64) *ZSliceCmd

// 分片遍历全部成员
func (c Client) ZScan(ctx context.Context, key string, cursor uint64, match string, count int64) *ScanCmd
```

Range 支持不同的查询类型

- Rank，索引顺序
- Score，分值
- Lex，成员词法顺序

不同的查询类型，start 和 stop 表示的含义如下：

分值范围：

- 默认为闭区间: start, stop 表示 [start, stop]
- 可是设置任意一端为开区间，例如：(start,  (stop 表示 （start, stop)
- -inf, +inf 表示负无穷到正无穷

索引范围：

- 默认为闭区间索引，例如：1, 2 表示 [1, 2]
- 负数索引表示从集合的后边计数索引，例如 -1 表示最后一个元素
- 索引越界不会触发错误

词法范围：

- +, - 表示词法的正负无穷
- 有效的词法语法要以 （ 或 [ 开头，表示开或闭区间

### 删除成员

```go
// 删除一个或多个成员
func (c Client) ZRem(ctx context.Context, key string, members ...interface{}) *IntCmd
// 删除特定词法范围的成员
func (c Client) ZRemRangeByLex(ctx context.Context, key, min, max string) *IntCmd
// 删除特定排名范围的成员
func (c Client) ZRemRangeByRank(ctx context.Context, key string, start, stop int64) *IntCmd
// 删除特定分值范围的成员
func (c Client) ZRemRangeByScore(ctx context.Context, key, min, max string) *IntCmd
```

### Pop 操作

```go
// 获取并删除分值最高的 count 个成员
func (c Client) ZPopMax(ctx context.Context, key string, count ...int64) *ZSliceCmd
// 获取并删除分值最低的 count 个成员
func (c Client) ZPopMin(ctx context.Context, key string, count ...int64) *ZSliceCmd
```

### 统计

```go
// 统计成员数量
func (c Client) ZCard(ctx context.Context, key string) *IntCmd
// 统计分值在特定区间内的成员数量
func (c Client) ZCount(ctx context.Context, key, min, max string) *IntCmd
// 统计成员在某个词法区间的数量
func (c Client) ZLexCount(ctx context.Context, key, min, max string) *IntCmd
// 统计成员的索引位置
func (c Client) ZRank(ctx context.Context, key, member string) *IntCmd
// ZRank 的逆序版
func (c Client) ZRevRank(ctx context.Context, key, member string) *IntCmd
```

### 集合运算

```go
// 差集，返回成员列表
func (c Client) ZDiff(ctx context.Context, keys ...string) *StringSliceCmd
// 差集，返回成员及分数列表
func (c Client) ZDiffWithScores(ctx context.Context, keys ...string) *ZSliceCmd
// 差集并存储
func (c Client) ZDiffStore(ctx context.Context, destination string, keys ...string) *IntCmd

// 交集，返回成员列表
func (c Client) ZInter(ctx context.Context, store *ZStore) *StringSliceCmd
// 交集，返回成员及分数列表
func (c Client) ZInterWithScores(ctx context.Context, store *ZStore) *ZSliceCmd
// 交集并存储
func (c Client) ZInterStore(ctx context.Context, destination string, store *ZStore) *IntCmd

// 并集，返回成员列表
func (c Client) ZUnion(ctx context.Context, store ZStore) *StringSliceCmd
// 并集并存储
func (c Client) ZUnionStore(ctx context.Context, dest string, store *ZStore) *IntCmd
// 并集，返回成员及分数列表
func (c Client) ZUnionWithScores(ctx context.Context, store ZStore) *ZSliceCmd

```

### 支持阻塞

```go
//  BPopMax 的阻塞版，阻塞到其中一个pop成功
func (c Client) BZPopMax(ctx context.Context, timeout time.Duration, keys ...string) *ZWithKeyCmd
//  BPopMin 的阻塞版
func (c Client) BZPopMin(ctx context.Context, timeout time.Duration, keys ...string) *ZWithKeyCmd
```

# Hash哈希操作

## Hash 介绍

Redis Hash是作为字段值对的集合构造的记录类型。您可以使用哈希来表示基本对象和存储计数器分组等。

限制：

每个Hash的键值对存储上限为：2^32-1。

性能：

大多数的Hash操作的时间复杂度为：O(1)。

部分命令，HKEYS, HVALS, 和 HGETALL 的时间复杂度为：O(n)

## Client 支持的 Hash 的命令

```go
// 删除字段，支持多个
func (c Client) HDel(ctx context.Context, key string, fields ...string) *IntCmd
// 判断某个字段是否存在
func (c Client) HExists(ctx context.Context, key, field string) *BoolCmd
// 获取特定字段值
func (c Client) HGet(ctx context.Context, key, field string) *StringCmd
// 获取全部字段及值
func (c Client) HGetAll(ctx context.Context, key string) *MapStringStringCmd
// 字段递增
func (c Client) HIncrBy(ctx context.Context, key, field string, incr int64) *IntCmd
// 字段递增，浮点数
func (c Client) HIncrByFloat(ctx context.Context, key, field string, incr float64) *FloatCmd
// 获取全部字段
func (c Client) HKeys(ctx context.Context, key string) *StringSliceCmd
// 获取Hash长度
func (c Client) HLen(ctx context.Context, key string) *IntCmd
// 获取多个字段值
func (c Client) HMGet(ctx context.Context, key string, fields ...string) *SliceCmd
// 设置字段值，支持多个（建议使用 HSet 代替，该命令已过期）
func (c Client) HMSet(ctx context.Context, key string, values ...interface{}) *BoolCmd
// 获取随机字段，支持多个
func (c Client) HRandField(ctx context.Context, key string, count int) *StringSliceCmd
// 获取随机字段及值，支持多个
func (c Client) HRandFieldWithValues(ctx context.Context, key string, count int) *KeyValueSliceCmd
// 遍历字段及值
func (c Client) HScan(ctx context.Context, key string, cursor uint64, match string, count int64) *ScanCmd
// 设置字段值，支持多个
func (c Client) HSet(ctx context.Context, key string, values ...interface{}) *IntCmd
// 设置特定字段值，前提是该字段不存在
func (c Client) HSetNX(ctx context.Context, key, field string, value interface{}) *BoolCmd
// 获取全部字段值
func (c Client) HVals(ctx context.Context, key string) *StringSliceCmd
```

不支持的命令:

```go
# 获取特定字段长度
HSTRLEN key field

func (c *Client) Do(ctx context.Context, "HTRRLEN", "key", "field") *Cmd
```

示例：

```go
fmt.Println(
    client.HSet(ctx, "students", "42001", "马士兵", "42002", "GoLang", "42003", "Redis").
    	Result()) // 3 <nil>

fmt.Println(
    client.HGetAll(ctx, "students").
   		Result()) // map[42001:马士兵 42002:GoLang 42003:Redis] <nil>

fmt.Println(
    client.HLen(ctx, "students").
    	Result()) // 3 <nil>

fmt.Println(
    client.HRandField(ctx, "students", 2).
    	Result()) // [42003 42001] <nil>, [42002 42003] <nil>
fmt.Println(
    client.HRandField(ctx, "students", -2).
    	Result()) // [42003 42001] <nil>, [42003 42003] <nil>
fmt.Println(
    client.HRandFieldWithValues(ctx, "students", 2).
    	Result()) // [{42002 GoLang} {42003 Redis}] <nil>

fmt.Println(
    client.HExists(ctx, "students", "42003").
    	Result()) // true <nil>
fmt.Println(
    client.HExists(ctx, "students", "42006").
   		Result()) // false <nil>

fmt.Println(
    client.HDel(ctx, "students", "42003").
    	Result()) // 1 <nil>
fmt.Println(
    client.HExists(ctx, "students", "42003").
    	Result()) // true <nil>

fmt.Println(
    client.Do(ctx, "HSTRLEN", "students", "42001").
    	Result()) // 9 <nil>


```

HScan 操作用于增量遍历Hash中的字段，与其他集合类型的用法一致，参见 SCAN 命令。

# Bitmap位图操作

## Bitmap 介绍

Redis位图是字符串数据类型的扩展，可以将字符串视为位向量。您还可以对一个或多个字符串执行按位操作。位图用例的一些示例包括：

- 集合的成员对应于整数0-N的情况下的有效集合表示。
- 对象权限，其中每个位代表一个特定的权限，类似于文件系统存储权限的方式。

性能：

- SETBIT 和 GETBIT 时间复杂度是 O(1)
- BITOP 时间复杂度是 O(n)

## 操作方法说明

### 设置位

```go
// 设置特定位
func (c Client) SetBit(ctx context.Context, key string, offset int64, value int) *IntCmd
```

### 获取位

```go
// 获取特定位
func (c Client) GetBit(ctx context.Context, key string, offset int64) *IntCmd
// 统计位数量
func (c Client) BitCount(ctx context.Context, key string, bitCount *BitCount) *IntCmd
// 获取特定位的位置
func (c Client) BitPos(ctx context.Context, key string, bit int64, pos ...int64) *IntCmd
```

### 位操作

```go
// 位与
func (c Client) BitOpAnd(ctx context.Context, destKey string, keys ...string) *IntCmd
// 位非
func (c Client) BitOpNot(ctx context.Context, destKey string, key string) *IntCmd
// 位或
func (c Client) BitOpOr(ctx context.Context, destKey string, keys ...string) *IntCmd
// 位异或
func (c Client) BitOpXor(ctx context.Context, destKey string, keys ...string) *IntCmd
```

```go
// 
func (c Client) BitField(ctx context.Context, key string, args ...interface{}) *IntSliceCmd
```

# HyperLogLog 草图操作

## HyperLogLog 介绍

HyperLogLog 是一种用于估计集合基数的数据结构。作为一种概率数据结构，HyperLogLog 以完美的准确性换取高效的空间利用。
Redis HyperLogLog 实现最多使用 12 KB，并提供 0.81% 的标准错误。

限制：

草图结构最多可以估计 2^64个成员。

性能：

写PFADD和读PFCOUNT操作会在常量时间和空间内完成。

合并操作时间复杂度：O(n)

## 操作方法介绍

### 添加成员

```go
// 添加成员
func (c Client) PFAdd(ctx context.Context, key string, els ...interface{}) *IntCmd
```

### 统计成员数量

```go
// count
func (c Client) PFCount(ctx context.Context, keys ...string) *IntCmd
```

### 合并

```go
// 合并多个 HLL 到一个
func (c Client) PFMerge(ctx context.Context, dest string, keys ...string) *StatusCmd
```

# Geospatial Indices地理空间索引操作

## Geospatial 说明

Redis 地理空间索引允许存储坐标并搜索它们。此数据结构对于查找给定半径或边界框内的附近点非常有用。

地理索引结构是一个排序集合。

每个地理位置项由：经度、纬度、名称(longitude, latitude, name)组成。

## 操作方法说明

### 添加地点成员

```go
// 添加
func (c Client) GeoAdd(ctx context.Context, key string, geoLocation ...*GeoLocation) *IntCmd
```

### 获取地点成员信息

```go
// 获取两个地点间的距离
func (c Client) GeoDist(ctx context.Context, key string, member1, member2, unit string) *FloatCmd
// 获取地点成员，返回Hash值
func (c Client) GeoHash(ctx context.Context, key string, members ...string) *StringSliceCmd
// 获取地点成员的坐标
func (c Client) GeoPos(ctx context.Context, key string, members ...string) *GeoPosCmd
```

### 搜索地点成员

```go
// 搜索地点
func (c Client) GeoSearch(ctx context.Context, key string, q *GeoSearchQuery) *StringSliceCmd
// 搜索地点，同时返回位置、距离、Hash等信息
func (c Client) GeoSearchLocation(ctx context.Context, key string, q *GeoSearchLocationQuery) *GeoSearchLocationCmd
// 搜索地点并存储
func (c Client) GeoSearchStore(ctx context.Context, key, store string, q *GeoSearchStoreQuery) *IntCmd
```

# Stream 流操作

## Stream 介绍

Redis 流是一种数据结构，其作用类似于附加日志。您可以使用流实时记录和同时联合事件。 Redis 流用例示例包括：

- 事件溯源（例如，跟踪用户操作、点击等）
- 传感器监控（例如，现场设备的读数）
- 通知（例如，将每个用户的通知记录存储在单独的流中）

Redis 为每个流条目生成一个唯一的 ID。您可以使用这些 ID 稍后检索它们的关联条目，或者读取和处理流中的所有后续条目。

Redis 流支持多种修剪策略（以防止流无限制地增长）和不止一种消费策略（请参阅 XREAD、XREADGROUP 和 XRANGE）。

## 操作方法说明

```
func (c Client) XAck(ctx context.Context, stream, group string, ids ...string) *IntCmd
func (c Client) XAdd(ctx context.Context, a *XAddArgs) *StringCmd
func (c Client) XAutoClaim(ctx context.Context, a *XAutoClaimArgs) *XAutoClaimCmd
func (c Client) XAutoClaimJustID(ctx context.Context, a *XAutoClaimArgs) *XAutoClaimJustIDCmd
func (c Client) XClaim(ctx context.Context, a *XClaimArgs) *XMessageSliceCmd
func (c Client) XClaimJustID(ctx context.Context, a *XClaimArgs) *StringSliceCmd
func (c Client) XDel(ctx context.Context, stream string, ids ...string) *IntCmd
func (c Client) XGroupCreate(ctx context.Context, stream, group, start string) *StatusCmd
func (c Client) XGroupCreateConsumer(ctx context.Context, stream, group, consumer string) *IntCmd
func (c Client) XGroupCreateMkStream(ctx context.Context, stream, group, start string) *StatusCmd
func (c Client) XGroupDelConsumer(ctx context.Context, stream, group, consumer string) *IntCmd
func (c Client) XGroupDestroy(ctx context.Context, stream, group string) *IntCmd
func (c Client) XGroupSetID(ctx context.Context, stream, group, start string) *StatusCmd
func (c Client) XInfoConsumers(ctx context.Context, key string, group string) *XInfoConsumersCmd
func (c Client) XInfoGroups(ctx context.Context, key string) *XInfoGroupsCmd
func (c Client) XInfoStream(ctx context.Context, key string) *XInfoStreamCmd
func (c Client) XInfoStreamFull(ctx context.Context, key string, count int) *XInfoStreamFullCmd
func (c Client) XLen(ctx context.Context, stream string) *IntCmd
func (c Client) XPending(ctx context.Context, stream, group string) *XPendingCmd
func (c Client) XPendingExt(ctx context.Context, a *XPendingExtArgs) *XPendingExtCmd
func (c Client) XRange(ctx context.Context, stream, start, stop string) *XMessageSliceCmd
func (c Client) XRangeN(ctx context.Context, stream, start, stop string, count int64) *XMessageSliceCmd
func (c Client) XRead(ctx context.Context, a *XReadArgs) *XStreamSliceCmd
func (c Client) XReadGroup(ctx context.Context, a *XReadGroupArgs) *XStreamSliceCmd
func (c Client) XReadStreams(ctx context.Context, streams ...string) *XStreamSliceCmd
func (c Client) XRevRange(ctx context.Context, stream, start, stop string) *XMessageSliceCmd
func (c Client) XRevRangeN(ctx context.Context, stream, start, stop string, count int64) *XMessageSliceCmd
func (c Client) XTrimMaxLen(ctx context.Context, key string, maxLen int64) *IntCmd
func (c Client) XTrimMaxLenApprox(ctx context.Context, key string, maxLen, limit int64) *IntCmd
func (c Client) XTrimMinID(ctx context.Context, key string, minID string) *IntCmd
func (c Client) XTrimMinIDApprox(ctx context.Context, key string, minID string, limit int64) *IntCmd
```

# Transaction 事务操作

## Redis 事务介绍

Redis 事务允许在单个步骤中执行一组命令，它们以命令 MULTI、EXEC、DISCARD 和 WATCH 为中心。 Redis 事务有两个重要的保证：

事务中的所有命令都被序列化并按顺序执行。另一个客户端发送的请求永远不会在 Redis 事务执行过程中得到处理。这保证了命令作为单个独立操作执行。

EXEC 命令触发事务中所有命令的执行，因此如果客户端在调用 EXEC 命令之前在事务上下文中丢失与服务器的连接，则不会执行任何操作，相反，如果调用 EXEC 命令，执行所有操作。使用 append-only file 时，Redis 确保使用单个 write(2) 系统调用将事务写入磁盘。但是，如果 Redis 服务器崩溃或被系统管理员以某种硬方式杀死，则可能只注册了部分操作。 Redis 会在重启时检测到这种情况，并会报错退出。使用 redis-check-aof 工具可以修复将删除部分事务的仅附加文件，以便服务器可以重新启动。

在交易过程中，可能会遇到两种命令错误：

- 一个命令可能无法入队，所以在调用EXEC之前可能会出错。例如，该命令可能在语法上是错误的（参数数量错误，命令名称错误，...），或者可能存在一些临界条件，如内存不足条件（如果服务器配置为使用 maxmemory 具有内存限制指示）
- 调用 EXEC 后命令可能会失败，例如，因为我们对具有错误值的键执行操作（例如对字符串值调用列表操作）

# 监控

# 深入

## 连接池的设计

## 执行命令过程
