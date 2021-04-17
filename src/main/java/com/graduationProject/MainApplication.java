package com.graduationProject;

import org.hyperledger.fabric.gateway.*;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

import java.io.IOException;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.util.Map;
import java.util.concurrent.TimeoutException;

/**
 * @类名 : com.graduationProject.MainApplication
 * @说明 : 主程序
 * @创建日期 : 2021/3/23
 * @作者 : Niaowuuu
 * @版本 : 1.0
 */

@SpringBootApplication
public class MainApplication {

//    public static void main(String[] args) {
//        SpringApplication.run(MainApplication.class, args);
//    }

    public static final String ENVKEY = "CONTRACT_NAME";

//    public static void main(String[] args) {
//        String contractName = "mycc";
//
//        Map<String, String> envvar = System.getenv();
//        if (envvar.containsKey(ENVKEY)){
//            contractName=envvar.get(ENVKEY);
//        }
//
//        Gateway.Builder builder = Gateway.createBuilder();
//
//        try {
//            // A wallet stores a collection of identities
//            Path walletPath = Paths.get(".", "wallet");
//            Wallet wallet = Wallets.newFileSystemWallet(walletPath);
//            System.out.println("Read wallet info from: " + walletPath.toString());
//
//            String userName = "User1@org2.example.com";
//
//            Path connectionProfile = Paths.get("..",  "gateway", "connection-org2.yaml");
//
//            // Set connection options on the gateway builder
//            builder.identity(wallet, userName).networkConfig(connectionProfile).discovery(false);
//
//            // Connect to gateway using application specified parameters
//            try(Gateway gateway = builder.connect()) {
//
//                // Access PaperNet network
//                System.out.println("Use network channel: mychannel.");
//                Network network = gateway.getNetwork("mychannel");
//
//                // Get addressability to commercial paper contract
//                System.out.println("Use org.papernet.commercialpaper smart contract.");
//                Contract contract = network.getContract(contractName, "org.papernet.commercialpaper");
//
//                // Issue commercial paper
//                System.out.println("Submit commercial paper issue transaction.");
//                byte[] response = contract.submitTransaction("issue", "MagnetoCorp", "00001", "2020-05-31", "2020-11-30", "5000000");
//
//                // Process response
//                System.out.println("Process issue transaction response.");
//                CommercialPaper paper = CommercialPaper.deserialize(response);
//                System.out.println(paper);
//            }
//        } catch (GatewayException | IOException | TimeoutException | InterruptedException e) {
//            e.printStackTrace();
//            System.exit(-1);
//        }
//    }
}
