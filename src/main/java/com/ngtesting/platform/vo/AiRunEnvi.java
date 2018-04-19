/**
  * Copyright 2017 bejson.com
  */
package com.ngtesting.platform.vo;

public class AiRunEnvi {
	private String site;
	private String testType;
	private String audioType;
	private String productId;
	private String aliasKey;
	private String resParam;
	private Integer startIndex;
	private Integer numbToRun;
	private String isFuse;

	public void setSite(String site) {
		this.site = site;
	}

	public String getSite() {
		return site;
	}
	public void setIsFuse(String isFuse) {
		this.isFuse = isFuse;
	}

	public String getIsFuse() {
		return isFuse;
	}
	public void setResParam(String resParam) {
		this.resParam = resParam;
	}

	public String getResParam() {
		return resParam;
	}

	public void setAliasKey(String aliasKey) {
		this.aliasKey = aliasKey;
	}

	public String getAliasKey() {
		return aliasKey;
	}

	public void setTestType(String testType) {
		this.testType = testType;
	}

	public String getTestType() {
		return testType;
	}

	public void setProductId(String productId) {
		this.productId = productId;
	}

	public String getProductId() {
		return productId;
	}

	public void setStartIndex(Integer startIndex) {
		this.startIndex = startIndex;
	}

	public Integer getStartIndex() {
		return startIndex;
	}

	public void setNumbToRun(Integer numbToRun) {
		this.numbToRun = numbToRun;
	}

	public Integer getNumbToRun() {
		return numbToRun;
	}

	public String getAudioType() {
		return audioType;
	}

	public void setAudioType(String autoType) {
		this.audioType = autoType;
	}
}
