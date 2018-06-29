package com.ngtesting.platform.model;

public class TstCaseLight extends BaseModel {
	private static final long serialVersionUID = -5955583523485410239L;

	private String name;
    private String type;
    private Integer ordr;
	private Integer pId;

    public TstCaseLight() {
        super();
    }

    public TstCaseLight(Integer id, Integer pId, String name) {
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

	public Integer getpId() {
		return pId;
	}

	public void setpId(Integer pId) {
		this.pId = pId;
	}


}
