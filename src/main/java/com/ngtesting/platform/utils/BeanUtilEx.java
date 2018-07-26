package com.ngtesting.platform.utils;

import org.apache.commons.beanutils.ConvertUtils;
import org.springframework.beans.BeanUtils;

public class BeanUtilEx extends BeanUtils {

	static {

	}

	public static void copyProperties(Object source, Object target) {
        ConvertUtils.register(new DateConverter(null), java.util.Date.class);

		if (source == null) {
			return;
		}

		try {
			BeanUtils.copyProperties(source, target);
		} catch (Exception e) {
			e.printStackTrace();
		}
	}

}
