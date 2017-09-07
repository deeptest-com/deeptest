package com.ngtesting.platform.vo;

public class TestCaseLightVo extends BaseVo {
	private static final long serialVersionUID = -5955583523485410239L;

	private String name;
    private String type;
    private Integer ordr;
	private Long pId;

    public TestCaseLightVo() {
        super();
    }

    public TestCaseLightVo(Long id, Long pId, String name) {
        super();

		this.id = id;
		this.pId = pId;
        this.name = name;
    }

    public Integer getOrdr() {
        return ordr;
    }

    public void setOrdr(Integer ordr) {
        this.ordr = ordr;
    }

    public String getType() {
        return this.type;
    }

    public void setType(String type) {
		this.type = type;
    }

    public String getName() {
		return name;
	}

	public void setName(String name) {
		this.name = name;
	}

	public Long getpId() {
		return pId;
	}

	public void setpId(Long pId) {
		this.pId = pId;
	}


}
