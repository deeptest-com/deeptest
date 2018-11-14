package com.ngtesting.platform.model;

public class IsuWorkflowTransition extends BaseModel {
	private static final long serialVersionUID = 595122500409917470L;

    String srcStatusId;
    String srcStatusName;
    String dictStatusId;
    String dictStatusName;

    String actionName;
    Integer actionPageId;
    String actionPageName;

    Integer workflowId;

    Integer orgId;

    public String getSrcStatusId() {
        return srcStatusId;
    }

    public void setSrcStatusId(String srcStatusId) {
        this.srcStatusId = srcStatusId;
    }

    public String getSrcStatusName() {
        return srcStatusName;
    }

    public void setSrcStatusName(String srcStatusName) {
        this.srcStatusName = srcStatusName;
    }

    public String getDictStatusId() {
        return dictStatusId;
    }

    public void setDictStatusId(String dictStatusId) {
        this.dictStatusId = dictStatusId;
    }

    public String getDictStatusName() {
        return dictStatusName;
    }

    public void setDictStatusName(String dictStatusName) {
        this.dictStatusName = dictStatusName;
    }

    public String getActionName() {
        return actionName;
    }

    public void setActionName(String actionName) {
        this.actionName = actionName;
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
