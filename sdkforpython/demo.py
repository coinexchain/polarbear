import ctypes


def demo():
    key_name = "alice".encode("utf-8")
    password = "12345678".encode("utf-8")
    bip39_password = "11111111".encode("utf-8")

    lib = ctypes.CDLL('./wallet.so')

    # sdk init
    lib.BearInit('tmp'.encode("utf-8"))

    # create key
    create_key = lib.CreateKey
    create_key.restype = ctypes.c_char_p
    key = create_key(key_name, password, bip39_password, 0, 0)
    assert b'coinex' in key

    # list keys
    list_keys = lib.ListKeys
    list_keys.restype = ctypes.c_char_p
    keys = list_keys()
    assert key_name in keys

    # get pubkey
    get_pubkey = lib.GetPubKey
    get_pubkey.restype = ctypes.c_char_p
    pubkey = get_pubkey(key_name)
    assert b'coinexpub' in pubkey

    # sign
    sign = lib.Sign
    sign.restype = ctypes.c_char_p
    signature = sign(key_name, password, "hello, that polar bear")
    assert b"signature" in signature

    # delete key
    delete = lib.DeleteKey
    delete.restype = ctypes.c_char_p
    res = delete(key_name, password)
    assert res == b''

    # delete success
    keys = list_keys()
    assert keys == b'[]'
    print("The polar bear live")


if __name__ == "__main__":
    demo()
