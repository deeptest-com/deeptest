package com.ngtesting.platform.entity;

import java.util.Date;

import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.FetchType;
import javax.persistence.JoinColumn;
import javax.persistence.ManyToOne;
import javax.persistence.Table;

@Entity
@Table(name = "evt_schedule_item")
public class EvtScheduleItem extends BaseEntity {
	private static final long serialVersionUID = -8733892292675111415L;

	private String subject;
    @Column(name = "descr", length = 10000)
    private String descr;

    private Date startDatetime;
    private Date endDatetime;
    
    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "event_id", insertable = false, updatable = false)
    private EvtEvent event;

    @Column(name = "event_id")
    private Long eventId;

    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "session_id", insertable = false, updatable = false)
    private EvtSession session;

    @Column(name = "session_id")
    private Long sessionId;

//    @ManyToOne(fetch = FetchType.LAZY)
//    @JoinColumn(name = "guest_id", insertable = false, updatable = false)
//    private EvtGuest guest;
//
//    @Column(name = "guest_id")
//    private Long guestId;

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

//    public EvtGuest getGuest() {
//        return guest;
//    }
//
//    public void setGuest(EvtGuest guest) {
//        this.guest = guest;
//    }
//
//    public Long getGuestId() {
//        return guestId;
//    }
//
//    public void setGuestId(Long guestId) {
//        this.guestId = guestId;
//    }

	public EvtSession getSession() {
		return session;
	}

	public void setSession(EvtSession session) {
		this.session = session;
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

	public Long getSessionId() {
		return sessionId;
	}

	public void setSessionId(Long sessionId) {
		this.sessionId = sessionId;
	}

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
}
