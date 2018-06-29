package com.ngtesting.platform.model;

import java.util.Date;
import java.util.LinkedList;
import java.util.List;

public class TstScheduleItem extends BaseModel {
	private static final long serialVersionUID = 777381642582610049L;
	private String name; // for session
	private String host; // for session

	private String subject;
    private String descr;
    private String address;
    private Date startDatetime;
    private Date endDatetime;

    private String startDatetimeStr;
    private String endDatetimeStr;

    private String startDate;
    private String endDate;
    private String startTime;
    private String endTime;

    private String itemType;

    private Long eventId;
    private Long sessionId;
//    private Long guestId;

    private List<TstScheduleItem> children = new LinkedList<TstScheduleItem>();

	public String getHost() {
		return host;
	}
	public void setHost(String host) {
		this.host = host;
	}
	public String getSubject() {
		return subject;
	}
	public void setSubject(String subject) {
		this.subject = subject;
	}
	public String getDescr() {
		return descr;
	}
	public void setDescr(String descr) {
		this.descr = descr;
	}

	public String getItemType() {
		return itemType;
	}
	public void setItemType(String itemType) {
		this.itemType = itemType;
	}
	public String getAddress() {
		return address;
	}
	public void setAddress(String address) {
		this.address = address;
	}

	public List<TstScheduleItem> getChildren() {
		return children;
	}
	public void setChildren(List<TstScheduleItem> children) {
		this.children = children;
	}

	public String getName() {
		return name;
	}
	public void setName(String name) {
		this.name = name;
	}
	public Long getEventId() {
		return eventId;
	}
	public void setEventId(Long eventId) {
		this.eventId = eventId;
	}
	public Long getSessionId() {
		return sessionId;
	}
	public void setSessionId(Long sessionId) {
		this.sessionId = sessionId;
	}
//	public Long getGuestId() {
//		return guestId;
//	}
//	public void setGuestId(Long guestId) {
//		this.guestId = guestId;
//	}
	public Date getStartDatetime() {
		return startDatetime;
	}
	public void setStartDatetime(Date startDatetime) {
		this.startDatetime = startDatetime;
	}
	public Date getEndDatetime() {
		return endDatetime;
	}
	public void setEndDatetime(Date endDatetime) {
		this.endDatetime = endDatetime;
	}
	public String getStartDate() {
		return startDate;
	}
	public void setStartDate(String startDate) {
		this.startDate = startDate;
	}
	public String getEndDate() {
		return endDate;
	}
	public void setEndDate(String endDate) {
		this.endDate = endDate;
	}
	public String getStartTime() {
		return startTime;
	}
	public void setStartTime(String startTime) {
		this.startTime = startTime;
	}
	public String getEndTime() {
		return endTime;
	}
	public void setEndTime(String endTime) {
		this.endTime = endTime;
	}
	public String getStartDatetimeStr() {
		return startDatetimeStr;
	}
	public void setStartDatetimeStr(String startDatetimeStr) {
		this.startDatetimeStr = startDatetimeStr;
	}
	public String getEndDatetimeStr() {
		return endDatetimeStr;
	}
	public void setEndDatetimeStr(String endDatetimeStr) {
		this.endDatetimeStr = endDatetimeStr;
	}

}
