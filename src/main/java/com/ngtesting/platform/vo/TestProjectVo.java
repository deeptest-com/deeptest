package com.ngtesting.platform.vo;


import java.util.HashMap;
import java.util.Map;

public class TestProjectVo extends BaseVo {
	private static final long serialVersionUID = 8069068510875783820L;

	private String name;
    private String descr;
    private Long orgId;

    private Long parentId;
    private String type;
    private Integer order;
    private Integer childrenNumb;

	private Boolean isLastestProjectGroup;

    private Boolean selected = false;
    private Boolean selecting = false;

	Map<String, Boolean> privs = new HashMap<>();

	public Boolean getLastestProjectGroup() {
		return isLastestProjectGroup;
	}

	public void setLastestProjectGroup(Boolean lastestProjectGroup) {
		isLastestProjectGroup = lastestProjectGroup;
	}

	public Map<String, Boolean> getPrivs() {
		return privs;
	}

	public void setPrivs(Map<String, Boolean> privs) {
		this.privs = privs;
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
	public Long getOrgId() {
		return orgId;
	}
	public void setOrgId(Long orgId) {
		this.orgId = orgId;
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
	public Boolean getSelected() {
		return selected;
	}
	public void setSelected(Boolean selected) {
		this.selected = selected;
	}
	public Boolean getSelecting() {
		return selecting;
	}
	public void setSelecting(Boolean selecting) {
		this.selecting = selecting;
	}

}
