package com.ngtesting.platform.vo;

import com.ngtesting.platform.util.NumberUtil;

/**
 * 文件上传进度信息
 *
 * @author xuxiang
 * @version $Id$
 * @since
 * @see
 */
public class ProgressEntity {
	/** 已读字节 **/
	private long pBytesRead = 0L;
	/** 总字节 **/
	private long pContentLength = 0L;
	private int pItems;
	/** 已读百分比 **/
	private String percent;

	public String getPercent() {
		percent = NumberUtil.getPercent(pBytesRead, pContentLength);
		return percent;
	}

	public void setPercent(String percent) {
		this.percent = percent;
	}

	public long getpBytesRead() {
		return pBytesRead;
	}

	public void setpBytesRead(long pBytesRead) {
		this.pBytesRead = pBytesRead;
	}

	public long getpContentLength() {
		return pContentLength;
	}

	public void setpContentLength(long pContentLength) {
		this.pContentLength = pContentLength;
	}

	public int getpItems() {
		return pItems;
	}

	public void setpItems(int pItems) {
		this.pItems = pItems;
	}

	@Override
	public String toString() {
		return "ProgressEntity [pBytesRead=" + pBytesRead + ", pContentLength="
				+ pContentLength + ", pItems=" + pItems + "]";
	}

}
