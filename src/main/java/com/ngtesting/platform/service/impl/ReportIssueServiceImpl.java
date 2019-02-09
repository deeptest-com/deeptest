package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.dao.ProjectDao;
import com.ngtesting.platform.dao.ReportIssueDao;
import com.ngtesting.platform.model.TstProject;
import com.ngtesting.platform.service.intf.ReportIssueService;
import org.apache.commons.lang3.StringUtils;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.*;

@Service
public class ReportIssueServiceImpl extends ReportServiceImpl implements ReportIssueService {
    @Autowired
    ReportIssueDao reportIssueDao;
    @Autowired
    ProjectDao projectDao;

    @Override
    public Map<String, List<Object>> chartIssueTrend(Integer id, TstProject.ProjectType type, Integer numb) {
        List<Map> createLs = reportIssueDao.chartIssueTrendCreate(id, type.toString(), numb);
        List<Map> finalLs = reportIssueDao.chartIssueTrendFinal(id, type.toString(), numb);

        Map<String, List<Object>> map = new LinkedHashMap<>();

        List<Object> xList = new LinkedList<>();
        List<Object> totalListCreate = new LinkedList<>();
        List<Object> totalListFinal = new LinkedList<>();

        Integer countCreate = null;
        Integer countFinal = null;
        for (int i = 0; i < createLs.size(); i++) {
            xList.add(createLs.get(i).get("date").toString());

            countCreate = Integer.valueOf(createLs.get(i).get("sum").toString());
            countFinal = Integer.valueOf(finalLs.get(i).get("sum").toString());

            totalListCreate.add(countCreate);
            totalListFinal.add(countFinal);
        }
        map.put("xList", xList);
        map.put("totalListCreate", totalListCreate);
        map.put("totalListFinal", totalListFinal);

        return map;
    }

    @Override
    public Map<String, List<Object>> chartIssueAgeByProject(Integer projectId, Integer numb, Integer orgId) {
        List<Map> ls = reportIssueDao.chartIssueAge(projectId, "project", numb);

        return countAgeByPriority(ls, numb, orgId, projectId);
    }

    @Override
    public Map<String, List<Object>> chartIssueAgeByOrgOrGroup(Integer id, TstProject.ProjectType type, Integer numb) {
        List<Map> ls = reportIssueDao.chartIssueAge(id, type.toString(), numb);

        return countAge(ls, numb);
    }

    @Override
    public List<Map<Object, Object>> chartIssueDistribByPriority(Integer id, TstProject.ProjectType type) {
        List<Map> ls = reportIssueDao.chartIssueDistribByPriority(id, type.toString());

        List<Map<Object, Object>> data2 = new LinkedList<>();
        for (Map item : ls) {
            Map<Object, Object> map2 = new HashMap();
            map2.put("name", item.get("label"));
            map2.put("value", item.get("count"));
            data2.add(map2);
        }

        return data2;
    }

    @Override
    public List<Map<Object, Object>> chartIssueDistribByStatus(Integer id, TstProject.ProjectType type) {
        List<Map> ls = reportIssueDao.chartIssueDistribByStatus(id, type.toString());

        List<Map<Object, Object>> data2 = new LinkedList<>();
        for (Map item : ls) {
            Map<Object, Object> map2 = new HashMap();
            map2.put("name", item.get("label"));
            map2.put("value", item.get("count"));
            data2.add(map2);
        }

        return data2;
    }

}

