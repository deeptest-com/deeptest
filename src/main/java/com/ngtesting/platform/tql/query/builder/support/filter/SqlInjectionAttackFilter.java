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

import com.ngtesting.platform.tql.query.builder.exception.FilterAddException;
import com.ngtesting.platform.tql.query.builder.exception.FilterException;
import com.ngtesting.platform.tql.query.builder.support.model.IRule;
import com.ngtesting.platform.tql.query.builder.support.model.JsonRule;
import com.ngtesting.platform.tql.query.builder.support.model.enums.EnumBuilderType;
import com.ngtesting.platform.tql.query.builder.support.model.enums.EnumDBType;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStream;
import java.io.InputStreamReader;
import java.util.HashSet;
import java.util.regex.Pattern;

/**
 * ---------------------------------------------------------------------------
 * SQL Injection Attack Filter
 * ---------------------------------------------------------------------------
 * @author: hewei
 * @time:2017/11/3 16:34
 * ---------------------------------------------------------------------------
 */
public class SqlInjectionAttackFilter implements IRuleFilter {
    private HashSet<String> keywords = new HashSet<>();
    private char beginningDelimiter;
    private char endingDelimiter;
    private EnumDBType dbType;

    /**
     * 构造函数
     * @param dbType
     */
    public SqlInjectionAttackFilter(EnumDBType dbType) {
        String file;
        this.dbType = dbType;
        if (EnumDBType.MYSQL.equals(dbType)) {
            this.beginningDelimiter = '`';
            this.endingDelimiter = '`';
            file = "keywords-mysql.txt";
        } else if (EnumDBType.ORACLE.equals(dbType)) {
            this.beginningDelimiter = '"';
            this.endingDelimiter = '"';
            file = "keywords-oracle.txt";
        } else if (EnumDBType.MS_SQL.equals(dbType)) {
            this.beginningDelimiter = '[';
            this.endingDelimiter = ']';
            file = "keywords-ms-sql.txt";
        } else if (EnumDBType.POSTGRE_SQL.equals(dbType)) {
            this.beginningDelimiter = '"';
            this.endingDelimiter = '"';
            file = "keywords-postgresql.txt";
        } else {
            // TODO not supports now
            throw new FilterAddException("Sorry not supports now");
        }


        try (InputStream inputStream = SqlInjectionAttackFilter.class.getClassLoader().getResourceAsStream(file);
             InputStreamReader inputStreamReader = new InputStreamReader(inputStream);
             BufferedReader bufferedReader = new BufferedReader(inputStreamReader)) {
            String line;
            while ((line = bufferedReader.readLine()) != null) {
                keywords.add(line);
            }
        } catch (IOException e) {
            e.printStackTrace();
        }
    }

    @Override
    public void doFilter(JsonRule jsonRule, EnumBuilderType type) throws FilterException {
        if (!jsonRule.isGroup()) {
            IRule rule = jsonRule.toRule();
            String field = rule.getField();
            if ((EnumDBType.MYSQL.equals(this.dbType) && field.length() > 64) || 
                (EnumDBType.ORACLE.equals(this.dbType) && field.length() > 30) || 
                (EnumDBType.MS_SQL.equals(this.dbType) && field.length() > 128)) {
                // field too long, MYSQL's max length is 64, ORACLE's max length is 30 and MS_SQL's max length is 128
                throw new FilterException("rule's field is too long for:" + jsonRule);
            }
            if (!Pattern.matches("^[A-Za-z0-9_]+$", field)) {
                // can not use Special word
                throw new FilterException("rule's field can only use [A-Za-z0-9_] for:" + jsonRule);
            }
            if (keywords.contains(field.toUpperCase())) {
                // keyword
                StringBuffer sb = new StringBuffer(field);
                sb.insert(0, beginningDelimiter);
                sb.append(endingDelimiter);
                rule.setField(sb.toString());
            }
        }
    }
}
