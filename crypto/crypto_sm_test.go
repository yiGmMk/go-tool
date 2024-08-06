package crypto

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/pkg/errors"
	"github.com/tjfoc/gmsm/sm2"
	"github.com/tjfoc/gmsm/sm4"
)

func checkKeyValid(pub, pri string) error {
	kPub, err := StringToPublicKey(pub)
	if err != nil {
		return err
	}

	kPri, err := StringToPrivateKey(pri, kPub)
	msg := "{\"name\":\"hexo\"}"
	encryText, err := kPub.EncryptAsn1([]byte(msg), rand.Reader)
	if err != nil {
		return errors.Wrapf(err, "加密失败")
	}
	decryText, err := kPri.DecryptAsn1(encryText)
	if err != nil {
		return errors.Wrapf(err, "解密失败")
	}
	if string(decryText) != string(msg) {
		return errors.Errorf("解密失败,解密内容和原始数据不一致")
	}
	fmt.Printf("验证通过,msg:%s,encry:%s\n", msg, hex.EncodeToString(encryText))
	return nil
}

func TestSm2(t *testing.T) {
	t.Run("gen", func(t *testing.T) {
		// 生成 SM2 密钥对
		privKey, err := sm2.GenerateKey(rand.Reader)
		if err != nil {
			fmt.Println("Error generating SM2 key pair:", err)
			return
		}
		pri, pub := Key2String(privKey)
		if err := checkKeyValid(pub, pri); err != nil {
			t.Error(err)
		}
		fmt.Printf("SM2 Private Key: %s\n", pri)
		fmt.Printf("SM2 Public  Key: %s\n", pub)
	})

	t.Run("load key", func(t *testing.T) {
		pub := "04b7afdce28e6f71c27ebb6bf366999494c59a8c0aa66a5c4f8ee2b500bc864ec6de292df1a94160b3a3f96c272576ce72a7b17b9d19055ffa2768edbadbe54800"
		pri := "f8adca86d964a5ff6d9211fc452d79d397153aec6386973751224f0bf7334dcd"
		if err := checkKeyValid(pub, pri); err != nil {
			t.Error(err)
		}
	})
}

const (
	// api秘钥
	apiPubKey = "04b7afdce28e6f71c27ebb6bf366999494c59a8c0aa66a5c4f8ee2b500bc864ec6de292df1a94160b3a3f96c272576ce72a7b17b9d19055ffa2768edbadbe54800"
	apiPriKey = "f8adca86d964a5ff6d9211fc452d79d397153aec6386973751224f0bf7334dcd"
	// 调用方秘钥
	myPubKey, myPriKey = "04067c410f396abd978381ca8a89f49659bbdd7c995ead8b2cd85d364b40d9a0cf92cdf433600db5da47f658abbe83dd9809d40b87544c3817a2f4500a585083b8",
		"9a9d44dcbbb490d41d051e29ea619ad6c7f310d5ecd705a555c647186a11439e"
)

func TestHutoolRequest(t *testing.T) {
	// 调用接口,只有服务端公钥和自己的密钥对
	apiPubKey := apiPubKey
	msg := "{\"name\":\"hexo\"}"

	// 1.公钥加密sm4 key, sm4 key加密数据
	sm4Key, sm4HexKey, err := GenerateSM4Key()
	if err != nil {
		t.Fatal(err, "sm4 key生成失败")
	}
	encryptData, err := sm4.Sm4Ecb(sm4Key, []byte(msg), true)
	if err != nil {
		t.Fatal(err, "sm4加密失败")
	}
	encryptDataHex := hex.EncodeToString(encryptData)

	sm2Pub, err := StringToPublicKey(apiPubKey)
	if err != nil {
		t.Fatal(err, "获取sm2公钥失败")
	}
	encryptKey, err := sm2.Encrypt(sm2Pub, []byte(sm4HexKey), rand.Reader, sm2.C1C3C2)
	if err != nil {
		t.Fatal(err, "sm2加密失败")
	}
	encryptKeyHex := hex.EncodeToString(encryptKey)
	// 2.调接口
	fmt.Printf("加密数据hex: %s\n加密key hex: %s\n", encryptDataHex, encryptKeyHex)
}
