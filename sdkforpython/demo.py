import ctypes, shutil


def demo():
    key_name = "alice".encode("utf-8")
    password = "12345678".encode("utf-8")
    bip39_password = "11111111".encode("utf-8")

    lib = ctypes.CDLL('./wallet_mac.so')

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

    #get address from WIF
    wif = b'5HueCGU8rMjxEXxiPuD5BDku4MkFqeZyd4dZ1jvhTVqvbTLvyTJ'
    get_address_from_wif = lib.GetAddressFromWIF
    get_address_from_wif.restype = ctypes.c_char_p
    address_wif = get_address_from_wif(wif)
    assert address_wif == b'coinex1my63mjadtw8nhzl69ukdepwzsyvv4yexfas4jz'

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

    shutil.rmtree('./tmp')

if __name__ == "__main__":
    demo()
