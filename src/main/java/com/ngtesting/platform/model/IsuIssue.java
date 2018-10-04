package com.ngtesting.platform.model;

import java.util.LinkedList;
import java.util.List;

public class IsuIssue extends BaseModel {
	private static final long serialVersionUID = 690125164201934534L;

    private String title;
    private String discr;

    private Integer projectId;

    private Integer typeId;
    private String typeName;

    private Integer statusId;
    private String statusName;

    private Integer priorityId;
    private String priorityName;

    private Integer creatorId;
    private String creatorName;
    private Integer assigneeId;
    private String assigneeName;

	private List<IsuAttachment> attachments = new LinkedList<>();
    private List<IsuComments> comments = new LinkedList<>();
    private List<IsuHistory> histories = new LinkedList<>();


}
