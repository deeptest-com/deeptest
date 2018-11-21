package com.ngtesting.platform.model;

import com.ngtesting.platform.config.ConstantIssue;

public class IsuPageSolutionItem extends BaseModel {
    private static final long serialVersionUID = -3872162780544592329L;

    private Integer typeId;
    private String typeName;

    private ConstantIssue.IssueOpt opt;

    private Integer pageId;
    private String pageName;

    private Integer solutionId;
    private Integer orgId;

    public Integer getOrgId() {
        return orgId;
    }

    public void setOrgId(Integer orgId) {
        this.orgId = orgId;
    }

    public ConstantIssue.IssueOpt getOpt() {
        return opt;
    }

    public void setOpt(ConstantIssue.IssueOpt opt) {
        this.opt = opt;
    }

    public Integer getPageId() {
        return pageId;
    }

    public void setPageId(Integer pageId) {
        this.pageId = pageId;
    }

    public String getPageName() {
        return pageName;
    }

    public void setPageName(String pageName) {
        this.pageName = pageName;
    }

    public Integer getSolutionId() {
        return solutionId;
    }

    public void setSolutionId(Integer solutionId) {
        this.solutionId = solutionId;
    }

    public Integer getTypeId() {
        return typeId;
    }

    public void setTypeId(Integer typeId) {
        this.typeId = typeId;
    }

    public String getTypeName() {
        return typeName;
    }

    public void setTypeName(String typeName) {
        this.typeName = typeName;
    }
}
