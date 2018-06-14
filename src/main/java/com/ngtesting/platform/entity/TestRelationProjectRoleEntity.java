package com.ngtesting.platform.entity;

import javax.persistence.*;

@Entity
@Table(name = "tst_r_project_role_entity")
public class TestRelationProjectRoleEntity extends BaseEntity {
	private static final long serialVersionUID = 5513768856000982338L;

    @Column(name = "org_id")
    private Long orgId;

	@Column(name = "project_id")
	private Long projectId;

    @Column(name = "project_role_id")
    private Long projectRoleId;

    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "project_role_id", insertable = false, updatable = false)
    private TestProjectRoleForOrg projectRole;

    @Column(name = "entity_id")
    private Long entityId;

    @Enumerated(EnumType.STRING)
    EntityType type;

    public TestRelationProjectRoleEntity() {

    }
    public TestRelationProjectRoleEntity(Long orgId, Long projectId, Long entityId, Long projectRoleId, String type) {
        this.orgId = orgId;
        this.projectId = projectId;
        this.entityId = entityId;
        this.projectRoleId = projectRoleId;
        this.type = EntityType.valueOf(type);
    }

    public enum EntityType {
        user("user"),
        group("group");

        EntityType(String val) {
            this.val = val;
        }

        private String val;
        public String toString() {
            return val;
        }
    }

    public Long getOrgId() {
        return orgId;
    }

    public void setOrgId(Long orgId) {
        this.orgId = orgId;
    }

    public EntityType getType() {
        return type;
    }

    public void setType(EntityType type) {
        this.type = type;
    }

    public Long getProjectId() {
		return projectId;
	}

	public void setProjectId(Long projectId) {
		this.projectId = projectId;
	}

	public Long getEntityId() {
		return entityId;
	}
	public void setEntityId(Long entityId) {
		this.entityId = entityId;
	}
	public Long getProjectRoleId() {
		return projectRoleId;
	}
	public void setProjectRoleId(Long projectRoleId) {
		this.projectRoleId = projectRoleId;
	}
	public TestProjectRoleForOrg getProjectRole() {
		return projectRole;
	}
	public void setProjectRole(TestProjectRoleForOrg projectRole) {
		this.projectRole = projectRole;
	}

}
