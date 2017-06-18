package com.ngtesting.platform.vo;

public class TestCasePropVo extends BaseVo {

    private String name;
    private String value;

    private CustomFieldVo field;

    public TestCasePropVo(Long id, String name, String value) {
        super();
        this.id = id;
        this.name = name;
        this.value = value;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
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
}
