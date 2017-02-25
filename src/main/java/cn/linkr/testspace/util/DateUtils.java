package cn.linkr.testspace.util;

import java.text.ParseException;
import java.text.SimpleDateFormat;
import java.util.ArrayList;
import java.util.Calendar;
import java.util.Date;
import java.util.List;

/**
 * 时间处理工具类
 *
 * @author xuxiang
 * @version $Id$
 * @see
 */
public class DateUtils {

    /**
     * 数字常量-1
     */
    private static final int NUMBER_MINE_ONE = -1;

    /**
     * 数字常量0
     */
    private static final int NUMBER_ZERO = 0;

    /**
     * 数字常量1
     */
    private static final int NUMBER_ONE = 1;

    /**
     * 数字常量2
     */
    private static final int NUMBER_TWO = 2;

    /**
     * 数字常量3
     */
    private static final int NUMBER_THREE = 3;

    /**
     * 数字常量4
     */
    private static final int NUMBER_FOUR = 4;

    /**
     * 数字常量5
     */
    private static final int NUMBER_FIVE = 5;

    /**
     * 数字常量6
     */
    private static final int NUMBER_SIX = 6;

    /**
     * 数字常量7
     */
    private static final int NUMBER_SEVEN = 7;

    /**
     * 数字常量8
     */
    private static final int NUMBER_EIGHT = 8;

    /**
     * 数字常量9
     */
    private static final int NUMBER_NINE = 9;

    /**
     * 数字常量10
     */
    private static final int NUMBER_TEN = 10;

    /**
     * 数字常量11
     */
    private static final int NUMBER_ELEVEN = 11;

    /**
     * 数字常量10
     */
    private static final int NUMBER_TWELVE = 12;

    /**
     * 数字常量28
     */
    private static final int NUMBER_TWENTY_EIGHT = 28;

    /**
     * 数字常量29
     */
    private static final int NUMBER_TWENTY_NINE = 29;

    /**
     * 数字常量30
     */
    private static final int NUMBER_THIRTY = 30;

    /**
     * 数字常量31
     */
    private static final int NUMBER_THIRTY_ONE = 31;

    /**
     * 数字常量100
     */
    private static final int NUMBER_ONE_HUNDRED = 100;

    /**
     * 数字常量400
     */
    private static final int NUMBER_FOUR_HUNDRED = 400;


    /**
     * 得到几天前的时间
     *
     * @param d
     * @param day
     * @return
     */
    public static Date getDateBefore(Date d, int day) {
        Calendar cal = Calendar.getInstance();
        cal.setTime(d);
        cal.set(Calendar.DATE, cal.get(Calendar.DATE) - day);
        return cal.getTime();
    }

    /**
     * 得到几天后的时间
     *
     * @param d
     * @param day
     * @return
     */
    public static Date getDateAfter(Date d, int day) {
        Calendar cal = Calendar.getInstance();
        cal.setTime(d);
        cal.set(Calendar.DATE, cal.get(Calendar.DATE) + day);
        return cal.getTime();
    }

    /**
     * 得到几分钟前的时间
     *
     * @param d
     * @param min
     * @return
     */
    public static Date getMinBefore(Date d, int min) {
        Calendar cal = Calendar.getInstance();
        cal.setTime(d);
        cal.set(Calendar.MINUTE, cal.get(Calendar.MINUTE) - min);
        return cal.getTime();
    }

    /**
     * 得到几分钟后的时间
     *
     * @param d
     * @param min
     * @return
     */
    public static Date getMinAfter(Date d, int min) {
        Calendar cal = Calendar.getInstance();
        cal.setTime(d);
        cal.set(Calendar.MINUTE, cal.get(Calendar.MINUTE) + min);
        return cal.getTime();
    }

