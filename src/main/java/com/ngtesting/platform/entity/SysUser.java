package com.ngtesting.platform.entity;

import java.util.Date;
import java.util.HashSet;
import java.util.Set;

import javax.persistence.CascadeType;
import javax.persistence.Entity;
import javax.persistence.ManyToMany;
import javax.persistence.Table;

import org.hibernate.annotations.Filter;

@Entity
@Table(name = "sys_user")
public class SysUser extends BaseEntity {
	private static final long serialVersionUID = 5110565175672074546L;
	
	private String email;
    private String phone;
    private String name;
    private String password;
    private String token;
    private String avatar;

    private String verifyCode;
    private Date lastLoginTime;
    
    private Long defaultOrgId;
    private Long defaultProjectId;
    
    @ManyToMany(cascade = {CascadeType.PERSIST, CascadeType.MERGE}, mappedBy = "userSet")
    private Set<SysRole> roleSet = new HashSet<SysRole>(0);
    
    @ManyToMany(cascade = {CascadeType.PERSIST, CascadeType.MERGE}, mappedBy = "userSet")
    private Set<SysOrg> orgSet = new HashSet<SysOrg>(0);

    public String getEmail() {
        return email;
    }

    public void setEmail(String email) {
        this.email = email;
    }

    public String getPhone() {
        return phone;
    }

    public void setPhone(String phone) {
        this.phone = phone;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public String getToken() {
        return token;
    }

    public void setToken(String token) {
        this.token = token;
    }

	public String getPassword() {
		return password;
	}

	public void setPassword(String password) {
		this.password = password;
	}

	public String getVerifyCode() {
		return verifyCode;
	}

	public void setVerifyCode(String verifyCode) {
		this.verifyCode = verifyCode;
	}

	public Date getLastLoginTime() {
		return lastLoginTime;
	}

	public void setLastLoginTime(Date lastLoginTime) {
		this.lastLoginTime = lastLoginTime;
	}

	public Set<SysRole> getRoleSet() {
		return roleSet;
	}

	public void setRoleSet(Set<SysRole> roleSet) {
		this.roleSet = roleSet;
	}

	public String getAvatar() {
		return avatar;
	}

	public void setAvatar(String avatar) {
		this.avatar = avatar;
	}

	public Long getDefaultOrgId() {
		return defaultOrgId;
	}

	public void setDefaultOrgId(Long defaultOrgId) {
		this.defaultOrgId = defaultOrgId;
	}

	public Set<SysOrg> getOrgSet() {
		return orgSet;
	}

	public void setOrgSet(Set<SysOrg> orgSet) {
		this.orgSet = orgSet;
	}

	public Long getDefaultProjectId() {
		return defaultProjectId;
	}

	public void setDefaultProjectId(Long defaultProjectId) {
		this.defaultProjectId = defaultProjectId;
	}
	
}
