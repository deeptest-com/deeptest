package com.ngtesting.platform.entity;

import javax.persistence.*;

@Entity
@Table(name = "tst_msg")
public class TestMsg extends BaseEntity {
    private static final long serialVersionUID = 530835958185680515L;

    private String title;
    @Column(name = "msg", length = 10000)
    private String descr;
    private String uri;

    private Long entityId;
    @Enumerated(EnumType.STRING)
    private TestAlert.AlertType type;

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

    public TestAlert.AlertType getType() {
        return type;
    }

    public void setType(TestAlert.AlertType type) {
        this.type = type;
    }

    public TestUser getOptUser() {
        return optUser;
    }

    public void setOptUser(TestUser optUser) {
        this.optUser = optUser;
    }

    public Long getEntityId() {
        return entityId;
    }

    public void setEntityId(Long entityId) {
        this.entityId = entityId;
    }

    public Long getOptUserId() {
        return optUserId;
    }

    public void setOptUserId(Long optUserId) {
        this.optUserId = optUserId;
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
