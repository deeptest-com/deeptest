package com.ngtesting.platform.vo;


public class SysPrivilegeVo extends BaseVo {

	private static final long serialVersionUID = 1473573544481704295L;
	private String code;
	private String name;
	private String action;
    private String descr;

    public SysPrivilegeVo() {
	}

	public SysPrivilegeVo(Long id, String code, String action, String name, String descr) {
		this.id = id;
		this.code = code;
		this.action = action;
		this.name = name;
		this.descr = descr;
	}

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

	public String getCode() {
		return code;
	}

	public void setCode(String code) {
		this.code = code;
	}

	public String getAction() {
		return action;
	}

	public void setAction(String action) {
		this.action = action;
	}
    
}
