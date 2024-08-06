package crypto

import (
	"encoding/hex"
	"math/big"

	"github.com/pkg/errors"
	"github.com/tjfoc/gmsm/sm2"
)

// sm2支持C1C2C3/C1C3C2 2种模式,Hutool 默认使用C1C3C2

// private  public
func Key2String(k *sm2.PrivateKey) (string, string) {
	kPri, kPub := k, k.PublicKey
	return hex.EncodeToString(kPri.D.Bytes()),
		"04" + hex.EncodeToString(kPub.X.Bytes()) + hex.EncodeToString(kPub.Y.Bytes())
}

// StringToPrivateKey 将十六进制编码的私钥字符串转换为 sm2.PrivateKey 对象
func StringToPrivateKey(privateKeyStr string, publicKey *sm2.PublicKey) (*sm2.PrivateKey, error) {
	// 解码十六进制字符串
	privateKeyBytes, err := hex.DecodeString(privateKeyStr)
	if err != nil {
		return nil, err
	}

	// 将字节切片转换为大整数
	d := new(big.Int).SetBytes(privateKeyBytes)

	// 验证私钥是否在曲线的有效范围内
	if d.Cmp(big.NewInt(0)) == 0 || d.Cmp(publicKey.Curve.Params().N) >= 0 {
		return nil, errors.New("invalid private key")
	}

	// 验证公钥是否与私钥匹配（可选）
	x, y := publicKey.Curve.ScalarBaseMult(privateKeyBytes)
	if x.Cmp(publicKey.X) != 0 || y.Cmp(publicKey.Y) != 0 {
		return nil, errors.New("public key does not match private key")
	}

	// 创建 sm2.PrivateKey 对象
	privateKey := &sm2.PrivateKey{
		PublicKey: *publicKey,
		D:         d,
	}

	return privateKey, nil
}

func StringToPublicKey(publicKeyStr string) (*sm2.PublicKey, error) {
	publicKeyBytes, err := hex.DecodeString(publicKeyStr)
	if err != nil {
		return nil, err
	}

	// 检查公钥是否以未压缩格式（0x04）开头
	if len(publicKeyBytes) == 0 || publicKeyBytes[0] != 0x04 {
		return nil, errors.New("invalid public key format")
	}

	// 提取 x 和 y 坐标字节切片
	curve := sm2.P256Sm2()
	byteLen := (curve.Params().BitSize + 7) / 8
	if len(publicKeyBytes) != 1+2*byteLen {
		return nil, errors.New("invalid public key length")
	}

	xBytes := publicKeyBytes[1 : byteLen+1]
	yBytes := publicKeyBytes[byteLen+1 : 2*byteLen+1]

	// 将字节切片转换为大整数
	x := new(big.Int).SetBytes(xBytes)
	y := new(big.Int).SetBytes(yBytes)

	// 创建 sm2.PublicKey 对象
	publicKey := &sm2.PublicKey{
		Curve: curve,
		X:     x,
		Y:     y,
	}

	// 验证公钥是否在曲线上
	if !curve.IsOnCurve(x, y) {
		return nil, errors.New("public key is not on the curve")
	}

	return publicKey, nil
}
