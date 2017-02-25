package cn.linkr.testspace.util;

import java.text.ParseException;
import java.util.ArrayList;
import java.util.Date;
import java.util.HashMap;
import java.util.Iterator;
import java.util.List;
import java.util.Map;

import net.sf.json.JSONArray;
import net.sf.json.JSONObject;
import net.sf.json.JsonConfig;
import net.sf.json.util.CycleDetectionStrategy;

/**
 * JsonUtils
 *
 * @author xuxiang
 * @version $Id$
 * @see
 */

public class JsonUtils {

    /**
     * 获得list的map类型
     *
     * @param src       list 集合
     * @param entityKey 实体关键词
     * @return 返回list 集合
     */
    public static List<Map<String, Object>> toMapList(List<Object> src, String entityKey) {

        List<Map<String, Object>> returnList = new ArrayList<Map<String, Object>>();
        for (Object object : src) {
            Map<String, Object> itemMap = new HashMap<String, Object>();
            itemMap.put(entityKey, object);
            returnList.add(itemMap);
        }
        return returnList;
    }

    /**
     * 将json字符串转成对象
     *
     * @param jsonString json字符串
     * @param pojoCalss  要转成的bean类型
     * @return 返回对象
     */
    public static Object getObject4JsonString(String jsonString, Class<?> pojoCalss) {
        JSONObject jsonObject = JSONObject.fromObject(jsonString);
        Object pojo = JSONObject.toBean(jsonObject, pojoCalss);
        return pojo;
    }

    /**
     * json字符串转成map
     *
     * @param jsonString json 字符串
     * @return 返回map
     */
    public static Map<String, Object> getMap4Json(String jsonString) {
        JSONObject jsonObject = JSONObject.fromObject(jsonString);
        @SuppressWarnings("unchecked")
        Iterator<String> keyIter = jsonObject.keys();
        Map<String, Object> valueMap = new HashMap<String, Object>();
        while (keyIter.hasNext()) {
            String key = (String) keyIter.next();
            Object value = jsonObject.get(key);
            valueMap.put(key, value);
        }
        return valueMap;
    }

    /**
     * json 字符串转成数组
     *
     * @param jsonString json字符串
     * @return 返回数组
     */
    public static Object[] getObjectArray4Json(String jsonString) {
        JSONArray jsonArray = JSONArray.fromObject(jsonString);
        return jsonArray.toArray();
    }

    /**
     * json 字符串转成数组
     *
     * @param jsonString json字符串
     * @param pojoClass  对象类
     * @return 返回list 集合
     */
    public static List<Object> getList4Json(String jsonString, Class<?> pojoClass) {
        JSONArray jsonArray = JSONArray.fromObject(jsonString);
        List<Object> list = new ArrayList<Object>();
        for (int i = 0; i < jsonArray.size(); i++) {
            JSONObject jsonObject = jsonArray.getJSONObject(i);
            Object pojoValue = JSONObject.toBean(jsonObject, pojoClass);
            list.add(pojoValue);
        }
        return list;
    }

    /**
     * json 字符串转成字符串数组
     *
     * @param jsonString json字符串
     * @return 返回字符串数组
     */
    public static String[] getStringArray4Json(String jsonString) {
        JSONArray jsonArray = JSONArray.fromObject(jsonString);
        String[] stringArray = new String[jsonArray.size()];
        for (int i = 0; i < jsonArray.size(); i++) {
            stringArray[i] = jsonArray.getString(i);
        }
        return stringArray;
    }

    /**
     * json 字符串转成Long数组
     *
     * @param jsonString json字符串
     * @return 返回Long数组
     */
    public static Long[] getLongArray4Json(String jsonString) {
        JSONArray jsonArray = JSONArray.fromObject(jsonString);
        Long[] longArray = new Long[jsonArray.size()];
        for (int i = 0; i < jsonArray.size(); i++) {
            longArray[i] = Long.valueOf(jsonArray.getLong(i));
        }
        return longArray;
    }

