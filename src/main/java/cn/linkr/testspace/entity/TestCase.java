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
@Table(name = "tst_case")
public class TestCase extends BaseEntity {
	private static final long serialVersionUID = -7253288259861070288L;

    private String title;
	private Integer priority;
	private Integer estimate;
    
	@Column(name = "descr", length = 1000)
    private String descr;

	@Enumerated(EnumType.STRING)
	private TreeNodeType type;
	private String path;
	@Transient
	private Integer level;
	private Integer orderInParent;
	
    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "parent_id", insertable = false, updatable = false)
    private TestCase parent;

    @Column(name = "parent_id")
    private Long parentId;
	
    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "project_id", insertable = false, updatable = false)
    private TestProject project;

    @Column(name = "project_id")
    private Long projectId;
    
    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "module_id", insertable = false, updatable = false)
    private TestModule module;

    @Column(name = "module_id")
    private Long moduleId;
    
	public Integer getLevel() {
		return getPath().split("/").length - 1;
	}

	public String getTitle() {
		return title;
	}

	public void setTitle(String title) {
		this.title = title;
	}

	public Integer getPriority() {
		return priority;
	}

	public void setPriority(Integer priority) {
		this.priority = priority;
	}

	public String getDescr() {
		return descr;
	}

	public void setDescr(String descr) {
		this.descr = descr;
	}

	public TestModule getModule() {
		return module;
	}

	public void setModule(TestModule module) {
		this.module = module;
	}

	public Long getModuleId() {
		return moduleId;
	}

	public void setModuleId(Long moduleId) {
		this.moduleId = moduleId;
	}

	public Integer getEstimate() {
		return estimate;
	}

	public void setEstimate(Integer estimate) {
		this.estimate = estimate;
	}

	public TreeNodeType getType() {
		return type;
	}

	public void setType(TreeNodeType type) {
		this.type = type;
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

	public TestCase getParent() {
		return parent;
	}

	public void setParent(TestCase parent) {
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

	public Integer getOrderInParent() {
		return orderInParent;
	}

	public void setOrderInParent(Integer orderInParent) {
		this.orderInParent = orderInParent;
	}
}
