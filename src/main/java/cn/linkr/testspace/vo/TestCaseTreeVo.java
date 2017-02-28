package cn.linkr.testspace.vo;

import java.io.Serializable;
import java.util.Date;
import java.util.LinkedHashSet;
import java.util.LinkedList;

import javax.persistence.Column;

public class TestCaseTreeVo implements Serializable {
	private static final long serialVersionUID = 1375843844627636495L;
	
	private Long id;
	private String value;
	private String path;
    
	LinkedHashSet<TestCaseTreeVo> children = new LinkedHashSet<TestCaseTreeVo>();
    
    public TestCaseTreeVo(Long id, String value, String path) {
		this.id = id;
		this.value = value;
		this.path = path;
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

	public String getPath() {
		return path;
	}

	public void setPath(String path) {
		this.path = path;
	}

	public String getValue() {
		return value;
	}

	public void setValue(String value) {
		this.value = value;
	}

}
