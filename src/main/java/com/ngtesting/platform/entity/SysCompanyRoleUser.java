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
@Table(name = "r_company_role_user")
public class SysCompanyRoleUser extends BaseEntity {
	private static final long serialVersionUID = 5513768856000982338L;

	@Column(name = "company_id")
    private Long companyId;

    @Column(name = "user_id")
    private Long userId;

    @Column(name = "company_role_id")
    private Long companyRoleId;
    
    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "company_role_id", insertable = false, updatable = false)
    private SysCompanyRole companyRole;
    
	public Long getUserId() {
		return userId;
	}
	public void setUserId(Long userId) {
		this.userId = userId;
	}
	public Long getCompanyRoleId() {
		return companyRoleId;
	}
	public void setCompanyRoleId(Long companyRoleId) {
		this.companyRoleId = companyRoleId;
	}
	public SysCompanyRole getCompanyRole() {
		return companyRole;
	}
	public void setCompanyRole(SysCompanyRole companyRole) {
		this.companyRole = companyRole;
	}
	public Long getCompanyId() {
		return companyId;
	}
	public void setCompanyId(Long companyId) {
		this.companyId = companyId;
	}
    
}
