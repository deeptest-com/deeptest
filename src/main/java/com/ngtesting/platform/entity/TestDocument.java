package com.ngtesting.platform.entity;

import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.EnumType;
import javax.persistence.Enumerated;
import javax.persistence.FetchType;
import javax.persistence.JoinColumn;
import javax.persistence.ManyToOne;
import javax.persistence.Table;

@Entity
@Table(name = "tst_document")
public class TestDocument extends BaseEntity {
	private static final long serialVersionUID = 2390019554025797778L;
	private String title;
    @Column(name = "descr", length = 10000)
    private String descr;
    private String uri;

    @Enumerated(EnumType.STRING)
    private DocType docType = TestDocument.DocType.file;

    @Column(name = "event_id")
    private Long eventId;

    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "user_id", insertable = false, updatable = false)
    private TestUser user;

    @Column(name = "user_id")
    private Long userId;

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

    public TestUser getUser() {
        return user;
    }

    public void setUser(TestUser user) {
        this.user = user;
    }

    public Long getUserId() {
        return userId;
    }

    public void setUserId(Long userId) {
        this.userId = userId;
    }
}
