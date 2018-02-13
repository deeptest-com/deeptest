package com.ngtesting.platform.entity;

import javax.persistence.*;
import java.util.HashSet;
import java.util.LinkedList;
import java.util.Set;

@Entity
@Table(name = "tst_custom_field")
public class TestCustomField extends BaseEntity {
	private static final long serialVersionUID = -1940351858441687302L;

	public TestCustomField() {

	}
	public void TestCustomField(String code, String label, String descr, Integer rows, Boolean global, Boolean required) {
		this.code = code;
        this.label = label;

		this.descr = descr;
		this.rows = rows;
		this.global = global;
		this.required = required;
	}

    private String code;
    private String label;
    private String descr;
	private String myColumn;

    @Enumerated(EnumType.STRING)
    private FieldApplyTo applyTo;

    @Enumerated(EnumType.STRING)
    private FieldType type;

    @Enumerated(EnumType.STRING)
    private FieldFormat format;

    private LinkedList<LinkedList<String>> options = new LinkedList();

    private Integer rows;
    private Boolean global;
    private Boolean required;
    private Boolean buildIn;

    private Integer ordr;

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
    	number("number"),
        dropdown("dropdown"),
        text("text");

//    	url("url"),
//
//    	radio("radio"),
//    	checkbox("checkbox"),
//
//
//    	multi_select("multi_select"),
//
//    	date("date"),
//
//    	user("user"),
//    	version("version"),
//    	step("step"),
//    	result("result");

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
        plain_text("plain_text"),
        rich_text("rich_text");

        private FieldFormat(String textVal) {
            this.textVal = textVal;
        }

        private String textVal;
        public String toString() {
            return textVal;
        }
    }

    public LinkedList<LinkedList<String>> getOptions() {
        return options;
    }

    public void setOptions(LinkedList<LinkedList<String>> options) {
        this.options = options;
    }

    public String getMyColumn() {
        return myColumn;
    }

    public void setMyColumn(String myColumn) {
        this.myColumn = myColumn;
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

    public Boolean getGlobal() {
        return this.global;
    }
    public void setGlobal(Boolean global) {
        this.global = global;
    }

    public Boolean getRequired() {
        return this.required;
    }
    public void setRequired(Boolean required) {
        this.required = required;
    }

    public Boolean getBuildIn() {
        return this.buildIn;
    }
    public void setBuildIn(Boolean buildIn) {
        this.buildIn = buildIn;
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

	public Integer getOrdr() {
		return ordr;
	}

	public void setOrdr(Integer ordr) {
		this.ordr = ordr;
	}
}
