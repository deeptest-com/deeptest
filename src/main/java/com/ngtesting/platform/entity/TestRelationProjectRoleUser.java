package com.ngtesting.platform.entity;

import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.FetchType;
import javax.persistence.JoinColumn;
import javax.persistence.ManyToOne;
import javax.persistence.Table;

@Entity
@Table(name = "tst_r_project_role_user")
public class TestRelationProjectRoleUser extends BaseEntity {
	private static final long serialVersionUID = 5513768856000982338L;

    private String projectRoleName;
    private String userName;

	@Column(name = "project_id")
	private Long projectId;
	
    @Column(name = "project_role_id")
    private Long projectRoleId;
    
    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "project_role_id", insertable = false, updatable = false)
    private TestProjectRole projectRole;

    @Column(name = "user_id")
    private Long userId;
    
    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "user_id", insertable = false, updatable = false)
    private TestUser user;

    public TestRelationProjectRoleUser() {

    }
    public TestRelationProjectRoleUser(Long projectId, Long userId, Long projectRoleId, String projectRoleName, String userName) {
        this.projectId = projectId;
        this.userId = userId;
        this.projectRoleId = projectRoleId;
        this.projectRoleName = projectRoleName;
        this.userName = userName;
    }

    public String getProjectRoleName() {
        return projectRoleName;
    }

    public void setProjectRoleName(String projectRoleName) {
        this.projectRoleName = projectRoleName;
    }

    public String getUserName() {
        return userName;
    }

    public void setUserName(String userName) {
        this.userName = userName;
    }

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
	public TestProjectRole getProjectRole() {
		return projectRole;
	}
	public void setProjectRole(TestProjectRole projectRole) {
		this.projectRole = projectRole;
	}
	public TestUser getUser() {
		return user;
	}
	public void setUser(TestUser user) {
		this.user = user;
	}
    
}
