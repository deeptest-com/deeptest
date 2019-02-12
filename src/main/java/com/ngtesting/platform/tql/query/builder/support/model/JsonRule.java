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

import com.fasterxml.jackson.databind.ObjectMapper;

import java.util.List;

/**
 * ---------------------------------------------------------------------------
 *
 * ---------------------------------------------------------------------------
 * @author: hewei
 * @time:2017/10/30 16:00
 * ---------------------------------------------------------------------------
 */
public class JsonRule implements IGroup, IRule {
    // --------------------------- Rule --------------------------------------
    private String id;
    private String field;
    private String type;
    private String input;
    private String operator;
    private Object value;
    private Object data;
    private Boolean buildIn;

    // --------------------------- Group --------------------------------------
    private String condition;
    private Boolean not;
    private List<JsonRule> rules;

    /**
     * 判断是否为group
     * @return
     */
    public boolean isGroup() {
        return condition != null;
    }

    /**
     * group
     * @return
     */
    public IGroup toGroup() {
        return this;
    }

    /**
     * rule
     * @return
     */
    public IRule toRule() {
        return this;
    }

    /**
     * to String
     * @return
     */
    @Override
    public String toString() {
        ObjectMapper mapper = new ObjectMapper();
        try {
            return mapper.writeValueAsString(this);
        } catch (Exception e) {
            return super.toString();
        }
    }

    /**
     * Getter method for property <tt>id</tt>.
     * @return property value of id
     * @author hewei
     */
    @Override
    public String getId() {
        return id;
    }

    /**
     * Setter method for property <tt>id</tt>.
     * @param id value to be assigned to property id
     * @author hewei
     */
    @Override
    public void setId(String id) {
        this.id = id;
    }

    /**
     * Getter method for property <tt>field</tt>.
     * @return property value of field
     * @author hewei
     */
    @Override
    public String getField() {
        return field;
    }

    /**
     * Setter method for property <tt>field</tt>.
     * @param field value to be assigned to property field
     * @author hewei
     */
    @Override
    public void setField(String field) {
        this.field = field;
    }

    /**
     * Getter method for property <tt>type</tt>.
     * @return property value of type
     * @author hewei
     */
    @Override
    public String getType() {
        return type;
    }

    /**
     * Setter method for property <tt>type</tt>.
     * @param type value to be assigned to property type
     * @author hewei
     */
    @Override
    public void setType(String type) {
        this.type = type;
    }

    /**
     * Getter method for property <tt>input</tt>.
     * @return property value of input
     * @author hewei
     */
    @Override
    public String getInput() {
        return input;
    }

    /**
     * Setter method for property <tt>input</tt>.
     * @param input value to be assigned to property input
     * @author hewei
     */
    @Override
    public void setInput(String input) {
        this.input = input;
    }

    /**
     * Getter method for property <tt>operator</tt>.
     * @return property value of operator
     * @author hewei
     */
    @Override
    public String getOperator() {
        return operator;
    }

    /**
     * Setter method for property <tt>operator</tt>.
     * @param operator value to be assigned to property operator
     * @author hewei
     */
    @Override
    public void setOperator(String operator) {
        this.operator = operator;
    }

    /**
     * Getter method for property <tt>value</tt>.
     * @return property value of value
     * @author hewei
     */
    @Override
    public Object getValue() {
        return value;
    }

    /**
     * Setter method for property <tt>value</tt>.
     * @param value value to be assigned to property value
     * @author hewei
     */
    @Override
    public void setValue(Object value) {
        this.value = value;
    }

    /**
     * Getter method for property <tt>data</tt>.
     * @return property value of data
     * @author hewei
     */
    @Override
    public Object getData() {
        return data;
    }

    /**
     * Setter method for property <tt>data</tt>.
     * @param data value to be assigned to property data
     * @author hewei
     */
    @Override
    public void setData(Object data) {
        this.data = data;
    }

    /**
     * Getter method for property <tt>condition</tt>.
     * @return property value of condition
     * @author hewei
     */
    @Override
    public String getCondition() {
        return condition;
    }

    /**
     * Setter method for property <tt>condition</tt>.
     * @param condition value to be assigned to property condition
     * @author hewei
     */
    @Override
    public void setCondition(String condition) {
        this.condition = condition;
    }

    /**
     * Getter method for property <tt>not</tt>.
     * @return property value of not
     * @author hewei
     */
    @Override
    public Boolean getNot() {
        return not;
    }

    /**
     * Setter method for property <tt>not</tt>.
     * @param not value to be assigned to property not
     * @author hewei
     */
    @Override
    public void setNot(Boolean not) {
        this.not = not;
    }

    /**
     * Getter method for property <tt>rules</tt>.
     * @return property value of rules
     * @author hewei
     */
    @Override
    public List<JsonRule> getRules() {
        return rules;
    }

    /**
     * Setter method for property <tt>rules</tt>.
     * @param rules value to be assigned to property rules
     * @author hewei
     */
    @Override
    public void setRules(List<JsonRule> rules) {
        this.rules = rules;
    }

    @Override
    public Boolean getBuildIn() {
        return this.buildIn;
    }
    @Override
    public void setBuildIn(Boolean buildIn) {
        this.buildIn = buildIn;
    }
}
