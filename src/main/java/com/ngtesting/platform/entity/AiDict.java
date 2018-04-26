package com.ngtesting.platform.entity;

import javax.persistence.Entity;
import javax.persistence.Table;

@Entity
@Table(name = "ai_dict")
public class AiDict extends BaseEntity {
    private static final long serialVersionUID = 8036841692633236687L;

    private String skillId;
    private String category;
	private String phrase;
    private String synonym;

    public AiDict() {
    }

    public AiDict(String skillId, String category, String phrase, String synonym) {
    	this.skillId = skillId;
        this.category = category;
        this.phrase = phrase;
        this.synonym = synonym;
    }

    public String getSkillId() {
		return skillId;
	}

	public void setSkillId(String skillId) {
		this.skillId = skillId;
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

    public String getSynonym() {
        return synonym;
    }

    public void setSynonym(String synonym) {
        this.synonym = synonym;
    }

}
