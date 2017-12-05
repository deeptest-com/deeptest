package com.ngtesting.platform.entity;

import javax.persistence.*;

@Entity
@Table(name = "tst_case_priority")
public class TestCasePriority extends BaseEntity {
	private static final long serialVersionUID = -7863740464740341072L;

	private String code;
	private String name;
	private String descr;
    private Integer displayOrder;
    private Boolean isDefault;

    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "org_id", insertable = false, updatable = false)
    private TestOrg org;

    @Column(name = "org_id")
    private Long orgId;

	public String getName() {
		return name;
	}
	public void setName(String name) {
		this.name = name;
	}
	public Boolean getIsDefault() {
		return isDefault;
	}
	public void setIsDefault(Boolean isDefault) {
		this.isDefault = isDefault;
	}
	public String getCode() {
		return code;
	}
	public void setCode(String code) {
		this.code = code;
	}
	public TestOrg getOrg() {
		return org;
	}
	public void setOrg(TestOrg org) {
		this.org = org;
	}
	public Long getOrgId() {
		return orgId;
	}
	public void setOrgId(Long orgId) {
		this.orgId = orgId;
	}
	public Integer getDisplayOrder() {
		return displayOrder;
	}
	public void setDisplayOrder(Integer displayOrder) {
		this.displayOrder = displayOrder;
	}
	public String getDescr() {
		return descr;
	}
	public void setDescr(String descr) {
		this.descr = descr;
	}

}
