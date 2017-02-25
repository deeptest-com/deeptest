package cn.linkr.testspace.vo;


public class AroundVo extends BaseVo {
	private static final long serialVersionUID = 2536439866508621469L;
	private String subject;
    private String descr;
    private String type;
    private Long eventId;
    
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
	public String getType() {
		return type;
	}
	public void setType(String type) {
		this.type = type;
	}
	public Long getEventId() {
		return eventId;
	}
	public void setEventId(Long eventId) {
		this.eventId = eventId;
	}

}
