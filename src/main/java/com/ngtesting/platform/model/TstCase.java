package com.ngtesting.platform.model;

import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONObject;

import java.util.LinkedList;
import java.util.List;

public class TstCase extends BaseModel {
	private static final long serialVersionUID = -5955583523485410239L;

	private String name;

	private Integer typeId;
    private String typeName;
	private Integer priorityId;
    private String priorityName;

	private Integer estimate = 10;
	private String objective;
    private String descr;
	private Boolean isParent;
    private Integer ordr;
	private Integer pId;
	private Integer projectId;

	private Integer createById;
	private Integer updateById;

    private Boolean checked;

	private CaseContentType contentType = CaseContentType.steps;
	private String content;
	private String key;
	private Boolean reviewResult;

	private Integer level;

	private String extProp;

    public static enum CaseContentType {
        steps("steps"),
        richText("richText");

        CaseContentType(String textVal) {
            this.textVal = textVal;
        }

        private String textVal;

        public String toString() {
            return textVal;
        }
    }

    public JSONObject getJsonProp() {
        return JSON.parseObject(extProp);
    }
    public void setExtProp(String extProp) {
        this.extProp = extProp;
    }

    public Boolean getIsParent() {
        return isParent;
    }

    public void setIsParent(Boolean parent) {
        isParent = parent;
    }

    private List<TstCase> children = new LinkedList<>();
	private List<TstCaseStep> steps = new LinkedList<>();
	private List<TstCaseComments> comments = new LinkedList<>();
	private List<TstCaseHistory> histories = new LinkedList<>();

	private List<TstCaseAttachment> attachments = new LinkedList<>();

	public List<TstCaseAttachment> getAttachments() {
		return attachments;
	}

	public void setAttachments(List<TstCaseAttachment> attachments) {
		this.attachments = attachments;
	}

	public TstCase(String name, Integer pId, Integer projectId, Integer typeId, Integer priorityId,
				   Integer createById, Boolean isParent, Integer ordr) {
		this.name = name;
		this.pId = pId;
		this.projectId = projectId;
		this.typeId = typeId;
		this.priorityId = priorityId;
		this.createById = createById;
		this.isParent = isParent;
		this.ordr = ordr;
	}
	public TstCase() {
    }

	public List<TstCaseHistory> getHistories() {
		return histories;
	}

	public void setHistories(List<TstCaseHistory> histories) {
		this.histories = histories;
	}

	public List<TstCaseComments> getComments() {
		return comments;
	}

	public void setComments(List<TstCaseComments> comments) {
		this.comments = comments;
	}

	public Boolean getReviewResult() {
		return reviewResult;
	}

	public void setReviewResult(Boolean reviewResult) {
		this.reviewResult = reviewResult;
	}

	public String getKey() {
		return key;
	}

	public void setKey(String key) {
		this.key = key;
	}

	public CaseContentType getContentType() {
		return contentType;
	}

	public void setContentType(CaseContentType contentType) {
		this.contentType = contentType;
	}

	public String getContent() {
		return content;
	}

	public void setContent(String content) {
		this.content = content;
	}

	public Boolean getChecked() {
        return checked;
    }

    public void setChecked(Boolean checked) {
        this.checked = checked;
    }

    public Integer getCreateById() {
        return createById;
    }

    public void setCreateById(Integer createById) {
        this.createById = createById;
    }

    public Integer getUpdateById() {
        return updateById;
    }

    public void setUpdateById(Integer updateById) {
        this.updateById = updateById;
    }

    public List<TstCase> getChildren() {
        return children;
    }

    public void setChildren(List<TstCase> children) {
        this.children = children;
    }

    public Integer getProjectId() {
		return projectId;
	}

	public void setProjectId(Integer projectId) {
		this.projectId = projectId;
	}

	public Integer getOrdr() {
        return ordr;
    }

    public void setOrdr(Integer ordr) {
        this.ordr = ordr;
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
	public Integer getPriorityId() {
		return priorityId;
	}
	public void setPriorityId(Integer priorityId) {
		this.priorityId = priorityId;
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

	public List<TstCaseStep> getSteps() {
		return steps;
	}

	public void setSteps(List<TstCaseStep> steps) {
		this.steps = steps;
	}

	public Integer getLevel() {
		return level;
	}

	public void setLevel(Integer level) {
		this.level = level;
	}

    public String getTypeName() {
        return typeName;
    }

    public void setTypeName(String typeName) {
        this.typeName = typeName;
    }

    public String getPriorityName() {
        return priorityName;
    }

    public void setPriorityName(String priorityName) {
        this.priorityName = priorityName;
    }
}
