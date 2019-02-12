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

package com.ngtesting.platform.tql.query.builder.support.builder;

import com.fasterxml.jackson.databind.DeserializationFeature;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.ngtesting.platform.tql.query.builder.exception.ParserNotFoundException;
import com.ngtesting.platform.tql.query.builder.support.filter.IRuleFilter;
import com.ngtesting.platform.tql.query.builder.support.model.JsonRule;
import com.ngtesting.platform.tql.query.builder.support.model.enums.EnumBuilderType;
import com.ngtesting.platform.tql.query.builder.support.parser.IGroupParser;
import com.ngtesting.platform.tql.query.builder.support.parser.IRuleParser;
import com.ngtesting.platform.tql.query.builder.support.parser.JsonRuleParser;

import java.io.IOException;
import java.util.List;

/**
 * ---------------------------------------------------------------------------
 * 构造类
 * ---------------------------------------------------------------------------
 * @author: hewei
 * @time:2017/10/30 15:44
 * ---------------------------------------------------------------------------
 */
public abstract class AbstractBuilder {
    private static ObjectMapper mapper; // object mapper
    private IGroupParser groupParser;   // group parser
    private List<IRuleParser> ruleParsers;  // rule ruleParsers
    private List<IRuleFilter> ruleFilters;  // rule filters

    static {
        // object mapper
        mapper = new ObjectMapper();

        mapper.configure(DeserializationFeature.FAIL_ON_UNKNOWN_PROPERTIES, false);
    }

    /**
     * 构造函数
     * @param groupParser
     * @param ruleParsers
     * @param ruleFilters
     */
    public AbstractBuilder(IGroupParser groupParser, List<IRuleParser> ruleParsers, List<IRuleFilter> ruleFilters) {
        this.groupParser = groupParser;
        this.ruleParsers = ruleParsers;
        this.ruleFilters = ruleFilters;
    }

    /**
     * 构建
     * @param query
     * @return
     */
    protected Object parse(String query) throws IOException {
        JsonRule jsonRule = mapper.readValue(query, JsonRule.class);
        // json rule parse
        return new JsonRuleParser(getBuilderType(), groupParser, ruleParsers, ruleFilters).parse(jsonRule);
    }

    /**
     * 执行构建
     * @param query
     * @return
     */
    public abstract Object build(String query) throws IOException, ParserNotFoundException;

    /**
     * get builder type
     * @return
     */
    protected abstract EnumBuilderType getBuilderType();
}
