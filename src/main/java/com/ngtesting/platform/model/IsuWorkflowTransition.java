package com.ngtesting.platform.model;

public class IsuWorkflowTransition extends BaseModel {
	private static final long serialVersionUID = 595122500409917470L;

    String name;
    Integer actionPageId;
    String actionPageName;

    Integer srcStatusId;
    String srcStatusName;
    Integer dictStatusId;
    String dictStatusName;

    Integer workflowId;

    Integer orgId;

    public IsuWorkflowTransition(){}

    public IsuWorkflowTransition(Integer workflowId, Integer srcId, Integer dictId) {
        super();
    }

    public Integer getSrcStatusId() {
        return srcStatusId;
    }

    public void setSrcStatusId(Integer srcStatusId) {
        this.srcStatusId = srcStatusId;
    }

    public String getSrcStatusName() {
        return srcStatusName;
    }

    public void setSrcStatusName(String srcStatusName) {
        this.srcStatusName = srcStatusName;
    }

    public Integer getDictStatusId() {
        return dictStatusId;
    }

    public void setDictStatusId(Integer dictStatusId) {
        this.dictStatusId = dictStatusId;
    }

    public String getDictStatusName() {
        return dictStatusName;
    }

    public void setDictStatusName(String dictStatusName) {
        this.dictStatusName = dictStatusName;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public Integer getActionPageId() {
        return actionPageId;
    }

    public void setActionPageId(Integer actionPageId) {
        this.actionPageId = actionPageId;
    }

    public String getActionPageName() {
        return actionPageName;
    }

    public void setActionPageName(String actionPageName) {
        this.actionPageName = actionPageName;
    }

    public Integer getOrgId() {
        return orgId;
    }

    public void setOrgId(Integer orgId) {
        this.orgId = orgId;
    }

    public Integer getWorkflowId() {
        return workflowId;
    }

    public void setWorkflowId(Integer workflowId) {
        this.workflowId = workflowId;
    }
}
