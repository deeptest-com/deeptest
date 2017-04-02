package com.ngtesting.platform.entity;

import java.util.HashSet;
import java.util.Set;

import javax.persistence.CascadeType;
import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.FetchType;
import javax.persistence.JoinColumn;
import javax.persistence.JoinTable;
import javax.persistence.ManyToMany;
import javax.persistence.ManyToOne;
import javax.persistence.Table;

@Entity
@Table(name = "sys_org")
public class SysOrg extends BaseEntity {
	private static final long serialVersionUID = -970910958057582029L;
	
	private String name;
    private String website;
    
    @Column(name = "admin_id")
    private Long adminId;
    
    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "admin_id", insertable = false, updatable = false)
    private SysUser admin;
    
    @ManyToMany(fetch = FetchType.LAZY, cascade = CascadeType.ALL)
	@JoinTable(name = "r_org_user", joinColumns = { 
			@JoinColumn(name = "org_id", nullable = false, updatable = false) }, 
			inverseJoinColumns = { @JoinColumn(name = "user_id", 
					nullable = false, updatable = false) })
    private Set<SysUser> userSet = new HashSet<SysUser>(0);

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

	public String getWebsite() {
		return website;
	}

	public void setWebsite(String website) {
		this.website = website;
	}

	public Set<SysUser> getUserSet() {
		return userSet;
	}

	public void setUserSet(Set<SysUser> userSet) {
		this.userSet = userSet;
	}

	public Long getAdminId() {
		return adminId;
	}

	public void setAdminId(Long adminId) {
		this.adminId = adminId;
	}

	public SysUser getAdmin() {
		return admin;
	}

	public void setAdmin(SysUser admin) {
		this.admin = admin;
	}

}
