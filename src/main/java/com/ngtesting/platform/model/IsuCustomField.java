package com.ngtesting.platform.model;

import com.ngtesting.platform.config.ConstantIssue;

import java.util.LinkedList;
import java.util.List;

public class IsuCustomField extends BaseModel {
	private static final long serialVersionUID = 1279085732086560549L;

	private String code;
    private String label;
    private String myColumn;
    private Boolean required = false;

    ConstantIssue.IssueType type;
    ConstantIssue.IssueInput input;

    private String descr;

	private List<IsuCustomFieldOption> options = new LinkedList<>();

    private Integer rows = 3;

    private ConstantIssue.TextFormat textFormat = ConstantIssue.TextFormat.plain_text;
    private Boolean global = true;
    private Boolean buildIn = false;

    private Integer ordr;
    private Integer orgId;

    public List<IsuCustomFieldOption> getOptions() {
        return options;
    }

    public void setOptions(List<IsuCustomFieldOption> options) {
        this.options = options;
    }

    public Boolean getRequired() {
        return required;
    }

    public void setRequired(Boolean required) {
        this.required = required;
    }

    public Boolean getGlobal() {
        return global;
    }

    public void setGlobal(Boolean global) {
        this.global = global;
    }

    public Boolean getBuildIn() {
        return buildIn;
    }

    public void setBuildIn(Boolean buildIn) {
        this.buildIn = buildIn;
    }

    public String getMyColumn() {
        return myColumn;
    }

    public void setMyColumn(String myColumn) {
        this.myColumn = myColumn;
    }

    public String getLabel() {
		return label;
	}

	public void setLabel(String label) {
		this.label = label;
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

	public ConstantIssue.IssueType getType() {
		return type;
	}

	public void setType(ConstantIssue.IssueType type) {
		this.type = type;
	}

	public Integer getRows() {
		return rows;
	}

	public void setRows(Integer rows) {
		this.rows = rows;
	}

    public ConstantIssue.TextFormat getTextFormat() {
        return textFormat;
    }

    public void setTextFormat(ConstantIssue.TextFormat textFormat) {
        this.textFormat = textFormat;
    }

    public Integer getOrdr() {
		return ordr;
	}

	public void setOrdr(Integer ordr) {
		this.ordr = ordr;
	}

    public Integer getOrgId() {
        return orgId;
    }

    public void setOrgId(Integer orgId) {
        this.orgId = orgId;
    }

    public ConstantIssue.IssueInput getInput() {
        return input;
    }

    public void setInput(ConstantIssue.IssueInput input) {
        this.input = input;
    }

}
