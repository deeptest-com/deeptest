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

package com.ngtesting.platform.tql.query.builder.support.model.sql;

/**
 * ---------------------------------------------------------------------------
 * 操作
 * ---------------------------------------------------------------------------
 * @author: hewei
 * @time:2017/10/31 17:15
 * ---------------------------------------------------------------------------
 */
public class Operation {
    private StringBuffer operate;
    private Object value;
    private Boolean hasValue = true;

    /**
     * 构造函数
     * @param operate
     * @param value
     */
    public Operation(StringBuffer operate, Object value) {
        this.operate = operate;
        this.value = value;
    }

    /**
     * Getter method for property <tt>hasValue</tt>.
     * @return property value of hasValue
     * @author hewei
     */
    public Boolean getHasValue() {
        return hasValue;
    }

    /**
     * Setter method for property <tt>hasValue</tt>.
     * @param hasValue value to be assigned to property hasValue
     * @author hewei
     */
    public void setHasValue(Boolean hasValue) {
        this.hasValue = hasValue;
    }

    /**
     * Getter method for property <tt>operate</tt>.
     * @return property value of operate
     * @author hewei
     */
    public StringBuffer getOperate() {
        return operate;
    }

    /**
     * Setter method for property <tt>operate</tt>.
     * @param operate value to be assigned to property operate
     * @author hewei
     */
    public void setOperate(StringBuffer operate) {
        this.operate = operate;
    }

    /**
     * Getter method for property <tt>value</tt>.
     * @return property value of value
     * @author hewei
     */
    public Object getValue() {
        return value;
    }

    /**
     * Setter method for property <tt>value</tt>.
     * @param value value to be assigned to property value
     * @author hewei
     */
    public void setValue(Object value) {
        this.value = value;
    }
}
