package com.ngtesting.platform.vo;

import com.ngtesting.platform.entity.TestCustomField;

import java.util.LinkedList;
import java.util.List;

public class CustomFieldVo extends BaseVo {
	private static final long serialVersionUID = -2071266644244632484L;

    private String code;
    private String label;
    private String myColumn;

    private String descr;
    private String applyTo;
    private String type;

	private List<CustomFieldOptionVo> options = new LinkedList<>();

    private Integer rows = 3;

    private String format = TestCustomField.FieldFormat.plain_text.toString();
    private Boolean required;
    private Boolean global = true;
    private Boolean buildIn = false;

    private Integer ordr;

    public List<CustomFieldOptionVo> getOptions() {
        return options;
    }

    public void setOptions(List<CustomFieldOptionVo> options) {
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


}
