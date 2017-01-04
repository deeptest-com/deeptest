package cn.mobiu.events.vo;


public class UserVo extends BaseVo {
	private static final long serialVersionUID = 1988353599211843484L;
	private String email;
    private String phone;
    private String name;
    private String token;
	public String getEmail() {
		return email;
	}
	public void setEmail(String email) {
		this.email = email;
	}
	public String getPhone() {
		return phone;
	}
	public void setPhone(String phone) {
		this.phone = phone;
	}
	public String getName() {
		return name;
	}
	public void setName(String name) {
		this.name = name;
	}
	public String getToken() {
		return token;
	}
	public void setToken(String token) {
		this.token = token;
	}

}
