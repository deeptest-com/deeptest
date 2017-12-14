package com.ngtesting.platform.entity;

import javax.persistence.*;
import java.util.Date;

@Entity
@Table(name = "tst_alert")
public class TestAlert extends BaseEntity {
    private static final long serialVersionUID = 4639102366457159222L;

    private String title;
    @Column(name = "descr", length = 10000)
    private String descr;
    private String uri;

    private Date startTime;
    private Date dueTime;

    private Boolean isRead = false;
    private Boolean sent = false;

    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "opt_user_id", insertable = false, updatable = false)
    private TestUser optUser;

    @Column(name = "opt_user_id")
    private Long optUserId;

    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "user_id", insertable = false, updatable = false)
    private TestUser user;

    @Column(name = "user_id")
    private Long userId;

    public Boolean getSent() {
        return sent;
    }

    public void setSent(Boolean sent) {
        this.sent = sent;
    }

    public Boolean getRead() {
        return isRead;
    }

    public void setRead(Boolean read) {
        isRead = read;
    }

    public Date getStartTime() {
        return startTime;
    }

    public void setStartTime(Date startTime) {
        this.startTime = startTime;
    }


    public TestUser getOptUser() {
        return optUser;
    }

    public void setOptUser(TestUser optUser) {
        this.optUser = optUser;
    }

    public Long getOptUserId() {
        return optUserId;
    }

    public void setOptUserId(Long optUserId) {
        this.optUserId = optUserId;
    }

    public Date getDueTime() {
        return dueTime;
    }

    public void setDueTime(Date dueTime) {
        this.dueTime = dueTime;
    }

    public String getDescr() {
        return descr;
    }

    public void setDescr(String descr) {
        this.descr = descr;
    }

    public String getUri() {
        return uri;
    }

    public void setUri(String uri) {
        this.uri = uri;
    }

	public String getTitle() {
		return title;
	}

	public void setTitle(String title) {
		this.title = title;
	}

    public TestUser getUser() {
        return user;
    }

    public void setUser(TestUser user) {
        this.user = user;
    }

    public Long getUserId() {
        return userId;
    }

    public void setUserId(Long userId) {
        this.userId = userId;
    }

}
