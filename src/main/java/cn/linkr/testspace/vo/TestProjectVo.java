package cn.linkr.testspace.vo;

import java.util.LinkedList;

public class TestProjectVo extends BaseVo {
	private static final long serialVersionUID = -9069520320732281911L;
	private String name;
    private String descr;
    private Long companyId;
    
    private Long parentId;
    private String type;
    private Integer order;
    private Integer childrenNumb;
    
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

	public Long getParentId() {
		return parentId;
	}
	public void setParentId(Long parentId) {
		this.parentId = parentId;
	}
	
	public String getType() {
		return type;
	}
	public void setType(String type) {
		this.type = type;
	}

	public Integer getOrder() {
		return order;
	}
	public void setOrder(Integer order) {
		this.order = order;
	}
	public Integer getChildrenNumb() {
		return childrenNumb;
	}
	public void setChildrenNumb(Integer childrenNumb) {
		this.childrenNumb = childrenNumb;
	}

}
