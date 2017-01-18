package cn.linkr.events.entity;

import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.FetchType;
import javax.persistence.JoinColumn;
import javax.persistence.ManyToOne;
import javax.persistence.Table;

@Entity
@Table(name = "clt_company")
public class EvtCompany extends BaseEntity {
	private static final long serialVersionUID = -970910958057582029L;
	
	private String name;
    private String website;

    private String province;
    private String city;
    private String address;

    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "admin_id", insertable = false, updatable = false)
    private EvtClient admin;

    @Column(name = "admin_id")
    private Long adminId;

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public EvtClient getAdmin() {
        return admin;
    }

    public void setAdmin(EvtClient admin) {
        this.admin = admin;
    }

    public Long getAdminId() {
        return adminId;
    }

    public void setAdminId(Long adminId) {
        this.adminId = adminId;
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

	public String getWebsite() {
		return website;
	}

	public void setWebsite(String website) {
		this.website = website;
	}
}
