package com.ngtesting.platform.entity;

import javax.persistence.*;

@Entity
@Table(name = "tst_case_prop")
public class TestCaseProp extends BaseEntity {

	private static final long serialVersionUID = -6655450841375690187L;
	private String name;
    private String value;

	@ManyToOne(fetch = FetchType.LAZY)
	@JoinColumn(name = "test_case_id", insertable = false, updatable = false)
	private TestCase testCase;

	@Column(name = "test_case_id")
	private Long testCaseId;

	@ManyToOne(fetch = FetchType.LAZY)
	@JoinColumn(name = "field_id", insertable = false, updatable = false)
	private TestCustomField field;

	@Column(name = "field_id")
	private Long fieldId;

	public Long getId() {
		return id;
	}

	public void setId(Long id) {
		this.id = id;
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

	public TestCase getTestCase() {
		return testCase;
	}

	public void setTestCase(TestCase testCase) {
		this.testCase = testCase;
	}

	public Long getTestCaseId() {
		return testCaseId;
	}

	public void setTestCaseId(Long testCaseId) {
		this.testCaseId = testCaseId;
	}

	public TestCustomField getField() {
		return field;
	}

	public void setField(TestCustomField field) {
		this.field = field;
	}

	public Long getFieldId() {
		return fieldId;
	}

	public void setFieldId(Long fieldId) {
		this.fieldId = fieldId;
	}
}
