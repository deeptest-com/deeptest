package com.ngtesting.platform.entity;

import java.util.Date;
import java.util.LinkedList;
import java.util.List;

import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.EnumType;
import javax.persistence.Enumerated;
import javax.persistence.FetchType;
import javax.persistence.JoinColumn;
import javax.persistence.ManyToOne;
import javax.persistence.OneToMany;
import javax.persistence.OrderBy;
import javax.persistence.Table;

import org.hibernate.annotations.Filter;

@Entity
@Table(name = "tst_project_access_history")
public class TestProjectAccessHistory extends BaseEntity {
	private static final long serialVersionUID = -8980336892866667072L;
	
	public TestProjectAccessHistory() {
	}
	
	public TestProjectAccessHistory(Long orgId, Long userId, Long projectId, String projectName) {
		this.orgId = orgId;
		this.userId = userId;
		this.projectId = projectId;
		this.projectName = projectName;
	}

	private Date lastAccessTime = new Date();
	
    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "org_id", insertable = false, updatable = false)
    private TestUser org;

    @Column(name = "org_id")
    private Long orgId;

    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "user_id", insertable = false, updatable = false)
    private TestUser user;

    @Column(name = "user_id")
    private Long userId;
    
    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "project_id", insertable = false, updatable = false)
    private TestProject project;

    @Column(name = "project_id")
    private Long projectId;
    
    private String projectName;

	public Date getLastAccessTime() {
		return lastAccessTime;
	}

	public void setLastAccessTime(Date lastAccessTime) {
		this.lastAccessTime = lastAccessTime;
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

	public TestProject getProject() {
		return project;
	}

	public void setProject(TestProject project) {
		this.project = project;
	}

	public Long getProjectId() {
		return projectId;
	}

	public void setProjectId(Long projectId) {
		this.projectId = projectId;
	}

	public TestUser getOrg() {
		return org;
	}

	public void setOrg(TestUser org) {
		this.org = org;
	}

	public Long getOrgId() {
		return orgId;
	}

	public void setOrgId(Long orgId) {
		this.orgId = orgId;
	}

	public String getProjectName() {
		return projectName;
	}

	public void setProjectName(String projectName) {
		this.projectName = projectName;
	}
    
}
