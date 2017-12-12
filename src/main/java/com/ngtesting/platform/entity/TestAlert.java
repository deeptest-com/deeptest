package com.ngtesting.platform.entity;

import javax.persistence.*;
import java.util.Date;

@Entity
@Table(name = "tst_alert")
public class TestAlert extends BaseEntity {
    private static final long serialVersionUID = 4639102366457159222L;

    private String title;
    @Column(name = "msg", length = 10000)
    private String descr;
    private String uri;

    private Date startTime;
    private Date dueTime;

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

    public enum AlertType {
        run_start("run_start", 1),
        run_end("run_end", 1);

        AlertType(String code, Integer remindDay) {
            this.code = code;
            this.remindDay = remindDay;
        }

        public String code;
        public Integer remindDay;
        public String toString() {
            return code;
        }
    }
    public enum AlertAction {
        create("create", "创建"),
        update("update", "更新");

        AlertAction(String code, String msg) {
            this.code = code;
            this.msg = msg;
        }

        public String code;
        public String msg;
        public String toString() {
            return code;
        }
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
