# polarbear
**polarbear is a cold wallet sdk，it is used for secret key generation, management and transaction signature**

[**API manual**](https://github.com/coinexchain/polarbear/blob/master/doc/api.md)

- go语言可以直接调用package wallet或者keybase下的API

- python用户可以通过下面的命令将sdk编译为动态库

```
go build -o wallet.so -buildmode=c-shared sdkforpython/wallet_for_python.go
```

通过python的ctypes库使用该sdk，参考[demo](https://github.com/coinexchain/polarbear/blob/master/sdkforpython/demo.py)

- android用户可以通过下面的命令生成aar包

```
GOOS=android GOARCH=arm CGO_CFLAGS=“-I/Users/helldealer/Library/Android/sdk/ndk-bundle/sysroot/usr/include -I/Users/helldealer/Library/Android/sdk/ndk-bundle/sysroot/usr/include/arm-linux-androideabi” CGO_LDFLAGS=“-L/Users/helldealer/Library/Android/sdk/ndk-bundle/sysroot/usr/lib” gomobile bind -target=android/arm  github.com/coinexchain/polarbear
```

[**使用手册**](https://github.com/coinexchain/polarbear/blob/master/doc/manual.md)

