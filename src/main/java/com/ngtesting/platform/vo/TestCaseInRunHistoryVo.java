package com.ngtesting.platform.vo;

import java.util.Date;

public class TestCaseInRunHistoryVo extends BaseVo {

    private static final long serialVersionUID = 6709823091860707370L;

    private String title;
    private String descr;
    private Long testCaseId;

    public TestCaseInRunHistoryVo() {
    }

    public TestCaseInRunHistoryVo(Long id, String title, String descr, Long testCaseId, Date createTime) {
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

    public Long getTestCaseId() {
        return testCaseId;
    }

    public void setTestCaseId(Long testCaseId) {
        this.testCaseId = testCaseId;
    }
}
