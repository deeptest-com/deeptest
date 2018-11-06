package com.ngtesting.platform.model;


import com.ngtesting.platform.config.ConstantIssue;

import java.io.Serializable;

public class IsuField implements Serializable {

	private static final long serialVersionUID = 3168995179136496564L;

    private Integer id;
	private String code;
	private String label;
    private String key;
    private ConstantIssue.IssueType type;
    private ConstantIssue.IssueInput input;
    private Integer ordr;
    private Boolean fullLine;

    private Boolean display;

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

    public ConstantIssue.IssueType getType() {
        return type;
    }

    public void setType(ConstantIssue.IssueType type) {
        this.type = type;
    }

    public ConstantIssue.IssueInput getInput() {
        return input;
    }

    public void setInput(ConstantIssue.IssueInput input) {
        this.input = input;
    }

    public Integer getOrdr() {
        return ordr;
    }

    public void setOrdr(Integer ordr) {
        this.ordr = ordr;
    }

    public Integer getId() {
        return id;
    }

    public void setId(Integer id) {
        this.id = id;
    }

    public String getKey() {
        return key;
    }

    public void setKey(String key) {
        this.key = key;
    }

    public Boolean getFullLine() {
        return fullLine;
    }

    public void setFullLine(Boolean fullLine) {
        this.fullLine = fullLine;
    }
}
