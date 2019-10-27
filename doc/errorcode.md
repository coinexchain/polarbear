## sdk错误码

由于该sdk支持跨语言和跨平台，所以所有接口返回值均为字符串，返回值数量为1

#### 格式

为了区分正确的返回值和错误信息，sdk的错误格式如下所示：

`错误统一前缀:具体的api接口:接口内部错误详情`

**错误码前缀**：`POLARBEARError`

**Example**：`POLARBEARError:Add key:Cannot overwrite key bob`  

该错误表示导入私钥时发现本地已有名为bob的私钥对。