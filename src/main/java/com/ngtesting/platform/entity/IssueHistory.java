package com.ngtesting.platform.entity;

import javax.persistence.*;

@Entity
@Table(name = "isu_history")
public class IssueHistory extends BaseEntity {

    private static final long serialVersionUID = 1545026595607666706L;
    private String title;
    @Column(name = "descr", length = 1000)
    private String descr;

    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "issue_id", insertable = false, updatable = false)
    private Issue issue;

    @Column(name = "issue_id")
    private Long issueId;

    public Issue getIssue() {
        return issue;
    }

    public void setIssue(Issue issue) {
        this.issue = issue;
    }

    public Long getIssueId() {
        return issueId;
    }

    public void setIssueId(Long issueId) {
        this.issueId = issueId;
    }

    public String getDescr() {
        return descr;
    }

    public void setDescr(String descr) {
        this.descr = descr;
    }

    public String getTitle() {
        return title;
    }

    public void setTitle(String title) {
        this.title = title;
    }
}
