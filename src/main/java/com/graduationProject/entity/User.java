package com.graduationProject.entity;

import com.graduationProject.mynetwork.UserContext;
import com.graduationProject.utils.CAUtils;
import lombok.AccessLevel;
import lombok.Data;
import lombok.Getter;
import lombok.Setter;
import org.hyperledger.fabric.gateway.*;
import org.hyperledger.fabric.sdk.Enrollment;
import org.hyperledger.fabric_ca.sdk.Attribute;
import org.hyperledger.fabric_ca.sdk.EnrollmentRequest;
import org.hyperledger.fabric_ca.sdk.HFCAClient;
import org.hyperledger.fabric_ca.sdk.RegistrationRequest;
import org.hyperledger.fabric_ca.sdk.exception.InvalidArgumentException;

import java.io.IOException;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.security.PrivateKey;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.concurrent.TimeoutException;

/**
 * @author : Niaowuuu
 * @ClassName : User
 * @说明 : 用户实体类
 * @创建日期 : 2021/4/11
 * @since : 1.0
 */
@Data
public class User {
    // static {
    //     System.setProperty("org.hyperledger.fabric.sdk.service_discovery.as_localhost", "true");
    // }

    //    用户身份认证
    private Enrollment enrollment;
    //    用户所属组织名
    private String orgName;
    private String caName;
    //    用户名
    private String userName;
    //    密码
    @Getter(AccessLevel.NONE)
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
    //    网络配置文件路径
    private Path networkConfigPath;

    public User(String userName, String userSecret, String orgName) throws IOException {
        this.userName = userName;
        this.userSecret = userSecret;
        this.orgName = orgName;
        this.caName = "ca-" + orgName;
        this.CA_CERT_PATH = "profiles/" + orgName + "/tls/" + "ca." + orgName.toLowerCase() + ".mynetwork.com-cert.pem";
        this.orgMSP = orgName + "MSP";
        this.wallet = Wallets.newFileSystemWallet(Paths.get("wallet", orgName));
        this.networkConfigPath = Paths.get("profiles", orgName, "connection.json");
    }

    private void setVIP(Boolean isVIP) throws Exception {
        wallet.remove(userName);
        final EnrollmentRequest enrollmentRequestTLS = new EnrollmentRequest();
        enrollmentRequestTLS.addHost("52.82.52.96");
        enrollmentRequestTLS.setProfile("tls");
        if (isVIP) {
            enrollmentRequestTLS.addAttrReq("strategy.role");
        }
        HFCAClient caClient = CAUtils.getCAClient(caName, orgName, CA_CERT_PATH);
        this.enrollment = caClient.enroll(userName, userSecret, enrollmentRequestTLS);
        Identity user = Identities.newX509Identity(orgMSP, enrollment);
        wallet.put(userName, user);
    }

    public Boolean doEnroll() throws Exception {
        X509Identity identity = (X509Identity) wallet.get(userName);
        if (identity != null) {
            System.out.println("An identity for the user " + userName +"@" + orgName + "\" already exists in the " +
                    "wallet");
            enrollment = new Enrollment() {
                @Override
                public PrivateKey getKey() {
                    return identity.getPrivateKey();
                }

                @Override
                public String getCert() {
                    return Identities.toPemString(identity.getCertificate());
                }
            };
            return true;
        }
        return this.login();
    }

    public Boolean login() throws Exception {
        final EnrollmentRequest enrollmentRequestTLS = new EnrollmentRequest();
        enrollmentRequestTLS.addHost("52.82.52.96");
        enrollmentRequestTLS.setProfile("tls");
        // 登录用户
        HFCAClient caClient = CAUtils.getCAClient(caName, orgName, CA_CERT_PATH);
        this.enrollment = caClient.enroll(userName, userSecret, enrollmentRequestTLS);
        Identity user = Identities.newX509Identity(orgMSP, enrollment);
        wallet.put(userName, user);
        // System.out.println(Identities.newX509Identity(orgMSP, enrollment).getCertificate());
        System.out.println("Successfully enrolled user \"" + userName + "@" + orgName + "\" and imported into the wallet");
        return true;
    }

