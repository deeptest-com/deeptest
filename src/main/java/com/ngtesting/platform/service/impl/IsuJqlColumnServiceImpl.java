package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.config.ConstantIssue;
import com.ngtesting.platform.dao.IssueFieldDao;
import com.ngtesting.platform.model.IsuFieldDefine;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.IsuJqlColumnService;
import com.ngtesting.platform.service.UserService;
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

    @Autowired
    UserService userService;

    @Autowired
    IssueFieldDao fieldDao;

    @Override
    @Transactional
    public List<IsuJqlColumn> loadColumns(TstUser user) {
        String columnsStr = user.getIssueColumns();
        if (StringUtils.isEmpty(columnsStr)) {
            columnsStr = buildDefaultColStr(user);
        }

        List<String> ls = new ArrayList<>(Arrays.asList(columnsStr.split(",")));
        List<IsuJqlColumn> vos = new LinkedList<>();

        List<IsuFieldDefine> cols = fieldDao.listFileds();
        int i = 0;
        for (IsuFieldDefine col : cols) {
            String code = col.getCode();
            String label = col.getLabel();
            ConstantIssue.IssueFilterType type = col.getType();

            Boolean enable;
            if (ls.size() > 0) {
                if (ls.contains(code)) {
                    enable = true;
                } else {
                    enable = false;
                }
            } else {
                enable = i++ < 5;
            }

            IsuJqlColumn vo = new IsuJqlColumn();
            vo.setCode(code);
            vo.setLabel(label);
            vo.setType(type);

            vo.setDisplay(enable);

            vos.add(vo);
        }

        return vos;
    }

    @Override
    @Transactional
    public String buildDefaultColStr(TstUser user) {
        String ret = "";

        List<IsuFieldDefine> cols = fieldDao.listFileds();
        int i = 0;
        for (IsuFieldDefine col : cols) {
            String code = col.getCode();

            if (i++ > 4) {
                break;
            }

            if (!StringUtils.isEmpty(ret)) {
                ret += ",";
            }
            ret += code;
        }

        userService.saveIssueColumns(ret, user);
        return ret;
    }

}
