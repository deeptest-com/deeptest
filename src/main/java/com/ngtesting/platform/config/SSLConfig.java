package com.ngtesting.platform.config;

import org.apache.catalina.Context;
import org.apache.catalina.connector.Connector;
import org.apache.tomcat.util.descriptor.web.SecurityCollection;
import org.apache.tomcat.util.descriptor.web.SecurityConstraint;
import org.apache.tomcat.websocket.server.WsSci;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.boot.web.embedded.tomcat.TomcatContextCustomizer;
import org.springframework.boot.web.embedded.tomcat.TomcatServletWebServerFactory;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.core.env.Environment;

@Configuration
public class SSLConfig {
    @Value("${http.port}")
    Integer httpPort;

    @Value("${server.port}")
    Integer httpsPort;

    @Autowired
    private Environment env;

    @Bean
    public TomcatServletWebServerFactory servletContainer() { //springboot2 新变化
        String profile = env.getActiveProfiles()[0];

        if ("docker-domain".equals(profile)) {
            TomcatServletWebServerFactory tomcat = new TomcatServletWebServerFactory() {

                @Override
                protected void postProcessContext(Context context) {
                    SecurityConstraint securityConstraint = new SecurityConstraint();
                    securityConstraint.setUserConstraint("CONFIDENTIAL");
                    SecurityCollection collection = new SecurityCollection();
                    collection.addPattern("/*");
                    securityConstraint.addCollection(collection);
                    context.addConstraint(securityConstraint);
                }

            };
            tomcat.addAdditionalTomcatConnectors(initiateHttpConnector());

            return tomcat;
        } else {
            return new TomcatServletWebServerFactory();
        }
    }

    private Connector initiateHttpConnector() {
        Connector connector = new Connector("org.apache.coyote.http11.Http11NioProtocol");
        connector.setScheme("http");
        connector.setPort(httpPort);
        connector.setSecure(false);
        connector.setRedirectPort(httpsPort);
        return connector;
    }

    @Bean
    public TomcatContextCustomizer tomcatContextCustomizer() {
        String profile = env.getActiveProfiles()[0];

        if ("docker-domain".equals(profile)) {
            return new TomcatContextCustomizer() {
                @Override
                public void customize(Context context) {
                    System.out.println("init customize");
                    context.addServletContainerInitializer(new WsSci(), null);
                }

            };
        } else {
            return new TomcatContextCustomizer() {
                @Override
                public void customize(Context context) {}
            };
        }
    }
}
