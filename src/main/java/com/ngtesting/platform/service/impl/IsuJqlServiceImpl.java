package com.ngtesting.platform.service.impl;

import com.alibaba.fastjson.JSON;
import com.itfsw.query.builder.support.model.JsonRule;
import com.itfsw.query.builder.support.model.enums.EnumOperator;
import com.itfsw.query.builder.support.model.enums.EnumRuleType;
import com.itfsw.query.builder.support.model.result.SqlQueryResult;
import com.ngtesting.platform.dao.IsuTqlDao;
import com.ngtesting.platform.service.IsuJqlBuildService;
import com.ngtesting.platform.service.IsuJqlFilterService;
import com.ngtesting.platform.service.IsuJqlService;
import com.ngtesting.platform.vo.IsuJqlFilter;
import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.io.UnsupportedEncodingException;
import java.net.URLDecoder;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

@Service
public class IsuJqlServiceImpl extends BaseServiceImpl implements IsuJqlService {
    Log logger = LogFactory.getLog(IsuJqlServiceImpl.class);

    @Autowired
    IsuJqlBuildService isuJqlBuildService;

    @Autowired
    IsuJqlFilterService isuJqlFilterService;

    @Autowired
    IsuTqlDao isuTqlDao;

    @Override
    public Map<String, Object> query(String jql, Integer orgId, Integer projectId) {
        Map<String, Object> result = new HashMap();

        if ("".equals(jql)) {
            jql = buildDefaultJql(orgId, projectId);
        } else {
            try {
                jql = URLDecoder.decode(jql, "UTF-8");
            } catch (UnsupportedEncodingException e) {
                e.printStackTrace();
            }
        }

        SqlQueryResult sqlQueryResult = isuJqlBuildService.buildSqlQuery(jql);
        // TODO: 执行查询

        List<IsuJqlFilter> filters = isuJqlFilterService.buildUiFilters(jql, orgId, projectId);

        result.put("result", null);
        result.put("jql", JSON.parseObject(jql));
        result.put("filters", filters);

        return result;
    }

    @Override
    public String buildDefaultJql(Integer orgId, Integer projectId) {
        JsonRule ret = isuJqlBuildService.genJsonRuleRoot();

        JsonRule projectRule = isuJqlBuildService.genJsonRule(
                "project", "projectId", "select", projectId.toString(),
                EnumOperator.EQUAL, EnumRuleType.INTEGER);

        ret.getRules().add(projectRule);

        return JSON.toJSONString(ret);
    }

}
