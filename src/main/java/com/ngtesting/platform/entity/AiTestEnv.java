package com.ngtesting.platform.entity;

import javax.persistence.*;

@Entity
@Table(name = "ai_test_env")
public class AiTestEnv extends BaseEntity {
    private static final long serialVersionUID = 5021996669215836192L;
    private String name;

	@Column(name = "descr", length = 1000)
    private String descr;

    private Integer displayOrder;

    public Integer getDisplayOrder() {
        return displayOrder;
    }

    public void setDisplayOrder(Integer displayOrder) {
        this.displayOrder = displayOrder;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public String getDescr() {
        return descr;
    }

    public void setDescr(String descr) {
        this.descr = descr;
    }

}
