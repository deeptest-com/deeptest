package com.ngtesting.platform.model;

public class TstModule extends BaseModel {

    private static final long serialVersionUID = -2027022855596143227L;

    private String name;

    private String descr;
    private Integer projectId;

    private Integer ordr;

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

    public Integer getProjectId() {
        return projectId;
    }

    public void setProjectId(Integer projectId) {
        this.projectId = projectId;
    }

    public Integer getOrdr() {
        return ordr;
    }

    public void setOrdr(Integer ordr) {
        this.ordr = ordr;
    }
}
