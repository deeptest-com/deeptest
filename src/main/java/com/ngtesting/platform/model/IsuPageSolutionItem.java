package com.ngtesting.platform.model;

import com.ngtesting.platform.config.ConstantIssue;

public class IsuPageSolutionItem extends BaseModel {
    private static final long serialVersionUID = -3872162780544592329L;

    private Integer issueTypeId;
    private String issueTypeName;

    private ConstantIssue.IssueOpt opt;

    private Integer issuePageId;
    private String issuePageName;

    private Integer pageSolutionId;
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

    public Integer getIssueTypeId() {
        return issueTypeId;
    }

    public void setIssueTypeId(Integer issueTypeId) {
        this.issueTypeId = issueTypeId;
    }

    public Integer getIssuePageId() {
        return issuePageId;
    }

    public void setIssuePageId(Integer issuePageId) {
        this.issuePageId = issuePageId;
    }

    public Integer getPageSolutionId() {
        return pageSolutionId;
    }

    public void setPageSolutionId(Integer pageSolutionId) {
        this.pageSolutionId = pageSolutionId;
    }

    public String getIssueTypeName() {
        return issueTypeName;
    }

    public void setIssueTypeName(String issueTypeName) {
        this.issueTypeName = issueTypeName;
    }

    public String getIssuePageName() {
        return issuePageName;
    }

    public void setIssuePageName(String issuePageName) {
        this.issuePageName = issuePageName;
    }
}
