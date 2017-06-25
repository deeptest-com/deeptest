package com.ngtesting.platform.vo;

public class TestCaseStepVo extends BaseVo {

    private String opt;
    private String expect;

    private Integer ordr;

    private Long testCaseId;

    public TestCaseStepVo(Long id, String opt, String expect, Integer ordr) {
        this.id = id;
        this.opt = opt;
        this.expect = expect;
        this.ordr = ordr;
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
}
