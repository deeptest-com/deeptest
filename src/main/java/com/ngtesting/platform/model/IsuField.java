package com.ngtesting.platform.model;


import java.io.Serializable;
import java.util.LinkedList;
import java.util.List;
import java.util.Map;

public class IsuField implements Serializable {

	private static final long serialVersionUID = 3168995179136496564L;

    private Integer fieldId;
	private String colCode;
	private String label;
    private String key;
    private String type;
    private String input;
    private Integer ordr;
    private Boolean fullLine;
    private Boolean required;
    private Boolean readonly;
    private Boolean buildIn;

    private Boolean display;

    private List<Map> options = new LinkedList<>();

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

    public Integer getOrdr() {
        return ordr;
    }

    public void setOrdr(Integer ordr) {
        this.ordr = ordr;
    }

    public Integer getFieldId() {
        return fieldId;
    }

    public void setFieldId(Integer fieldId) {
        this.fieldId = fieldId;
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

    public Boolean getRequired() {
        return required;
    }

    public void setRequired(Boolean required) {
        this.required = required;
    }

    public Boolean getReadonly() {
        return readonly;
    }

    public void setReadonly(Boolean readonly) {
        this.readonly = readonly;
    }

    public List<Map> getOptions() {
        return options;
    }

    public void setOptions(List<Map> options) {
        this.options = options;
    }

    public String getColCode() {
        return colCode;
    }

    public void setColCode(String colCode) {
        this.colCode = colCode;
    }

    public Boolean getBuildIn() {
        return buildIn;
    }

    public void setBuildIn(Boolean buildIn) {
        this.buildIn = buildIn;
    }
}
