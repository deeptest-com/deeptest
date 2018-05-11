package com.ngtesting.platform.entity;

import javax.persistence.*;

@Entity
@Table(name = "isu_status_transition")
public class IssueStatusTransition extends BaseEntity {
    private static final long serialVersionUID = 7260005873110268288L;

    IssueStatus src;
    IssueStatus dict;
    String action;

    private Integer ordr;

    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "workflow_id", insertable = false, updatable = false)
    private IssueWorkflow workflow;

    @Column(name = "workflow_id")
    private Long workflowId;

    public Integer getOrdr() {
        return ordr;
    }

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

    public void setOrdr(Integer ordr) {
        this.ordr = ordr;
    }

    public IssueWorkflow getWorkflow() {
        return workflow;
    }

    public void setWorkflow(IssueWorkflow workflow) {
        this.workflow = workflow;
    }

    public Long getWorkflowId() {
        return workflowId;
    }

    public void setWorkflowId(Long workflowId) {
        this.workflowId = workflowId;
    }
}
