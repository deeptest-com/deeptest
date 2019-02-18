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

package com.ngtesting.platform.tql.query.builder.support.model.result;

import java.sql.Time;
import java.text.SimpleDateFormat;
import java.util.Date;
import java.util.List;

/**
 * ---------------------------------------------------------------------------
 *
 * ---------------------------------------------------------------------------
 * @author: hewei
 * @time:2017/11/1 20:34
 * ---------------------------------------------------------------------------
 */
public class SqlQueryResult extends AbstractResult {
    private String query;
    private List<Object> params;

    /**
     * 构造函数
     * @param queryJson
     * @param query
     * @param params
     */
    public SqlQueryResult(String queryJson, String query, List<Object> params) {
        this.queryJson = queryJson;
        this.query = query;
        this.params = params;
    }

    /**
     * Getter method for property <tt>query</tt>.
     * @return property value of query
     * @author hewei
     */
    @Override
    public String getQuery() {
        return query;
    }

    /**
     * 获取查询语句
     * @param withParams
     * @return
     */
    public String getQuery(boolean withParams) {
        if (withParams) {
            StringBuffer sql = new StringBuffer(query);
            for (Object param : params) {
                int index = sql.indexOf("?");

                StringBuffer str;
                if (param instanceof Time) { // 日期
                    str = new StringBuffer(new SimpleDateFormat("HH:mm:ss").format(param));
                } else if (param instanceof Date || param instanceof java.sql.Date) { // 日期时间
                    str = new StringBuffer(new SimpleDateFormat("yyyy-MM-dd HH:mm:ss").format(param));
                } else {
                    str = new StringBuffer(param.toString());
                }

                // 非数字
                if (!(param instanceof Number)) {
                    str.insert(0, "'");
                    str.append("'");
                }
                sql.replace(index, index + 1, str.toString());
            }
            return sql.toString();
        } else {
            return getQuery();
        }
    }

    /**
     * Getter method for property <tt>params</tt>.
     * @return property value of params
     * @author hewei
     */
    public List<Object> getParams() {
        return params;
    }

    /**
     * to string
     * @return
     */
    @Override
    public String toString() {
        return getQuery(true);
    }
}
