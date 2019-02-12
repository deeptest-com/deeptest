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
 * @time:2017/10/31 18:22
 * ---------------------------------------------------------------------------
 */
public enum EnumCondition {
    AND("AND"),
    OR("OR");

    private final String value;

    /**
     * 构造函数
     * @param value
     */
    EnumCondition(String value) {
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
     * 判断相等
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
