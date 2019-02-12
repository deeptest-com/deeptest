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

package com.ngtesting.platform.tql.query.builder.config;

import com.ngtesting.platform.tql.query.builder.support.model.enums.EnumDBType;

/**
 * ---------------------------------------------------------------------------
 *
 * ---------------------------------------------------------------------------
 * @author: hewei
 * @time:2017/11/3 20:43
 * ---------------------------------------------------------------------------
 */
public class SqlQueryBuilderConfig {
    private EnumDBType dbType = EnumDBType.POSTGRE_SQL;   // 数据库类型

    /**
     * Getter method for property <tt>dbType</tt>.
     * @return property value of dbType
     * @author hewei
     */
    public EnumDBType getDbType() {
        return dbType;
    }

    /**
     * Setter method for property <tt>dbType</tt>.
     * @param dbType value to be assigned to property dbType
     * @author hewei
     */
    public void setDbType(EnumDBType dbType) {
        this.dbType = dbType;
    }
}
