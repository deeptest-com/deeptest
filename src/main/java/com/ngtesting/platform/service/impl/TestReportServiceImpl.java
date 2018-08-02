package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.dao.TestReportDao;
import com.ngtesting.platform.model.TstProject;
import com.ngtesting.platform.service.TestReportService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.*;

@Service
public class TestReportServiceImpl extends BaseServiceImpl implements TestReportService {
    @Autowired
    TestReportDao reportDao;

    @Override
    public Map<String, List<Object>> chartDesignProgressByProject(Integer projectId, TstProject.ProjectType type, Integer numb) {
        Map<String, List<Object>> map = new LinkedHashMap<>();

        List<Object> xList = new LinkedList<>();
        List<Object> numbList = new LinkedList<>();
        List<Object> totalList = new LinkedList<>();

        List<Map> ls = reportDao.chartDesignProgressByProject(projectId, type.toString(), numb);
        Integer sum = null;
        for (Map record : ls) {
            if(sum == null) {
                sum = Integer.valueOf(record.get("sum").toString());
            }
            xList.add(record.get("date").toString());
            numbList.add(record.get("numb"));

            sum += Integer.valueOf(record.get("numb").toString());
            totalList.add(sum);
        }
        map.put("xList", xList);
        map.put("numbList", numbList);
        map.put("totalList", totalList);

        return map;
    }

    @Override
    public Map<String, List<Object>> chartExcutionProcessByProject(Integer projectId, TstProject.ProjectType type, Integer numb) {
        List<Map> ls = reportDao.chartExecutionProcessByProject(projectId, type.toString(), numb);

        return countByStatus(ls);
    }

    @Override
    public List<Map<Object, Object>> chartExecutionResultByPlan(Integer planId) {
        List<Map> ls = reportDao.chartExecutionResultByPlan(planId);

        Map<String, String> map = new HashMap();
        for (Map item : ls) {
            map.put(item.get("status").toString(), item.get("count").toString());
        }

        List<Map<Object, Object>> data = orderByStatus(map);
        return data;
    }

    @Override
    public Map<String, List<Object>> chartExecutionProcessByPlan(Integer planId, Integer numb) {
        List<Map> ls = reportDao.chartExecutionProcessByPlan(planId, numb);

        return countByStatus(ls);
    }

    @Override
    public Map<String, Object> chartExecutionProcessByPlanUser(Integer planId, Integer numb) {
        List<Map> ls = reportDao.chartExecutionProcessByPlanUser(planId, numb);

        return countByUser(ls);
    }

    @Override
    public Map<String, Object> chartExecutionProgressByPlan(Integer planId, Integer numb) {
        Map<String, Object> map = new LinkedHashMap<>();
        Map<String, List<Object>> series = new LinkedHashMap<>();

        List<Object> xList = new LinkedList<>();
        List<Object> numbList = new LinkedList<>();

        List<Map> ls = reportDao.chartExecutionProgressByPlan(planId, numb);
        Integer exeSum = 0;
        int i = 0;
        for (Map item : ls) {
            xList.add(item.get("date").toString());

            Integer totalNumb = Integer.valueOf(item.get("total").toString());
            Integer exeNumb = item.get("numb")==null?0:Integer.valueOf(item.get("numb").toString());
            exeSum += exeNumb;
            numbList.add(totalNumb - exeSum);
        }
        map.put("xList", xList);

        map.put("series", series);
        series.put("剩余用例", numbList);

        return map;
    }

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
        for (Map arr : ls) {
            String dayTemp = arr.get("date").toString();

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
            if (arr.get("status") != null) {
                String status = arr.get("status").toString();
                String numb = arr.get("numb").toString();
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

    @Override
    public List<Map<Object, Object>> orderByStatus(Map map) {
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

