package com.ngtesting.platform.vo;


import java.util.List;
import java.util.Map;

public class IsuJqlFilter {
    private static final long serialVersionUID = -5923944030125754321L;

    String id;
    String label;
    String type;
    String input;
    Map values;
    List<String> operators;

    public IsuJqlFilter(String id, String label, Map values) {
        this.id = id;
        this.label = label;
        this.values = values;
    }

    public String getId() {
        return id;
    }

    public void setId(String id) {
        this.id = id;
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

    public Map getValues() {
        return values;
    }

    public void setValues(Map values) {
        this.values = values;
    }

    public List<String> getOperators() {
        return operators;
    }

    public void setOperators(List<String> operators) {
        this.operators = operators;
    }
}
