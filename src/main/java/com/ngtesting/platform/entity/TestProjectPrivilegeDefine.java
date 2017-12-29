package com.ngtesting.platform.entity;

import javax.persistence.Entity;
import javax.persistence.EnumType;
import javax.persistence.Enumerated;
import javax.persistence.Table;

@Entity
@Table(name = "tst_project_privilege_define")
public class TestProjectPrivilegeDefine extends BaseEntity {
	private static final long serialVersionUID = -5510206858644860272L;

	@Enumerated(EnumType.STRING)
    private ProjectPrivilegeCode code;

    private String name;
    private String descr;

    @Enumerated(EnumType.STRING)
    private PrivilegeAction action;

    public static enum ProjectPrivilegeCode {
    	req("req"),
    	cases("cases"),
    	plan("plan"),
    	round("round"),
    	result("result"),
    	report("report");

        private ProjectPrivilegeCode(String textVal) {
            this.textVal = textVal;
        }

        private String textVal;
        public String toString() {
            return textVal;
        }
    }

    public static enum PrivilegeAction {
		create("create"),
        update("update"),
        remove("remove"),
        close("close");

        private PrivilegeAction(String textVal) {
            this.textVal = textVal;
        }

        private String textVal;
        public String toString() {
            return textVal;
        }

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

	public ProjectPrivilegeCode getCode() {
		return code;
	}

	public void setCode(ProjectPrivilegeCode code) {
		this.code = code;
	}

	public PrivilegeAction getAction() {
		return action;
	}

	public void setAction(PrivilegeAction action) {
		this.action = action;
	}

}
