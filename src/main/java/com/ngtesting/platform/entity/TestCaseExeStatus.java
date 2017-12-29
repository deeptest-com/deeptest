package com.ngtesting.platform.entity;

import javax.persistence.*;

@Entity
@Table(name = "tst_case_exe_status")
public class TestCaseExeStatus extends BaseEntity {
	private static final long serialVersionUID = 4775052158868753948L;

	private String name;
    private String code;
    private String descr;
    private Integer displayOrder;
    private Boolean isFinal;

    private Boolean isBuildIn = false;

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

	public String getCode() {
		return code;
	}

	public void setCode(String code) {
		this.code = code;
	}

	public String getDescr() {
		return descr;
	}

	public void setDescr(String descr) {
		this.descr = descr;
	}

	public Integer getDisplayOrder() {
		return displayOrder;
	}

	public void setDisplayOrder(Integer displayOrder) {
		this.displayOrder = displayOrder;
	}

	public Boolean getIsFinal() {
		return isFinal;
	}
	public void setIsFinal(Boolean isFinal) {
		this.isFinal = isFinal;
	}

	public Boolean getBuildIn() {
		return isBuildIn;
	}
	public void setBuildIn(Boolean buildIn) {
		isBuildIn = buildIn;
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



}
