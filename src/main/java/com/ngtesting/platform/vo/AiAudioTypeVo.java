package com.ngtesting.platform.vo;

import com.ngtesting.platform.entity.BaseEntity;

public class AiAudioTypeVo extends BaseEntity {

    private static final long serialVersionUID = 6365977684914795772L;
    private String name;

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
