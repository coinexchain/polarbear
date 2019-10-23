# API接口

polarbear sdk提供下列接口供离线钱包使用

#### sdk初始化

```go
func BearInit(root string)
```

BearInit初始化sdk并在用户指定的路径下创建钱包目录

#### 创建秘钥对和地址

```go
func CreateKey(name, password, bip39Passphrase string, account, index int) string
```

CreateKey创建一个新的秘钥对，参数分别是秘钥的名字（不允许重复），加密存储秘钥所需要的密码，生成助记词所需要的盐值，BIP44 path中的account和address_index。

返回值格式为generated address string + mnemonic，即新生成地址的bench32表示和助记词，中间用+连接。

助记词只有创建时产生，不在sdk中保存。

#### 删除秘钥对

```go
func DeleteKey(name, password string) string
```

DeleteKey删除本地保存的秘钥对，需要提供保存私钥使用的密码。

#### 恢复秘钥对

```go
func RecoverKey(name, mnemonic, password, bip39Passphrase string, account, index int) string
```

通过提供助记词和生成该助记词时输入的盐值，以及account和index来恢复秘钥对，提供新的password来加密恢复出来的秘钥对，这个password可以和之前创建秘钥对时提供的密码不一致。

#### 导入秘钥对

```go
func AddKey(name, armor string) string
```

从外部导入秘钥对到本地。armor可以由本sdk的秘钥对导出功能生成，或者通过cetcli keys export *keyname* 导出 

#### 导出秘钥对

```go
func ExportKey(name string) string
```

导出指定秘钥对，返回值为armor编码的秘钥对信息，该字符串可以用于秘钥对导入，例如

```
-----BEGIN TENDERMINT PRIVATE KEY-----
kdf: bcrypt
salt: 1190F48AAAC67A595040129049BDE6C7

JiooF+ymHFABgUV3w1y1LkG8OmeMZv7igpKi5dRALOQEhpocz6L2mOD6/b3nMj9T
VtKPPbga5NUQ2F2JM1WNZOS+XKIXzGCHhoRiofs=
=Elr7
-----END TENDERMINT PRIVATE KEY-----
```

#### 列出全部秘钥对信息

```go
func ListKeys() string
```

列出本地全部秘钥对信息，格式为json字符串，发生错误时返回空字符串

```
[{"name":"default","type":"local","address":"coinex1njh7ahpvrvq2d2vnlcuhzlnj0whe2kzgnwe8c6","pubkey":"coinexpub1addwnpepqt80h38na9rfl8t763cpg9kqdwz3kujr6y4adj985k0meyuv7cng7zvfy0a"}]
```

#### 重置秘钥对存储密码

```go
func ResetPassword(name, password, newPassword string) string
```

需要旧的密码和新的密码，成功返回空字符串，否则返回错误内容

#### 获取指定秘钥对的地址

```go
func GetAddress(name string) string
```

name为秘钥名，返回值为bench32编码后的地址

#### 获取指定秘钥对的公钥

```go
func GetPubKey(name string) string
```

name为秘钥名，返回值为bench32编码后的公钥

#### 获取签名者

```
func GetSigner(signerInfo string) string
```

返回值为签名者的秘钥的名字

#### 签名

```go
func Sign(name, password, tx string) string 
```

name是key的名字，password是存储私钥的密码，tx是如下结构的json序列化后的字符串

```
type StdSignDoc struct {
	AccountNumber uint64            `json:"account_number" yaml:"account_number"`
	ChainID       string            `json:"chain_id" yaml:"chain_id"`
	Fee           json.RawMessage   `json:"fee" yaml:"fee"`
	Memo          string            `json:"memo" yaml:"memo"`
	Msgs          []json.RawMessage `json:"msgs" yaml:"msgs"`
	Sequence      uint64            `json:"sequence" yaml:"sequence"`
}
```

返回如下格式的json字符串

```
{"pub_key":{"type":"tendermint/PubKeySecp256k1","value":"A/GNrOR+zS7bvomsMG+BEIiGB8H+EvIDWGqfMOO5GVVV"},"signature":"pxUS22oste4S3Bmix5LDYgns27Lf5NxZH1duGdT+Yu5Fvz2kYOieeb5j/nxvjdM1TQ5wQUPo47vnWW+1fnjuiQ=="}
```

#### 签名交易

```
func SignStdTx(name, password, tx, chainId string, accountNum, sequence uint64) string
```

name是秘钥名字，password是秘钥的存储密码，chainId是链的id，accountNum和sequence是账户相关字段，可以通过下面接口从节点获取

```
GET /auth/accounts/{address}
```

tx是下面结构的json序列化字符串

```
type StdTx struct {
	Msgs       []sdk.Msg      `json:"msg" yaml:"msg"`
	Fee        StdFee         `json:"fee" yaml:"fee"`
	Signatures []StdSignature `json:"signatures" yaml:"signatures"`
	Memo       string         `json:"memo" yaml:"memo"`
}
```

该字符串可以通过节点提供的各个交易类型的rest接口返回的响应体中获取，例如转账交易可以通过下面的接口获取：

```
POST /bank/accounts/{address}/transfers
```

该接口返回值同上面的Sign方法

#### 签名交易并返回可直接用于广播的字符串

```
func SignAndBuildBroadcast(name, password, tx, chainId, mode string, accountNum, sequence uint64) string
```

参数字段同上，mode为节点rest接口：POST /txs中的mode字段取值，如block，sync，async。

返回值为下面结构的json序列化字符串，可直接用于POST /txs接口调用

```
type BroadcastReq struct {
	Tx   types.StdTx `json:"tx" yaml:"tx"`
	Mode string      `json:"mode" yaml:"mode"`
}
```

**example：**

输入的tx：

```
{"type":"auth/StdTx","value":{"msg":[{"type":"bankx/MsgSetMemoRequired","value":{"address":"coinex10kvrwz96tw5f6r2mg6l60qdqjc3dqr0kp5pn2l","required":true}}],"fee":{"amount":[{"denom":"cet","amount":"50"}],"gas":"200000"},"signatures":null,"memo":"Sent with example memo"}}
```

返回值：

```
{"tx":{"msg":[{"type":"bankx/MsgSetMemoRequired","value":{"address":"coinex10kvrwz96tw5f6r2mg6l60qdqjc3dqr0kp5pn2l","required":true}}],"fee":{"amount":[{"denom":"cet","amount":"50"}],"gas":"200000"},"signatures":[{"pub_key":{"type":"tendermint/PubKeySecp256k1","value":"A7xquC4+chG2jNu97GepRh/XQZqdRWkszaLt5OhXPYZ7"},"signature":"v9bAdhkVhBqfm7dKcH1Pteza3HjAcZZ7qzBKoaIOc81g2ngnxEt1G03G9X9I6zDpqBVAfWQK3+UoWLNzSB/i3A=="}],"memo":"Sent with example memo"},"mode":"sync"}
```