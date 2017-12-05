package com.ngtesting.platform.entity;

import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.FetchType;
import javax.persistence.JoinColumn;
import javax.persistence.ManyToOne;
import javax.persistence.Table;

@Entity
@Table(name = "tst_r_org_group_user")
public class TestRelationOrgGroupUser extends BaseEntity {
	private static final long serialVersionUID = 8667364655327450367L;

	public TestRelationOrgGroupUser() {
	}
	public TestRelationOrgGroupUser(Long orgId, Long orgGroupId, Long userId) {
		this.orgId = orgId;
		this.orgGroupId = orgGroupId;
		this.userId = userId;
	}

	@Column(name = "org_id")
    private Long orgId;

	private String orgGroupName;

    @Column(name = "org_group_id")
    private Long orgGroupId;

    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "org_group_id", insertable = false, updatable = false)
    private TestOrgGroup orgGroup;

    @Column(name = "user_id")
    private Long userId;

    private String userName;

    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "user_id", insertable = false, updatable = false)
    private TestUser user;

	public Long getUserId() {
		return userId;
	}

	public void setUserId(Long userId) {
		this.userId = userId;
	}
	public Long getOrgGroupId() {
		return orgGroupId;
	}
	public void setOrgGroupId(Long orgGroupId) {
		this.orgGroupId = orgGroupId;
	}
	public TestOrgGroup getOrgGroup() {
		return orgGroup;
	}
	public void setOrgGroup(TestOrgGroup orgGroup) {
		this.orgGroup = orgGroup;
	}
	public TestUser getUser() {
		return user;
	}
	public void setUser(TestUser user) {
		this.user = user;
	}
	public String getOrgGroupName() {
		return orgGroupName;
	}
	public void setOrgGroupName(String orgGroupName) {
		this.orgGroupName = orgGroupName;
	}
	public String getUserName() {
		return userName;
	}
	public void setUserName(String userName) {
		this.userName = userName;
	}

}
