package com.ngtesting.platform.model;


public class IsuFilterSelectItem extends BaseModel {

	private static final long serialVersionUID = -1459784755438950248L;
	private String value;
	private String label;

	public String getValue() {
		return value;
	}

	public void setValue(String value) {
		this.value = value;
	}

	public String getLabel() {
		return label;
	}

	public void setLabel(String label) {
		this.label = label;
	}
}