    /**
     * json 字符串转成Integer数组
     *
     * @param jsonString json字符串
     * @return 返回Integer数组
     */
    public static Integer[] getIntegerArray4Json(String jsonString) {
        JSONArray jsonArray = JSONArray.fromObject(jsonString);
        Integer[] integerArray = new Integer[jsonArray.size()];
        for (int i = 0; i < jsonArray.size(); i++) {
            integerArray[i] = Integer.valueOf(jsonArray.getInt(i));
        }
        return integerArray;
    }

    /**
     * json 字符串转成Date数组
     *
     * @param jsonString json字符串
     * @param dataFormat 时间格式
     * @return 返回Date数组
     * @throws ParseException 异常
     */
    public static Date[] getDateArray4Json(String jsonString, String dataFormat) throws ParseException {
        JSONArray jsonArray = JSONArray.fromObject(jsonString);
        Date[] dateArray = new Date[jsonArray.size()];
        for (int i = 0; i < jsonArray.size(); i++) {
            String dateString = jsonArray.getString(i);
            Date date = DateUtils.str2Date(dateString, dataFormat);
            dateArray[i] = date;
        }
        return dateArray;
    }

    /**
     * json 字符串转成Double数组
     *
     * @param jsonString json字符串
     * @return 返回Double数组
     */
    public static Double[] getDoubleArray4Json(String jsonString) {
        JSONArray jsonArray = JSONArray.fromObject(jsonString);
        Double[] doubleArray = new Double[jsonArray.size()];
        for (int i = 0; i < jsonArray.size(); i++) {
            doubleArray[i] = Double.valueOf(jsonArray.getDouble(i));
        }

        return doubleArray;
    }

    /**
     * java对象转成json字符串
     *
     * @param javaObj 对象
     * @return 返回字符串
     */
    public static String getJsonString4JavaPOJO(Object javaObj) {
        JSONObject json = JSONObject.fromObject(javaObj);
        return json.toString();
    }

    /**
     * java对象转成json字符串
     *
     * @param javaObj    对象
     * @param dataFormat 时间格式
     * @return 返回字符串
     */
    public static String getJsonString4JavaPOJO(Object javaObj, String dataFormat) {
        JsonConfig jsonConfig = configJson(dataFormat);
        JSONObject json = JSONObject.fromObject(javaObj, jsonConfig);
        return json.toString();
    }

    /**
     * json 配置属性，并转化时间格式
     *
     * @param datePattern 时间格式
     * @return 返回JsonConfig
     */
    public static JsonConfig configJson(String datePattern) {
        JsonConfig jsonConfig = new JsonConfig();
        jsonConfig.setIgnoreDefaultExcludes(false);
        jsonConfig.setCycleDetectionStrategy(CycleDetectionStrategy.LENIENT);
        jsonConfig.registerJsonValueProcessor(Date.class, new JsonDateValueProcessor(datePattern));
        return jsonConfig;
    }

    /**
     * json 配置属性，如果有日期类型的，将日期类型格式成字符串，以及排除相关属性
     *
     * @param excludes    排除相关属性
     * @param datePattern 时间格式
     * @return 返回JsonConfig
     */
    public static JsonConfig configJson(String[] excludes, String datePattern) {
        JsonConfig jsonConfig = new JsonConfig();
        jsonConfig.setExcludes(excludes);
        jsonConfig.setIgnoreDefaultExcludes(true);
        jsonConfig.setCycleDetectionStrategy(CycleDetectionStrategy.LENIENT);
        jsonConfig.registerJsonValueProcessor(Date.class, new JsonDateValueProcessor(datePattern));
        return jsonConfig;
    }

    /**
     * json 配置属性，排除相关属性
     *
     * @param excludes 排除相关属性
     * @return 返回JsonConfig
     */
    public static JsonConfig configJson(String[] excludes) {
        JsonConfig jsonConfig = new JsonConfig();
        jsonConfig.setExcludes(excludes);
        jsonConfig.setIgnoreDefaultExcludes(true);
        jsonConfig.setCycleDetectionStrategy(CycleDetectionStrategy.LENIENT);
        jsonConfig.registerJsonValueProcessor(Date.class, new JsonDateValueProcessor("yyyy-MM-dd"));
        return jsonConfig;
    }
}
