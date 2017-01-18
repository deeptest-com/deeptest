package cn.linkr.events.entity;

import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.FetchType;
import javax.persistence.JoinColumn;
import javax.persistence.ManyToOne;
import javax.persistence.Table;

@Entity
@Table(name = "evt_qa")
public class EvtQa extends BaseEntity {
	private static final long serialVersionUID = 7155986999219056232L;

    @Column(name = "question", length = 10000)
    private String question;
    
    @Column(name = "answer", length = 10000)
	private String answer;
    
    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "event_id", insertable = false, updatable = false)
    private EvtEvent event;

    @Column(name = "event_id")
    private Long eventId;

    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "author_id", insertable = false, updatable = false)
    private EvtClient author;

    @Column(name = "author_id")
    private Long authorId;

    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "parent_id", insertable = false, updatable = false)
    private EvtQa parent;

    @Column(name = "parent_id")
    private Long parentId;

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

    public EvtClient getAuthor() {
        return author;
    }

    public void setAuthor(EvtClient author) {
        this.author = author;
    }

    public Long getAuthorId() {
        return authorId;
    }

    public void setAuthorId(Long authorId) {
        this.authorId = authorId;
    }

    public EvtQa getParent() {
        return parent;
    }

    public void setParent(EvtQa parent) {
        this.parent = parent;
    }

    public Long getParentId() {
        return parentId;
    }

    public void setParentId(Long parentId) {
        this.parentId = parentId;
    }

	public String getQuestion() {
		return question;
	}

	public void setQuestion(String question) {
		this.question = question;
	}

	public String getAnswer() {
		return answer;
	}

	public void setAnswer(String answer) {
		this.answer = answer;
	}

}
