package com.ngtesting.platform.vo;

import java.util.HashMap;
import java.util.Map;

public class TestCaseVo extends BaseVo {
	private static final long serialVersionUID = -5955583523485410239L;
	
	private String title;
	private Integer priority;
	private Integer estimate;
	private String objective;
    private String descr;

	private Integer path;
	private String type;

	private Map<String, String> extend = new HashMap<String, String>();

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

	public Integer getEstimate() {
		return estimate;
	}

	public void setEstimate(Integer estimate) {
		this.estimate = estimate;
	}

	public String getDescr() {
		return descr;
	}

	public void setDescr(String descr) {
		this.descr = descr;
	}

	public Integer getPath() {
		return path;
	}

	public void setPath(Integer path) {
		this.path = path;
	}

	public String getType() {
		return type;
	}

	public void setType(String type) {
		this.type = type;
	}

	public String getObjective() {
		return objective;
	}

	public void setObjective(String objective) {
		this.objective = objective;
	}

	public Map<String, String> getExtend() {
		return extend;
	}

	public void setExtend(Map<String, String> extend) {
		this.extend = extend;
	}

//    private Long moduleId;

//    LinkedList<TestCaseVo> children = new LinkedList<TestCaseVo>();

//	public Long getModuleId() {
//		return moduleId;
//	}
//
//	public void setModuleId(Long moduleId) {
//		this.moduleId = moduleId;
//	}

//	public LinkedList<TestCaseVo> getChildren() {
//		return children;
//	}
//
//	public void setChildren(LinkedList<TestCaseVo> children) {
//		this.children = children;
//	}

}
