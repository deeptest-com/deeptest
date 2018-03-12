package com.ngtesting.platform.entity;

import javax.persistence.*;

@Entity
@Table(name = "tst_case_step")
public class TestCaseStep extends BaseEntity {

	private static final long serialVersionUID = 1860144344660852859L;

	@Column(name = "opt", length = 10000)
	private String opt;
	@Column(name = "expect", length = 10000)
    private String expect;

	private Integer ordr;

	@ManyToOne(fetch = FetchType.LAZY)
	@JoinColumn(name = "test_case_id", insertable = false, updatable = false)
	private TestCase testCase;

	@Column(name = "test_case_id")
	private Long testCaseId;


	public TestCaseStep() {
		super();
	}
	public TestCaseStep(Long caseId, String opt, String expect, Integer ordr) {
		super();
		this.testCaseId = caseId;
		this.opt = opt;
		this.expect = expect;
		this.ordr = ordr;
	}

	public String getOpt() {
		return opt;
	}

	public void setOpt(String opt) {
		this.opt = opt;
	}

	public String getExpect() {
		return expect;
	}

	public void setExpect(String expect) {
		this.expect = expect;
	}

	public Integer getOrdr() {
		return ordr;
	}

	public void setOrdr(Integer ordr) {
		this.ordr = ordr;
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


}
