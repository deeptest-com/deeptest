package cn.mobiu.events.entity;

import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.EnumType;
import javax.persistence.Enumerated;
import javax.persistence.FetchType;
import javax.persistence.JoinColumn;
import javax.persistence.ManyToOne;
import javax.persistence.Table;

@Entity
@Table(name = "evt_around")
public class EvtAround extends BaseEntity {
	private static final long serialVersionUID = 8438417416461388880L;

	private String subject;

    @Column(name = "descr", length = 10000)
    private String descr;

    @Enumerated(EnumType.STRING)
    private AroundType type;

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

    public static enum AroundType {
        food("food"),
        accommodation("accommodation"),
        transportation("transportation"),
        tour("tour"),
        shopping("shopping"),
        entertainment("entertainment");

        private AroundType(String textVal) {
            this.textVal = textVal;
        }

        private String textVal;
        public String toString() {
            return textVal;
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

    public AroundType getType() {
        return type;
    }

    public void setType(AroundType type) {
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
}
