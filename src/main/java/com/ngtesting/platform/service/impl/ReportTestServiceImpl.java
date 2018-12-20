package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.dao.ReportTestDao;
import com.ngtesting.platform.model.TstProject;
import com.ngtesting.platform.service.intf.ReportTestService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.*;

@Service
public class ReportTestServiceImpl extends ReportServiceImpl implements ReportTestService {
    @Autowired
    ReportTestDao reportTestDao;

    @Override
    public Map<String, List<Object>> chartDesignProgress(Integer projectId, TstProject.ProjectType type, Integer numb) {
        Map<String, List<Object>> map = new LinkedHashMap<>();

        List<Object> xList = new LinkedList<>();
        List<Object> numbList = new LinkedList<>();
        List<Object> totalList = new LinkedList<>();

        List<Map> ls = reportTestDao.chartDesignProgressByProject(projectId, type.toString(), numb);
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
    public Map<String, List<Object>> chartExcutionProcess(Integer projectId, TstProject.ProjectType type, Integer numb) {
        List<Map> ls = reportTestDao.chartExecutionProcessByProject(projectId, type.toString(), numb);

        return countByStatus(ls);
    }

    @Override
    public List<Map<Object, Object>> chartExecutionResultByPlan(Integer planId) {
        List<Map> ls = reportTestDao.chartExecutionResultByPlan(planId);

        Map<String, String> map = new HashMap();
        for (Map item : ls) {
            map.put(item.get("status").toString(), item.get("count").toString());
        }

        List<Map<Object, Object>> data = orderByExeStatus(map);
        return data;
    }

    @Override
    public Map<String, List<Object>> chartExecutionProcessByPlan(Integer planId, Integer numb) {
        List<Map> ls = reportTestDao.chartExecutionProcessByPlan(planId, numb);

        return countByStatus(ls);
    }

    @Override
    public Map<String, Object> chartExecutionProcessByPlanUser(Integer planId, Integer numb) {
        List<Map> ls = reportTestDao.chartExecutionProcessByPlanUser(planId, numb);

        return countByUser(ls);
    }

    @Override
    public Map<String, Object> chartExecutionProgressByPlan(Integer planId, Integer numb) {
        Map<String, Object> map = new LinkedHashMap<>();
        Map<String, List<Object>> series = new LinkedHashMap<>();

        List<Object> xList = new LinkedList<>();
        List<Object> numbList = new LinkedList<>();

        List<Map> ls = reportTestDao.chartExecutionProgressByPlan(planId, numb);
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

}

