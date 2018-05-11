package com.ngtesting.platform.vo;

import java.util.LinkedList;
import java.util.List;

public class IssueWorkflowVo extends BaseVo {
	private static final long serialVersionUID = -132330064734288116L;

	private String name;
    private String code;
    private String descr;
    private Integer displayOrder;
    private Boolean isFinal;
    private Boolean isBuildIn = false;
    private Long orgId;

	private List<IssueStatusTransitionVo> statusTransitions = new LinkedList<>();

	public List<IssueStatusTransitionVo> getStatusTransitions() {
		return statusTransitions;
	}

	public void setStatusTransitions(List<IssueStatusTransitionVo> statusTransitions) {
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
	public Long getOrgId() {
		return orgId;
	}
	public void setOrgId(Long orgId) {
		this.orgId = orgId;
	}

}
