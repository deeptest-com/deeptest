package com.ngtesting.platform.util;

import java.util.HashMap;
import java.util.Map;
import java.util.Properties;

import org.springframework.beans.BeansException;
import org.springframework.beans.factory.config.ConfigurableListableBeanFactory;
import org.springframework.beans.factory.config.PropertyPlaceholderConfigurer;

public class PropertyConfig extends PropertyPlaceholderConfigurer {

    private static Map<String, Object> CTXPROPERTIESMAP;

    @Override
    protected void processProperties(ConfigurableListableBeanFactory beanFactory,
                                     Properties props) throws BeansException {

        super.processProperties(beanFactory, props);

        CTXPROPERTIESMAP = new HashMap<String, Object>();
        for (Object key : props.keySet()) {
            String keyStr = key.toString();
            String value = props.getProperty(keyStr);
            CTXPROPERTIESMAP.put(keyStr, value);
        }

        // 初始化常量
//        Constant.WORK_DIR = PropertyConfig.getConfig("dir.base");

    }

    public static Object getContextProperty(String name) {
        return CTXPROPERTIESMAP.get(name);
    }

    public static String getConfig(String name) {
        return CTXPROPERTIESMAP.get(name).toString();
    }

}
