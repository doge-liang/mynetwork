package com.graduationProject.entity;

import com.graduationProject.mynetwork.UserContext;
import com.graduationProject.utils.CAUtils;
import lombok.AccessLevel;
import lombok.Data;
import lombok.Setter;
import org.hyperledger.fabric.gateway.*;
import org.hyperledger.fabric.sdk.Enrollment;
import org.hyperledger.fabric_ca.sdk.HFCAClient;
import org.hyperledger.fabric_ca.sdk.RegistrationRequest;

import java.io.IOException;
import java.nio.file.Paths;
import java.security.PrivateKey;

/**
 * @类名 : User
 * @说明 : 用户实体类
 * @创建日期 : 2021/4/11
 * @作者 : Niaowuuu
 * @版本 : 1.0
 */
@Data
public class User {

    //    用户ID
    private String identity;
    //    用户身份认证
    private Enrollment enrollment;
    //    用户所属组织名
    private String orgName;
    //    用户名
    private String userName;
    //    密码
    private String userSecret;
    //    用户凭证
    @Setter(AccessLevel.NONE)
    private Enrollment key;
    //    组织 MSPID
    @Setter(AccessLevel.NONE)
    private String orgMSP;
    //    CA 公钥
    @Setter(AccessLevel.NONE)
    private String CA_CERT_PATH;
    //    本地钱包
    private Wallet wallet;

    public void setOrgName(String orgName) throws IOException {
        orgName = orgName;
        CA_CERT_PATH = "profiles/" + orgName + "/tls/" + "ca." + orgName.toLowerCase() + ".mynetwork.com-cert.pem";
        orgMSP = orgName + "MSP";
        wallet = Wallets.newFileSystemWallet(Paths.get("wallet", orgName));
    }

    public Boolean doRegisterUser(String orgMSP, String adminName, String userName, String userSecret) throws Exception {
        //Connect CA
        HFCAClient caClient = CAUtils.getCAClient(orgName, CA_CERT_PATH);

        Wallet wallet = Wallets.newFileSystemWallet(Paths.get("wallet", orgName));

        //Check admin existence in wallet
        X509Identity adminIdentity = (X509Identity) wallet.get(adminName);
        if (adminIdentity == null) {
            System.out.println("\"" + adminName + "@" + orgName + "\" needs to be enrolled and added to the wallet first");
            return false;
        }

        //Check admin to be created existence in wallet
        if (wallet.get(userName) != null) {
            System.out.println("An identity for the user \"" + userName + "@" + orgName + "\" already exists in the wallet");
            return true;
        }

        //Get admin's UserContext
        Enrollment adminKeys = new Enrollment() {
            @Override
            public PrivateKey getKey() {
                return adminIdentity.getPrivateKey();
            }

            @Override
            public String getCert() {
                return Identities.toPemString(adminIdentity.getCertificate());
            }
        };
        org.hyperledger.fabric.sdk.User admin = new UserContext(adminName, orgMSP, adminKeys);

        //Register user
        RegistrationRequest registrationRequest = new RegistrationRequest(userName);
        registrationRequest.setSecret(userSecret);
        caClient.register(registrationRequest, admin);
        try {
            doEnroll();
        } catch (Exception e) {
            System.out.println("登录失败");
            e.printStackTrace();
            return false;
        }
        return true;
    }

    public Boolean doEnroll() throws Exception {
        HFCAClient caClient = CAUtils.getCAClient(orgName, CA_CERT_PATH);
        //Enroll user
        Enrollment enrollment = caClient.enroll(userName, userSecret);
        Identity user = Identities.newX509Identity(orgMSP, enrollment);
        wallet.put(userName, user);
        System.out.println("Successfully enrolled user \"" + userName + "@" + orgName + "\" and imported into the wallet");
        return true;
    }

}
