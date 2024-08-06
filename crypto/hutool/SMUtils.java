package sm;

import cn.hutool.core.util.HexUtil;
import cn.hutool.crypto.SmUtil;
import cn.hutool.crypto.asymmetric.KeyType;
import cn.hutool.crypto.asymmetric.SM2;
import cn.hutool.crypto.symmetric.SM4;
import org.bouncycastle.jcajce.provider.asymmetric.ec.BCECPrivateKey;
import org.bouncycastle.jcajce.provider.asymmetric.ec.BCECPublicKey;
import org.bouncycastle.jce.provider.BouncyCastleProvider;
import org.bouncycastle.math.ec.ECPoint;
import org.bouncycastle.pqc.math.linearalgebra.ByteUtils;

import javax.crypto.KeyGenerator;
import java.math.BigInteger;
import java.security.SecureRandom;
import java.security.Security;

public class SMUtils {
    private static final int DEFAULT_KEY_SIZE = 128;
    private static final String SM4_ALGORITHM_NAME = "SM4";

    static {
        Security.addProvider(new BouncyCastleProvider());
    }

    /**
     * 生成hex的SM2密钥
     */
    public static void generateSM2Key() {
        SM2 sm2 = SmUtil.sm2();
        BCECPrivateKey ecPri = (BCECPrivateKey) sm2.getPrivateKey();
        BCECPublicKey ecPub = (BCECPublicKey) sm2.getPublicKey();
        BigInteger privateKey = ecPri.getD();
        ECPoint publicKey = ecPub.getQ();
        String privateKeyHex = HexUtil.encodeHexStr(privateKey.toByteArray());
        String publicKeyHex = HexUtil.encodeHexStr(publicKey.getEncoded(false));
        System.out.println(privateKeyHex);
        System.out.println(publicKeyHex);
    }

    /**
     * 生成SM4密钥字符串
     * @return
     * @throws Exception
     */
    public static String generateSM4Key() throws Exception {
        return HexUtil.encodeHexStr(generateSM4Key(DEFAULT_KEY_SIZE));
    }

    /**
     * 生成SM4密钥字节
     * @param keySize
     * @return
     * @throws Exception
     */
    private static byte[] generateSM4Key(int keySize) throws Exception {
        KeyGenerator keyGenerator = KeyGenerator.getInstance(SM4_ALGORITHM_NAME, BouncyCastleProvider.PROVIDER_NAME);
        keyGenerator.init(keySize, new SecureRandom());
        return keyGenerator.generateKey().getEncoded();
    }

    public static void main(String[] args) throws Exception {
        String privateSM2KeyHex = "f8adca86d964a5ff6d9211fc452d79d397153aec6386973751224f0bf7334dcd";
        String publicSM2KeyHex = "04b7afdce28e6f71c27ebb6bf366999494c59a8c0aa66a5c4f8ee2b500bc864ec6de292df1a94160b3a3f96c272576ce72a7b17b9d19055ffa2768edbadbe54800";

        String jsonRequest ="{\"name\":\"hexo\"}";

        // 生成sm4算法密钥
        String sm4Key = SMUtils.generateSM4Key();
        SM4 sm4 = SmUtil.sm4(ByteUtils.fromHexString(sm4Key));

        // 对请求数据进行sm4加密
        String encryptData = sm4.encryptHex(jsonRequest);

        // 对sm4字符串密钥进行SM2加密
        SM2 sm2 = SmUtil.sm2(null, publicSM2KeyHex);
        String encryptKey = sm2.encryptHex(sm4Key, KeyType.PublicKey);

        System.out.println("请求数据：" + jsonRequest);
        System.out.println("sm4Key字符串：" + sm4Key);
        System.out.println("加密后数据：" + encryptData);
        System.out.println("加密后密钥数据：" + encryptKey);

        // 根据SM2密钥对sm4密钥进行解密
        SM2 decryptSM2 = SmUtil.sm2(privateSM2KeyHex, null);
        String decryptSM4Key = decryptSM2.decryptStr(encryptKey, KeyType.PrivateKey);

        SM4 decryptSM4 = SmUtil.sm4(ByteUtils.fromHexString(decryptSM4Key));
        String decryptData = decryptSM4.decryptStr(encryptData);

        System.out.println("解密后sm4Key字符串：" + decryptSM4Key);
        System.out.println("解密后数据：" + decryptData);
    }
}
