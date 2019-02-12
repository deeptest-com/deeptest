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

package com.ngtesting.platform.tql.query.builder.support.parser.sql;

import com.ngtesting.platform.tql.query.builder.support.model.IGroup;
import com.ngtesting.platform.tql.query.builder.support.model.enums.EnumCondition;
import com.ngtesting.platform.tql.query.builder.support.model.sql.Operation;
import com.ngtesting.platform.tql.query.builder.support.parser.IGroupParser;
import com.ngtesting.platform.tql.query.builder.support.parser.JsonRuleParser;

import java.util.ArrayList;
import java.util.Collection;
import java.util.List;

/**
 * ---------------------------------------------------------------------------
 *
 * ---------------------------------------------------------------------------
 * @author: hewei
 * @time:2017/11/9 11:03
 * ---------------------------------------------------------------------------
 */
public class DefaultGroupParser implements IGroupParser {

    /**
     * 解析
     * @param group
     * @param parser
     * @return
     */
    @Override
    public Object parse(IGroup group, JsonRuleParser parser) {
        StringBuffer operate = new StringBuffer();

        // NOT
        if (group.getNot() != null && group.getNot()) {
            operate.append("( NOT ");
        }

        if (group.getRules().size() > 0) {
            operate.append("( ");
        }

        // rules
        List<Object> params = new ArrayList<Object>();
        for (int i = 0; i < group.getRules().size(); i++) {
            // json parse
            Operation operation = (Operation) parser.parse(group.getRules().get(i));

            // operate
            operate.append(operation.getOperate());
            if (i < group.getRules().size() - 1) {
                operate.append(EnumCondition.AND.equals(group.getCondition()) ? " AND " : " OR ");
            }
            // params
            if (operation.getHasValue()) {
                if (operation.getValue() instanceof List) {
                    params.addAll((Collection<?>) operation.getValue());
                } else {
                    params.add(operation.getValue());
                }
            }
        }

        if (group.getRules().size() > 0) {
            operate.append(" )");
        }
        if (group.getNot() != null && group.getNot()) {
            operate.append(" )");
        }

        return new Operation(operate, params);
    }
}
