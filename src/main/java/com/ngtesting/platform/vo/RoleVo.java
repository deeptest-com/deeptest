package com.ngtesting.platform.vo;


public class RoleVo extends BaseVo {
	private static final long serialVersionUID = 3252975025180725858L;
	
	private String code;
    private String name;
    private String descr;
    
    private Long companyId;

	public String getCode() {
		return code;
	}

	public void setCode(String code) {
		this.code = code;
	}

	public String getName() {
		return name;
	}

	public void setName(String name) {
		this.name = name;
	}

	public Long getCompanyId() {
		return companyId;
	}

	public void setCompanyId(Long companyId) {
		this.companyId = companyId;
	}

	public String getDescr() {
		return descr;
	}

	public void setDescr(String descr) {
		this.descr = descr;
	}
    
    

}
