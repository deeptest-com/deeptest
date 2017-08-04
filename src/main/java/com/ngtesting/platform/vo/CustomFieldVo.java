package com.ngtesting.platform.vo;

import com.ngtesting.platform.entity.TestCustomField;

public class CustomFieldVo extends BaseVo {
	private static final long serialVersionUID = -2071266644244632484L;

    private String code;
    private String label;
	private String column;

    private String descr;
    private String applyTo;
    private String type;

    private Integer rows = 3;
    private Boolean isRequired;
    private String format = TestCustomField.FieldFormat.plain_text.toString();
    private Boolean isGlobal = true;
    private Boolean isBuildIn = false;
    
    private Integer ordr;

	public String getColumn() {
		return column;
	}

	public void setColumn(String column) {
		this.column = column;
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

	public Boolean getIsGlobal() {
		return isGlobal;
	}

	public void setIsGlobal(Boolean isGlobal) {
		this.isGlobal = isGlobal;
	}

	public Boolean getIsBuildIn() {
		return isBuildIn;
	}

	public void setIsBuildIn(Boolean isBuildIn) {
		this.isBuildIn = isBuildIn;
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

	public Boolean getIsRequired() {
		return isRequired;
	}

	public void setIsRequired(Boolean isRequired) {
		this.isRequired = isRequired;
	}

	public String getFormat() {
		return format;
	}

	public void setFormat(String format) {
		this.format = format;
	}

	public Boolean getRequired() {
		return isRequired;
	}

	public void setRequired(Boolean required) {
		isRequired = required;
	}

	public Boolean getGlobal() {
		return isGlobal;
	}

	public void setGlobal(Boolean global) {
		isGlobal = global;
	}

	public Boolean getBuildIn() {
		return isBuildIn;
	}

	public void setBuildIn(Boolean buildIn) {
		isBuildIn = buildIn;
	}

	public Integer getOrdr() {
		return ordr;
	}

	public void setOrdr(Integer ordr) {
		this.ordr = ordr;
	}
}
