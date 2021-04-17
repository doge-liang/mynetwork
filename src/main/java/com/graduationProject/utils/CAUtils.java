package com.graduationProject.utils;

import org.hyperledger.fabric.sdk.NetworkConfig;
import org.hyperledger.fabric.sdk.security.CryptoSuite;
import org.hyperledger.fabric.sdk.security.CryptoSuiteFactory;
import org.hyperledger.fabric_ca.sdk.HFCAClient;

import java.io.File;
import java.nio.file.Paths;
import java.util.Properties;

/**
 * @类名 : CAUtils
 * @说明 : Fabric CA 工具类
 * @创建日期 : 2021/4/18
 * @作者 : Niaowuuu
 * @版本 : 1.0
 */
public class CAUtils  {

//    获取 CA 客户端对象
    public static HFCAClient getCAClient(String orgName, String CA_CERT_PATH) throws Exception {
        // 加载连接文件
        String filePath = Paths.get( "profiles", orgName, "connection.json").toString();
        NetworkConfig config = NetworkConfig.fromJsonFile(new File(filePath));
        NetworkConfig.CAInfo caInfo = config.getOrganizationInfo(orgName).getCertificateAuthorities().get(0);
        String caURL = caInfo.getUrl();
        Properties props = new Properties();
        props.put("pemFile", CA_CERT_PATH);
        props.put("allowAllHostNames", "true");
        // 连接到 CA
        HFCAClient caClient = HFCAClient.createNewInstance(caURL, props);
        CryptoSuite cryptoSuite = CryptoSuiteFactory.getDefault().getCryptoSuite();
        caClient.setCryptoSuite(cryptoSuite);
        return caClient;
    }
}
