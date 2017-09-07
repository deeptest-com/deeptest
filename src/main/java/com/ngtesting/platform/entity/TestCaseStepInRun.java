package com.ngtesting.platform.entity;

import javax.persistence.*;

@Entity
@Table(name = "tst_case_step_in_run")
public class TestCaseStepInRun extends BaseEntity {

	private static final long serialVersionUID = -4777916667107661049L;
	private String opt;
    private String expect;

	private Integer ordr;

	@ManyToOne(fetch = FetchType.LAZY)
	@JoinColumn(name = "test_case_in_run_id", insertable = false, updatable = false)
	private TestCaseInRun testCaseInRun;

	@Column(name = "test_case_in_run_id")
	private Long testCaseInRunId;

	public TestCaseInRun getTestCaseInRun() {
		return testCaseInRun;
	}

	public void setTestCaseInRun(TestCaseInRun testCaseInRun) {
		this.testCaseInRun = testCaseInRun;
	}

	public Long getTestCaseInRunId() {
		return testCaseInRunId;
	}

	public void setTestCaseInRunId(Long testCaseInRunId) {
		this.testCaseInRunId = testCaseInRunId;
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

}
