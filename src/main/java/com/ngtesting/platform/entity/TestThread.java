package com.ngtesting.platform.entity;

import javax.persistence.*;
import java.util.Date;

@Entity
@Table(name = "evt_thread")
public class TestThread extends BaseEntity {
	private static final long serialVersionUID = 5786241404855669174L;

    @Column(name = "content", length = 10000)
    private String content;

    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "author_id", insertable = false, updatable = false)
    private TestUser author;

    @Column(name = "author_id")
    private Long authorId;

    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "parent_id", insertable = false, updatable = false)
    private TestThread parent;

    @Column(name = "parent_id")
    private Long parentId;

    public TestThread(Long eventId, Long clientId, Long parentId,
                      String content) {

		this.authorId = clientId;
		this.content = content;

		if (parentId != null) {
			this.parentId = parentId;
		}
	}

	@Override
    public Date getCreateTime() {
        return createTime;
    }

    @Override
    public void setCreateTime(Date createTime) {
        this.createTime = createTime;
    }

    public Long getAuthorId() {
        return authorId;
    }

    public void setAuthorId(Long authorId) {
        this.authorId = authorId;
    }

    public TestThread getParent() {
        return parent;
    }

    public void setParent(TestThread parent) {
        this.parent = parent;
    }

    public Long getParentId() {
        return parentId;
    }

    public void setParentId(Long parentId) {
        this.parentId = parentId;
    }

	public String getContent() {
		return content;
	}

	public void setContent(String content) {
		this.content = content;
	}

    public TestUser getAuthor() {
        return author;
    }

    public void setAuthor(TestUser author) {
        this.author = author;
    }
}
