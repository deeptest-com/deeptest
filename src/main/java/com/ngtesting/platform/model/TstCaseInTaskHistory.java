package com.ngtesting.platform.model;

public class TstCaseInTaskHistory extends BaseModel {

    private static final long serialVersionUID = 6709823091860707370L;

    private String title;
    private String descr;
    private Integer caseId;
    private Integer caseInTaskId;

    public Integer getCaseId() {
        return caseId;
    }

    public void setCaseId(Integer caseId) {
        this.caseId = caseId;
    }

    public Integer getCaseInTaskId() {
        return caseInTaskId;
    }

    public void setCaseInTaskId(Integer caseInTaskId) {
        this.caseInTaskId = caseInTaskId;
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

}
