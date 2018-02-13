package com.ngtesting.platform.entity;

import javax.persistence.*;

@Entity
@Table(name = "tst_custom_field_options")
public class TestCustomFieldOptions extends BaseEntity {
    private static final long serialVersionUID = -5005831075946958149L;

    public TestCustomFieldOptions() {

	}
	public void TestCustomField(String code, String label, String descr, Integer order) {
		this.code = code;
        this.label = label;
		this.descr = descr;
		this.order = order;
	}

    private String code;
    private String label;
    private String descr;
    private Integer order;

    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "field_id", insertable = false, updatable = false)
    private TestCustomField field;

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

    public String getDescr() {
        return descr;
    }

    public void setDescr(String descr) {
        this.descr = descr;
    }

    public Integer getOrder() {
        return order;
    }

    public void setOrder(Integer order) {
        this.order = order;
    }

    public TestCustomField getField() {
        return field;
    }

    public void setField(TestCustomField field) {
        this.field = field;
    }
}
