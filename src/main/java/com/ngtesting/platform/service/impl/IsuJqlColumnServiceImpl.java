package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.config.ConstantIssue;
import com.ngtesting.platform.dao.IssueFieldDao;
import com.ngtesting.platform.model.IsuFieldDefine;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.intf.IsuJqlColumnService;
import com.ngtesting.platform.service.intf.UserService;
import com.ngtesting.platform.vo.IsuJqlColumn;
import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;
import org.springframework.util.StringUtils;

import java.util.*;

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

        List<IsuFieldDefine> cols = fieldDao.listDefaultField();

        if (ls.size() > 0) {
            Map<String, IsuFieldDefine> map = new HashMap<>();
            for (IsuFieldDefine col : cols) {
                map.put(col.getCode(), col);
            }

            for (String colCode : ls) {
                IsuFieldDefine col = map.get(colCode);
                IsuJqlColumn vo = new IsuJqlColumn(col.getCode(), col.getLabel(), col.getType(), true);

                vos.add(vo);
            }
        } else {
            int i = 0;
            for (IsuFieldDefine col : cols) {
                Boolean display = i++ < 5;
                IsuJqlColumn vo = new IsuJqlColumn(col.getCode(), col.getLabel(), col.getType(), display);

                vos.add(vo);
            }
        }

        return vos;
    }

    @Override
    @Transactional
    public String buildDefaultColStr(TstUser user) {
        String ret = "";

        List<IsuFieldDefine> cols = fieldDao.listDefaultField();
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
