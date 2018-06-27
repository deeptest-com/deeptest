package com.ngtesting.platform.vo;

public class IssueStatusTransitionVo extends BaseVo {
	private static final long serialVersionUID = 595122500409917470L;

    String src;
    String dict;
    String action;

    private Integer ordr;

    private Long workflowId;

    public String getSrc() {
        return src;
    }

    public void setSrc(String src) {
        this.src = src;
    }

    public String getDict() {
        return dict;
    }

    public void setDict(String dict) {
        this.dict = dict;
    }

    public String getAction() {
        return action;
    }

    public void setAction(String action) {
        this.action = action;
    }

    public Integer getOrdr() {
        return ordr;
    }

    public void setOrdr(Integer ordr) {
        this.ordr = ordr;
    }

    public Long getWorkflowId() {
        return workflowId;
    }

    public void setWorkflowId(Long workflowId) {
        this.workflowId = workflowId;
    }
}
