package cn.linkr.testspace.entity;

import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.EnumType;
import javax.persistence.Enumerated;
import javax.persistence.FetchType;
import javax.persistence.JoinColumn;
import javax.persistence.ManyToOne;
import javax.persistence.Table;

@Entity
@Table(name = "evt_organizer")
public class EvtOrganizer extends BaseEntity {
	private static final long serialVersionUID = 6632775542350557852L;
	private String name;
    private String url;
    
    @Enumerated(EnumType.STRING)
    private OrganizerType type;

    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "event_id", insertable = false, updatable = false)
    private EvtEvent event;
    
    public static enum OrganizerType {
    	host("host"), // 主办
    	organizer("organizer"), // 承办
    	co_organizer("co_organizer"), // 协办
    	
    	title_sponsor("title_sponsor"), // 冠名
    	sponsor("sponsor"); // 赞助
    	
        private OrganizerType(String textVal) {
            this.textVal = textVal;
        }

        private String textVal;
        public String toString() {
            return textVal;
        }
    }

    @Column(name = "event_id")
    private Long eventId;

	public String getName() {
		return name;
	}

	public void setName(String name) {
		this.name = name;
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

	public String getUrl() {
		return url;
	}

	public void setUrl(String url) {
		this.url = url;
	}

	public OrganizerType getType() {
		return type;
	}

	public void setType(OrganizerType type) {
		this.type = type;
	}

}
