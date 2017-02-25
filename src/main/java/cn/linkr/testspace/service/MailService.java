package cn.linkr.testspace.service;

import java.io.File;
import java.util.Map;

import cn.linkr.testspace.service.impl.MailServiceImpl;

public interface MailService extends BaseService {

	void send(String subject, String text, String toEmail);
	void sendTemplateMail(String subject, String templateName, String toEmail, Map<String, String> map);

	String getAppPath(Class<?> cls);
	String getFileName(String path);
	String getFilePath();

}