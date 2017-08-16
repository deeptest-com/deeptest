package com.ngtesting.platform.entity;

import javax.persistence.*;

@Entity
@Table(name = "tst_r_project_role_entity")
public class TestRelationProjectRoleEntity extends BaseEntity {
	private static final long serialVersionUID = 5513768856000982338L;

    private String projectRoleName;
    private String entityName;

	@Column(name = "project_id")
	private Long projectId;
	
    @Column(name = "project_role_id")
    private Long projectRoleId;
    
    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "project_role_id", insertable = false, updatable = false)
    private TestProjectRole projectRole;

    @Column(name = "entity_id")
    private Long entityId;

    @Enumerated(EnumType.STRING)
    EntityType type;

    public TestRelationProjectRoleEntity() {

    }
    public TestRelationProjectRoleEntity(Long projectId, Long entityId, Long projectRoleId,
                                         String projectRoleName, String entityName, String type) {
        this.projectId = projectId;
        this.entityId = entityId;
        this.projectRoleId = projectRoleId;
        this.projectRoleName = projectRoleName;
        this.entityName = entityName;
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

    public EntityType getType() {
        return type;
    }

    public void setType(EntityType type) {
        this.type = type;
    }

    public String getProjectRoleName() {
        return projectRoleName;
    }

    public void setProjectRoleName(String projectRoleName) {
        this.projectRoleName = projectRoleName;
    }

    public String getEntityName() {
        return entityName;
    }

    public void setEntityName(String entityName) {
        this.entityName = entityName;
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
	public TestProjectRole getProjectRole() {
		return projectRole;
	}
	public void setProjectRole(TestProjectRole projectRole) {
		this.projectRole = projectRole;
	}
    
}
