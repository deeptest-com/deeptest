package com.ngtesting.platform.model;


public class IsuLinkReason extends BaseModel {
    private static final long serialVersionUID = -8861420715111683476L;

    private String label;
    private String value;

    public String getLabel() {
        return label;
    }

    public void setLabel(String label) {
        this.label = label;
    }

    public String getValue() {
        return value;
    }

    public void setValue(String value) {
        this.value = value;
    }
}
