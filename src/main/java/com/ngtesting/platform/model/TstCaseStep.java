package com.ngtesting.platform.model;

public class TstCaseStep extends BaseModel {

    private static final long serialVersionUID = -4163163556800840532L;

    private String opt;
    private String expect;

    private Integer ordr;

    private Integer testCaseId;

    public TstCaseStep() {
    }

    public TstCaseStep(Integer id, String opt, String expect, Integer ordr, Integer testCaseId) {
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

    public Integer getTestCaseId() {
        return testCaseId;
    }

    public void setTestCaseId(Integer testCaseId) {
        this.testCaseId = testCaseId;
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
