package com.ngtesting.platform.entity;

import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.Table;
import java.util.Date;

@Entity
@Table(name = "tst_verify_code")
public class TestVerifyCode extends BaseEntity {
	private static final long serialVersionUID = -7404422286952704677L;

	private String code;
	private Date expireTime;

	@Column(name="ref_id")
	private Long refId;

	public String getCode() {
		return code;
	}

	public void setCode(String code) {
		this.code = code;
	}

	public Date getCreateTime() {
		return createTime;
	}

	public void setCreateTime(Date createTime) {
		this.createTime = createTime;
	}

	public Date getExpireTime() {
		return expireTime;
	}

	public void setExpireTime(Date expireTime) {
		this.expireTime = expireTime;
	}

	public Long getRefId() {
		return refId;
	}

	public void setRefId(Long refId) {
		this.refId = refId;
	}

}
