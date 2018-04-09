package com.ngtesting.platform.entity;

import javax.persistence.*;

@Entity
@Table(name = "tst_custom_field_option")
public class IssueCustomFieldOption extends BaseEntity {
    private static final long serialVersionUID = -5005831075946958149L;

    public IssueCustomFieldOption() {

	}
	public void TestCustomField(String value, String label, String descr, Integer ordr) {
		this.value = value;
        this.label = label;
		this.descr = descr;
		this.ordr = ordr;
	}

    private String value;
    private String label;
    private String descr;
    private Integer ordr;

    @Column(name = "field_id")
    private Long fieldId;

    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "field_id", insertable = false, updatable = false)
    private TestCustomField field;

    public Long getFieldId() {
        return fieldId;
    }

    public void setFieldId(Long fieldId) {
        this.fieldId = fieldId;
    }

    public String getLabel() {
        return label;
    }

    public void setLabel(String label) {
        this.label = label;
    }

    public String getDescr() {
        return descr;
    }

    public void setDescr(String descr) {
        this.descr = descr;
    }

    public Integer getOrdr() {
        return ordr;
    }

    public void setOrdr(Integer ordr) {
        this.ordr = ordr;
    }

    public TestCustomField getField() {
        return field;
    }

    public void setField(TestCustomField field) {
        this.field = field;
    }

    public String getValue() {
        return value;
    }

    public void setValue(String value) {
        this.value = value;
    }
}
