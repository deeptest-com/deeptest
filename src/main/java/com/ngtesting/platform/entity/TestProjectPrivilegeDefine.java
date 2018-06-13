package com.ngtesting.platform.entity;

import javax.persistence.Entity;
import javax.persistence.Table;

@Entity
@Table(name = "tst_project_privilege_define")
public class TestProjectPrivilegeDefine extends BaseEntity {
	private static final long serialVersionUID = -5510206858644860272L;

    private String code;

    private String name;
    private String descr;

    private String action;
	private String actionName;

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

	public String getCode() {
		return code;
	}

	public void setCode(String code) {
		this.code = code;
	}

	public String getAction() {
		return action;
	}

	public void setAction(String action) {
		this.action = action;
	}

	public String getActionName() {
		return actionName;
	}

	public void setActionName(String actionName) {
		this.actionName = actionName;
	}
}
