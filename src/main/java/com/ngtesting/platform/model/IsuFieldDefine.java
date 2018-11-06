package com.ngtesting.platform.model;


import com.ngtesting.platform.config.ConstantIssue;

public class IsuFieldDefine extends BaseModel {
	private static final long serialVersionUID = 8734289343612127207L;

	private String code;
	private String label;

	private ConstantIssue.IssueType type;
	private ConstantIssue.IssueInput input;

	private Boolean defaultShowInFilters;
	private Integer filterOrdr;

	private Boolean defaultShowInColumns;
	private Integer columnOrdr;

	private Boolean defaultShowInPage;
	private Integer elemOrdr;
    private Boolean fullLine;

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

    public Boolean getDefaultShowInPage() {
        return defaultShowInPage;
    }

    public void setDefaultShowInPage(Boolean defaultShowInPage) {
        this.defaultShowInPage = defaultShowInPage;
    }

    public Integer getElemOrdr() {
        return elemOrdr;
    }

    public void setElemOrdr(Integer elemOrdr) {
        this.elemOrdr = elemOrdr;
    }

    public Boolean getFullLine() {
        return fullLine;
    }

    public void setFullLine(Boolean fullLine) {
        this.fullLine = fullLine;
    }
}
