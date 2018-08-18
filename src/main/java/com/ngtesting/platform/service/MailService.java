package com.ngtesting.platform.service;

import java.util.Map;

public interface MailService extends BaseService {

	void send(String subject, String text, String toEmail);
	void sendTemplateMail(String subject, String templateName, String toEmail, Map<String, String> map);

}
