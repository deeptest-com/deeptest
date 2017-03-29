package com.ngtesting.platform.vo;

import java.util.Date;
import java.util.HashMap;
import java.util.LinkedList;
import java.util.List;
import java.util.Map;

public class EventVo extends BaseVo {
	private static final long serialVersionUID = -3558813526627228681L;
	
	private String title;
    private String descr;

    private String host;

    private String sponsor;
    
    private String country;
    private String province;
    private String city;
    private String address;
    private String place;

    private String email;
    private String website;
    private String phone;
    private String fax;
    private String qq;
    private String weibo;
    private String wechat;
    private Boolean hasParallelSessin;

    private Date startDatetime;
    private Date endDatetime;
    
    private Date registerStartDatetime;
    private Date registerEndDatetime;   
    
    private Date signStartDatetime;
    private Date signEndDatetime;
    
    // 微信小程序显示用
    private String startDatetimeStr;
    private String endDatetimeStr;
    
    private String registerStartDayStr;
    private String registerStartDatetimeStr;
    private String registerEndDatetimeStr; 
    
    private String signStartDatetimeStr;
    private String signEndDatetimeStr;
    
    // 接受页面数据用
    private String startDate;
    private String endDate;
    private String startTime;
    private String endTime;
    
    private String registerStartDate;
    private String registerEndDate;
    private String registerStartTime;
    private String registerEndTime;
    
    private String signStartDate;
    private String signEndDate;
    private String signStartTime;
    private String signEndTime;
    
    private Integer signBefore;

    private String status;

    private Long companyId;
    private Long creatorId;
    
    private List<BannerVo> banners = new LinkedList<BannerVo>();
    private List<DocumentVo> documents = new LinkedList<DocumentVo>();
    private Map<String, List<OrganizerVo>> organizers = new HashMap<String, List<OrganizerVo>>();

    public String getDescr() {
        return descr;
    }

    public void setDescr(String descr) {
        this.descr = descr;
    }

    public String getHost() {
        return host;
    }

    public void setHost(String host) {
        this.host = host;
    }

    public String getStatus() {
        return status;
    }

    public void setStatus(String status) {
        this.status = status;
    }

    public Long getCompanyId() {
        return companyId;
    }

    public void setCompanyId(Long companyId) {
        this.companyId = companyId;
    }

    public Long getCreatorId() {
        return creatorId;
    }

    public void setCreatorId(Long creatorId) {
        this.creatorId = creatorId;
    }

    public List<DocumentVo> getDocuments() {
        return documents;
    }

    public void setDocuments(List<DocumentVo> documents) {
        this.documents = documents;
    }

	public String getTitle() {
		return title;
	}

	public void setTitle(String title) {
		this.title = title;
	}

	public String getEmail() {
		return email;
	}

	public void setEmail(String email) {
		this.email = email;
	}

	public String getWebsite() {
		return website;
	}

	public void setWebsite(String website) {
		this.website = website;
	}

	public String getQq() {
		return qq;
	}

	public void setQq(String qq) {
		this.qq = qq;
	}

	public String getWeibo() {
		return weibo;
	}

	public void setWeibo(String weibo) {
		this.weibo = weibo;
	}

	public String getWechat() {
		return wechat;
	}

	public void setWechat(String wechat) {
		this.wechat = wechat;
	}

	public Map<String, List<OrganizerVo>> getOrganizers() {
		return organizers;
	}

	public void setOrganizers(Map<String, List<OrganizerVo>> organizers) {
		this.organizers = organizers;
	}

	public String getSponsor() {
		return sponsor;
	}

	public void setSponsor(String sponsor) {
		this.sponsor = sponsor;
	}

	public Boolean getHasParallelSessin() {
		return hasParallelSessin;
	}

	public void setHasParallelSessin(Boolean hasParallelSessin) {
		this.hasParallelSessin = hasParallelSessin;
	}

	public String getPhone() {
		return phone;
	}

	public void setPhone(String phone) {
		this.phone = phone;
	}

	public String getFax() {
		return fax;
	}

	public void setFax(String fax) {
		this.fax = fax;
	}

	public String getCountry() {
		return country;
	}

	public void setCountry(String country) {
		this.country = country;
	}

	public String getProvince() {
		return province;
	}

	public void setProvince(String province) {
		this.province = province;
	}

	public String getCity() {
		return city;
	}

	public void setCity(String city) {
		this.city = city;
	}

	public String getAddress() {
		return address;
	}

	public void setAddress(String address) {
		this.address = address;
	}

	public String getPlace() {
		return place;
	}

	public void setPlace(String place) {
		this.place = place;
	}

	public Integer getSignBefore() {
		return signBefore;
	}

	public void setSignBefore(Integer signBefore) {
		this.signBefore = signBefore;
	}

