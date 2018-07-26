package com.ngtesting.platform.model;

public class TstCaseHistory extends BaseModel {

    private static final long serialVersionUID = 8174711284511001943L;

    private String title;
    private String descr;
    private Integer caseId;

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

    public Integer getCaseId() {
        return caseId;
    }

    public void setCaseId(Integer caseId) {
        this.caseId = caseId;
    }
}
