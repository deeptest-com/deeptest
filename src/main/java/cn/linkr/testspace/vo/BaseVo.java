package cn.linkr.testspace.vo;

import java.io.Serializable;
import java.util.Date;


public abstract class BaseVo implements Serializable {
	protected Long id;
    protected Boolean deleted;
    protected Boolean disabled;
    protected Integer version;
    
	public Long getId() {
		return id;
	}
	public void setId(Long id) {
		this.id = id;
	}

	public Boolean getDeleted() {
		return deleted;
	}
	public void setDeleted(Boolean deleted) {
		this.deleted = deleted;
	}
	public Boolean getDisabled() {
		return disabled;
	}
	public void setDisabled(Boolean disabled) {
		this.disabled = disabled;
	}
	public Integer getVersion() {
		return version;
	}
	public void setVersion(Integer version) {
		this.version = version;
	}
    
}