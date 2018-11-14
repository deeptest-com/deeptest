package com.ngtesting.platform.model;

import java.util.LinkedList;
import java.util.List;

public class IsuWorkflow extends BaseModel {
	private static final long serialVersionUID = -132330064734288116L;

	private String name;
    private String code;
    private String descr;
    private Integer displayOrder;
    private Boolean isFinal;
    private Boolean isBuildIn = false;
    private Integer orgId;

	private List<IsuWorkflowTransition> statusTransitions = new LinkedList<>();

	public List<IsuWorkflowTransition> getStatusTransitions() {
		return statusTransitions;
	}

	public void setStatusTransitions(List<IsuWorkflowTransition> statusTransitions) {
		this.statusTransitions = statusTransitions;
	}

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
	public Boolean getIsBuildIn() {
		return isBuildIn;
	}
	public void setIsBuildIn(Boolean isBuildIn) {
		this.isBuildIn = isBuildIn;
	}
	public Integer getOrgId() {
		return orgId;
	}
	public void setOrgId(Integer orgId) {
		this.orgId = orgId;
	}

}
