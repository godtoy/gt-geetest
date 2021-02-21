package geetest

import (
    "fmt"
    "math/rand"
    "regexp"
    "time"
    "zhaojunlike/common/encrypt"
)

const rsaKey = "00C1E3934D1614465B33053E7F48EE4EC87B14B95EF88947713D25EECBFF7E74C7977D02DC1D9451F79DD5D1C10C29ACB6A9B4D6FB7D0A0279B6719E1772565F09AF627715919221AEF91899CAE08C0D686D748B20A3603BE2318CA6BC2B59706592A9219D0BF05C9F65023A21D2330807252AE0066D59CEEFA5F2748EA80BAB81"

var geeRSA = encrypt.NewRsa()

func init() {
    _ = geeRSA.SetPublicKey(rsaKey, 10001)
}

// 解析jsonp
func JsonpParse(name string, respStr string) string {
    reg, _ := regexp.Compile(fmt.Sprintf(`^%v\((.*)\)$`, name))
    return reg.ReplaceAllString(respStr, "$1")
}
func RandomBytes(size int) (resp []byte) {
    rand.Seed(time.Now().UnixNano())
    blk := make([]byte, size)
    _, err := rand.Read(blk)
    if err != nil {
        panic(err)
    }
    return blk
}

func pad(s string, c int) string {
    str := ""
    for i := 0; i < c; i++ {
        str += s
    }
    return str
}