    /**
     * 时间格式化
     *
     * @param time  时间秒
     * @param style 格式
     * @return 返回string
     */
    public static String formatDate(long time, String style) {
        Date date = new Date(time);
        SimpleDateFormat outFormat = new SimpleDateFormat(style);
        return outFormat.format(date);
    }

    /**
     * 时间格式化
     *
     * @param date 日期
     * @return 返回string
     */
    public static String formatDate(Date date) {
        if (date == null) {
            return "";
        }
        SimpleDateFormat outFormat = new SimpleDateFormat("yyyy-MM-dd");
        return outFormat.format(date);
    }

    /**
     * 时间格式化
     *
     * @return 返回string(yyyyMMdd)
     */
    public static String getDateNoSeparator() {
        Date date = new Date();
        SimpleDateFormat outFormat = new SimpleDateFormat("yyyyMMdd");
        return outFormat.format(date);
    }


    /**
     * 时间格式化
     *
     * @return 返回string(yyyyMMdd)
     */
    public static String getSmallDateNoSeparator() {
        Date date = new Date();
        SimpleDateFormat outFormat = new SimpleDateFormat("yyyyMMddHHmmssSSS");
        return outFormat.format(date);
    }

    /**
     * 时间格式化
     *
     * @param date   日期
     * @param format 格式
     * @return 返回string
     */
    public static String formatDate(Date date, String format) {
        if (date == null) {
            return "";
        }
        SimpleDateFormat outFormat = new SimpleDateFormat(format);
        return outFormat.format(date);
    }

    /**
     * 时间格式化
     *
     * @param date 日期
     * @return 返回string
     */
    public static String formatDateTime(Date date) {
        SimpleDateFormat outFormat = new SimpleDateFormat("yyyy-MM-dd HH:mm:ss");
        return outFormat.format(date);
    }

    /**
     * 字符串时间转日期
     *
     * @param dt     日期字符串
     * @param format 格式
     * @return 返回日期
     */
    public static Date str2Date(String dt, String format) {
        SimpleDateFormat df = new SimpleDateFormat(format);
        Date date = null;
        try {
            date = df.parse(dt);
            return date;
        } catch (ParseException e) {
            e.printStackTrace();
        }
        return date;
    }

    /**
     * 获得年份
     *
     * @return 年份
     */
    public static int getYear() {
        Calendar cld = Calendar.getInstance();
        cld.setTime(new Date());
        return cld.get(1);
    }

    /**
     * 获得月份
     *
     * @return 月份
     */
    public static int getMonth() {
        Calendar cld = Calendar.getInstance();
        cld.setTime(new Date());
        return cld.get(Calendar.MONTH) + 1;
    }

    /**
     * 获得日期
     *
     * @return 日期day
     */
    public static int getDay() {
        Calendar cld = Calendar.getInstance();
        cld.setTime(new Date());
        return cld.get(Calendar.DAY_OF_MONTH);
    }

    /**
     * 获得小时
     *
     * @return 小时
     */
    public static int getHour() {
        Calendar cld = Calendar.getInstance();
        cld.setTime(new Date());
        return cld.get(Calendar.HOUR_OF_DAY);
    }

    /**
     * 获得年份
     *
     * @param t 时间
     * @return 年
     */
    public static int getYear(long t) {
        Calendar cld = Calendar.getInstance();
        if (t > 0L) {
            cld.setTime(new Date(t));
        }
        return cld.get(1);
    }

    /**
     * 获得月份
     *
     * @param t 时间
     * @return 月
     */
    public static int getMonth(long t) {
        Calendar cld = Calendar.getInstance();
        if (t > 0L) {
            cld.setTime(new Date(t));
        }
        return cld.get(Calendar.MONTH) + 1;
    }

    /**
     * 获得日期
     *
     * @param t 时间
     * @return 日期day
     */
    public static int getDay(long t) {
        Calendar cld = Calendar.getInstance();
        if (t > 0L) {
            cld.setTime(new Date(t));
        }
        return cld.get(Calendar.DAY_OF_MONTH);
    }

