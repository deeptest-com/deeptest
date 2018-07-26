package com.ngtesting.platform.model;


public class TstCaseAttachment extends BaseModel {
	private static final long serialVersionUID = 8077431586614975953L;

	private String name;
	private String descr;
	private String uri;

	private String docType;
	private Integer caseId;
	private Integer userId;

    public TstCaseAttachment() {
    }
    public TstCaseAttachment(String name, String path, Integer caseId, Integer userId) {
        super();
        this.name = name;
        this.uri = path;
        this.caseId = caseId;
        this.userId = userId;
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

	public String getUri() {
		return uri;
	}

	public void setUri(String uri) {
		this.uri = uri;
	}

	public String getDocType() {
		return docType;
	}

	public void setDocType(String docType) {
		this.docType = docType;
	}

    public Integer getCaseId() {
        return caseId;
    }

    public void setCaseId(Integer caseId) {
        this.caseId = caseId;
    }

    public Integer getUserId() {
		return userId;
	}

	public void setUserId(Integer userId) {
		this.userId = userId;
	}
}
