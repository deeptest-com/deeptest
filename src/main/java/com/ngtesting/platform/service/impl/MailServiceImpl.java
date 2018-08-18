package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.service.MailService;
import com.ngtesting.platform.utils.StringUtil;
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
import java.io.File;
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
		Template template = null;
		Configuration freeMarkerConfig = null;
		HtmlEmail mail = new HtmlEmail();
		try {
			String dir = getFilePath();
			freeMarkerConfig = new Configuration();
			freeMarkerConfig.setDirectoryForTemplateLoading(new File(dir));

			String file = getFileName(templateName);
			template = freeMarkerConfig.getTemplate(file, new Locale("Zh_cn"), "UTF-8");

			String htmlText = FreeMarkerTemplateUtils.processTemplateIntoString(template, map);
			mail.setMsg(htmlText);
			send(subject, htmlText, toEmail);
			logger.info("至" + toEmail + "的邮件发送成功");
		} catch (Exception e) {
			logger.info("邮件发送错误：" + e.getMessage());
		}
	}

	@Override
	public String getClassesPath(Class<?> cls) {
		return cls.getClassLoader().getResource("").getPath();
	}

	@Override
	public String getAppPath(Class<?> cls) {
		// 检查用户传入的参数是否为空
		if (cls == null)
			throw new IllegalArgumentException("参数不能为空！");
		ClassLoader loader = cls.getClassLoader();
		// 获得类的全名，包括包名
		String clsName = cls.getName() + ".class";
		// 获得传入参数所在的包
		Package pack = cls.getPackage();
		String path = "";
		// 如果不是匿名包，将包名转化为路径
		if (pack != null) {
			String packName = pack.getName();
			// 此处简单判定是否是Java基础类库，防止用户传入JDK内置的类库
			if (packName.startsWith("java.") || packName.startsWith("javax."))
				throw new IllegalArgumentException("不要传送系统类！");
			// 在类的名称中，去掉包名的部分，获得类的文件名
			clsName = clsName.substring(packName.length() + 1);
			// 判定包名是否是简单包名，如果是，则直接将包名转换为路径，
			if (packName.indexOf(".") < 0)
				path = packName + "/";
			else {// 否则按照包名的组成部分，将包名转换为路径
				int start = 0, end = 0;
				end = packName.indexOf(".");
				while (end != -1) {
					path = path + packName.substring(start, end) + "/";
					start = end + 1;
					end = packName.indexOf(".", start);
				}
				path = path + packName.substring(start) + "/";
			}
		}
		// 调用ClassLoader的getResource方法，传入包含路径信息的类文件名
		java.net.URL url = loader.getResource(path + clsName);
		// 从URL对象中获取路径信息
		String realPath = url.getPath();
		// 去掉路径信息中的协议名"file:"
		int pos = realPath.indexOf("file:");
		if (pos > -1)
			realPath = realPath.substring(pos + 5);
		// 去掉路径信息最后包含类文件信息的部分，得到类所在的路径
		pos = realPath.indexOf(path + clsName);
		realPath = realPath.substring(0, pos - 1);
		// 如果类文件被打包到JAR等文件中时，去掉对应的JAR等打包文件名
		if (realPath.endsWith("!"))
			realPath = realPath.substring(0, realPath.lastIndexOf("/"));
		/*------------------------------------------------------------
		 ClassLoader的getResource方法使用了utf-8对路径信息进行了编码，当路径
		  中存在中文和空格时，他会对这些字符进行转换，这样，得到的往往不是我们想要
		  的真实路径，在此，调用了URLDecoder的decode方法进行解码，以便得到原始的
		  中文及空格路径
		-------------------------------------------------------------*/
		try {
			realPath = java.net.URLDecoder.decode(realPath, "utf-8");
		} catch (Exception e) {
			throw new RuntimeException(e);
		}
		return realPath;
	}

	@Override
	public String getFileName(String path) {
		path = path.replace("\\", "/");
		return path.substring(path.lastIndexOf("/") + 1);
	}

	@Override
	public String getFilePath() {
		String path = getClassesPath(MailServiceImpl.class);
		path = path + "mail-template" + File.separator;
		path = path.replace("\\", "/");
		return path;
	}

}
