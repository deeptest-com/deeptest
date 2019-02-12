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
import com.ngtesting.platform.tql.query.builder.support.model.IRule;
import com.ngtesting.platform.tql.query.builder.support.model.JsonRule;
import com.ngtesting.platform.tql.query.builder.support.model.enums.EnumBuilderType;
import com.ngtesting.platform.tql.query.builder.support.model.enums.EnumRuleType;
import com.ngtesting.platform.tql.query.builder.support.utils.spring.NumberUtils;

import java.util.ArrayList;
import java.util.List;

/**
 * ---------------------------------------------------------------------------
 *
 * ---------------------------------------------------------------------------
 * @author: hewei
 * @time:2017/11/2 10:03
 * ---------------------------------------------------------------------------
 */
public class DefaultValueConvertFilter implements IRuleFilter {
    @Override
    public void doFilter(JsonRule jsonRule, EnumBuilderType type) throws FilterException {
        if (!jsonRule.isGroup()) {
            IRule rule = jsonRule.toRule();

            if (rule.getValue() != null) {
                try {
                    if (rule.getValue() instanceof List) {
                        List<Object> list = new ArrayList<>();
                        for (Object value : (List) rule.getValue()) {
                            list.add(convert(value, rule.getType()));
                        }
                        rule.setValue(list);
                    } else {
                        rule.setValue(convert(rule.getValue(), rule.getType()));
                    }
                } catch (Exception e) {
                    throw new FilterException(e.getMessage() + " for:" + rule );
                }
            }
        }
    }

    /**
     * convert
     * @param value
     * @param type
     * @return
     */
    private Object convert(Object value, String type) {
        if (EnumRuleType.STRING.equals(type)) {
            if (!(value instanceof String)) {
                return String.valueOf(value);
            }
        } else if (EnumRuleType.DOUBLE.equals(type)) {
            if (!(value instanceof Double)) {
                if (value instanceof Number) {
                    return NumberUtils.convertNumberToTargetClass((Number) value, Double.class);
                } else {
                    return NumberUtils.parseNumber(value.toString(), Double.class);
                }
            }
        } else if (EnumRuleType.INTEGER.equals(type)) {
            if (!(value instanceof Integer)) {
                if (value instanceof Number) {
                    return NumberUtils.convertNumberToTargetClass((Number) value, Integer.class);
                } else {
                    return NumberUtils.parseNumber(value.toString(), Integer.class);
                }
            }
        }
        return value;
    }
}