    /**
     * 获得小时
     *
     * @param t 时间
     * @return 小时
     */
    public static int getHour(long t) {
        Calendar cld = Calendar.getInstance();
        if (t > 0L) {
            cld.setTime(new Date(t));
        }
        return cld.get(Calendar.HOUR_OF_DAY);
    }

    /**
     * 获得分钟
     *
     * @param t 时间
     * @return 分钟
     */
    public static int getMinute(long t) {
        Calendar cld = Calendar.getInstance();
        if (t > 0L) {
            cld.setTime(new Date(t));
        }
        return cld.get(Calendar.MINUTE);
    }

    /**
     * 获得秒
     *
     * @param t 时间
     * @return 秒
     */
    public static int getSecond(long t) {
        Calendar cld = Calendar.getInstance();
        if (t > 0L) {
            cld.setTime(new Date(t));
        }
        return cld.get(Calendar.MILLISECOND);
    }

    /**
     * 获得年
     *
     * @param date 日期
     * @return 年
     */
    public static int getYear(Date date) {
        Calendar cld = Calendar.getInstance();
        cld.setTime(date);
        return cld.get(1);
    }

    /**
     * 获得月
     *
     * @param date 日期
     * @return 月
     */
    public static int getMonth(Date date) {
        Calendar cld = Calendar.getInstance();
        cld.setTime(date);
        return cld.get(Calendar.MONTH) + 1;
    }

    /**
     * 获得日期day
     *
     * @param date 日期
     * @return 日期day
     */
    public static int getDay(Date date) {
        Calendar cld = Calendar.getInstance();
        cld.setTime(date);
        return cld.get(Calendar.DAY_OF_MONTH);
    }

    /**
     * 获得小时
     *
     * @param date 日期
     * @return 小时
     */
    public static int getHour(Date date) {
        Calendar cld = Calendar.getInstance();
        cld.setTime(date);
        return cld.get(Calendar.HOUR_OF_DAY);
    }

    /**
     * 获得分钟
     *
     * @param date 日期
     * @return 分钟
     */
    public static int getMinute(Date date) {
        Calendar cld = Calendar.getInstance();
        cld.setTime(date);
        return cld.get(Calendar.MINUTE);
    }

    /**
     * 获得秒
     *
     * @param date 日期
     * @return 秒
     */
    public static int getSecond(Date date) {
        Calendar cld = Calendar.getInstance();
        cld.setTime(date);
        return cld.get(Calendar.MILLISECOND);
    }

    /**
     * 天数加
     *
     * @param date 日期
     * @param days 天数
     * @return 日期
     */
    public static Date addDays(Date date, int days) {
        return org.apache.commons.lang.time.DateUtils.addDays(date, days);
    }

    /**
     * 年数加
     *
     * @param date  日期
     * @param years 年数
     * @return 日期
     */
    public static Date addYears(Date date, int years) {
        return org.apache.commons.lang.time.DateUtils.addYears(date, years);
    }

    /**
     * 月数加
     *
     * @param date   日期
     * @param months 月数
     * @return 日期
     */
    public static Date addMonths(Date date, int months) {
        return org.apache.commons.lang.time.DateUtils.addMonths(date, months);
    }

    /**
     * 周数加
     *
     * @param date  日期
     * @param weeks 周数
     * @return 日期
     */
    public static Date addWeeks(Date date, int weeks) {
        return org.apache.commons.lang.time.DateUtils.addWeeks(date, weeks);
    }

    /**
     * 是否为闰年
     *
     * @param year 年
     * @return boolea
     */
    public static boolean checkLeapyear(int year) {
        boolean isLeapyear = false;
        if ((year % NUMBER_FOUR == 0) && (year % NUMBER_ONE_HUNDRED != 0)) {
            isLeapyear = true;
        }
        if (year % NUMBER_FOUR_HUNDRED == 0) {
            isLeapyear = true;
        } else if (year % NUMBER_FOUR != 0) {
            isLeapyear = false;
        }
        return isLeapyear;
    }

