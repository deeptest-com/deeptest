package com.ngtesting.platform.entity;

import javax.persistence.*;

@Entity
@Table(name = "dict")
public class Dict extends BaseEntity {
	private static final long serialVersionUID = 5841438770589651847L;

    private String category;
	private String phrase;

    @Column(name = "synonym_id")
    private Long synonymId;
    
    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "synonym_id", insertable = false, updatable = false)
    private Dict synonym;

    public Dict(String category, String phrase) {
        this.category = category;
        this.phrase = phrase;
    }

    public Dict(String category, String phrase, Long synonymId) {
        this.category = category;
        this.phrase = phrase;
        this.synonymId = synonymId;
    }

    public Long getSynonymId() {
        return synonymId;
    }

    public void setSynonymId(Long synonymId) {
        this.synonymId = synonymId;
    }

    public String getCategory() {
        return category;
    }

    public void setCategory(String category) {
        this.category = category;
    }

    public String getPhrase() {
        return phrase;
    }

    public void setPhrase(String phrase) {
        this.phrase = phrase;
    }

    public Dict getSynonym() {
        return synonym;
    }

    public void setSynonym(Dict synonym) {
        this.synonym = synonym;
    }

}
