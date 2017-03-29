package com.ngtesting.platform.vo;

import java.io.Serializable;
import java.util.Date;
import java.util.LinkedHashSet;
import java.util.LinkedList;

import javax.persistence.Column;

public class TestCaseTreeVo implements Serializable {
	private static final long serialVersionUID = 1375843844627636495L;
	
	private Long id;
	private String value;
	private Long pid;
	
	private String type;
    
	LinkedHashSet<TestCaseTreeVo> children = new LinkedHashSet<TestCaseTreeVo>();
    
    public TestCaseTreeVo(Long id, String value, String type, Long pid) {
		this.id = id;
		this.value = value;
		this.type = type;
		this.pid = pid;
	}

	public Long getId() {
		return id;
	}

	public void setId(Long id) {
		this.id = id;
	}

	public LinkedHashSet<TestCaseTreeVo> getChildren() {
		return children;
	}

	public void setChildren(LinkedHashSet<TestCaseTreeVo> children) {
		this.children = children;
	}

	public String getValue() {
		return value;
	}

	public void setValue(String value) {
		this.value = value;
	}

	public Long getPid() {
		return pid;
	}

	public void setPid(Long pid) {
		this.pid = pid;
	}

	public String getType() {
		return type;
	}

	public void setType(String type) {
		this.type = type;
	}

}