    /**
     * 获取每月的天数
     *
     * @param month 月份
     * @param year  年
     * @return boolea
     */
    public static int getDaysOfmonth(int month, int year) {
        int dates = NUMBER_ZERO;
        if ((month < NUMBER_ZERO) || (month > Calendar.MINUTE)) {
            System.out.println("month Error");
        }
        if ((month == NUMBER_ONE) || (month == NUMBER_THREE) || (month == NUMBER_FIVE) || (month == NUMBER_SEVEN) || (month == NUMBER_EIGHT)
                || (month == NUMBER_TEN) || (month == NUMBER_TWELVE)) {
            dates = NUMBER_THIRTY_ONE;
        }
        if ((month == NUMBER_TWO) && (checkLeapyear(year))) {
            dates = NUMBER_TWENTY_NINE;
        }
        if ((month == NUMBER_TWO) && (!checkLeapyear(year))) {
            dates = NUMBER_TWENTY_EIGHT;
        }
        if ((month == NUMBER_FOUR) || (month == NUMBER_SIX) || (month == NUMBER_NINE) || (month == NUMBER_ELEVEN)) {
            dates = NUMBER_THIRTY_ONE;
        }
        return dates;
    }

    /**
     * 获得上一个月的结束日期
     * 〈详细描述〉
     *
     * @return 返回日期
     */
    public static Date getLastMonthEndDate() {
        Calendar cal = Calendar.getInstance();
        Date date = new Date();
        cal.setTime(date);
        int year = 0;
        int month = cal.get(NUMBER_TWO);
        int day = cal.getActualMaximum(NUMBER_FIVE);
        if (month == NUMBER_ZERO) {
            year = cal.get(Calendar.YEAR) - NUMBER_ONE;
            month = NUMBER_TWELVE;
        } else {
            year = cal.get(Calendar.YEAR);
        }
        String endDate = year + "-" + month + "-" + day;
        return str2Date(endDate, "yyyy-MM-dd");
    }

    /**
     * 获得上一个月的开始时间
     * 〈详细描述〉
     *
     * @return 返回日期
     */
    public static Date getLastMonthStartDate() {
        Calendar cal = Calendar.getInstance();
        Date date = new Date();
        cal.setTime(date);
        int year = NUMBER_ZERO;
        int month = cal.get(Calendar.MONTH);
        int day = cal.getActualMinimum(Calendar.DAY_OF_MONTH);
        if (month == NUMBER_ZERO) {
            year = cal.get(Calendar.YEAR) - NUMBER_ONE;
            month = Calendar.MINUTE;
        } else {
            year = cal.get(Calendar.YEAR);
        }
        String startDate = year + "-" + month + "-" + day;
        return str2Date(startDate, "yyyy-MM-dd");
    }

    /**
     * 获得当前季度的开始时间
     *
     * @return 返回日期
     */
    public static Date getCurrentQuarterStartTime() {
        SimpleDateFormat shortSdf = new SimpleDateFormat("yyyy-MM-dd");
        SimpleDateFormat longSdf = new SimpleDateFormat("yyyy-MM-dd HH:mm:ss");
        Calendar c = Calendar.getInstance();
        int currentMonth = c.get(Calendar.MONTH) + NUMBER_ONE;
        Date now = null;
        try {
            if ((currentMonth >= NUMBER_ONE) && (currentMonth <= NUMBER_THREE)) {
                c.set(Calendar.MONTH, NUMBER_ZERO);
            } else if ((currentMonth >= NUMBER_FOUR) && (currentMonth <= NUMBER_SIX)) {
                c.set(Calendar.MONTH, NUMBER_THREE);
            } else if ((currentMonth >= NUMBER_SEVEN) && (currentMonth <= NUMBER_NINE)) {
                c.set(Calendar.MONTH, NUMBER_SIX);
            } else if ((currentMonth >= NUMBER_TEN) && (currentMonth <= NUMBER_TWELVE)) {
                c.set(Calendar.MONTH, NUMBER_NINE);
                c.set(Calendar.DAY_OF_MONTH, NUMBER_ONE);
            }
            now = longSdf.parse(shortSdf.format(c.getTime()) + " 00:00:00");
        } catch (Exception e) {
            e.printStackTrace();
        }
        return now;
    }

