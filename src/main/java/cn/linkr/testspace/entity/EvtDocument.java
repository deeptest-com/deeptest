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
@Table(name = "evt_document")
public class EvtDocument extends BaseEntity {
	private static final long serialVersionUID = 2390019554025797778L;
	private String title;
    @Column(name = "msg", length = 10000)
    private String descr;
    private String uri;

    @Enumerated(EnumType.STRING)
    private DocType docType = EvtDocument.DocType.file;

    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "event_id", insertable = false, updatable = false)
    private EvtEvent event;

    @Column(name = "event_id")
    private Long eventId;

    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "client_id", insertable = false, updatable = false)
    private EvtClient client;

    @Column(name = "client_id")
    private Long clientId;

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

	public DocType getDocType() {
		return docType;
	}

	public void setDocType(DocType docType) {
		this.docType = docType;
	}

	public String getTitle() {
		return title;
	}

	public void setTitle(String title) {
		this.title = title;
	}

	public EvtClient getClient() {
		return client;
	}

	public void setClient(EvtClient client) {
		this.client = client;
	}

	public Long getClientId() {
		return clientId;
	}

	public void setClientId(Long clientId) {
		this.clientId = clientId;
	}
}
