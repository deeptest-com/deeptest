package com.ngtesting.platform.model;

public class AiAsrLangModel extends BaseModel {

    private static final long serialVersionUID = 4967023618883171581L;
    private String name;
    private String descr;

    private Integer ordr;

    public Integer getOrdr() {
        return ordr;
    }

    public void setOrdr(Integer ordr) {
        this.ordr = ordr;
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
