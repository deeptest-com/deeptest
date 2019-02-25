/*
 * Copyright (c) 2017.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package com.ngtesting.platform.tql.query.builder.support.filter;

import com.ngtesting.platform.tql.query.builder.exception.FilterException;
import com.ngtesting.platform.tql.query.builder.support.model.IGroup;
import com.ngtesting.platform.tql.query.builder.support.model.IRule;
import com.ngtesting.platform.tql.query.builder.support.model.JsonRule;
import com.ngtesting.platform.tql.query.builder.support.model.enums.EnumBuilderType;
import com.ngtesting.platform.tql.query.builder.support.model.enums.EnumOperator;
import com.ngtesting.platform.tql.query.builder.support.utils.spring.CollectionUtils;
import com.ngtesting.platform.tql.query.builder.support.utils.spring.StringUtils;

import java.util.List;

/**
 * ---------------------------------------------------------------------------
 * 验证数据
 * ---------------------------------------------------------------------------
 * @author: hewei
 * @time:2017/11/1 11:32
 * ---------------------------------------------------------------------------
 */
public class ValidateFilter implements IRuleFilter {
    /**
     * 执行拦截器
     * @param jsonRule
     * @throws FilterException
     */
    public void doFilter(JsonRule jsonRule, EnumBuilderType type) throws FilterException {
        if (jsonRule.isGroup()) {
            IGroup group = jsonRule.toGroup();
            if (CollectionUtils.isEmpty(group.getRules())) {
                throw new FilterException("group's rules can not be empty for: " + group);
            }
        } else {
            IRule rule = jsonRule.toRule();
            // field
            if (StringUtils.isEmpty(rule.getField())) {
                throw new FilterException("rule's field can not be empty for:" + rule);
            }
            // operator
            if (StringUtils.isEmpty(rule.getOperator())) {
                throw new FilterException("rule's operator can not be empty for:" + rule);
            }
            // type
            if (StringUtils.isEmpty(rule.getType())) {
                throw new FilterException("rule's type can not be empty for:" + rule);
            }
            // equal
            if (EnumOperator.EQUAL.equals(rule.getOperator())) {
                if (rule.getValue() instanceof List) {
                    List<Object> list = (List<Object>) rule.getValue();
                    if (list.size() > 1) {
                        throw new FilterException("Operator \"equal\" cannot accept multiple values for:" + rule);
                    }
                    rule.setValue(list.get(0));
                }
            }

            // must be listAll
            if (EnumOperator.IN.equals(rule.getOperator()) || EnumOperator.NOT_IN.equals(rule.getOperator())
                    || EnumOperator.BETWEEN.equals(rule.getOperator()) || EnumOperator.NOT_BETWEEN.equals(rule.getOperator())) {
                // listAll
                if (!(rule.getValue() instanceof List)) {
                    throw new FilterException("rule's value must be Array for:" + rule);
                }

                // size
                if (EnumOperator.BETWEEN.equals(rule.getOperator()) || EnumOperator.NOT_BETWEEN.equals(rule.getOperator())) {
                    List list = (List) rule.getValue();
                    if (list.size() != 2) {
                        throw new FilterException("rule's value size must be 2 for:" + rule);
                    }
                }
            }
        }
    }
}
