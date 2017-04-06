package com.ngtesting.platform.entity;

import java.util.HashSet;
import java.util.Set;

import javax.persistence.CascadeType;
import javax.persistence.Entity;
import javax.persistence.EnumType;
import javax.persistence.Enumerated;
import javax.persistence.FetchType;
import javax.persistence.JoinColumn;
import javax.persistence.JoinTable;
import javax.persistence.ManyToMany;
import javax.persistence.Table;

@Entity
@Table(name = "sys_case_exe_status")
public class SysCaseExeStatus extends BaseEntity {
	private static final long serialVersionUID = 4775052158868753948L;
	
	private String name;
    private String code;
    private String descr;
    private Integer displayOrder;
    private Boolean isFinal;
    
    private Boolean isBuildIn = false;

	public String getName() {
		return name;
	}

	public void setName(String name) {
		this.name = name;
	}

	public String getCode() {
		return code;
	}

	public void setCode(String code) {
		this.code = code;
	}

	public String getDescr() {
		return descr;
	}

	public void setDescr(String descr) {
		this.descr = descr;
	}

	public Integer getDisplayOrder() {
		return displayOrder;
	}

	public void setDisplayOrder(Integer displayOrder) {
		this.displayOrder = displayOrder;
	}

	public Boolean getIsFinal() {
		return isFinal;
	}

	public void setIsFinal(Boolean isFinal) {
		this.isFinal = isFinal;
	}

	public Boolean getIsBuildIn() {
		return isBuildIn;
	}

	public void setIsBuildIn(Boolean isBuildIn) {
		this.isBuildIn = isBuildIn;
	}

    
	
}
