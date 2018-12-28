package com.ngtesting.platform.vo;


import java.io.Serializable;
import java.util.Map;

public class IsuJqlColumn implements Serializable {

	private static final long serialVersionUID = -7413029715796093478L;
	private String code;
	private String label;
    private String input;
	private String type;
    Boolean buildIn;

    private Boolean display;

	public IsuJqlColumn() {
	}

    public IsuJqlColumn(Map field) {
        this.code = field.get("colCode").toString();
        this.label = field.get("label").toString();
        this.type = field.get("type").toString();
        this.input = field.get("input").toString();
        this.buildIn = "1".equals(field.get("buildIn").toString());
        this.display = field.get("defaultShowInColumns") != null?
                Boolean.valueOf(field.get("defaultShowInColumns").toString()): null;
    }

    public String getInput() {
        return input;
    }

    public void setInput(String input) {
        this.input = input;
    }

    public Boolean getBuildIn() {
        return buildIn;
    }

    public void setBuildIn(Boolean buildIn) {
        this.buildIn = buildIn;
    }

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
}
