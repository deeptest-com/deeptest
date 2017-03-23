package cn.linkr.testspace.entity;

import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.EnumType;
import javax.persistence.Enumerated;
import javax.persistence.FetchType;
import javax.persistence.JoinColumn;
import javax.persistence.ManyToOne;
import javax.persistence.Table;
import javax.persistence.Transient;

import cn.linkr.testspace.util.Constant.TreeNodeType;

@Entity
@Table(name = "tst_project")
public class TestProject extends BaseEntity {
	private static final long serialVersionUID = 7813647435255173689L;
	private String name;
    
	@Column(name = "descr", length = 1000)
    private String descr;
	
	private Boolean isActive;
	@Enumerated(EnumType.STRING)
	private TreeNodeType type;
	private String path;
	@Transient
	private Integer level;
	private Integer oderInParent;
	
    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "parent_id", insertable = false, updatable = false)
    private TestProject parent;

    @Column(name = "parent_id")
    private Long parentId;
    
    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "company_id", insertable = false, updatable = false)
    private SysCompany company;

    @Column(name = "company_id")
    private Long companyId;
    
	public Integer getLevel() {
		return getPath().split("/").length - 1;
	}

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

	public TestProject getParent() {
		return parent;
	}

	public void setParent(TestProject parent) {
		this.parent = parent;
	}

	public Long getParentId() {
		return parentId;
	}

	public void setParentId(Long parentId) {
		this.parentId = parentId;
	}

	public String getPath() {
		return path;
	}

	public void setPath(String path) {
		this.path = path;
	}

	public TreeNodeType getType() {
		return type;
	}

	public void setType(TreeNodeType type) {
		this.type = type;
	}

	public Integer getOderInParent() {
		return oderInParent;
	}

	public void setOderInParent(Integer oderInParent) {
		this.oderInParent = oderInParent;
	}
    
}
