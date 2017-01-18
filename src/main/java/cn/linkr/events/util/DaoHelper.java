package cn.linkr.events.util;

import java.util.regex.Matcher;
import java.util.regex.Pattern;

import org.apache.commons.lang.StringUtils;
import org.hibernate.FetchMode;
import org.hibernate.criterion.DetachedCriteria;
import org.hibernate.criterion.Order;
import org.hibernate.criterion.Restrictions;
import org.springframework.jdbc.core.RowMapper;
import org.springframework.jdbc.core.simple.ParameterizedBeanPropertyRowMapper;

@SuppressWarnings("unchecked")
public class DaoHelper {

    public static RowMapper resultBeanMapper(Class clazz) {
        return ParameterizedBeanPropertyRowMapper.newInstance(clazz);
    }

    public static String removeSelect(String hql) {
        int beginPos = hql.toLowerCase().indexOf("from");
        return hql.substring(beginPos);
    }

    public static String removeOrders(String hql) {
        Pattern p = Pattern.compile("order\\s*by[\\w|\\W|\\s|\\S]*",
                Pattern.CASE_INSENSITIVE);
        Matcher m = p.matcher(hql);
        StringBuffer sb = new StringBuffer();
        while (m.find()) {
            m.appendReplacement(sb, "");
        }
        m.appendTail(sb);
        return sb.toString();
    }

    public static DetachedCriteria includeClass(Class clasz, String... includes) {

        DetachedCriteria dc = DetachedCriteria.forClass(clasz);
        if (includes != null) {
            for (String propertyName : includes) {
                dc.setFetchMode(propertyName, FetchMode.JOIN);
            }
        }
        return dc;
    }


    public static DetachedCriteria forClass(Class clasz, String orderBy,
                                            Boolean isAsc, String sqlWhere, String... includes) {
        DetachedCriteria dc = includeClass(clasz, includes);
        if (!StringUtils.isEmpty(sqlWhere)) {
            dc.add(Restrictions.sqlRestriction(sqlWhere));
        }
        if (isAsc != null) {
            if (isAsc)
                dc.addOrder(Order.asc(orderBy));
            else
                dc.addOrder(Order.desc(orderBy));
        }

        return dc;
    }
}
