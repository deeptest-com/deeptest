package com.ngtesting.platform.vo;


import java.io.Serializable;
import java.util.List;
import java.util.Map;

public class IsuJqlFilter implements Serializable {
    private static final long serialVersionUID = -5923944030125754321L;

    private String code;
    private String label;
    private String type;
    private String input;
//    Map values;
    private List<String> operators;
    private Boolean buildIn;
    private Boolean display;

    public IsuJqlFilter(Map field) {
        this.code = field.get("colCode").toString();
        this.label = field.get("label").toString();
        this.type = field.get("type").toString();
        this.input = field.get("input").toString();
        this.buildIn = "1".equals(field.get("buildIn").toString());
        this.display = field.get("defaultShowInFilters") != null?
                Boolean.valueOf(field.get("defaultShowInFilters").toString()): null;
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

//    public Map getValues() {
//        return values;
//    }
//
//    public void setValues(Map values) {
//        this.values = values;
//    }

    public List<String> getOperators() {
        return operators;
    }

    public void setOperators(List<String> operators) {
        this.operators = operators;
    }

    public Boolean getDisplay() {
        return display;
    }

    public void setDisplay(Boolean display) {
        this.display = display;
    }
}
