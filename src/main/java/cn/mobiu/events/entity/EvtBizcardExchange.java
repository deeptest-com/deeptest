package cn.mobiu.events.entity;

import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.FetchType;
import javax.persistence.JoinColumn;
import javax.persistence.ManyToOne;
import javax.persistence.Table;

@Entity
@Table(name = "evt_business_card_exchange")
public class EvtBizcardExchange extends BaseEntity {
	private static final long serialVersionUID = -5246173253025911985L;

	private String remark;

    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "src_id", insertable = false, updatable = false)
    private EvtClient src;

    @Column(name = "src_id")
    private Long srcId;

    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "dist_id", insertable = false, updatable = false)
    private EvtClient dist;

    @Column(name = "dist_id")
    private Long distId;
    
    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "event_id", insertable = false, updatable = false)
    private EvtEvent event;

    @Column(name = "event_id")
    private Long eventId;

    public String getRemark() {
        return remark;
    }

    public void setRemark(String remark) {
        this.remark = remark;
    }

    public EvtClient getSrc() {
        return src;
    }

    public void setSrc(EvtClient src) {
        this.src = src;
    }

    public Long getSrcId() {
        return srcId;
    }

    public void setSrcId(Long srcId) {
        this.srcId = srcId;
    }

    public EvtClient getDist() {
        return dist;
    }

    public void setDist(EvtClient dist) {
        this.dist = dist;
    }

    public Long getDistId() {
        return distId;
    }

    public void setDistId(Long distId) {
        this.distId = distId;
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
}
