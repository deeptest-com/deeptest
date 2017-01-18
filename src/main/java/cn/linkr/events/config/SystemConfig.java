package cn.linkr.events.config;

import java.util.Properties;

import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.mail.SimpleMailMessage;
import org.springframework.mail.javamail.JavaMailSender;
import org.springframework.mail.javamail.JavaMailSenderImpl;

/**
 * 手机命令执行返回命令监听RabbitMQ配置
 *
 * @author Ethan
 */
@Configuration
public class SystemConfig {

    @Value("${mail.username}")
    private String mailUserName;

    @Value("${mail.password}")
    private String mailPassword;

    @Value("${mail.host}")
    private String mailHost;

    @Value("${mail.port}")
    private Integer mailPort;

    @Bean
    public SimpleMailMessage simpleMailMessage() {
        SimpleMailMessage mailMessage = new SimpleMailMessage();
        mailMessage.setFrom(mailUserName);
        mailMessage.setSubject("[服务]Tinypace");
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
}