    /**
     * 获得当前季度的结束时间
     *
     * @return 返回日期
     */
    public static Date getCurrentQuarterEndTime() {
        SimpleDateFormat shortSdf = new SimpleDateFormat("yyyy-MM-dd");
        SimpleDateFormat longSdf = new SimpleDateFormat("yyyy-MM-dd HH:mm:ss");
        Calendar c = Calendar.getInstance();
        int currentMonth = c.get(Calendar.MONTH) + NUMBER_ONE;
        Date now = null;
        try {
            if ((currentMonth >= NUMBER_ONE) && (currentMonth <= NUMBER_THREE)) {
                c.set(Calendar.MONTH, NUMBER_TWO);
                c.set(Calendar.DAY_OF_MONTH, NUMBER_THIRTY_ONE);
            } else if ((currentMonth >= NUMBER_FOUR) && (currentMonth <= NUMBER_SIX)) {
                c.set(Calendar.MONTH, NUMBER_FIVE);
                c.set(Calendar.DAY_OF_MONTH, NUMBER_THIRTY);
            } else if ((currentMonth >= NUMBER_SEVEN) && (currentMonth <= NUMBER_NINE)) {
                c.set(Calendar.MONTH, NUMBER_EIGHT);
                c.set(Calendar.DAY_OF_MONTH, NUMBER_THIRTY);
            } else if ((currentMonth >= NUMBER_TEN) && (currentMonth <= Calendar.MINUTE)) {
                c.set(Calendar.MONTH, Calendar.HOUR_OF_DAY);
                c.set(Calendar.DAY_OF_MONTH, NUMBER_THIRTY_ONE);
            }
            now = longSdf.parse(shortSdf.format(c.getTime()) + " 23:59:59");
        } catch (Exception e) {
            e.printStackTrace();
        }
        return now;
    }

    /**
     * 获得上个季度的开始时间
     *
     * @return 返回日期
     */
    public static Date getLastQuarterStartTime() {
        Date date = getCurrentQuarterStartTime();
        Date endDate = addDays(date, NUMBER_MINE_ONE);
        Date startDate = addMonths(endDate, -NUMBER_THREE);
        int year = getYear(startDate);
        int month = getMonth(startDate) + NUMBER_ONE;
        String nextDateString = year + "-" + (month >= NUMBER_TEN ? Integer.valueOf(month) : new StringBuilder("0").append(month).toString()) + "-01 00:00:00";
        Date d = str2Date(nextDateString, "yyyy-MM-dd HH:mm:ss");
        return d;
    }

    /**
     * 获得上个季度的结束时间
     *
     * @return 返回日期
     */
    public static Date getLastQuarterEndTime() {
        Date date = getCurrentQuarterStartTime();
        Date endDate = addDays(date, NUMBER_MINE_ONE);
        return endDate;
    }

    /**
     * 获得上周的开始时间
     *
     * @return 返回日期
     */
    public static Date getLastWeekDayStartTime() {
        SimpleDateFormat shortSdf = new SimpleDateFormat("yyyy-MM-dd");
        SimpleDateFormat longSdf = new SimpleDateFormat("yyyy-MM-dd HH:mm:ss");
        Calendar c = Calendar.getInstance();
        c.add(Calendar.WEEK_OF_YEAR, -1);
        try {
            int weekday = c.get(Calendar.DAY_OF_WEEK) - NUMBER_TWO;
            c.add(Calendar.DAY_OF_MONTH, -weekday);
            c.setTime(longSdf.parse(shortSdf.format(c.getTime()) + " 00:00:00"));
        } catch (Exception e) {
            e.printStackTrace();
        }
        return c.getTime();
    }

