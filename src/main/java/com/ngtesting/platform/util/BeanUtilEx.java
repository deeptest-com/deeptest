package com.ngtesting.platform.util;

import org.apache.commons.beanutils.BeanUtils;
import org.apache.commons.beanutils.ConvertUtils;

public class BeanUtilEx extends BeanUtils {

	static {

	}

	public static void copyProperties(Object target, Object source) {
        ConvertUtils.register(new DateConverter(null), java.util.Date.class);

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
