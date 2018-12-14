package com.ngtesting.platform.model;


public class IsuAttachment extends BaseModel {
	private static final long serialVersionUID = -7769659393127598472L;

    private String name;
	private String descr;
	private String uri;

	private String docType;
	private Integer issueId;
	private Integer userId;

	public IsuAttachment() {
	}
	public IsuAttachment(String name, String path, Integer issueId, Integer userId) {
		super();
		this.name = name;
		this.uri = path;
		this.issueId = issueId;
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

    public Integer getIssueId() {
        return issueId;
    }

    public void setIssueId(Integer issueId) {
        this.issueId = issueId;
    }

    public Integer getUserId() {
		return userId;
	}

	public void setUserId(Integer userId) {
		this.userId = userId;
	}

}
