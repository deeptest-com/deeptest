package com.ngtesting.platform.model;

import java.util.LinkedList;
import java.util.List;

public class IsuIssue extends BaseModel {
	private static final long serialVersionUID = 690125164201934534L;

    private String title;
    private String discr;

    private Integer projectId;

    private Integer isuTypeId;
    private String isuTypeName;

    private Integer isuStatusId;
    private String isuStatusName;

    private Integer isuPriorityId;
    private String isuPriorityName;

    private Integer creatorId;
    private String creatorName;
    private Integer assigneeId;
    private String assigneeName;

	private List<IsuAttachment> attachments = new LinkedList<>();
    private List<IsuComments> comments = new LinkedList<>();
    private List<IsuHistory> histories = new LinkedList<>();


}
