package com.ngtesting.platform.entity;

import javax.persistence.*;
import java.util.HashSet;
import java.util.Set;

@Entity
@Table(name = "tst_custom_field")
public class TestCustomField extends BaseEntity {
	private static final long serialVersionUID = -1940351858441687302L;
	
	public TestCustomField() {
		
	}
	public void TestCustomField(String code, String label, String descr, Integer rows, Boolean isGlobal, Boolean isRequired) {
		this.code = code;
        this.label = label;

		this.descr = descr;
		this.rows = rows;
		this.isGlobal = isGlobal;
		this.isRequired = isRequired;
	}

    private String code;
    private String label;
    private String descr;
	private String column;
    
    @Enumerated(EnumType.STRING)
    private FieldApplyTo applyTo;
    
    @Enumerated(EnumType.STRING)
    private FieldType type;
    
    @Enumerated(EnumType.STRING)
    private FieldFormat format;

    private Integer rows;
    private Boolean isGlobal;
    private Boolean isRequired;
    private Boolean isBuildIn;
    
    private Integer ordr;
    
    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "org_id", insertable = false, updatable = false)
    private TestOrg org;

    @Column(name = "org_id")
    private Long orgId;

	public String getColumn() {
		return column;
	}

	public void setColumn(String column) {
		this.column = column;
	}

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
    	step("step"),
    	result("result");

        private FieldType(String textVal) {
            this.textVal = textVal;
        }

        private String textVal;
        public String toString() {
            return textVal;
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
    }

    public String getLabel() {
        return label;
    }

    public void setLabel(String label) {
        this.label = label;
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

	public Boolean getGlobal() {
		return isGlobal;
	}

	public void setGlobal(Boolean global) {
		isGlobal = global;
	}

	public Boolean getRequired() {
		return isRequired;
	}

	public void setRequired(Boolean required) {
		isRequired = required;
	}

	public Boolean getBuildIn() {
		return isBuildIn;
	}

	public void setBuildIn(Boolean buildIn) {
		isBuildIn = buildIn;
	}

	public Integer getOrdr() {
		return ordr;
	}

	public void setOrdr(Integer ordr) {
		this.ordr = ordr;
	}
}
