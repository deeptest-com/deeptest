package com.ngtesting.platform.entity;

import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.FetchType;
import javax.persistence.JoinColumn;
import javax.persistence.ManyToOne;
import javax.persistence.Table;

import org.hibernate.annotations.Cache;
import org.hibernate.annotations.CacheConcurrencyStrategy;
import org.hibernate.annotations.DynamicInsert;
import org.hibernate.annotations.DynamicUpdate;

@Entity
@Table(name = "r_project_role_user")
@DynamicInsert @DynamicUpdate
@Cache(usage = CacheConcurrencyStrategy.READ_WRITE)
public class SysProjectRoleUser extends BaseEntity {
	private static final long serialVersionUID = 5513768856000982338L;

	@Column(name = "project_id")
    private Long projectId;

    @Column(name = "user_id")
    private Long userId;

    @Column(name = "project_role_id")
    private Long projectRoleId;
    
    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "role_id", insertable = false, updatable = false)
    private SysProjectRole projectRole;
    
	public Long getProjectId() {
		return projectId;
	}
	public void setProjectId(Long projectId) {
		this.projectId = projectId;
	}
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
    
}
