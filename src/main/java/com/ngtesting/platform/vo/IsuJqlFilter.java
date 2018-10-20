package com.ngtesting.platform.vo;


import com.ngtesting.platform.config.ConstantIssue;
import com.ngtesting.platform.model.IsuFieldDefine;

import java.io.Serializable;
import java.util.List;
import java.util.Map;

public class IsuJqlFilter implements Serializable {
    private static final long serialVersionUID = -5923944030125754321L;

    String code;
    String label;
    ConstantIssue.IssueFilterType type;
    ConstantIssue.IssueFilterInput input;
    Map values;
    List<String> operators;
    Boolean display;

    public IsuJqlFilter(IsuFieldDefine field) {
        this.code = field.getCode();
        this.label = field.getLabel();
        this.type = field.getType();
        this.input = field.getInput();
        this.display = field.getDefaultShowInFilters();

        if (ConstantIssue.IssueFilterInput.string.equals(type)) {
            this.operators = ConstantIssue.OperatorsForString;
        } else if (ConstantIssue.IssueFilterInput.date.equals(type)) {
            this.operators = ConstantIssue.OperatorsForDate;
        }
    }

    public IsuJqlFilter(IsuFieldDefine field, Map values) {
        this.code = field.getCode();
        this.label = field.getLabel();
        this.type = field.getType();
        this.input = field.getInput();
        this.display = field.getDefaultShowInFilters();

        this.values = values;
        if (ConstantIssue.IssueFilterInput.string.equals(input)) {
            this.operators = ConstantIssue.OperatorsForString;
        } else if (ConstantIssue.IssueFilterInput.select.equals(input)) {
            this.operators = ConstantIssue.OperatorsForSelect;
        } else if (ConstantIssue.IssueFilterInput.date.equals(input)) {
            this.operators = ConstantIssue.OperatorsForDate;
        }
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

    public ConstantIssue.IssueFilterType getType() {
        return type;
    }

    public void setType(ConstantIssue.IssueFilterType type) {
        this.type = type;
    }

    public ConstantIssue.IssueFilterInput getInput() {
        return input;
    }

    public void setInput(ConstantIssue.IssueFilterInput input) {
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

    public Boolean getDisplay() {
        return display;
    }

    public void setDisplay(Boolean display) {
        this.display = display;
    }
}
