package com.ngtesting.platform.entity;

import java.util.Date;

import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.EnumType;
import javax.persistence.Enumerated;
import javax.persistence.FetchType;
import javax.persistence.JoinColumn;
import javax.persistence.ManyToOne;
import javax.persistence.Table;

@Entity
@Table(name = "evt_event")
public class EvtEvent extends BaseEntity {
	private static final long serialVersionUID = 3428179208641340812L;
	
	private String title;

    // @Lob
    @Column(name = "descr", length = 10000)
    private String descr;

    private String sponsor;

    private String email;
    private String website;
    private String phone;
    private String fax;
    private String qq;
    private String weibo;
    private String wechat;
    
    private String country;
    private String province;
    private String city;
    private String address;
    private String place;

    private Date startDatetime;
    private Date endDatetime;

    private Date registerStartDatetime;
    private Date registerEndDatetime;
    
    private Date signStartDatetime;
    private Date signEndDatetime;
    
    private Integer signBefore;
    
    private Boolean hasParallelSessin = false;

    @Enumerated(EnumType.STRING)
    private EventStatus status;

    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "company_id", insertable = false, updatable = false)
    private SysCompany company;

    @Column(name = "company_id")
    private Long companyId;
    
    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "creator_id", insertable = false, updatable = false)
    private SysUser creator;

    @Column(name = "creator_id")
    private Long creatorId;

    public static enum EventStatus {
        not_start("not_start"),
        register("register"),
        sign("sign"),
        in_progress("in_progress"),
        end("end"),
        cancel("cancel");

        private EventStatus(String textVal) {
            this.textVal = textVal;
        }

        private String textVal;
        public String toString() {
            return textVal;
        }
        
        public static EventStatus getValue(String str) {
        	EventStatus status = null;
        	switch(str) { 
            	case "not_start": status = EventStatus.not_start; break;
            	case "register": status = EventStatus.register; break;
            	case "sign": status = EventStatus.sign; break;
            	case "in_progress": status = EventStatus.in_progress; break;
            	case "end": status = EventStatus.end; break;
            	case "cancel": status = EventStatus.cancel; break;
            }
        	
        	return status;
        }
    }

    public String getDescr() {
        return descr;
    }

    public void setDescr(String descr) {
        this.descr = descr;
    }

    public EventStatus getStatus() {
    	
		long now = new Date().getTime();
		EventStatus status = null;
		if (now > this.getEndDatetime().getTime()){
			status = EventStatus.end;
		} else if (now >= this.getStartDatetime().getTime()
				&& now <= this.getEndDatetime().getTime()) {
			status = EventStatus.in_progress;
		} else if (this.getSignStartDatetime() != null && now >= this.getSignStartDatetime().getTime() 
				&& this.getSignEndDatetime() !=null && now <= this.getSignEndDatetime().getTime()) {
			status = EventStatus.sign;
		} else if (this.getRegisterStartDatetime() != null && now >= this.getRegisterStartDatetime().getTime() 
				&& this.getRegisterEndDatetime() != null && now <= this.getRegisterEndDatetime().getTime()) {
			status = EventStatus.register;
		} else {
			status = EventStatus.not_start;
		}
		
//		if (!status.equals(this.getStatus())) {
//			this.setStatus(status);
//			saveOrUpdate(this);
//		}
    	
        return status;
    }

    public void setStatus(EventStatus status) {
        this.status = status;
    }

    public SysCompany getCompany() {
        return company;
    }

    public void setCompany(SysCompany company) {
        this.company = company;
    }

    public Long getCompanyId() {
        return companyId;
    }

    public void setCompanyId(Long companyId) {
        this.companyId = companyId;
    }

    public SysUser getCreator() {
        return creator;
    }

    public void setCreator(SysUser creator) {
        this.creator = creator;
    }

    public Long getCreatorId() {
        return creatorId;
    }

    public void setCreatorId(Long creatorId) {
        this.creatorId = creatorId;
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

	public String getWebsite() {
		return website;
	}

	public void setWebsite(String website) {
		this.website = website;
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
}
