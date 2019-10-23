## 使用手册

#### 钱包创建

用户可以指定任意目录，调用`func BearInit(root string)`将钱包创建到该指定目录下，用户创建或导入的秘钥都存储于该目录下。

#### 秘钥生成

用户可以通过`func CreateKey(name, password, bip39Passphrase string, account, index int) string`创建新的秘钥对，这个秘钥对后续可以用来签名交易。

秘钥对支持删除，恢复，导入，导出，查询等操作，具体参考[API文档](https://github.com/coinexchain/polarbear/blob/master/doc/api.md)

#### 秘钥导入

如果用户之前已经通过cetcli生成了秘钥，那么可以通过`cetcli keys export <keyname>`将秘钥导出，如下所示

```
$cetcli keys export bob
Enter passphrase to decrypt your key:
Enter passphrase to encrypt the exported key:
-----BEGIN TENDERMINT PRIVATE KEY-----
salt: 94D311953AC8C1CB5CEFA477C1B446D8
kdf: bcrypt

A93g12JrGEpsbwrTZ9x7UUjWFGzexu3X7SA8bl5U6IPTVDCBKJkHgRyWS2frKFtE
XmJL2BoGvDdBWUu8VULicnvNSVRyaQpgfsqIxRk=
=57oH
-----END TENDERMINT PRIVATE KEY-----
```

然后将命令行生成的字符串传入到`func AddKey(name, armor string) string`的`armor`参数。

秘钥导入完成

#### 秘钥导出

用户通过调用`func ExportKey(name string) string`来导出指定的秘钥，导出字符串为上面**秘钥导入**部分讲到的armor字符串，如下所示：

```
-----BEGIN TENDERMINT PRIVATE KEY-----
salt: 94D311953AC8C1CB5CEFA477C1B446D8
kdf: bcrypt

A93g12JrGEpsbwrTZ9x7UUjWFGzexu3X7SA8bl5U6IPTVDCBKJkHgRyWS2frKFtE
XmJL2BoGvDdBWUu8VULicnvNSVRyaQpgfsqIxRk=
=57oH
-----END TENDERMINT PRIVATE KEY-----
```

#### 签名

假设用户向发送一笔转账到节点，那么用户首先查询自己的账户信息，接口如下：

```
GET /auth/accounts/{address}
```

从该接口的响应获取账户的accountNum和sequence两个字断。

接下来，用户发送一笔转账交易到节点，通过调用如下接口：

```
POST /bank/accounts/{address}/transfers
```

将该接口的返回值作为参数tx传入到sdk的`func SignAndBuildBroadcast(name, password, tx, chainId, mode string, accountNum, sequence uint64) string`中，api返回值可以直接用于构造`广播交易`rest接口的请求体。广播交易接口如下：

```
POST /txs
```



#### 其他

详见[API接口](https://github.com/coinexchain/polarbear/blob/master/doc/api.md)

