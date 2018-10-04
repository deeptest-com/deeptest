package com.ngtesting.platform.model;


public class Document extends BaseModel {
	private static final long serialVersionUID = 5013730864709651144L;

	private String title;
    private String descr;
    private String uri;
    private String type;
    private Integer eventId;
    private Integer authorId;

	public static enum DocType {
		file("file"),
		audio("audio"),
		video("video"),
		image("image"),
		link("link");

		private DocType(String textVal) {
			this.textVal = textVal;
		}

		private String textVal;
		public String toString() {
			return textVal;
		}
	}

	public String getDescr() {
		return descr;
	}
	public void setDescr(String descr) {
		this.descr = descr;
	}
	public String getUri() {
		return uri;
	}
	public void setUri(String uri) {
		this.uri = uri;
	}
	public String getType() {
		return type;
	}
	public void setType(String type) {
		this.type = type;
	}
	public Integer getEventId() {
		return eventId;
	}
	public void setEventId(Integer eventId) {
		this.eventId = eventId;
	}
	public Integer getAuthorId() {
		return authorId;
	}
	public void setAuthorId(Integer authorId) {
		this.authorId = authorId;
	}
	public String getTitle() {
		return title;
	}
	public void setTitle(String title) {
		this.title = title;
	}

}
