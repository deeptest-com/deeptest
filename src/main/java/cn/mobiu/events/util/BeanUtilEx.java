package cn.mobiu.events.util;

import java.lang.reflect.InvocationTargetException;

import org.apache.commons.beanutils.BeanUtils;
import org.apache.commons.beanutils.ConvertUtils;

public class BeanUtilEx extends BeanUtils {


	static {
	      ConvertUtils.register(new DateConverter(null), java.util.Date.class);
	}

	public static void copyProperties(Object target, Object source) {
		if (source == null) {
			return;
		}
		
		try {
			BeanUtils.copyProperties(target, source);
		} catch (Exception e) {
			e.printStackTrace();
		}
	}

}
