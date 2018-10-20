package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.config.ConstantIssue;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.*;
import com.ngtesting.platform.vo.IsuJqlColumn;
import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;
import org.springframework.util.StringUtils;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.LinkedList;
import java.util.List;

@Service
public class IsuJqlColumnServiceImpl extends BaseServiceImpl implements IsuJqlColumnService {
    Log logger = LogFactory.getLog(IsuJqlColumnServiceImpl.class);

    public static String[] defaultFilters = new String[] {"project", };

    @Autowired
    IsuJqlBuildService isuJqlBuildService;
    @Autowired
    IssueService issueService;

    @Autowired
    IsuJqlFilterService isuJqlFilterService;

    @Autowired
    UserService userService;

    @Override
    @Transactional
    public List<IsuJqlColumn> loadColumns(TstUser user) {
        String columnsStr = user.getIssueColumns();
        if (StringUtils.isEmpty(columnsStr)) {
            columnsStr = buildDefault(user);
        }

        List<String> ls = new ArrayList<>(Arrays.asList(columnsStr.split(",")));
        List<IsuJqlColumn> cols = new LinkedList<>();

        int i = 0;
        for (String id : ConstantIssue.IssueColumns.keySet()) {
            Boolean enable;
            if (ls.size() > 0) {
                if (ls.contains(id)) {
                    enable = true;
                } else {
                    enable = false;
                }
            } else {
                enable = i++ < 5;
            }

            IsuJqlColumn col = new IsuJqlColumn();
            col.setLabel(ConstantIssue.IssueColumns.get(id));
            col.setId(id);
            col.setDisplay(enable);

            cols.add(col);
        }

        return cols;
    }

    @Override
    @Transactional
    public String buildDefault(TstUser user) {
        String ret = "";
        int i = 0;
        for (String id : ConstantIssue.IssueColumns.keySet()) {
            if (i++ > 4) {
                break;
            }

            if (!StringUtils.isEmpty(ret)) {
                ret += ",";
            }
            ret += id;
        }

        userService.saveIssueColumns(ret, user);
        return ret;
    }

}
