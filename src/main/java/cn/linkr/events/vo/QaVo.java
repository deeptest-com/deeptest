package cn.linkr.events.vo;


public class QaVo extends BaseVo {
	private static final long serialVersionUID = 758955223182231290L;
	private String question;
	private String answer;
	
	public String getQuestion() {
		return question;
	}
	public void setQuestion(String question) {
		this.question = question;
	}
	public String getAnswer() {
		return answer;
	}
	public void setAnswer(String answer) {
		this.answer = answer;
	}

}
