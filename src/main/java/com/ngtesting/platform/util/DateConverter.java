package com.ngtesting.platform.util;

import org.apache.commons.beanutils.ConversionException;
import org.apache.commons.beanutils.Converter;
import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;

public final class DateConverter implements Converter {

    private Log log = LogFactory.getLog(this.getClass());

    public DateConverter() {
        this.defaultValue = null;
        this.useDefault = false;
    }

    public DateConverter(Object defaultValue) {
        this.defaultValue = defaultValue;
        this.useDefault = true;
    }

    private Object defaultValue = null;
    private boolean useDefault = true;

    public Object convert(Class type, Object value) {
        if (value == null || "".equals(value)) {
            if (useDefault) {
                return (defaultValue);
            } else {
                throw new ConversionException("No value specified");
            }
        }

        if (value instanceof java.util.Date) {
            return (value);
        } else  if (value instanceof java.sql.Date) {
            return new java.util.Date(((java.sql.Date)value).getTime());
        } else  if (value instanceof java.sql.Timestamp) {
            return new java.util.Date(((java.sql.Timestamp)value).getTime());
        }

        return (defaultValue);
    }
}