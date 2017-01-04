package cn.mobiu.events.vo;


public class ServiceVo extends BaseVo {
	private static final long serialVersionUID = -9008705460037884864L;
	private String subject;
    private String descr;
    private String type;
    private String typeName;
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
	public String getTypeName() {
		return typeName;
	}
	public void setTypeName(String typeName) {
		this.typeName = typeName;
	}

}
