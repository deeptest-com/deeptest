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

package com.ngtesting.platform.tql.query.builder.support.model.enums;

/**
 * ---------------------------------------------------------------------------
 *
 * ---------------------------------------------------------------------------
 * @author: hewei
 * @time:2017/10/31 17:28
 * ---------------------------------------------------------------------------
 */
public enum EnumOperator {
    FULLTEXT("fulltext"),
    EQUAL("equal"),
    NOT_EQUAL("not_equal"),
    IN("in"),
    NOT_IN("not_in"),
    LESS("less"),
    LESS_OR_EQUAL("less_or_equal"),
    GREATER("greater"),
    GREATER_OR_EQUAL("greater_or_equal"),
    BETWEEN("between"),
    NOT_BETWEEN("not_between"),
    BEGINS_WITH("begins_with"),
    NOT_BEGINS_WITH("not_begins_with"),
    CONTAINS("contains"),
    NOT_CONTAINS("not_contains"),
    ENDS_WITH("ends_with"),
    NOT_ENDS_WITH("not_ends_with"),
    IS_EMPTY("is_empty"),
    IS_NOT_EMPTY("is_not_empty"),
    IS_NULL("is_null"),
    IS_NOT_NULL("is_not_null");


    private final String value;

    /**
     * 构造函数
     * @param value
     */
    EnumOperator(String value) {
        this.value = value;
    }

    /**
     * Getter method for property <tt>value</tt>.
     * @return property value of value
     * @author hewei
     */
    public String getValue() {
        return value;
    }

    /**
     * Getter method for property <tt>value</tt>.
     * @return property value of value
     * @author hewei
     */
    public String value() {
        return value;
    }

    /**
     * 比较
     * @param value
     * @return
     */
    public boolean equals(String value){
        if (value == null){
            return false;
        }
        return this.value.equals(value);
    }
}
