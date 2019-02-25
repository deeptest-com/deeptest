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

package com.ngtesting.platform.tql.query.builder;

import com.ngtesting.platform.tql.query.builder.exception.FilterAddException;
import com.ngtesting.platform.tql.query.builder.exception.ParserAddException;
import com.ngtesting.platform.tql.query.builder.support.builder.AbstractBuilder;
import com.ngtesting.platform.tql.query.builder.support.filter.*;
import com.ngtesting.platform.tql.query.builder.support.parser.IGroupParser;
import com.ngtesting.platform.tql.query.builder.support.parser.IRuleParser;

import java.util.ArrayList;
import java.util.List;

/**
 * ---------------------------------------------------------------------------
 *
 * ---------------------------------------------------------------------------
 * @author: hewei
 * @time:2017/11/2 14:49
 * ---------------------------------------------------------------------------
 */
public abstract class AbstractQueryBuilderFactory {
    protected IGroupParser groupParser; // group parser
    protected List<IRuleFilter> filters;    // filters
    protected List<IRuleParser> ruleParsers;   // rule parser

    /**
     * 构造函数
     */
    public AbstractQueryBuilderFactory() {
        this.filters = new ArrayList<>();
        this.ruleParsers = new ArrayList<>();

        // -------------------------- filter -----------------------------
        filters.add(new ValidateFilter());
        filters.add(new DefaultValueConvertFilter());
        filters.add(new DatetimeConvertFilter());
    }

    /**
     * 获取builder
     * @return
     */
    public abstract AbstractBuilder builder();

    /**
     * add filter
     * @param filter
     */
    public void addFilter(IRuleFilter filter) {
        this.filters.add(filter);
    }

    /**
     * add filter before
     * @param filter
     * @param beforeFilter
     */
    public void addFilterBefore(IRuleFilter filter, Class<? extends IRuleFilter> beforeFilter) {
        int index = getIndexOfClass(filters, beforeFilter);
        if (index == -1) {
            throw new FilterAddException("filter " + beforeFilter.getSimpleName() + " has not been added");
        }
        filters.add(index, filter);
    }

    /**
     * add filter after
     * @param filter
     * @param afterFilter
     */
    public void addFilterAfter(IRuleFilter filter, Class<? extends IRuleFilter> afterFilter) {
        int index = getIndexOfClass(filters, afterFilter);
        if (index == -1) {
            throw new FilterAddException("filter " + afterFilter.getSimpleName() + " has not been added");
        }
        filters.add(index + 1, filter);
    }

    /**
     * replace filter
     * @param filter
     * @param atFilter
     */
    public void addFilterAt(IRuleFilter filter, Class<? extends IRuleFilter> atFilter) {
        int index = getIndexOfClass(filters, atFilter);
        if (index == -1) {
            throw new FilterAddException("filter " + atFilter.getSimpleName() + " has not been added");
        }
        filters.remove(index);
        filters.add(index, filter);
    }

    /**
     * add parser
     * @param parser
     */
    public void addRuleParser(IRuleParser parser) {
        this.ruleParsers.add(parser);
    }

    /**
     * add parser before
     * @param parser
     * @param beforeParser
     */
    public void addRuleParserBefore(IRuleParser parser, Class<? extends IRuleParser> beforeParser) {
        int index = getIndexOfClass(ruleParsers, beforeParser);
        if (index == -1) {
            throw new ParserAddException("parser " + beforeParser.getSimpleName() + " has not been added");
        }
        ruleParsers.add(index, parser);
    }

    /**
     * add parser after
     * @param parser
     * @param afterParser
     */
    public void addRuleParserAfter(IRuleParser parser, Class<? extends IRuleParser> afterParser) {
        int index = getIndexOfClass(ruleParsers, afterParser);
        if (index == -1) {
            throw new ParserAddException("parser " + afterParser.getSimpleName() + " has not been added");
        }
        ruleParsers.add(index + 1, parser);
    }

    /**
     * replace parser
     * @param parser
     * @param atParser
     */
    public void addRuleParserAt(IRuleParser parser, Class<? extends IRuleParser> atParser) {
        int index = getIndexOfClass(ruleParsers, atParser);
        if (index == -1) {
            throw new ParserAddException("parser " + atParser.getSimpleName() + " has not been added");
        }
        ruleParsers.remove(index);
        ruleParsers.add(index, parser);
    }


    /**
     * Getter method for property <tt>filters</tt>.
     * @return property value of filters
     * @author hewei
     */
    public List<IRuleFilter> getFilters() {
        return filters;
    }

    /**
     * Getter method for property <tt>ruleParsers</tt>.
     * @return property value of ruleParsers
     * @author hewei
     */
    public List<IRuleParser> getRuleParsers() {
        return ruleParsers;
    }

    /**
     * Getter method for property <tt>groupParser</tt>.
     * @return property value of groupParser
     * @author hewei
     */
    public IGroupParser getGroupParser() {
        return groupParser;
    }

    /**
     * Setter method for property <tt>groupParser</tt>.
     * @param groupParser value to be assigned to property groupParser
     * @author hewei
     */
    public void setGroupParser(IGroupParser groupParser) {
        this.groupParser = groupParser;
    }

    /**
     * get class index of listAll
     * @param list
     * @param cls
     * @return
     */
    private int getIndexOfClass(List list, Class cls) {
        for (int i = 0; i < list.size(); i++) {
            if (list.get(i).getClass().equals(cls)) {
                return i;
            }
        }
        return -1;
    }
}
