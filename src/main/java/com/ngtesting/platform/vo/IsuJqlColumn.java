package com.ngtesting.platform.vo;


import java.io.Serializable;

public class IsuJqlColumn implements Serializable {

	private static final long serialVersionUID = -7413029715796093478L;
	private String code;
	private String label;
	private Boolean display;
	private String type;

	public IsuJqlColumn() {
	}
	public IsuJqlColumn(String code, String label, String type, Boolean display) {
		this.code = code;
		this.label = label;
		this.type = type;
		this.display = display;
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
