package com.ngtesting.platform.vo;


public class OrganizerVo extends BaseVo {
	private static final long serialVersionUID = -5512498247399666456L;
	private String name;
    private String url;
    private String type;
    
	public String getUrl() {
		return url;
	}
	public void setUrl(String url) {
		this.url = url;
	}
	public String getName() {
		return name;
	}
	public void setName(String name) {
		this.name = name;
	}
	public String getType() {
		return type;
	}
	public void setType(String type) {
		this.type = type;
	}

}
