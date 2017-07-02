package com.ngtesting.platform.vo;

public class TestSuiteVo extends BaseVo {

	private static final long serialVersionUID = 7077320758714477534L;

	private String title;
	private Integer priority;
	private Integer estimate;
	private String objective;
    private String descr;

	private String path;
	private String type;

    public TestSuiteVo(Long id, String title, Integer priority, Integer estimate, String objective, String descr, String path, String type) {
        super();

		this.id = id;
        this.title = title;
        this.priority = priority;
        this.estimate = estimate;
        this.objective = objective;
        this.descr = descr;
        this.path = path;
        this.type = type;

    }

    public String getTitle() {
		return title;
	}

	public void setTitle(String title) {
		this.title = title;
	}

	public Integer getPriority() {
		return priority;
	}

	public void setPriority(Integer priority) {
		this.priority = priority;
	}

	public Integer getEstimate() {
		return estimate;
	}

	public void setEstimate(Integer estimate) {
		this.estimate = estimate;
	}

	public String getDescr() {
		return descr;
	}

	public void setDescr(String descr) {
		this.descr = descr;
	}

	public String getPath() {
		return path;
	}

	public void setPath(String path) {
		this.path = path;
	}

	public String getType() {
		return type;
	}

	public void setType(String type) {
		this.type = type;
	}

	public String getObjective() {
		return objective;
	}

	public void setObjective(String objective) {
		this.objective = objective;
	}

}
