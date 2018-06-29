package com.ngtesting.platform.service.inf;

import java.util.Map;

public interface MailService extends BaseService {

	void send(String subject, String text, String toEmail);
	void sendTemplateMail(String subject, String templateName, String toEmail, Map<String, String> map);

	String getAppPath(Class<?> cls);
	String getFileName(String path);
	String getFilePath();

}
