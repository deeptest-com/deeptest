package com.ngtesting.platform.model;

import java.util.Date;

public class TstCaseHistory extends BaseModel {

    private static final long serialVersionUID = 8174711284511001943L;

    private String title;
    private String descr;
    private Integer testCaseId;

    public TstCaseHistory() {
    }

    public TstCaseHistory(Integer id, String title, String descr, Integer testCaseId, Date createTime) {
        this.id = id;
        this.title = title;
        this.descr = descr;
        this.testCaseId = testCaseId;
        this.createTime = createTime;
    }

    public String getTitle() {
        return title;
    }

    public void setTitle(String title) {
        this.title = title;
    }

    public String getDescr() {
        return descr;
    }

    public void setDescr(String descr) {
        this.descr = descr;
    }

    public Integer getTestCaseId() {
        return testCaseId;
    }

    public void setTestCaseId(Integer testCaseId) {
        this.testCaseId = testCaseId;
    }
}
