package com.ngtesting.platform.model;


import java.util.HashMap;
import java.util.LinkedList;
import java.util.List;
import java.util.Map;

public class TstProject extends BaseModel {
	private static final long serialVersionUID = 8069068510875783820L;

	private String name;
    private String descr;
    private Integer orgId;

    private Integer parentId;
    private ProjectType type;
    private Integer childrenNumb;

	private Boolean isLastestProjectGroup;

    private Boolean selected = false;
    private Boolean selecting = false;

    List<TstProject> children = new LinkedList();
	Map<String, Boolean> privs = new HashMap<>();

	public static enum ProjectType {
		org("org"),
		group("group"),
		project("project");

		ProjectType(String textVal) {
			this.textVal = textVal;
		}

		private String textVal;

		public String toString() {
			return textVal;
		}
	}

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
	public Integer getOrgId() {
		return orgId;
	}
	public void setOrgId(Integer orgId) {
		this.orgId = orgId;
	}

	public Integer getParentId() {
		return parentId;
	}
	public void setParentId(Integer parentId) {
		this.parentId = parentId;
	}

	public ProjectType getType() {
		return type;
	}
	public void setType(ProjectType type) {
		this.type = type;
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

    public List<TstProject> getChildren() {
        return children;
    }

    public void setChildren(List<TstProject> children) {
        this.children = children;
    }
}
