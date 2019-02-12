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

import com.ngtesting.platform.tql.query.builder.exception.FilterException;
import com.ngtesting.platform.tql.query.builder.support.model.JsonRule;
import com.ngtesting.platform.tql.query.builder.support.model.enums.EnumBuilderType;

/**
 * ---------------------------------------------------------------------------
 * 拦截器
 * ---------------------------------------------------------------------------
 * @author: hewei
 * @time:2017/10/31 16:31
 * ---------------------------------------------------------------------------
 */
public interface IRuleFilter {
    /**
     * 执行拦截器
     * @param jsonRule
     * @param type
     * @throws FilterException
     */
    void doFilter(JsonRule jsonRule, EnumBuilderType type) throws FilterException;
}
