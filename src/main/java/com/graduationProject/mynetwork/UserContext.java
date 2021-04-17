package com.graduationProject.mynetwork;

import org.hyperledger.fabric.sdk.Enrollment;
import org.hyperledger.fabric.sdk.User;

import java.util.Set;

/**
 * @类名 : UserContext
 * @说明 : 用户上下文
 * @创建日期 : 2021/4/17
 * @作者 : Niaowuuu
 * @版本 : 1.0
 */
public class UserContext implements User {
    private String name;
    private String mspID;
    private Enrollment enrollment;

    public UserContext(String name, String mspID) {
        super();
        this.name = name;
        this.mspID = mspID;
    }

    public UserContext(String name, String mspID, Enrollment enrollment) {
        this(name, mspID);
        this.enrollment = enrollment;
    }

    @Override
    public String getName() {
        return this.name;
    }

    @Override
    public Set<String> getRoles() {
        return null;
    }

    @Override
    public String getAccount() {
        return null;
    }

    @Override
    public String getAffiliation() {
        return null;
    }

    @Override
    public Enrollment getEnrollment() {
        return this.enrollment;
    }

    @Override
    public String getMspId() {
        return this.mspID;
    }
}