    /**
     * 获得上周的结束时间
     *
     * @return 返回日期
     */
    public static Date getLastWeekDayEndTime() {
        SimpleDateFormat shortSdf = new SimpleDateFormat("yyyy-MM-dd");
        SimpleDateFormat longSdf = new SimpleDateFormat("yyyy-MM-dd HH:mm:ss");
        Calendar c = Calendar.getInstance();
        c.add(Calendar.WEEK_OF_YEAR, NUMBER_MINE_ONE);
        try {
            int weekday = c.get(Calendar.DAY_OF_WEEK);
            c.add(Calendar.DAY_OF_MONTH, NUMBER_EIGHT - weekday);
            c.setTime(longSdf.parse(shortSdf.format(c.getTime()) + " 23:59:59"));
        } catch (Exception e) {
            e.printStackTrace();
        }
        return c.getTime();
    }

    /**
     * 获得星期几
     *
     * @param d 日期 如果日期为空，则返回当前的时间
     * @return 字符串
     */
    public static String getWeekDay(Date d) {
        Calendar calendar = Calendar.getInstance();
        if (d != null) {
            calendar.setTime(d);
        }
        int dayOfWeek = calendar.get(Calendar.DAY_OF_WEEK);
        String weekDayString = "";
        switch (dayOfWeek) {
            case Calendar.SUNDAY:
                weekDayString = "星期日";
                break;
            case Calendar.MONDAY:
                weekDayString = "星期一";
                break;
            case Calendar.TUESDAY:
                weekDayString = "星期二";
                break;
            case Calendar.WEDNESDAY:
                weekDayString = "星期三";
                break;
            case Calendar.THURSDAY:
                weekDayString = "星期四";
                break;
            case Calendar.FRIDAY:
                weekDayString = "星期五";
                break;
            case Calendar.SATURDAY:
                weekDayString = "星期六";
                break;
            default:
                break;
        }
        return weekDayString;
    }

    public static Long getNow() {
        Calendar ca = Calendar.getInstance();
        ca.setTime(new Date());
        return ca.getTimeInMillis();
    }

    public static Date getTomorrow(Date dt) {
        Calendar cal = Calendar.getInstance();
        cal.setTime(dt);

        cal.set(Calendar.MINUTE, 0);
        cal.set(Calendar.SECOND, 0);
        cal.set(Calendar.MILLISECOND, 0);
        cal.add(Calendar.HOUR, 1);

        return cal.getTime();
    }

    public static Date getStartTimeOfDay(Date dt) {
        Calendar cal = Calendar.getInstance();
        cal.setTime(dt);

        Date todayEnd = DateUtils.getEndTimeOfDay(new Date());
        if (dt.before(todayEnd)) { // 当天，取下一小时

        } else { // 取那一天的0点
            cal.set(Calendar.HOUR_OF_DAY, 0);
            cal.set(Calendar.MINUTE, 0);
            cal.set(Calendar.SECOND, 0);
            cal.set(Calendar.MILLISECOND, 0);
        }

        return cal.getTime();
    }

    public static Date getEndTimeOfDay(Date dt) { // 到第二天之前
        Calendar cal = Calendar.getInstance();
        cal.setTime(dt);
        cal.set(Calendar.HOUR_OF_DAY, 0);
        cal.set(Calendar.MINUTE, 0);
        cal.set(Calendar.SECOND, 0);
        cal.set(Calendar.MILLISECOND, 0);
        cal.add(Calendar.DAY_OF_MONTH, 1);
        cal.add(Calendar.MILLISECOND, -1);
        return cal.getTime();
    }

