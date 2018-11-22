package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.service.intf.MailService;
import com.ngtesting.platform.utils.StringUtil;
import freemarker.cache.ClassTemplateLoader;
import freemarker.cache.TemplateLoader;
import freemarker.template.Configuration;
import freemarker.template.Template;
import org.apache.commons.mail.HtmlEmail;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.mail.SimpleMailMessage;
import org.springframework.mail.javamail.JavaMailSender;
import org.springframework.mail.javamail.MimeMessageHelper;
import org.springframework.stereotype.Service;
import org.springframework.ui.freemarker.FreeMarkerTemplateUtils;

import javax.mail.internet.MimeMessage;
import java.util.Locale;
import java.util.Map;

@Service
public class MailServiceImpl extends BaseServiceImpl implements MailService {
	final static Logger logger  =  LoggerFactory.getLogger(MailServiceImpl.class );

	public final static String ENCODE = "UTF-8";

	@Autowired
	private SimpleMailMessage mailMessage;

	@Autowired
	private JavaMailSender mailSender;

	public void send(String subject, String text, String toEmail) {
		MimeMessage mimeMessage = mailSender.createMimeMessage();

		MimeMessageHelper messageHelper = null;
		messageHelper = new MimeMessageHelper(mimeMessage, ENCODE);

		if (StringUtil.isEmpty(subject)) {
			subject = mailMessage.getSubject();
		}
		try {
			messageHelper.setSubject(subject);
			messageHelper.setText(text, true);
			messageHelper.setTo(toEmail);
			messageHelper.setFrom(mailMessage.getFrom());
		} catch (Exception e) {
			e.printStackTrace();
		}
		mailSender.send(mimeMessage);
	}

	@Override
	public void sendTemplateMail(String subject, String templateName, String toEmail, Map<String, String> map) {
        Configuration freeMarkerConfig = null;
        HtmlEmail mail = new HtmlEmail();
        try {
            freeMarkerConfig = new Configuration();
            TemplateLoader c1 = new ClassTemplateLoader(TemplateLoader.class, "/mail-template");
            freeMarkerConfig.setTemplateLoader(c1);
            Template template = freeMarkerConfig.getTemplate(templateName, new Locale("Zh_cn"), "UTF-8");

            String htmlText = FreeMarkerTemplateUtils.processTemplateIntoString(template, map);

            mail.setMsg(htmlText);
            send(subject, htmlText, toEmail);
            logger.info("至" + toEmail + "的邮件发送成功");
        } catch (Exception e) {
            logger.info("邮件发送错误：" + e.getMessage());
        }
    }

}
