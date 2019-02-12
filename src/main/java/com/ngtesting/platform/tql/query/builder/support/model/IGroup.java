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

package com.ngtesting.platform.tql.query.builder.support.model;

import java.util.List;

/**
 * ---------------------------------------------------------------------------
 *
 * ---------------------------------------------------------------------------
 * @author: hewei
 * @time:2017/10/31 13:10
 * ---------------------------------------------------------------------------
 */
public interface IGroup {
    /**
     * Getter method for property <tt>condition</tt>.
     * @return property value of condition
     * @author hewei
     */
    String getCondition();

    /**
     * Setter method for property <tt>condition</tt>.
     * @param condition value to be assigned to property condition
     * @author hewei
     */
    void setCondition(String condition);

    /**
     * Getter method for property <tt>not</tt>.
     * @return property value of not
     * @author hewei
     */
    Boolean getNot();

    /**
     * Setter method for property <tt>not</tt>.
     * @param not value to be assigned to property not
     * @author hewei
     */
    void setNot(Boolean not);

    /**
     * Getter method for property <tt>rules</tt>.
     * @return property value of rules
     * @author hewei
     */
    List<JsonRule> getRules();

    /**
     * Setter method for property <tt>rules</tt>.
     * @param rules value to be assigned to property rules
     * @author hewei
     */
    void setRules(List<JsonRule> rules);

    /**
     * Getter method for property <tt>data</tt>.
     * @return property value of value
     * @author hewei
     */
    Object getData();

    /**
     * Setter method for property <tt>data</tt>.
     * @param value value to be assigned to property value
     * @author hewei
     */
    void setData(Object value);
}
