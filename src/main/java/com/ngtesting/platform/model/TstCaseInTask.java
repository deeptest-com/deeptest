package com.ngtesting.platform.model;

import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONObject;

import java.util.Date;
import java.util.LinkedList;
import java.util.List;

public class TstCaseInTask extends BaseModel {
	private static final long serialVersionUID = -5955583523485410239L;

	private String name;

	private Integer typeId;
	private String typeName;
	private Integer priorityId;
	private String priorityName;

	private Integer estimate;
	private String objective;
    private String descr;
	private String result;
	private Integer pId;
	private Integer ordr;
	private Boolean isParent;

    private Integer entityId; // 真正是实体Id
	private Integer projectId;
	private Integer planId;
	private Integer taskId;

	private Integer createBy;
	private Integer updateBy;

	private Date exeTime;
	private Integer exeBy;

    private CaseExtStatus status;

	private String contentType;
	private String content;

	private String extProp;

	private String key;

	public static enum CaseExtStatus {
		untest("untest"),
		pass("pass"),
		fail("fail"),
		block("block");

		CaseExtStatus(String val) {
			this.val = val;
		}

		private String val;
		public String toString() {
			return val;
		}
	}

	private List<TstCaseStep> steps = new LinkedList<>();
	private List<TstCaseInTaskComments> comments = new LinkedList<>();
	private List<TstCaseInTaskAttachment> attachments = new LinkedList<>();
	private List<TstCaseInTaskIssue> issues = new LinkedList<>();

    private List<TstCaseInTaskHistory> histories = new LinkedList<>();

    public TstCaseInTask(Integer issueId, Integer caseInTaskId) {
        super();
    }

	public JSONObject getJsonProp() {
		return JSON.parseObject(extProp);
	}
	public void setExtProp(String extProp) {
		this.extProp = extProp;
	}

    public Integer getProjectId() {
		return projectId;
	}

	public void setProjectId(Integer projectId) {
		this.projectId = projectId;
	}

	public Integer getPlanId() {
		return planId;
	}

	public void setPlanId(Integer planId) {
		this.planId = planId;
	}

	public Integer getTaskId() {
		return taskId;
	}

	public void setTaskId(Integer taskId) {
		this.taskId = taskId;
	}

	public Integer getOrdr() {
		return ordr;
	}

	public void setOrdr(Integer ordr) {
		this.ordr = ordr;
	}

	public Boolean getIsParent() {
		return isParent;
	}

	public void setIsParent(Boolean parent) {
		isParent = parent;
	}

	public TstCaseInTask() {
        super();
    }

	public String getKey() {
		return key;
	}

	public void setKey(String key) {
		this.key = key;
	}

	public String getResult() {
		return result;
	}

	public void setResult(String result) {
		this.result = result;
	}

	public String getContentType() {
		return contentType;
	}

	public void setContentType(String contentType) {
		this.contentType = contentType;
	}

	public String getContent() {
		return content;
	}

	public void setContent(String content) {
		this.content = content;
	}

	public Integer getEntityId() {
        return entityId;
    }

    public void setEntityId(Integer entityId) {
        this.entityId = entityId;
    }

    public CaseExtStatus getStatus() {
        return status;
    }

    public void setStatus(CaseExtStatus status) {
        this.status = status;
    }

    public Integer getCreateBy() {
		return createBy;
	}

	public void setCreateBy(Integer createById) {
		this.createBy = createBy;
	}

	public Integer getUpdateBy() {
		return updateBy;
	}

	public void setUpdateBy(Integer updateBy) {
		this.updateBy = updateBy;
	}

	public String getName() {
		return name;
	}

	public void setName(String name) {
		this.name = name;
	}

	public Integer getpId() {
		return pId;
	}

	public void setpId(Integer pId) {
		this.pId = pId;
	}

	public Integer getTypeId() {
		return typeId;
	}

	public void setTypeId(Integer typeId) {
		this.typeId = typeId;
	}

	public String getTypeName() {
		return typeName;
	}

	public void setTypeName(String typeName) {
		this.typeName = typeName;
	}

	public Integer getPriorityId() {
		return priorityId;
	}

	public void setPriorityId(Integer priorityId) {
		this.priorityId = priorityId;
	}

	public String getPriorityName() {
		return priorityName;
	}

	public void setPriorityName(String priorityName) {
		this.priorityName = priorityName;
	}

	public Boolean getParent() {
		return isParent;
	}

	public void setParent(Boolean parent) {
		isParent = parent;
	}

	public Integer getEstimate() {
		return estimate;
	}

	public void setEstimate(Integer estimate) {
		this.estimate = estimate;
	}

	public String getDescr() {
		return descr;
	}

	public void setDescr(String descr) {
		this.descr = descr;
	}

	public String getObjective() {
		return objective;
	}

	public void setObjective(String objective) {
		this.objective = objective;
	}

//    public List<TestCasePropVo> getProps() {
//        return props;
//    }
//
//    public void setProps(List<TestCasePropVo> props) {
//        this.props = props;
//    }

	public List<TstCaseStep> getSteps() {
		return steps;
	}

	public void setSteps(List<TstCaseStep> steps) {
		this.steps = steps;
	}

    public List<TstCaseInTaskHistory> getHistories() {
        return histories;
    }

    public void setHistories(List<TstCaseInTaskHistory> histories) {
        this.histories = histories;
    }

	public List<TstCaseInTaskComments> getComments() {
		return comments;
	}

	public void setComments(List<TstCaseInTaskComments> comments) {
		this.comments = comments;
	}

	public List<TstCaseInTaskAttachment> getAttachments() {
		return attachments;
	}

	public void setAttachments(List<TstCaseInTaskAttachment> attachments) {
		this.attachments = attachments;
	}

	public List<TstCaseInTaskIssue> getIssues() {
		return issues;
	}

	public void setIssues(List<TstCaseInTaskIssue> issues) {
		this.issues = issues;
	}

	public Date getExeTime() {
		return exeTime;
	}

	public void setExeTime(Date exeTime) {
		this.exeTime = exeTime;
	}

	public Integer getExeBy() {
		return exeBy;
	}

	public void setExeBy(Integer exeBy) {
		this.exeBy = exeBy;
	}
}
