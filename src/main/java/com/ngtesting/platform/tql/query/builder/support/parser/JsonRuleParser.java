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

package com.ngtesting.platform.tql.query.builder.support.parser;

import com.ngtesting.platform.tql.query.builder.exception.ParserNotFoundException;
import com.ngtesting.platform.tql.query.builder.support.filter.IRuleFilter;
import com.ngtesting.platform.tql.query.builder.support.model.JsonRule;
import com.ngtesting.platform.tql.query.builder.support.model.enums.EnumBuilderType;

import java.util.List;

/**
 * ---------------------------------------------------------------------------
 * Json rule 解析
 * ---------------------------------------------------------------------------
 * @author: hewei
 * @time:2017/11/9 11:20
 * ---------------------------------------------------------------------------
 */
public class JsonRuleParser {
    private EnumBuilderType builderType;    // builder type
    private IGroupParser groupParser;   // group parser
    private List<IRuleParser> ruleParsers;  // rule ruleParsers
    private List<IRuleFilter> ruleFilters;  // rule filter

    /**
     * 构造函数
     * @param groupParser
     * @param ruleParsers
     */
    public JsonRuleParser(EnumBuilderType builderType, IGroupParser groupParser, List<IRuleParser> ruleParsers, List<IRuleFilter> ruleFilters) {
        this.builderType = builderType;
        this.groupParser = groupParser;
        this.ruleParsers = ruleParsers;
        this.ruleFilters = ruleFilters;
    }

    /**
     * 解析
     * @param jsonRule
     * @return
     */
    public Object parse(JsonRule jsonRule) {
        // filter
        for (IRuleFilter ruleFilter : ruleFilters) {
            ruleFilter.doFilter(jsonRule, builderType);
        }

        // parse
        if (jsonRule.isGroup()) {
            return groupParser.parse(jsonRule, this);
        } else {
            for (IRuleParser ruleParser : ruleParsers) {
                if (ruleParser.canParse(jsonRule)) {
                    return ruleParser.parse(jsonRule, this);
                }
            }

            throw new ParserNotFoundException("Can't found rule parser for:" + jsonRule);
        }
    }
}
