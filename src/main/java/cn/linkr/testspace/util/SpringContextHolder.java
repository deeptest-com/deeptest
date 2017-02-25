package cn.linkr.testspace.util;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.BeansException;
import org.springframework.context.ApplicationContext;
import org.springframework.context.ApplicationContextAware;

/**
 * SpringContextHolder spring 工具类
 *
 * @author xuxiang
 * @version $Id$
 * @see
 */

public class SpringContextHolder implements ApplicationContextAware {

    /**
     * log
     */
    private static Logger LOGGER = LoggerFactory.getLogger(SpringContextHolder.class);

    /**
     *
     */
    private static ApplicationContext APPLICATIONCONTEXT;

    @Override
    public void setApplicationContext(ApplicationContext arg0) throws BeansException {
        LOGGER.debug("注入ApplicationContext到SpringContextHolder:{}", APPLICATIONCONTEXT);
        APPLICATIONCONTEXT = arg0;
    }

    /**
     * getApplicationContext()
     *
     * @return 返回 ApplicationContext 对象
     */
    public static ApplicationContext getApplicationContext() {
        checkApplicationContext();
        return APPLICATIONCONTEXT;
    }

    /**
     * 根据名称获得指定的bean
     *
     * @param name bean 名称
     * @param <T>  bean
     * @return 返回相关的bean
     */
    public static <T> T getBean(String name, Class<T> reqiredType) {
        checkApplicationContext();
        return APPLICATIONCONTEXT.getBean(name, reqiredType);
    }

    /**
     * 〈简述〉
     * 〈详细描述〉
     *
     * @param clazz bean类型
     * @param <T>   bean
     * @return 返回相关的bean
     */
    public static <T> T getBean(Class<T> clazz) {
        checkApplicationContext();
        return APPLICATIONCONTEXT.getBean(clazz);
    }

    /**
     * 检查ApplicationContext不为空.
     */
    public static void checkApplicationContext() {
        if (APPLICATIONCONTEXT == null) {
            throw new IllegalStateException("applicationContext未注入，请在applicationContext.xml中配置");
        }
    }
}
