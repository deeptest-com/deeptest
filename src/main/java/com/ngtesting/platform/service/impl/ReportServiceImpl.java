package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.service.ReportService;
import org.springframework.stereotype.Service;

import java.util.*;

@Service
public class ReportServiceImpl extends BaseServiceImpl implements ReportService {
    @Override
    public Map<String, List<Object>> chart_design_progress_by_project(Long projectId, Integer numb) {
        Map<String, List<Object>> map = new LinkedHashMap<>();

        List<Object> xList = new LinkedList<>();
        List<Object> numbList = new LinkedList<>();
        List<Object> totalList = new LinkedList<>();

        List<Object[]> ls = getDao().getListBySQL("{call chart_design_progress_by_project(?,?)}",
                projectId, numb);
        for (Object[] arr : ls) {
            xList.add(arr[0].toString());
            numbList.add(arr[1]==null?0:arr[1]);
            totalList.add(arr[2]==null?0:arr[2]);
        }
        map.put("xList", xList);
        map.put("numbList", numbList);
        map.put("totalList", totalList);

        return map;
    }

    @Override
    public Map<String, List<Object>> chart_excution_process_by_project(Long projectId, Integer numb) {
        List<Object[]> ls = getDao().getListBySQL("{call chart_excution_process_by_project(?,?)}",
                projectId, numb);

        return countByStatus(ls);
    }

    @Override
    public Map<String, List<Object>> chart_execution_result_by_plan(Long planId, Integer numb) {
        Map<String, List<Object>> map = new LinkedHashMap<>();

        List<Object> xList = new LinkedList<>();
        List<Object> numbList = new LinkedList<>();

        List<Object[]> ls = getDao().getListBySQL("{call chart_design_progress_by_project(?,?)}",
                planId, numb);
        for (Object[] arr : ls) {
            xList.add(arr[0].toString());
            numbList.add(arr[1]==null?0:arr[1]);
        }
        map.put("xList", xList);
        map.put("numbList", numbList);

        return map;
    }

    @Override
    public Map<String, List<Object>> chart_execution_process_by_plan(Long planId, Integer numb) {
        List<Object[]> ls = getDao().getListBySQL("{call chart_execution_process_by_plan(?,?)}",
                planId, numb);

        return countByStatus(ls);
    }

    @Override
    public Map<String, List<Object>> chart_execution_progress_by_plan(Long planId, Integer numb) {
        Map<String, List<Object>> map = new LinkedHashMap<>();

        List<Object> xList = new LinkedList<>();
        List<Object> numbList = new LinkedList<>();

        List<Object[]> ls = getDao().getListBySQL("{call chart_execution_progress_by_plan(?,?)}",
                planId, numb);
        Integer exeSum = 0;
        for (Object[] arr : ls) {
            xList.add(arr[0].toString());

            Integer totalNumb = Integer.valueOf(arr[2].toString());
            Integer exeNumb = arr[1]==null?0:Integer.valueOf(arr[1].toString());
            exeSum += exeNumb;
            numbList.add(totalNumb - exeSum);
        }
        map.put("xList", xList);
        map.put("numbList", numbList);

        return map;
    }

    @Override
    public Map<String, List<Object>> countByStatus(List<Object[]> ls) {
        Map<String, List<Object>> map = new LinkedHashMap<>();

        List<Object> xList = new LinkedList<>();
        List<Object> passList = new LinkedList<>();
        List<Object> failList = new LinkedList<>();
        List<Object> blockList = new LinkedList<>();

        String day = null;
        Map<String, Object> dayStatus = new HashMap();
        for (Object[] arr : ls) {
            String dayTemp = arr[0].toString();

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

            String status = arr[1].toString();
            if ("pass".equals(status)) {
                dayStatus.put("pass", arr[2]);
            } else if("fail".equals(status)) {
                dayStatus.put("fail", arr[2]);
            } else if("block".equals(status)) {
                dayStatus.put("block", arr[2]);
            }
            day = dayTemp;
        }

        map.put("xList", xList);
        map.put("passList", passList);
        map.put("failList", failList);
        map.put("blockList", blockList);

        return map;
    }

}

