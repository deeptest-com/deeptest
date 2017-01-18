package cn.linkr.events.vo;


public class NoticeVo extends BaseVo {
	private static final long serialVersionUID = -2534782330543962540L;
	private String subject;
    private String descr;
    private String link;
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
	public String getLink() {
		return link;
	}
	public void setLink(String link) {
		this.link = link;
	}

}