	public Date getStartDatetime() {
		return startDatetime;
	}

	public void setStartDatetime(Date startDatetime) {
		this.startDatetime = startDatetime;
	}

	public Date getEndDatetime() {
		return endDatetime;
	}

	public void setEndDatetime(Date endDatetime) {
		this.endDatetime = endDatetime;
	}

	public Date getRegisterStartDatetime() {
		return registerStartDatetime;
	}

	public void setRegisterStartDatetime(Date registerStartDatetime) {
		this.registerStartDatetime = registerStartDatetime;
	}

	public Date getRegisterEndDatetime() {
		return registerEndDatetime;
	}

	public void setRegisterEndDatetime(Date registerEndDatetime) {
		this.registerEndDatetime = registerEndDatetime;
	}

	public Date getSignStartDatetime() {
		return signStartDatetime;
	}

	public void setSignStartDatetime(Date signStartDatetime) {
		this.signStartDatetime = signStartDatetime;
	}

	public Date getSignEndDatetime() {
		return signEndDatetime;
	}

	public void setSignEndDatetime(Date signEndDatetime) {
		this.signEndDatetime = signEndDatetime;
	}

	public String getStartDate() {
		return startDate;
	}

	public void setStartDate(String startDate) {
		this.startDate = startDate;
	}

	public String getEndDate() {
		return endDate;
	}

	public void setEndDate(String endDate) {
		this.endDate = endDate;
	}

	public String getStartTime() {
		return startTime;
	}

	public void setStartTime(String startTime) {
		this.startTime = startTime;
	}

	public String getEndTime() {
		return endTime;
	}

	public void setEndTime(String endTime) {
		this.endTime = endTime;
	}

	public String getRegisterStartDate() {
		return registerStartDate;
	}

	public void setRegisterStartDate(String registerStartDate) {
		this.registerStartDate = registerStartDate;
	}

	public String getRegisterEndDate() {
		return registerEndDate;
	}

	public void setRegisterEndDate(String registerEndDate) {
		this.registerEndDate = registerEndDate;
	}

	public String getRegisterStartTime() {
		return registerStartTime;
	}

	public void setRegisterStartTime(String registerStartTime) {
		this.registerStartTime = registerStartTime;
	}

	public String getRegisterEndTime() {
		return registerEndTime;
	}

	public void setRegisterEndTime(String registerEndTime) {
		this.registerEndTime = registerEndTime;
	}

	public String getSignStartDate() {
		return signStartDate;
	}

	public void setSignStartDate(String signStartDate) {
		this.signStartDate = signStartDate;
	}

	public String getSignEndDate() {
		return signEndDate;
	}

	public void setSignEndDate(String signEndDate) {
		this.signEndDate = signEndDate;
	}

	public String getSignStartTime() {
		return signStartTime;
	}

	public void setSignStartTime(String signStartTime) {
		this.signStartTime = signStartTime;
	}

	public String getSignEndTime() {
		return signEndTime;
	}

	public void setSignEndTime(String signEndTime) {
		this.signEndTime = signEndTime;
	}

	public List<BannerVo> getBanners() {
		return banners;
	}

	public void setBanners(List<BannerVo> banners) {
		this.banners = banners;
	}

	public String getStartDatetimeStr() {
		return startDatetimeStr;
	}

	public void setStartDatetimeStr(String startDatetimeStr) {
		this.startDatetimeStr = startDatetimeStr;
	}

	public String getEndDatetimeStr() {
		return endDatetimeStr;
	}

	public void setEndDatetimeStr(String endDatetimeStr) {
		this.endDatetimeStr = endDatetimeStr;
	}

	public String getRegisterStartDatetimeStr() {
		return registerStartDatetimeStr;
	}

	public void setRegisterStartDatetimeStr(String registerStartDatetimeStr) {
		this.registerStartDatetimeStr = registerStartDatetimeStr;
	}

	public String getRegisterEndDatetimeStr() {
		return registerEndDatetimeStr;
	}

	public void setRegisterEndDatetimeStr(String registerEndDatetimeStr) {
		this.registerEndDatetimeStr = registerEndDatetimeStr;
	}

	public String getSignStartDatetimeStr() {
		return signStartDatetimeStr;
	}

	public void setSignStartDatetimeStr(String signStartDatetimeStr) {
		this.signStartDatetimeStr = signStartDatetimeStr;
	}

	public String getSignEndDatetimeStr() {
		return signEndDatetimeStr;
	}

	public void setSignEndDatetimeStr(String signEndDatetimeStr) {
		this.signEndDatetimeStr = signEndDatetimeStr;
	}

	public String getRegisterStartDayStr() {
		return registerStartDayStr;
	}

	public void setRegisterStartDayStr(String registerStartDayStr) {
		this.registerStartDayStr = registerStartDayStr;
	}

}
