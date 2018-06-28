package com.ngtesting.platform.vo;

import com.ngtesting.platform.config.Constant;

import java.util.List;

public class Page<T> {

    /**
     * 下一页步长
     */
    private static final int NEXT_PAGE_INDEX = 2;

    public Page() {

    }

    @SuppressWarnings("rawtypes")
    public Page(int start, int limit) {
        this.start = start;
        this.limit = limit;
    }

    @SuppressWarnings("rawtypes")
    public Page(int start, int limit, int total, List items) {
        this.start = start;
        this.total = total;
        this.items = items;
        this.limit = limit;
    }

    /**
     * 分页大小
     */
    private int limit = Constant.PAGE_SIZE;

    /**
     * 分页开始索引
     */
    private int start;

    /**
     * 总记录数
     */
    private int total;

    /**
     * 分页记录
     */
    @SuppressWarnings("rawtypes")
    private List items;

    /**
     * 计算总页数.
     *
     * @return 返回int
     */
    public int getTotalPages() {
        if (total == -1) {
            return -1;
        }

        int count = total / limit;
        if (total % limit > 0) {
            count++;
        }
        return count;
    }

    /**
     * 是否还有下一页.
     *
     * @return boolea
     */
    public boolean hasNextPage() {
        return getTotalPages() > start / limit + 1;
    }

    /**
     * 获取页数,从1开始计数
     *
     * @return int
     */
    public int getNextPage() {
        if (hasNextPage()) {
            return start / limit + NEXT_PAGE_INDEX;
        } else {
            return start / limit + 1;
        }
    }

    /**
     * 是否还有上一页.
     *
     * @return boolean
     */
    public boolean hasPrePage() {
        return start > 0;
    }

    /**
     * 返回上页的页号,序号从1开始.
     *
     * @return 返回int
     */
    public int getPrePage() {
        if (hasPrePage()) {
            return start / limit;
        } else {
            return 1;
        }
    }

    public int getCurrentPage() {
        if (hasPrePage()) {
            return start / limit + 1;
        } else {
            return 1;
        }
    }

    public int getFirstPage() {
        return 1;
    }

    public int getLastPage() {
        return getTotalPages();
    }

    public int getLimit() {
        return limit;
    }

    public void setLimit(int limit) {
        this.limit = limit;
    }

    public int getStart() {
        return start;
    }

    public void setStart(int start) {
        this.start = start;
    }

    public int getTotal() {
        return total;
    }

    public void setTotal(int total) {
        this.total = total;
    }

    @SuppressWarnings("rawtypes")
    public List getItems() {
        return items;
    }

    @SuppressWarnings("rawtypes")
    public void setItems(List items) {
        this.items = items;
    }

}
