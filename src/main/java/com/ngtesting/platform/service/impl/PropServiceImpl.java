package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.service.PropService;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.Bean;
import org.springframework.mail.SimpleMailMessage;
import org.springframework.mail.javamail.JavaMailSender;
import org.springframework.mail.javamail.JavaMailSenderImpl;
import org.springframework.stereotype.Service;

import java.util.Properties;

@Service
public class PropServiceImpl implements PropService {
    @Value("${sys.work.dir}")
    private String workDir;

    @Value("${mail.smtp.username}")
    private String mailUserName;

    @Value("${mail.smtp.password}")
    private String mailPassword;

    @Value("${mail.smtp.host}")
    private String mailHost;

    @Value("${mail.smtp.port}")
    private Integer mailPort;

    @Value("${sys.name}")
    private String sysName;

    @Value("${sys.url.login}")
    private String urlLogin;

    @Value("${sys.url.resetPassword}")
    private String urlResetPassword;

    @Bean
    public SimpleMailMessage simpleMailMessage() {
        SimpleMailMessage mailMessage = new SimpleMailMessage();
        mailMessage.setFrom(mailUserName);
        mailMessage.setSubject("来自ngtesting.com的邮件");
        return mailMessage;
    }

    @Bean
    public JavaMailSender javaMailSender() {
        JavaMailSenderImpl javaMailSender = new JavaMailSenderImpl();
        javaMailSender.setHost(mailHost);
        javaMailSender.setPort(mailPort);
        javaMailSender.setUsername(mailUserName);
        javaMailSender.setPassword(mailPassword);
        Properties p = new Properties();
        p.setProperty("mail.smtp.auth", "true");
        p.setProperty("mail.smtp.socketFactory.class", "javax.net.ssl.SSLSocketFactory");
        javaMailSender.setJavaMailProperties(p);
        return javaMailSender;
    }

    @Override
    public String getSysName() {
        return sysName;
    }

    @Override
    public String getUrlLogin() {
        return urlLogin;
    }

    @Override
    public String getUrlResetPassword() {
        return urlResetPassword;
    }

    @Override
    public String getWorkDir() {
        return workDir;
    }
}
