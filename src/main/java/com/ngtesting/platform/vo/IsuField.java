package com.ngtesting.platform.vo;


import com.ngtesting.platform.config.ConstantIssue;

import java.io.Serializable;

public class IsuField implements Serializable {

	private static final long serialVersionUID = 3168995179136496564L;

	private String code;
	private String label;
	private Boolean display;
	ConstantIssue.IssueFilterType type;
    ConstantIssue.IssueFilterInput input;

    public String getCode() {
        return code;
    }

    public void setCode(String code) {
        this.code = code;
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

    public ConstantIssue.IssueFilterInput getInput() {
        return input;
    }

    public void setInput(ConstantIssue.IssueFilterInput input) {
        this.input = input;
    }
}