    public static Date getSartTimeOfMonth(Date dt) { // 第二天开始
        Calendar cal = Calendar.getInstance();
        cal.setTime(dt);

        if (DateUtils.isToday(dt)) { // 当天，取第二天
            cal.set(Calendar.HOUR_OF_DAY, 0);
            cal.set(Calendar.MINUTE, 0);
            cal.set(Calendar.SECOND, 0);
            cal.set(Calendar.MILLISECOND, 0);
            cal.add(Calendar.DAY_OF_MONTH, 1);
        } else { // 取那个月的第一天
            cal.set(Calendar.HOUR_OF_DAY, 0);
            cal.set(Calendar.MINUTE, 0);
            cal.set(Calendar.SECOND, 0);
            cal.set(Calendar.MILLISECOND, 0);
            cal.set(Calendar.DAY_OF_MONTH, 0); // 到月末
            cal.add(Calendar.DAY_OF_MONTH, 1);
        }
        return cal.getTime();
    }

    public static Date getEndTimeOfMonth(Date dt) { // 到月末
        Calendar cal = Calendar.getInstance();
        cal.setTime(dt);
        cal.set(Calendar.HOUR_OF_DAY, 0);
        cal.set(Calendar.MINUTE, 0);
        cal.set(Calendar.SECOND, 0);
        cal.set(Calendar.MILLISECOND, 0);
        cal.set(Calendar.DAY_OF_MONTH, cal.getActualMaximum(Calendar.DAY_OF_MONTH)); // 到月末
        cal.add(Calendar.DAY_OF_MONTH, 1);
        cal.add(Calendar.MILLISECOND, -1);

        return cal.getTime();
    }

    public static List<Date> getDayList(Date start, Date end) {
        List<Date> list = new ArrayList<Date>();

        Calendar cal = Calendar.getInstance();
        cal.setTime(start);
        while (!cal.getTime().after(end)) {
            list.add(cal.getTime());
            cal.add(Calendar.DAY_OF_MONTH, 1);
        }

        return list;
    }

    public static List<Date> getHourList(Date start, Date end) {
        List<Date> list = new ArrayList<Date>();

        Calendar cal = Calendar.getInstance();
        cal.setTime(start);
        while (!cal.getTime().after(end)) {
            list.add(cal.getTime());
            cal.add(Calendar.MINUTE, 30);
        }

        return list;
    }

    public static List<Date> getDayList(Date dt) {
        Date start = DateUtils.getSartTimeOfMonth(dt);
        Date end = DateUtils.getEndTimeOfMonth(start);

        return DateUtils.getDayList(start, end);
    }

    public static List<Date> getHourList(Date dt) {
        Date start = DateUtils.getStartTimeOfDay(dt);
        Date end = DateUtils.getEndTimeOfDay(start);

        return DateUtils.getHourList(start, end);
    }

    public static Date timeAfterHalfHour(Date dt) {
        Calendar cal = Calendar.getInstance();
        cal.setTime(dt);
        cal.add(Calendar.MINUTE, 30);

        return cal.getTime();
    }

    public static Date timeAfterOneDay(Date dt) {
        Calendar cal = Calendar.getInstance();
        cal.setTime(dt);
        cal.add(Calendar.DAY_OF_MONTH, 1);

        return cal.getTime();
    }

    public static boolean isToday(Date dt) {
        return dt.getTime() < DateUtils.getEndTimeOfDay(new Date()).getTime();
    }

	public static String DiffStr(Date dt1,  Date dt2) {
		long day = 1000*3600*24;
		long hour = 1000*3600;
		long min = 1000*60;
		
		long diff = dt1.getTime() - dt2.getTime();
		if (diff < 0) {
			diff *= -1;
		}
		long days = diff / day;
		long hours = (diff - (days * day)) / hour;
		long mins = (diff - (days * day) - (hours * hour)) / min;
		
		return days + "天" + hours + "小时" + mins + "分钟";
	}

}
