package com.ngtesting.platform.vo;

import java.util.HashSet;
import java.util.Set;

import com.ngtesting.platform.entity.TestCustomField;
import com.ngtesting.platform.entity.TestProject;

public class CustomFieldVo extends BaseVo {
	private static final long serialVersionUID = -2071266644244632484L;
	
	private String name;
    private String label;
    private String code;
    private String descr;
    private String applyTo;
    private String type;

    private String configs;
    private Integer rows = 3;
    private Boolean isRequired;
    private String format = TestCustomField.FieldFormat.plain_text.toString();
    private Boolean isMulti;
    private Boolean isGlobal = true;
    private Boolean isBuildIn = false;
    
    private Integer displayOrder;

	public String getName() {
		return name;
	}

	public void setName(String name) {
		this.name = name;
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

	public String getConfigs() {
		return configs;
	}

	public void setConfigs(String configs) {
		this.configs = configs;
	}

	public Boolean getIsMulti() {
		return isMulti;
	}

	public void setIsMulti(Boolean isMulti) {
		this.isMulti = isMulti;
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

	public Integer getDisplayOrder() {
		return displayOrder;
	}

	public void setDisplayOrder(Integer displayOrder) {
		this.displayOrder = displayOrder;
	}

}
