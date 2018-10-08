package com.ngtesting.platform.model;


import java.util.LinkedList;
import java.util.List;

public class IsuFilter extends BaseModel {
	private static final long serialVersionUID = -2233361748401561326L;

	private String title;
	private String descr;
    private IsuFilterType type;

    private List<IsuFilterSelectItem> selectItems = new LinkedList<>();

    public static enum IsuFilterType {
        text("text"),
        select("select"),
        date("date");

        IsuFilterType(String textVal) {
            this.textVal = textVal;
        }

        private String textVal;
        public String toString() {
            return textVal;
        }
    }

    public String getTitle() {
        return title;
    }

    public void setTitle(String title) {
        this.title = title;
    }

    public String getDescr() {
        return descr;
    }

    public void setDescr(String descr) {
        this.descr = descr;
    }

    public IsuFilterType getType() {
        return type;
    }

    public void setType(IsuFilterType type) {
        this.type = type;
    }

    public List<IsuFilterSelectItem> getSelectItems() {
        return selectItems;
    }

    public void setSelectItems(List<IsuFilterSelectItem> selectItems) {
        this.selectItems = selectItems;
    }
}
