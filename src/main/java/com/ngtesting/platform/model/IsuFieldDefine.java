package com.ngtesting.platform.model;


import com.ngtesting.platform.config.ConstantIssue;

public class IsuFieldDefine extends BaseModel {

	private static final long serialVersionUID = 8734289343612127207L;

	private String code;
	private String label;

	private ConstantIssue.IssueFilterType type;
	private ConstantIssue.IssueFilterInput input;

	private Boolean defaultShowInFilters;
	private Integer filterOrdr;

	private Boolean defaultShowInColumns;
	private Integer columnOrdr;

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

	public Boolean getDefaultShowInFilters() {
		return defaultShowInFilters;
	}

	public void setDefaultShowInFilters(Boolean defaultShowInFilters) {
		this.defaultShowInFilters = defaultShowInFilters;
	}

	public Integer getFilterOrdr() {
		return filterOrdr;
	}

	public void setFilterOrdr(Integer filterOrdr) {
		this.filterOrdr = filterOrdr;
	}

	public Boolean getDefaultShowInColumns() {
		return defaultShowInColumns;
	}

	public void setDefaultShowInColumns(Boolean defaultShowInColumns) {
		this.defaultShowInColumns = defaultShowInColumns;
	}

	public Integer getColumnOrdr() {
		return columnOrdr;
	}

	public void setColumnOrdr(Integer columnOrdr) {
		this.columnOrdr = columnOrdr;
	}
}
