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
import org.hibernate.annotations.FilterDef;
import org.hibernate.annotations.ParamDef;

@Entity
@Table(name = "tst_project")
@FilterDef(name="filter_project_deleted", parameters={@ParamDef( name="isDeleted", type="boolean" ) })
public class TestProject extends BaseEntity {
	private static final long serialVersionUID = 7813647435255173689L;
	private String name;
    
	@Column(name = "descr", length = 1000)
    private String descr;
	
	@Enumerated(EnumType.STRING)
	private ProjectType type;

	private Integer ord;
	private Date lastAccessTime;
	
    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "parent_id", insertable = false, updatable = false)
    private TestProject parent;

    @Column(name = "parent_id")
    private Long parentId;
    
    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "org_id", insertable = false, updatable = false)
    private SysOrg org;

    @Column(name = "org_id")
    private Long orgId;
    
    @OneToMany(mappedBy="parent", fetch = FetchType.LAZY) 
    @OrderBy("id")
    @Filter(name="filter_project_deleted", condition="deleted = :isDeleted ")
    private List<TestProject> children = new LinkedList<TestProject>();
    
    public static enum ProjectType {
        group("group"),
        project("project");

        private ProjectType(String textVal) {
            this.textVal = textVal;
        }

        private String textVal;
        
        public static ProjectType getEnum(String str) {
        	ProjectType type = null;
        	switch(str) { 
            	case "group": type = ProjectType.group; break;
            	case "project": type = ProjectType.project; break;
            }
        	
        	return type;
        }

        public String toString() {
            return textVal;
        }
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

	public SysOrg getOrg() {
		return org;
	}

	public void setOrg(SysOrg org) {
		this.org = org;
	}

	public Long getOrgId() {
		return orgId;
	}

	public void setOrgId(Long orgId) {
		this.orgId = orgId;
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

	public ProjectType getType() {
		return type;
	}

	public void setType(ProjectType type) {
		this.type = type;
	}

	public List<TestProject> getChildren() {
		return children;
	}

	public void setChildren(List<TestProject> children) {
		this.children = children;
	}

	public Integer getOrd() {
		return ord;
	}

	public void setOrd(Integer ord) {
		this.ord = ord;
	}

	public Date getLastAccessTime() {
		return lastAccessTime;
	}

	public void setLastAccessTime(Date lastAccessTime) {
		this.lastAccessTime = lastAccessTime;
	}
    
}
