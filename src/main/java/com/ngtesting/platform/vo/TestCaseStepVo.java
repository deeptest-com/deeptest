package com.ngtesting.platform.vo;

public class TestCaseStepVo extends BaseVo {

    private static final long serialVersionUID = -4163163556800840532L;

    private String opt;
    private String expect;

    private Integer ordr;

    private Long testCaseId;

    public TestCaseStepVo() {
    }

    public TestCaseStepVo(Long id, String opt, String expect, Integer ordr, Long testCaseId) {
        this.id = id;
        this.opt = opt;
        this.expect = expect;
        this.ordr = ordr;
        this.testCaseId = testCaseId;
    }

    public String getOpt() {
        return opt;
    }

    public void setOpt(String opt) {
        this.opt = opt;
    }

    public String getExpect() {
        return expect;
    }

    public void setExpect(String expect) {
        this.expect = expect;
    }

    public Integer getOrdr() {
        return ordr;
    }

    public void setOrdr(Integer ordr) {
        this.ordr = ordr;
    }

    public Long getTestCaseId() {
        return testCaseId;
    }

    public void setTestCaseId(Long testCaseId) {
        this.testCaseId = testCaseId;
    }

    public Long getId() {
        return id;
    }

    public void setId(Long id) {
        this.id = id;
    }

    public Boolean getDeleted() {
        return deleted;
    }

    public void setDeleted(Boolean deleted) {
        this.deleted = deleted;
    }

    public Boolean getDisabled() {
        return disabled;
    }

    public void setDisabled(Boolean disabled) {
        this.disabled = disabled;
    }

    public Integer getVersion() {
        return version;
    }

    public void setVersion(Integer version) {
        this.version = version;
    }


}
