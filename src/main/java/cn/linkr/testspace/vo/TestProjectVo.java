package cn.linkr.testspace.vo;

public class TestProjectVo extends BaseVo {
	private static final long serialVersionUID = -9069520320732281911L;
	private String name;
    private String descr;
    private Long companyId;
	public String getName() {
		return name;
	}
	public void setName(String name) {
		this.name = name;
	}
	public String getDescr() {
		return descr;
	}
	public void setDescr(String descr) {
		this.descr = descr;
	}
	public Long getCompanyId() {
		return companyId;
	}
	public void setCompanyId(Long companyId) {
		this.companyId = companyId;
	}
}
