package com.ngtesting.platform.entity;

import com.ngtesting.platform.config.Constant;

import javax.persistence.*;

@Entity
@Table(name = "tst_msg")
public class TestMsg extends BaseEntity {
    private static final long serialVersionUID = 530835958185680515L;

    private String name;
    @Column(name = "descr", length = 10000)
    private String descr;
    private String uri;

    private Long entityId;
    @Enumerated(EnumType.STRING)
    private Constant.AlertType type;
    @Enumerated(EnumType.STRING)
    private Constant.MsgType action;

    private Boolean isRead = false;

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

    public Boolean getRead() {
        return isRead;
    }

    public void setRead(Boolean read) {
        isRead = read;
    }

    public Constant.MsgType getAction() {
        return action;
    }

    public void setAction(Constant.MsgType action) {
        this.action = action;
    }

    public Constant.AlertType getType() {
        return type;
    }

    public void setType(Constant.AlertType type) {
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

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
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
