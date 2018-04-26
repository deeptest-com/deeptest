package com.ngtesting.platform.vo;

import com.ngtesting.platform.entity.IssueStatus;

public class IssueStatusTransitionVo extends BaseVo {
	private static final long serialVersionUID = 595122500409917470L;

    IssueStatus src;
    IssueStatus dict;
    String action;

    private Integer ordr;

    private Long workflowId;

    public IssueStatus getSrc() {
        return src;
    }

    public void setSrc(IssueStatus src) {
        this.src = src;
    }

    public IssueStatus getDict() {
        return dict;
    }

    public void setDict(IssueStatus dict) {
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
