package cn.mobiu.events.vo;

import java.util.Date;

public class RegisterRecordVo extends BaseVo {
	private static final long serialVersionUID = 3893837658002665944L;
	private String subject;
    private String descr;
    private Date registerTime;
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
	public Date getRegisterTime() {
		return registerTime;
	}
	public void setRegisterTime(Date registerTime) {
		this.registerTime = registerTime;
	}

}
