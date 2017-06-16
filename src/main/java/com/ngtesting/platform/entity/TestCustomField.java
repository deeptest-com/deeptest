package com.ngtesting.platform.entity;

import java.util.HashSet;
import java.util.Set;

import javax.persistence.CascadeType;
import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.EnumType;
import javax.persistence.Enumerated;
import javax.persistence.FetchType;
import javax.persistence.JoinColumn;
import javax.persistence.JoinTable;
import javax.persistence.ManyToMany;
import javax.persistence.ManyToOne;
import javax.persistence.Table;

@Entity
@Table(name = "tst_custom_field")
public class TestCustomField extends BaseEntity {
	private static final long serialVersionUID = -1940351858441687302L;
	
	public TestCustomField() {
		
	}
	public void setValues(String name, String descr, Integer rows, Boolean isGlobal, Boolean isRequired) {
		this.name = name;
		this.descr = descr;
		this.rows = rows;
		this.isGlobal = isGlobal;
		this.isRequired = isRequired;
	}
	
	private String name;
    private String code;
    private String descr;
    
    @Enumerated(EnumType.STRING)
    private FieldApplyTo applyTo;
    
    @Enumerated(EnumType.STRING)
    private FieldType type;
    
    @Enumerated(EnumType.STRING)
    private FieldFormat format;

    private String configs;
    private Integer rows;
    private Boolean isGlobal;
    private Boolean isRequired;
    private Boolean isBuildIn;
    
    private Integer displayOrder;
    
    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "org_id", insertable = false, updatable = false)
    private TestOrg org;

    @Column(name = "org_id")
    private Long orgId;

    @ManyToMany(fetch = FetchType.LAZY, cascade = CascadeType.ALL)
	@JoinTable(name = "tst_r_custom_field_project", joinColumns = { 
			@JoinColumn(name = "custom_field_id", nullable = false, updatable = false) }, 
			inverseJoinColumns = { @JoinColumn(name = "project_id", 
					nullable = false, updatable = false) })
    private Set<TestProject> projectSet = new HashSet<TestProject>(0);
    
    public static enum FieldType {
    	string("string"),
    	text("text"),
    	number("number"),
    	url("url"),
    	
    	radio("radio"),
    	checkbox("checkbox"),
    	
    	dropdown("dropdown"),
    	multi_select("multi_select"),
    	
    	date("date"),

    	user("user"),
    	version("version"),
    	steps("steps"),
    	results("results");

        private FieldType(String textVal) {
            this.textVal = textVal;
        }

        private String textVal;
        public String toString() {
            return textVal;
        }
        
        public static FieldType getValue(String str) {
        	FieldType status = null;
        	switch(str) { 
            	case "number": status = FieldType.number; break;
            	case "string": status = FieldType.string; break;
            	case "text": status = FieldType.text; break;
            	
            	case "radio": status = FieldType.radio; break;
            	case "checkbox": status = FieldType.checkbox; break;
            	case "dropdown": status = FieldType.dropdown; break;
            	case "multi_select": status = FieldType.multi_select; break;
            	
            	case "date": status = FieldType.date; break;
            	case "url": status = FieldType.url; break;
            	
            	case "user": status = FieldType.user; break;
            	case "version": status = FieldType.version; break;
            	case "steps": status = FieldType.steps; break;
            	case "results": status = FieldType.results; break;
            }
        	
        	return status;
        }
    }
    
    public static enum FieldApplyTo {
    	test_case("test_case"),
        test_result("test_result");

        private FieldApplyTo(String textVal) {
            this.textVal = textVal;
        }

        private String textVal;
        public String toString() {
            return textVal;
        }
        
        public static FieldApplyTo getValue(String str) {
        	FieldApplyTo status = null;
        	switch(str) { 
            	case "test_case": status = FieldApplyTo.test_case; break;
            	case "test_result": status = FieldApplyTo.test_result; break;
            }
        	
        	return status;
        }
    }
    
    public static enum FieldFormat {
    	markdown("markdown"),
        plain_text("plain_text");

        private FieldFormat(String textVal) {
            this.textVal = textVal;
        }

        private String textVal;
        public String toString() {
            return textVal;
        }
        
        public static FieldFormat getValue(String str) {
        	FieldFormat status = null;
        	switch(str) { 
            	case "markdown": status = FieldFormat.markdown; break;
            	case "plain_text": status = FieldFormat.plain_text; break;
            }
        	
        	return status;
        }
    }

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

	public String getConfigs() {
		return configs;
	}

	public void setConfigs(String configs) {
		this.configs = configs;
	}

	public Boolean getIsGlobal() {
		return isGlobal;
	}

	public void setIsGlobal(Boolean isGlobal) {
		this.isGlobal = isGlobal;
	}

	public Boolean getIsBuildIn() {
		return isBuildIn;
	}

	public void setIsBuildIn(Boolean isBuildIn) {
		this.isBuildIn = isBuildIn;
	}

	public Set<TestProject> getProjectSet() {
		return projectSet;
	}

	public void setProjectSet(Set<TestProject> projectSet) {
		this.projectSet = projectSet;
	}

	public FieldApplyTo getApplyTo() {
		return applyTo;
	}

	public void setApplyTo(FieldApplyTo applyTo) {
		this.applyTo = applyTo;
	}

	public FieldType getType() {
		return type;
	}

	public void setType(FieldType type) {
		this.type = type;
	}

	public FieldFormat getFormat() {
		return format;
	}

	public void setFormat(FieldFormat format) {
		this.format = format;
	}

	public Integer getRows() {
		return rows;
	}

	public void setRows(Integer rows) {
		this.rows = rows;
	}

	public Boolean getIsRequired() {
		return isRequired;
	}

	public void setIsRequired(Boolean isRequired) {
		this.isRequired = isRequired;
	}

	public Integer getDisplayOrder() {
		return displayOrder;
	}

	public void setDisplayOrder(Integer displayOrder) {
		this.displayOrder = displayOrder;
	}

	public TestOrg getOrg() {
		return org;
	}

	public void setOrg(TestOrg org) {
		this.org = org;
	}

	public Long getOrgId() {
		return orgId;
	}

	public void setOrgId(Long orgId) {
		this.orgId = orgId;
	}
	
}
