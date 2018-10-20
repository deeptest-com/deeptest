package com.ngtesting.platform.vo;


import com.ngtesting.platform.config.ConstantIssue;

import java.io.Serializable;

public class IsuJqlColumn implements Serializable {

	private static final long serialVersionUID = -7413029715796093478L;
	private String id;
	private String label;
	private Boolean display;
	ConstantIssue.IssueFilterType type;

    public String getId() {
        return id;
    }

    public void setId(String id) {
        this.id = id;
    }

    public String getLabel() {
		return label;
	}

	public void setLabel(String label) {
		this.label = label;
	}

	public Boolean getDisplay() {
		return display;
	}

	public void setDisplay(Boolean display) {
		this.display = display;
	}

	public ConstantIssue.IssueFilterType getType() {
		return type;
	}

	public void setType(ConstantIssue.IssueFilterType type) {
		this.type = type;
	}
}
