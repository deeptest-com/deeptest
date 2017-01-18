package cn.linkr.events.entity;

import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.EnumType;
import javax.persistence.Enumerated;
import javax.persistence.FetchType;
import javax.persistence.JoinColumn;
import javax.persistence.ManyToOne;
import javax.persistence.Table;

@Entity
@Table(name = "evt_service")
public class EvtService extends BaseEntity {
	private static final long serialVersionUID = -3353122043605044415L;

	private String subject;

    @Column(name = "descr", length = 10000)
    private String descr;
    
    @Column(name = "html", length = 10000)
    private String html;

    @Enumerated(EnumType.STRING)
    private ServiceType type;

    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "event_id", insertable = false, updatable = false)
    private EvtEvent event;

    @Column(name = "event_id")
    private Long eventId;

    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "creator_id", insertable = false, updatable = false)
    private EvtClient creator;

    @Column(name = "creator_id")
    private Long creatorId;

    public static enum ServiceType {
        taxi("taxi", "叫车"),
        accommodation("accommodation", "住宿"),
        food("food", "餐饮"),
        wifi("wifi", "WIFI"),
        shopping("shopping", "购物"),
        print("print", "打印");

        private ServiceType(String textVal, String name) {
            this.textVal = textVal;
            this.name = name;
        }

        private String textVal;
        private String name;
        public String toString() {
            return textVal;
        }
        public String getName() {
            return name;
        }
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

    public ServiceType getType() {
        return type;
    }

    public void setType(ServiceType type) {
        this.type = type;
    }

    public EvtEvent getEvent() {
        return event;
    }

    public void setEvent(EvtEvent event) {
        this.event = event;
    }

    public Long getEventId() {
        return eventId;
    }

    public void setEventId(Long eventId) {
        this.eventId = eventId;
    }

    public EvtClient getCreator() {
        return creator;
    }

    public void setCreator(EvtClient creator) {
        this.creator = creator;
    }

    public Long getCreatorId() {
        return creatorId;
    }

    public void setCreatorId(Long creatorId) {
        this.creatorId = creatorId;
    }

	public String getHtml() {
		return html;
	}

	public void setHtml(String html) {
		this.html = html;
	}
}
