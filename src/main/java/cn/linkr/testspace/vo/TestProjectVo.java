package cn.linkr.testspace.vo;

import java.util.LinkedList;

public class TestProjectVo extends BaseVo {
	private static final long serialVersionUID = -9069520320732281911L;
	private String name;
    private String descr;
    private Long companyId;
    private Boolean isActive;
    private Boolean isFirstChild;
    
    private Long parentId;
    private String path;
    private String type;
    private Boolean isFilterOut;
    
    private Integer level;
    private Integer brotherNumb = 0;
    private Integer parentDescendantNumber = 0;
    
    LinkedList<TestProjectVo> children = new LinkedList<TestProjectVo>();
    
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

	public Long getParentId() {
		return parentId;
	}
	public void setParentId(Long parentId) {
		this.parentId = parentId;
	}

	public LinkedList<TestProjectVo> getChildren() {
		return children;
	}
	public void setChildren(LinkedList<TestProjectVo> children) {
		this.children = children;
	}

	public String getPath() {
		return path;
	}
	public void setPath(String path) {
		this.path = path;
	}
	public Integer getLevel() {
		return level;
	}
	public void setLevel(Integer level) {
		this.level = level;
	}
	public Boolean getIsFirstChild() {
		return isFirstChild;
	}
	public void setIsFirstChild(Boolean isFirstChild) {
		this.isFirstChild = isFirstChild;
	}
	public Integer getBrotherNumb() {
		return brotherNumb;
	}
	public void setBrotherNumb(Integer brotherNumb) {
		this.brotherNumb = brotherNumb;
	}

	public Integer getParentDescendantNumber() {
		return parentDescendantNumber;
	}
	public void setParentDescendantNumber(Integer parentDescendantNumber) {
		this.parentDescendantNumber = parentDescendantNumber;
	}
	public String getType() {
		return type;
	}
	public void setType(String type) {
		this.type = type;
	}
	public Boolean getIsFilterOut() {
		return isFilterOut;
	}
	public void setIsFilterOut(Boolean isFilterOut) {
		this.isFilterOut = isFilterOut;
	}

}
