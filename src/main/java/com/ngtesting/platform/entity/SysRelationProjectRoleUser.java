package com.ngtesting.platform.entity;

import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.FetchType;
import javax.persistence.JoinColumn;
import javax.persistence.ManyToOne;
import javax.persistence.Table;

@Entity
@Table(name = "r_project_role_user")
public class SysRelationProjectRoleUser extends BaseEntity {
	private static final long serialVersionUID = 5513768856000982338L;
	
    @Column(name = "project_role_id")
    private Long projectRoleId;
    
    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "project_role_id", insertable = false, updatable = false)
    private SysProjectRole projectRole;

    @Column(name = "user_id")
    private Long userId;
    
    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "user_id", insertable = false, updatable = false)
    private SysUser user;

	public Long getUserId() {
		return userId;
	}
	public void setUserId(Long userId) {
		this.userId = userId;
	}
	public Long getProjectRoleId() {
		return projectRoleId;
	}
	public void setProjectRoleId(Long projectRoleId) {
		this.projectRoleId = projectRoleId;
	}
	public SysProjectRole getProjectRole() {
		return projectRole;
	}
	public void setProjectRole(SysProjectRole projectRole) {
		this.projectRole = projectRole;
	}
	public SysUser getUser() {
		return user;
	}
	public void setUser(SysUser user) {
		this.user = user;
	}
    
}
