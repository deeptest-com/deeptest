package cn.linkr.testspace.util;

import java.util.Date;

import org.apache.commons.beanutils.converters.DateTimeConverter;

public class MyDateConverter extends DateTimeConverter {

    public MyDateConverter() {
    }

    public MyDateConverter(Object defaultValue) {
        super(defaultValue);
    }

    @SuppressWarnings("rawtypes")
    protected Class getDefaultType() {
        return Date.class;
    }

    @SuppressWarnings("rawtypes")
    @Override
    protected Object convertToType(Class arg0, Object arg1) throws Exception {
        if (arg1 == null) {
            return null;
        }
        String value = arg1.toString().trim();
        if (value.length() == 0) {
            return null;
        }
        return super.convertToType(arg0, arg1);
    }
}
