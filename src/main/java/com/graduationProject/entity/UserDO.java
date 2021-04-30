package com.graduationProject.entity;

import com.graduationProject.mynetwork.UserContext;
import com.graduationProject.utils.CAUtils;
import lombok.AccessLevel;
import lombok.Data;
import lombok.Setter;
import org.hyperledger.fabric.gateway.*;
import org.hyperledger.fabric.sdk.Enrollment;
import org.hyperledger.fabric.sdk.User;
import org.hyperledger.fabric_ca.sdk.HFCAClient;
import org.hyperledger.fabric_ca.sdk.RegistrationRequest;

import java.io.IOException;
import java.nio.file.Path;
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
public class UserDO {

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

    public UserDO(String userName, String userSecret, String orgName) throws IOException {
        this.userName = userName;
        this.userSecret = userSecret;
        this.orgName = orgName;
        this.CA_CERT_PATH = "profiles/" + orgName + "/tls/" + "ca." + orgName.toLowerCase() + ".mynetwork.com-cert.pem";
        this.orgMSP = orgName + "MSP";
        this.wallet = Wallets.newFileSystemWallet(Paths.get("wallet", orgName));
    }

    public Boolean doEnroll() throws Exception {
        HFCAClient caClient = CAUtils.getCAClient(orgName, CA_CERT_PATH);
        // 登录用户
        this.enrollment = caClient.enroll(userName, userSecret);
        Identity user = Identities.newX509Identity(orgMSP, enrollment);
        wallet.put(userName, user);
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
        HFCAClient caClient = CAUtils.getCAClient(orgName, CA_CERT_PATH);
        AdminDO admin = new AdminDO(adminName, adminSecret, orgName);

        // 检查管理员登录状态
        admin.doEnroll();
        if (!admin.doEnroll()) {
            System.out.println("\"" + adminName + "@" + orgName + "\" needs to be enrolled and added to the wallet first");
            return false;
        }
        System.out.println(admin.getAdminKeys().getCert());
        // 拿到管理员的上下文
        User adminContext = new UserContext(adminName, orgMSP, admin.getAdminKeys());

        // 注册用户
        RegistrationRequest registrationRequest = new RegistrationRequest(userName);
        registrationRequest.setSecret(userSecret);
        caClient.register(registrationRequest, adminContext);

        // 登录用户
        this.doEnroll();
        System.out.println("Successfully enrolled user \"" + userName + "@" + orgName + "\" and imported into the wallet");
        return true;
    }

    public String doQuery(String functionName, String key) throws IOException, ContractException {
        Identity identity = wallet.get(userName);
        // 如果未登录
        if (identity == null) {
            System.out.println("The identity \"" + userName + "@" + orgName + "\" doesn't exists in the wallet");
            return "";
        }
        //加载连接文件
        Path networkConfigPath = Paths.get("profiles", orgName, "connection.json");
        Gateway.Builder builder = Gateway.createBuilder();
        builder.identity(wallet, userName).networkConfig(networkConfigPath).discovery(true);

        // 建立连接
        try (Gateway gateway = builder.connect()) {

            // 获取合约和网络
            Network network = gateway.getNetwork("mychannel");
            Contract contract = network.getContract("strategy");

            byte[] result = contract.evaluateTransaction(functionName, key);
            System.out.println(new String(result));
            return new String(result);
        }
    }

    public static void main(String[] args) throws Exception {
        UserDO user = new UserDO("user1", "user1pw", "Subscriber");
        user.doEnroll();
        user.doQuery("GetAllStrategies", "");
    }

}
