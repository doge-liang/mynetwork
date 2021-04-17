package com.graduationProject.mynetwork;

import org.hyperledger.fabric.gateway.*;

import java.io.IOException;
import java.nio.file.Path;
import java.nio.file.Paths;

/**
 * @类名 : InvokeQuery
 * @说明 : Query请求
 * @创建日期 : 2021/4/17
 * @作者 : Niaowuuu
 * @版本 : 1.0
 */
public class InvokeQuery {
    static {
        System.setProperty("org.hyperledger.fabric.sdk.service_discovery.as_localhost", "true");
    }

//
    private String orgname_org1;
    private String username_org1;
    private String channel_name;
    private String contract_name;

    private void doQuery(String orgName, String userName, String functionName, String key)
            throws IOException, ContractException {
        //get user identity from wallet.
        Path walletPath = Paths.get("wallet", orgName);
        Wallet wallet = Wallets.newFileSystemWallet(walletPath);
        Identity identity = wallet.get(userName);

        //check identity existence in wallet
        if (identity == null) {
            System.out.println("The identity \"" + userName + "@"+ orgName + "\" doesn't exists in the wallet");
            return;
        }

        //load connection profile
        Path networkConfigPath = Paths.get( "profiles", orgName, "connection.json");
        Gateway.Builder builder = Gateway.createBuilder();
        builder.identity(wallet, userName).networkConfig(networkConfigPath).discovery(true);

        //create a gateway connection
        try (Gateway gateway = builder.connect()) {

            // get the network and contract
            Network network = gateway.getNetwork(channel_name);
            Contract contract = network.getContract(contract_name);

            byte[] result = contract.evaluateTransaction(functionName, key);
            System.out.println(new String(result));
        }
    }
}
