package com.graduationProject.entity;

import com.graduationProject.utils.CAUtils;
import lombok.AccessLevel;
import lombok.Data;
import lombok.Setter;
import org.hyperledger.fabric.gateway.*;
import org.hyperledger.fabric.sdk.Enrollment;
import org.hyperledger.fabric.sdk.security.CryptoSuite;
import org.hyperledger.fabric.sdk.security.CryptoSuiteFactory;
import org.hyperledger.fabric_ca.sdk.EnrollmentRequest;
import org.hyperledger.fabric_ca.sdk.HFCAClient;

import java.io.IOException;
import java.nio.file.Paths;
import java.security.PrivateKey;

/**
 * @类名 : EnrollAdmin
 * @说明 : 登录管理员
 * @创建日期 : 2021/4/17
 * @作者 : Niaowuuu
 * @版本 : 1.0
 */
@Data
public class AdminDO {

    static {
        System.setProperty("org.hyperledger.fabric.sdk.service_discovery.as_localhost", "true");
    }

//    private String orgname_provider = "Provider";
//    private String adminname_provider = "admin";
//    private String adminpwd_provider = "adminpw";
//    private String ca_cert_provider = "profiles/" + orgname_provider + "/tls/" + "ca.provider.mynetwork.com-cert.pem";
//    private String mspid_provider = "ProviderMSP";

    //    Admin 所属组织
    private String orgName;
    //    Admin 用户名
    private String name;
    //    Admin 密码
    private String pwd;
    //    钱包
    private Wallet wallet;
    //    Admin 用户凭证
    @Setter(AccessLevel.NONE)
    private Enrollment adminKeys;
    //    组织 MSPID
    @Setter(AccessLevel.NONE)
    private String orgMSP;
    //    CA 公钥
    @Setter(AccessLevel.NONE)
    private String CA_CERT_PATH;

    public AdminDO(String name, String pwd, String orgName) throws IOException {
        this.name = name;
        this.pwd = pwd;
        this.orgName = orgName;
        CA_CERT_PATH = "profiles/" + orgName + "/tls/" + "ca." + orgName.toLowerCase() + ".mynetwork.com-cert.pem";
        orgMSP = orgName + "MSP";
        wallet = Wallets.newFileSystemWallet(Paths.get("wallet", orgName));
    }

    public void setOrgName(String orgName) {
        this.orgName = orgName;
    }

    public void setAdminKeys(Enrollment adminKeys) {

    }

    public Boolean doEnroll() throws Exception {
        // 连接到 CA
        HFCAClient caClient = CAUtils.getCAClient(orgName, CA_CERT_PATH);
        CryptoSuite cryptoSuite = CryptoSuiteFactory.getDefault().getCryptoSuite();
        caClient.setCryptoSuite(cryptoSuite);
        // 检查 Wallet 内是否有 admin
        X509Identity adminIdentity = (X509Identity) wallet.get(name);
        if (adminIdentity != null) {
            System.out.println("An identity for the admin user \"admin@" + orgName + "\" already exists in the " +
                    "wallet");
            adminKeys = new Enrollment() {
                @Override
                public PrivateKey getKey() {
                    return adminIdentity.getPrivateKey();
                }

                @Override
                public String getCert() {
                    return Identities.toPemString(adminIdentity.getCertificate());
                }
            };
            return true;
        }
        // 登录账号
        final EnrollmentRequest enrollmentRequestTLS = new EnrollmentRequest();
        enrollmentRequestTLS.addHost("52.82.52.96");
        enrollmentRequestTLS.setProfile("tls");
        adminKeys = caClient.enroll(name, pwd, enrollmentRequestTLS);
        Identity user = Identities.newX509Identity(orgMSP, adminKeys);

        System.out.println(user.getMspId());
        System.out.println(this.adminKeys.getCert());
        wallet.put(name, user);
        System.out.println("Successfully enrolled user \"admin@" + orgName + "\" and imported into the wallet");
        return true;
    }


    public static void main(String[] args) throws Exception {
        AdminDO adminDO = new AdminDO("admin", "adminpw", "Subscriber");
        System.out.print(adminDO.doEnroll());
    }

}
