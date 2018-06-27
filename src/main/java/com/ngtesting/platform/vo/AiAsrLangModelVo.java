package com.ngtesting.platform.vo;

public class AiAsrLangModelVo extends BaseVo {

    private static final long serialVersionUID = 4967023618883171581L;
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
