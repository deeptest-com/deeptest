package cn.linkr.testspace.entity;

import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.FetchType;
import javax.persistence.JoinColumn;
import javax.persistence.ManyToOne;
import javax.persistence.Table;

import cn.linkr.testspace.entity.EvtEvent.EventStatus;

@Entity
@Table(name = "tst_project")
public class TestProject extends BaseEntity {
	private static final long serialVersionUID = 7813647435255173689L;
	private String name;
    
	@Column(name = "descr", length = 1000)
    private String descr;
	
	private Boolean isActive;
    
    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "company_id", insertable = false, updatable = false)
    private SysCompany company;

    @Column(name = "company_id")
    private Long companyId;

	public String getName() {
		return name;
	}

	public void setName(String name) {
		this.name = name;
	}

	public String getDescr() {
		return descr;
	}

	public void setDescr(String descr) {
		this.descr = descr;
	}

	public SysCompany getCompany() {
		return company;
	}

	public void setCompany(SysCompany company) {
		this.company = company;
	}

	public Long getCompanyId() {
		return companyId;
	}

	public void setCompanyId(Long companyId) {
		this.companyId = companyId;
	}

	public Boolean getIsActive() {
		return isActive;
	}

	public void setIsActive(Boolean isActive) {
		this.isActive = isActive;
	}
    
    
}
