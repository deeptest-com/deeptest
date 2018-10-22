package com.ngtesting.platform.model;

import java.util.LinkedList;
import java.util.List;

public class IsuCustomField extends BaseModel {
	private static final long serialVersionUID = 1279085732086560549L;

	private String code;
    private String label;
    private String myColumn;

    private String descr;
    private String applyTo;
    private String type;

	private List<IsuCustomFieldOption> options = new LinkedList<>();

    private Integer rows = 3;

    private String format = "plain_text";
    private Boolean required;
    private Boolean global = true;
    private Boolean buildIn = false;

    private Integer ordr;
    private Integer orgId;

	public static enum FieldType {
		string("string"),
		number("number"),
		dropdown("dropdown"),
		text("text");

//    	url("url"),
//
//    	radio("radio"),
//    	checkbox("checkbox"),
//
//
//    	multi_select("multi_select"),
//
//    	date("date"),
//
//    	user("user"),
//    	version("version"),
//    	step("step"),
//    	result("result");

		private FieldType(String textVal) {
			this.textVal = textVal;
		}

		private String textVal;
		public String toString() {
			return textVal;
		}
	}

    public static enum FieldFormat {
        plain_text("plain_text");
//        rich_text("rich_text");

        private FieldFormat(String textVal) {
            this.textVal = textVal;
        }

        private String textVal;
        public String toString() {
            return textVal;
        }
    }

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

	public String getType() {
		return type;
	}

	public void setType(String type) {
		this.type = type;
	}

    public String getApplyTo() {
		return applyTo;
	}

	public void setApplyTo(String applyTo) {
		this.applyTo = applyTo;
	}

	public Integer getRows() {
		return rows;
	}

	public void setRows(Integer rows) {
		this.rows = rows;
	}

	public String getFormat() {
		return format;
	}

	public void setFormat(String format) {
		this.format = format;
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
}