    public Boolean doRegister(String adminName, String adminSecret) throws Exception {
        // 查询钱包是否已经存在用户
        if (wallet.get(userName) != null) {
            System.out.println("An identity for the user \"" + userName + "@" + orgName + "\" already exists in the wallet");
            this.doEnroll();
            return true;
        }
        HFCAClient caClient = CAUtils.getCAClient(caName, orgName, CA_CERT_PATH);
        User admin = new User(adminName, adminSecret, orgName);

        // 检查管理员登录状态
        admin.doEnroll();
        System.out.println(admin.getEnrollment().getCert());
        // 拿到管理员的上下文
        org.hyperledger.fabric.sdk.User adminContext = new UserContext(adminName, orgMSP, admin.getEnrollment());

        // 注册用户
        RegistrationRequest registrationRequest = new RegistrationRequest(userName);
        registrationRequest.setSecret(userSecret);
        // registrationRequest.setAffiliation(orgName);
        registrationRequest.addAttribute(new Attribute("strategy.role", "vip"));
        caClient.register(registrationRequest, adminContext);

        // 登录用户
        this.doEnroll();
        System.out.println("Successfully enrolled user \"" + userName + "@" + orgName + "\" and imported into the wallet");
        return true;
    }

    public byte[] doQuery(String functionName, String... key) {
        Gateway.Builder builder = Gateway.createBuilder();
        try {

            Identity identity = wallet.get(userName);
            // 如果未登录
            if (identity == null) {
                System.out.println("The identity \"" + userName + "@" + orgName + "\" doesn't exists in the wallet");
                return new byte[0];
            }
            //加载连接文件
            builder.identity(wallet, userName).networkConfig(networkConfigPath).discovery(false);


            // 建立连接
            try (Gateway gateway = builder.connect()) {

                // 获取合约和网络
                Network network = gateway.getNetwork("mychannel");
                Contract contract = network.getContract("strategy", "org.mynetwork.strategy");


                return contract.evaluateTransaction(functionName, key);
            }
        } catch (IOException | ContractException e) {
            e.printStackTrace();
            return new byte[0];
        }
    }

    public byte[] doInvoke(String functionName, String... args) {

        Gateway.Builder builder = Gateway.createBuilder();
        try {
            if (wallet.get(userName) == null) {
                System.out.println("The identity \"" + userName + "@" + orgName + "\" doesn't exists in the wallet");
                return new byte[0];
            }
            builder.identity(wallet, userName).networkConfig(networkConfigPath).discovery(false);
            try (Gateway gateway = builder.connect()) {
                Network network = gateway.getNetwork("mychannel");
                Contract contract = network.getContract("strategy", "org.mynetwork.strategy");

                return contract.submitTransaction(functionName, args);
            }
        } catch (IOException | ContractException | TimeoutException | InterruptedException e) {
            e.printStackTrace();
            return new byte[0];
        }
    }

    /**
     * 测试方法
     */
    public static void main(String[] args) throws Exception {
        // User user = new User("user1", "user1pw", "Subscriber");
        User user = new User("vip", "vippw", "Subscriber");
        // User user = new User("admin", "adminpw", "Provider");
        // user.doRegister("admin", "adminpw");
        user.doEnroll();
        // System.out.println(user.getEnrollment().getCert());
        // Strategy newStrategy = Strategy.createInstance("1", "测试SDK创建策略", "", 0.0, 0.0, 0.0, 0, new ArrayList<>());
        // System.out.println(Strategy.serialize(newStrategy));
        // user.doInvoke("Distribute", Strategy.serialize(newStrategy));
        user.doInvoke("Subscribe", "2");
        user.setVIP(true);
        user.doInvoke("Subscribe", "2");
        // byte[] result = user.doQuery("GetAllStrategies", "");
        // if (result.length != 0) {
        //     List<Strategy> strategy = Strategy.deserializeList(result);
        //     System.out.println(strategy);
        // }
    }

}
