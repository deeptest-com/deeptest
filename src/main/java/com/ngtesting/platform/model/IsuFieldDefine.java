package com.ngtesting.platform.model;


public class IsuFieldDefine extends BaseModel {
	private static final long serialVersionUID = 8734289343612127207L;

	private String colCode;
	private String label;

	private String type;
	private String input;

	private Boolean defaultShowInFilters;
	private Integer filterOrdr;

	private Boolean defaultShowInColumns;
	private Integer columnOrdr;

	private Boolean defaultShowInPage;
	private Integer elemOrdr;
	private Boolean readonly;

	private Boolean fullLine;

	public String getColCode() {
		return colCode;
	}

	public void setColCode(String colCode) {
		this.colCode = colCode;
	}

	public String getLabel() {
		return label;
	}

	public void setLabel(String label) {
		this.label = label;
	}

	public String getType() {
		return type;
	}

	public void setType(String type) {
		this.type = type;
	}

	public String getInput() {
		return input;
	}

	public void setInput(String input) {
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

	public Boolean getReadonly() {
		return readonly;
	}

	public void setReadonly(Boolean readonly) {
		this.readonly = readonly;
	}
}
