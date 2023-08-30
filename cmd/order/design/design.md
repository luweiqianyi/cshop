# 订单
## 业务逻辑
* 客户端：用户查看商品后选择某个商品准备进行下单->填写订单信息->提交订单。
* 服务端：根据用户上传的提交订单请求的信息，调用相关服务完成订单的创建，以及写入数据库。订单创建后即可以写入数据库，其支付状态可以是未支付。

### 用户选择阶段
用户看中某个商品后，在提交订单信息之前会选定商品的数量，商品的属性，选择支付方式，配送方式。
> 商品属性是由商家在录入商品信息的时候指定，前端展示的商品属性都是从后端读取后才在前端显示的，前端在某个属性上设置为灰态表示当前该属性的商品没有货。

商品属性如下所示：
* 商品名称(唯一代表某类商品，比如iphone13 pro)
* 商品数量
* 商品属性1(比如 颜色)
* 商品属性2(比如 256G)
* 商品属性3(比如 版本)
> * 等到用户选好属性后，前端会把选择信息告知后端，后端给出一个商品总价返回给前端进行展现。一般商品的属性不要太多，因为商品的最终种类与商品属性的关系是指数级别，如一种商品有3属性，那么该对于该商品就可能有8种选择。
> * 一般对于超市中的商品来说，基本上提供一种属性应该可以满足绝大多数情况
#### 商品表
根据以上逻辑，可以简单推测出某类商品表(如手机类)的基本结构。
`phone_table`
|id|productName|color|storage|version|
|----|----|----|----|----|
|1|iphone13 pro|远峰蓝|256G|港版|
|2|iphone13 proMax|远峰蓝|256G|港版|
|3|Huawei Mate60|雅川清|512G|国内版|

那么查询某类颜色是否有相关机型(在前端就表现为有没有货，影响按钮是不是置灰，允不允许用户下单)：
```sql
select count(*) from `phone_table` where productName='iphone13 pro' and color = '远峰蓝'
```
> * 所以说对于那种大类杂类比较多的商品交易市场(比如京东、淘宝、天猫等)来说，涉及的表可能比较多，且各类商品的表结构可能不相同
> * 对于那种专注于某一类商品的商品交易市场来说(比如说懂车帝、汽车之家等)，表结构就比较固定了

### 用户提交订单
用户经过选择阶段之后，会向服务端发送提交订单请求。本次提交订单请求是以http的方式向服务端order-api服务发送提交信息，order-api服务，将本次请求发送给order-rpc服务来完成本次下单过程的处理。

#### api请求设计
前端以`post`方式向后端发送请求。参数设计如下所示：
| key | value | 数据类型 |描述|
|------|------|------|------|
| token | xasadkadhka85511 | string| 用户token，唯一标识某用户 |
| productName | iphone13 pro | string | 商品名称 |
| productNumber | 2 | int | 商品数量 |
| propertity1 | 1 | int | 商品属性1 |
| propertity2 | 2 | int | 商品属性2 |
| propertity3 | 0 | int | 商品属性3 |
| payMethod | 1 | int | 支付方式 |
| deliveryMethod | 1 | int | 配送方式 |
| expectedDeliveryTime| 2023-01-02,09:00-21:00 | string | 期望配送时间 |
> 价格的计算不放在前端计算，由后端的其他服务(可能是计算服务，可能是商品服务，可能是订单服务,暂时先不决定)来完成计算。这里只负责将用户的订购详情上传到服务端。
 
### 服务端设计
#### 订单属性
以下表的属性设计用于在用户对某个订单进行查询时，展现给用户的信息。
| 属性key |数据类型| 属性value |
|------|------|------|
| 订单号 | string | 279675205162 |
| 下单时间 | string | 2023-01-01 11:02:00 |
| 商品详情 | string | iphone13 pro, x2, 远峰蓝,256G, 大陆版 |
| 支付方式 | string | 支付宝支付 |
| 支付状态 | string | 已支付/已退款/待支付/已取消|
| 配送方式 | string | 京东配送/自建物流/ |
| 期望配送时间 | string | 2023-01-02,09:00-21:00 |
| 商品总价 | float64 | 17998 |
| 运费| float64 | 0 |
| 会员减免 | float64 | 300 |
| 实付款 | float64 | 17698 |
> * 商品详情是用户提交订单的请求到达服务端后，服务端调用相关模块，拼装组装成具体的字符串。
> * 商品总价的计算是用户提交订单的请求到达服务端后，由服务端根据用户请求内容调用服务端模块完成商品总价的计算。



采用“商品详情”的字段而不是下面的属性值作为订单表的属性是因为数据库表结构应该是固定的，不应该随着商品类别的不同而改变表结构
| 属性key |数据类型| 属性value |
|------|------|------|
| 商品名称 | string | iphone13 pro |
| 商品数量 | int | 2 |
| 颜色 | string | 远峰蓝 |
| 存储 | string | 256G |
| 版本 | string | 港版/美版/大陆版 |

#### 订单和订单创建者的关系
* 基本关系： 订单创建者可以创建多个订单，一个订单只能属于一个订单创建者
* 数据库表关系
    * 用户-订单表：哪个用户创建了哪个订单
    * 订单详细表

#### 数据库表设计
数据库表的设计要不要考虑创建时间和修改时间。

##### 用户订单表
userID可以取accountID。
| userID | orderID |
|----|----|
| 15700080001 | orderID-001|
| 15700080001 | orderID-002|
| 15700080002 | orderID-003|
> 如上所示：用户“15700080001”创建了两个订单，分别是订单“orderID-001”，“orderID-002”；用户“15700080002”创建了一个订单，为“orderID-003”

##### 订单详细表
|orderID|orderCreateTime|productDetail|payMethod|payStatus|deliveryMethod|expectedDeliveryTime|totalPrice|freight|memberDiscount|finalPrice|creatorID|
|----|----|----|----|----|----|----|----|----|----|----|----|
| orderID-001 | 2023-01-01 11:02:00 | iphone13 pro, x2, 远峰蓝,256G, 大陆版 |支付宝支付 | 已支付 | 京东配送 | 2023-01-02,09:00-21:00 | 17998 | 0 | 300 | 17698 | 15700080001 |
| orderID-002 | 2023-02-01 15:05:30 | iphone13 pro, x2, 远峰蓝,256G, 大陆版 |微信支付 | 待支付 |京东配送 | 2023-02-01,09:00-21:00 | 17998 | 0 | 300 | 17698 | 15700080001 |

### 订单表管理
对于小型商店来说，我并不关注要保存用户的历史订单信息，我只为其保存3个月或者6个月或者1年，待过期后全部进行删除。那这样的话存在的问题就是全表扫描，删除过期记录。
如果按照时间进行分表，那么查询用户的所有订单数据就涉及到多分表查询(多分表查询还要区别是在同主机还是在不同主机两种情况进行讨论)。