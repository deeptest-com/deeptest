package com.ngtesting.platform.service.impl;

import com.alibaba.fastjson.JSON;
import com.itfsw.query.builder.SqlQueryBuilderFactory;
import com.itfsw.query.builder.support.builder.SqlBuilder;
import com.itfsw.query.builder.support.model.JsonRule;
import com.itfsw.query.builder.support.model.enums.EnumCondition;
import com.itfsw.query.builder.support.model.enums.EnumOperator;
import com.itfsw.query.builder.support.model.enums.EnumRuleType;
import com.itfsw.query.builder.support.model.result.SqlQueryResult;
import com.ngtesting.platform.dao.IsuTqlDao;
import com.ngtesting.platform.service.IsuJqlService;
import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.io.IOException;
import java.io.UnsupportedEncodingException;
import java.net.URLDecoder;
import java.util.HashMap;
import java.util.LinkedList;
import java.util.List;
import java.util.Map;

@Service
public class IsuJqlServiceImpl extends BaseServiceImpl implements IsuJqlService {
    Log logger = LogFactory.getLog(IsuJqlServiceImpl.class);

    @Autowired
    IsuTqlDao isuTqlDao;

    @Override
    public Map<String, Object> query(String jql, Integer projectId) {
        Map<String, Object> result = new HashMap();

        if ("".equals(jql)) {
            jql = buildDefault(projectId);
        } else {
            try {
                jql = URLDecoder.decode(jql, "UTF-8");
            } catch (UnsupportedEncodingException e) {
                e.printStackTrace();
            }
        }

        SqlQueryBuilderFactory sqlQueryBuilderFactory = new SqlQueryBuilderFactory();
        SqlBuilder sqlBuilder = sqlQueryBuilderFactory.builder();

        SqlQueryResult sqlQueryResult = null;
        try {
            sqlQueryResult = sqlBuilder.build(jql);
        } catch (IOException e) {
            e.printStackTrace();
        }

        logger.info(sqlQueryResult.getQuery());
        logger.info(sqlQueryResult.getParams());

        result.put("result", null);
        result.put("jql", jql);

        return result;
    }

    @Override
    public String buildDefault(Integer projectId) {
        // String json = "{\"condition\":\"OR\",\"rules\":[{\"id\":\"name\",\"field\":\"username\",\"type\":\"string\",\"input\":\"text\",\"operator\":\"equal\",\"value\":\"Mistic\"}],\"not\":false,\"valid\":true}";

        JsonRule ret = genRootJsonRuleGroup();

        JsonRule projectRule = genJsonRule("project", "projectId", "select", projectId.toString(),
                EnumOperator.EQUAL, EnumRuleType.INTEGER);

        ret.getRules().add(projectRule);

        return JSON.toJSONString(ret);
    }

    @Override
    public JsonRule genJsonRule(String id, String field, String input, String val,
                                     EnumOperator operator, EnumRuleType type) {

        JsonRule rule = new JsonRule();

        rule.setId(id);
        rule.setField(field);
        rule.setInput(input);
        rule.setValue(val);

        rule.setType(type.value());
        rule.setOperator(operator.value());

        return rule;
    }

    @Override
    public JsonRule genJsonRuleGroup(String id, String field, String input, String val,
                                EnumCondition condition, EnumOperator operator, EnumRuleType type) {

        JsonRule rule = genJsonRule(id, field, input, val, operator, type);
        rule.setCondition(condition.value());
        rule.setRules(new LinkedList<>());

        return rule;
    }

    @Override
    public JsonRule genRootJsonRuleGroup() {
        JsonRule ret = new JsonRule();
        ret.setCondition("AND");

        List<JsonRule> rules = new LinkedList<>();
        ret.setRules(rules);

        return ret;
    }

}
