package com.ngtesting.platform.entity;

import java.util.HashSet;
import java.util.LinkedList;
import java.util.List;
import java.util.Set;

import javax.persistence.CascadeType;
import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.FetchType;
import javax.persistence.JoinColumn;
import javax.persistence.JoinTable;
import javax.persistence.ManyToMany;
import javax.persistence.ManyToOne;
import javax.persistence.OneToMany;
import javax.persistence.OrderBy;
import javax.persistence.Table;

import org.hibernate.annotations.Filter;

@Entity
@Table(name = "tst_org")
public class TestOrg extends BaseEntity {
	private static final long serialVersionUID = -970910958057582029L;

	private String name;
    private String website;

    @Column(name = "admin_id")
    private Long adminId;

    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "admin_id", insertable = false, updatable = false)
    private TestUser admin;

    @ManyToMany(fetch = FetchType.LAZY, cascade = CascadeType.ALL)
	@JoinTable(name = "tst_r_org_user", joinColumns = {
			@JoinColumn(name = "org_id", nullable = false, updatable = false) },
			inverseJoinColumns = { @JoinColumn(name = "user_id",
					nullable = false, updatable = false) })
    private Set<TestUser> userSet = new HashSet<TestUser>(0);

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

	public Set<TestUser> getUserSet() {
		return userSet;
	}

	public void setUserSet(Set<TestUser> userSet) {
		this.userSet = userSet;
	}

	public Long getAdminId() {
		return adminId;
	}

	public void setAdminId(Long adminId) {
		this.adminId = adminId;
	}

	public TestUser getAdmin() {
		return admin;
	}

	public void setAdmin(TestUser admin) {
		this.admin = admin;
	}

}
