package com.ngtesting.platform.vo;

public class TestCasePropVo extends BaseVo {

    private String code;
    private String label;
    private String value;

    private Long fieldId;
    private CustomFieldVo field;

    public TestCasePropVo(Long id, String code, String label, String value, Long fieldId) {
        super();
        this.id = id;
        this.code = code;
        this.label = label;
        this.value = value;

        this.fieldId = fieldId;
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

    public String getValue() {
        return value;
    }

    public void setValue(String value) {
        this.value = value;
    }

    public CustomFieldVo getField() {
        return field;
    }

    public void setField(CustomFieldVo field) {
        this.field = field;
    }

    public Long getFieldId() {
        return fieldId;
    }

    public void setFieldId(Long fieldId) {
        this.fieldId = fieldId;
    }
}
