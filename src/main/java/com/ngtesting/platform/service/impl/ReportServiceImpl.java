package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.service.intf.IssueDynamicFormService;
import com.ngtesting.platform.service.intf.ReportService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.*;

@Service
public class ReportServiceImpl extends BaseServiceImpl implements ReportService {

    @Autowired
    IssueDynamicFormService dynamicFormService;

    @Override
    public Map<String, List<Object>> countByStatus(List<Map> ls) {
        Map<String, List<Object>> map = new LinkedHashMap<>();

        List<Object> xList = new LinkedList<>();
        List<Object> passList = new LinkedList<>();
        List<Object> failList = new LinkedList<>();
        List<Object> blockList = new LinkedList<>();

        String day = null;
        Map<String, Object> dayStatus = new HashMap();

        Map last = new HashMap() {{
            put("date", "last");
            put("status", null);
            put("numb", null);
        }};
        ls.add(last);
        for (Map map1 : ls) {
            String dayTemp = map1.get("date").toString();

            if (!dayTemp.equals(day) && day != null) { // 新的一天
                xList.add(day);

                if (!dayStatus.containsKey("pass")) {
                    passList.add(0);
                } else {
                    passList.add(dayStatus.get("pass"));
                }

                if (!dayStatus.containsKey("fail")) {
                    failList.add(0);
                } else {
                    failList.add(dayStatus.get("fail"));
                }

                if (!dayStatus.containsKey("block")) {
                    blockList.add(0);
                } else {
                    blockList.add(dayStatus.get("block"));
                }

                dayStatus = new HashMap();
            }

            // 同一天，对多行内容进行统计
            if (map1.get("status") != null) {
                String status = map1.get("status").toString();
                String numb = map1.get("numb").toString();
                if ("pass".equals(status)) {
                    dayStatus.put("pass", numb);
                } else if("fail".equals(status)) {
                    dayStatus.put("fail", numb);
                } else if("block".equals(status)) {
                    dayStatus.put("block", numb);
                }
            }

            day = dayTemp;
        }

        map.put("xList", xList);
        map.put("passList", passList);
        map.put("failList", failList);
        map.put("blockList", blockList);

        return map;
    }

    @Override
    public Map<String, Object> countByUser(List<Map> ls) {
        Map<String, Object> map = new LinkedHashMap<>();

        List<Object> xList = new LinkedList<>();
        Map<String, List<Object>> byUserMap = new TreeMap<>();

        String day = null;

        Map last = new HashMap() {{
            put("date", "last");
            put("nickname", null);
            put("numb", null);
            put("sum", null);
        }};

        ls.add(last);

        for (Map item : ls) {
            if (item.get("name") != null && !item.get("name").equals("null")) {
                String userName = item.get("name").toString();

                if (!byUserMap.containsKey(userName)) {
                    byUserMap.put(userName, new LinkedList<>());
                }
            }
        }

        Map<String, Object> dayMap = new HashMap();
        for (Map item : ls) {
            String dayTemp = item.get("date").toString();
            String name = item.get("name")!=null?item.get("name").toString():null;
            Integer numb = item.get("numb")!=null?Integer.valueOf(item.get("numb").toString()):null;
            Integer sum =  item.get("sum")!=null?Integer.valueOf(item.get("sum").toString()):null;

            if (!dayTemp.equals(day) && day != null) { // 新的一天
                xList.add(day);

                for (String userName: byUserMap.keySet()) {
                    if (dayMap.containsKey(userName)) {
                        byUserMap.get(userName).add(dayMap.get(userName));
                    } else {
                        byUserMap.get(userName).add(0);
                    }
                }

                dayMap = new HashMap();
            }

            // 同一天，对多行内容进行统计
            if (name != null) {
                dayMap.put(name, numb);
            }

            day = dayTemp;
        }

        map.put("xList", xList);
        map.put("series", byUserMap);

        return map;
    }

    public Map<String, List<Object>> countAgeByPriority(List<Map> ls, Integer numb, Integer orgId, Integer prjId) {
        Map<String, List<Map>> issuePropMap = dynamicFormService.genIssuePropMap(orgId, prjId);

        Map<String, List<Object>> map = new LinkedHashMap<>();
        List<Object> xList = new LinkedList<>();
        map.put("xList", xList);

        Map<String, String> validMap = new HashMap();
        for (Map map1 : ls) {
            String category = map1.get("category").toString();
            String priority = map1.get("priority").toString();
            String number = map1.get("numb").toString();

            String key = category + "-" + priority;
            validMap.put(key, number);
        }

        int round = 0;
        for (Map option : issuePropMap.get("priorityId")) {
            String priority = option.get("label").toString();
            if (!map.containsKey(option)) {
                map.put(priority, new LinkedList());
            }

            for (int i = 1; i <= numb + 1; i++) {
                String category = i <= numb? "" + i: ">" + (i-1);

                if (round == 0) {
                    xList.add(category);
                }

                String key = category + "-" + option.get("label").toString();
                if (validMap.containsKey(key)) {
                    map.get(priority).add(validMap.get(key));
                } else {
                    map.get(priority).add(0);
                }
            }

            round++;
        }

        return map;
    }

    public Map<String, List<Object>> countAge(List<Map> ls, Integer numb) {
        Map<String, List<Object>> map = new LinkedHashMap<>();
        List<Object> xList = new LinkedList<>();
        map.put("xList", xList);

        map.put("问题", new LinkedList());

        Map<String, String> validMap = new HashMap();
        for (Map map1 : ls) {
            String category = map1.get("category").toString();
            String number = map1.get("numb").toString();
            validMap.put(category, number);
        }

        for (int i = 1; i <= numb + 1; i++) {
            String category = i <= numb? "" + i: ">" + (i-1);
            xList.add(category);

            if (validMap.containsKey(category)) {
                map.get("问题").add(validMap.get(category));
            } else {
                map.get("问题").add(0);
            }
        }

        return map;
    }

    @Override
    public List<Map<Object, Object>> orderByExeStatus(Map map) {
        List<Map<Object, Object>> data2 = new LinkedList<>();
        List<String> keys = Arrays.asList("pass","fail","block", "untest");

        for(String key : keys) {
            Map<Object, Object> map2 = new HashMap();
            map2.put("name", Constant.ExeStatus.get(key));
            map2.put("value", map.get(key)!=null?map.get(key): 0);
            data2.add(map2);
        }

        return data2;
    }

}

