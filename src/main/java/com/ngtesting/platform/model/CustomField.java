package com.ngtesting.platform.model;

import com.ngtesting.platform.config.ConstantIssue;

import java.util.LinkedList;
import java.util.List;

public class CustomField extends BaseModel {
	private static final long serialVersionUID = 1279085732086560549L;

    private String colCode;
    private String label;
    private Boolean required = false;
    private Boolean readonly = false;
    private Boolean buildIn = false;
    private Boolean fullLine = false;

    private String type;
    private String input;
    private ConstantIssue.TextFormat textFormat = ConstantIssue.TextFormat.plain_text;
    private FieldApplyTo applyTo;

    private String descr;

	private List<CustomFieldOption> options = new LinkedList<>();

    private Integer rows = 3;

    private Integer ordr;
    private Integer orgId;

    public static enum FieldApplyTo {
        test_case("test_case"),
        test_result("test_result"),
        issue("issue");

        FieldApplyTo(String textVal) {
            this.textVal = textVal;
        }

        private String textVal;

        public String toString() {
            return textVal;
        }
    }

    public FieldApplyTo getApplyTo() {
        return applyTo;
    }

    public void setApplyTo(FieldApplyTo applyTo) {
        this.applyTo = applyTo;
    }

    public List<CustomFieldOption> getOptions() {
        return options;
    }

    public void setOptions(List<CustomFieldOption> options) {
        this.options = options;
    }

    public Boolean getRequired() {
        return required;
    }

    public void setRequired(Boolean required) {
        this.required = required;
    }

    public Boolean getBuildIn() {
        return buildIn;
    }

    public void setBuildIn(Boolean buildIn) {
        this.buildIn = buildIn;
    }

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


	public String getDescr() {
		return descr;
	}

	public void setDescr(String descr) {
		this.descr = descr;
	}

	public String getType() {
		return type;
	}

	public void setType(String type) {
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

    public String getInput() {
        return input;
    }

    public void setInput(String input) {
        this.input = input;
    }

    public Boolean getReadonly() {
        return readonly;
    }

    public void setReadonly(Boolean readonly) {
        this.readonly = readonly;
    }

    public Boolean getFullLine() {
        return fullLine;
    }

    public void setFullLine(Boolean fullLine) {
        this.fullLine = fullLine;
    }
}
